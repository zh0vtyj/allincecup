package order

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/zh0vtyj/allincecup-server/internal/db"
	server "github.com/zh0vtyj/allincecup-server/internal/shopping"
)

type Storage interface {
	New(order Info) (uuid.UUID, error)
	GetUserOrders(userId int, createdAt string) ([]FullInfo, error)
	GetOrderById(orderId uuid.UUID) (FullInfo, error)
	GetAdminOrders(status string, lastOrderCreatedAt string) ([]Order, error)
	GetDeliveryTypes() ([]server.DeliveryType, error)
	GetPaymentTypes() ([]server.PaymentType, error)
	ChangeOrderStatus(orderId uuid.UUID, toStatus string) error
}

type storage struct {
	db *sqlx.DB
}

func NewOrdersPostgres(db *sqlx.DB) *storage {
	return &storage{db: db}
}

var orderInfoColumnsInsert = []string{
	"user_lastname",
	"user_firstname",
	"user_middle_name",
	"user_phone_number",
	"user_email",
	"order_comment",
	"order_sum_price",
	"delivery_type_id",
	"payment_type_id",
}

func (o *storage) New(order Info) (uuid.UUID, error) {
	tx, _ := o.db.Begin()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	var deliveryTypeId int
	queryGetDeliveryId := fmt.Sprintf("SELECT id FROM %s WHERE delivery_type_title=$1", db.DeliveryTypesTable)
	err := o.db.Get(&deliveryTypeId, queryGetDeliveryId, order.Order.DeliveryTypeTitle)
	if err != nil {
		return [16]byte{}, fmt.Errorf("failed to create order, delivery type not found %s, error: %v", order.Order.DeliveryTypeTitle, err)
	}

	var paymentTypeId int
	queryGetPaymentTypeId := fmt.Sprintf("SELECT id FROM %s WHERE payment_type_title=$1", db.PaymentTypesTable)
	err = o.db.Get(&paymentTypeId, queryGetPaymentTypeId, order.Order.PaymentTypeTitle)
	if err != nil {
		return [16]byte{}, fmt.Errorf("failed to create order, payment type not found %s, error: %v", order.Order.PaymentTypeTitle, err)
	}

	if order.Order.UserId != 0 {
		orderInfoColumnsInsert = append(orderInfoColumnsInsert, "user_id")
	}

	queryInsertOrder := psql.Insert(db.OrdersTable).Columns(orderInfoColumnsInsert...)

	if order.Order.UserId != 0 {
		queryInsertOrder = queryInsertOrder.Values(
			order.Order.UserLastName,
			order.Order.UserFirstName,
			order.Order.UserMiddleName,
			order.Order.UserPhoneNumber,
			order.Order.UserEmail,
			order.Order.OrderComment,
			order.Order.OrderSumPrice,
			deliveryTypeId,
			paymentTypeId,
			order.Order.UserId,
		)
	} else {
		queryInsertOrder = queryInsertOrder.Values(
			order.Order.UserLastName,
			order.Order.UserFirstName,
			order.Order.UserMiddleName,
			order.Order.UserPhoneNumber,
			order.Order.UserEmail,
			order.Order.OrderComment,
			order.Order.OrderSumPrice,
			deliveryTypeId,
			paymentTypeId,
		)
	}

	queryInsertOrderSql, args, err := queryInsertOrder.ToSql()
	if err != nil {
		return [16]byte{}, fmt.Errorf("failed to build sql query to insert order due to: %v", err)
	}

	var orderId uuid.UUID
	row := tx.QueryRow(queryInsertOrderSql+" RETURNING id", args...)
	if err = row.Scan(&orderId); err != nil {
		_ = tx.Rollback()
		return [16]byte{}, fmt.Errorf("failed to insert new order into table due to: %v", err)
	}

	for _, product := range order.Products {
		queryInsertProducts, args, err := psql.Insert(db.OrdersProductsTable).
			Columns("order_id", "product_id", "quantity", "price_for_quantity").
			Values(orderId, product.ProductId, product.Quantity, product.PriceForQuantity).
			ToSql()
		if err != nil {
			return [16]byte{}, err
		}
		_, err = tx.Exec(queryInsertProducts, args...)
		if err != nil {
			_ = tx.Rollback()
			return [16]byte{}, err
		}
	}

	for _, delivery := range order.Delivery {
		queryInsertDelivery, args, err := psql.Insert(db.OrdersDeliveryTable).
			Columns("order_id", "delivery_title", "delivery_description").
			Values(orderId, delivery.DeliveryTitle, delivery.DeliveryDescription).ToSql()

		_, err = tx.Exec(queryInsertDelivery, args...)
		if err != nil {
			_ = tx.Rollback()
			return [16]byte{}, err
		}
	}

	if order.Order.UserId != 0 {
		queryDeleteCartProducts, args, err := psql.Delete(db.CartsProductsTable).Where(sq.Eq{"cart_id": order.Order.UserId}).ToSql()
		if err != nil {
			_ = tx.Rollback()
			return [16]byte{}, err
		}
		_, err = tx.Exec(queryDeleteCartProducts, args...)
		if err != nil {
			_ = tx.Rollback()
			return [16]byte{}, err
		}
	}

	return orderId, tx.Commit()
}

func (o *storage) GetUserOrders(userId int, createdAt string) ([]FullInfo, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	var ordersAmount int
	queryOrdersAmount, args, err := psql.Select("count(*)").From(db.OrdersTable).Where(sq.Eq{"user_id": userId}).ToSql()
	if err != nil {
		return nil, err
	}

	if err = o.db.Get(&ordersAmount, queryOrdersAmount, args...); err != nil {
		return nil, err
	}

	ordersLimit := 12
	if ordersAmount <= 12 {
		ordersLimit = ordersAmount
	}

	orders := make([]FullInfo, ordersLimit)

	query := psql.Select(
		"order.id",
		"order.user_id",
		"order.user_lastname",
		"order.user_firstname",
		"order.user_middle_name",
		"order.user_phone_number",
		"order.user_email",
		"order.order_status",
		"order.order_comment",
		"order.order_sum_price",
		"delivery_types.delivery_type_title",
		"payment_types.payment_type_title",
		"order.created_at",
		"order.closed_at",
	).
		From(db.OrdersTable).
		LeftJoin(db.DeliveryTypesTable + " ON order.delivery_type_id=delivery_types.id").
		LeftJoin(db.PaymentTypesTable + " ON order.payment_type_id=payment_types.id").
		Where(sq.Eq{"user_id": userId})

	if createdAt != "" {
		query = query.Where(sq.Lt{"created_at": createdAt})
	}

	ordered := query.OrderBy("order.created_at DESC").Limit(12)

	querySql, args, err := ordered.ToSql()
	if err != nil {
		return nil, err
	}

	for i := 0; i < ordersLimit; i++ {
		err = o.db.Get(&orders[i].Info, querySql, args...)
		if err != nil {
			return nil, err
		}
	}

	// TODO "message": "sql: Scan error on column index 1, name \"user_id\": converting NULL to int is unsupported"
	for i := 0; i < ordersLimit; i++ {
		queryOrderProducts, args, err := psql.
			Select(
				"id",
				"order_id",
				"article",
				"product_title",
				"img_url",
				"amount_in_stock",
				"price",
				"units_in_package",
				"packages_in_box",
				"created_at",
				"quantity",
				"price_for_quantity",
			).
			From(db.OrdersProductsTable).
			LeftJoin(db.ProductsTable + " ON orders_products.product_id=products.id").
			Where(sq.Eq{"orders_products.order_id": orders[i].Info.Id}).
			ToSql()
		if err != nil {
			return nil, err
		}

		err = o.db.Select(&orders[i].Products, queryOrderProducts, args...)
		if err != nil {
			return nil, err
		}

		queryOrderDelivery, args, err := psql.
			Select("*").
			From(db.OrdersDeliveryTable).
			Where(sq.Eq{"order_id": orders[i].Info.Id}).
			ToSql()
		if err != nil {
			return nil, err
		}

		err = o.db.Select(&orders[i].Delivery, queryOrderDelivery, args...)
		if err != nil {
			return nil, err
		}
	}

	return orders, err
}

func (o *storage) GetOrderById(orderId uuid.UUID) (FullInfo, error) {
	var order FullInfo

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	queryOrderInfo := psql.
		Select(
			"order.id",
			"order.user_id",
			"order.user_lastname",
			"order.user_firstname",
			"order.user_middle_name",
			"order.user_phone_number",
			"order.user_email",
			"order.order_status",
			"order.order_comment",
			"order.order_sum_price",
			"delivery_types.delivery_type_title",
			"payment_types.payment_type_title",
			"order.created_at",
			"order.closed_at",
		).
		From(db.OrdersTable).
		LeftJoin(db.DeliveryTypesTable + " ON order.delivery_type_id=delivery_types.id").
		LeftJoin(db.PaymentTypesTable + " ON order.payment_type_id=payment_types.id").
		Where(sq.Eq{"order.id": orderId})

	queryOrderInfoSql, args, err := queryOrderInfo.ToSql()
	if err != nil {
		return FullInfo{}, err
	}

	err = o.db.Get(&order.Info, queryOrderInfoSql, args...)
	if err != nil {
		return FullInfo{}, err
	}

	queryProducts := psql.
		Select(
			"orders_products.quantity",
			"orders_products.price_for_quantity",
			"products.id",
			"products.article",
			"products.product_title",
			"products.img_url",
			"products.amount_in_stock",
			"products.price",
			"products.units_in_package",
			"products.packages_in_box",
			"products.created_at",
		).
		From(db.OrdersProductsTable).
		LeftJoin(db.ProductsTable + " ON products.id=orders_products.product_id").
		Where(sq.Eq{"order_id": orderId})

	queryProductsSql, args, err := queryProducts.ToSql()
	if err != nil {
		return FullInfo{}, err
	}

	err = o.db.Select(&order.Products, queryProductsSql, args...)
	if err != nil {
		return FullInfo{}, err
	}

	queryDelivery := psql.Select("*").From(db.OrdersDeliveryTable).Where(sq.Eq{"order_id": orderId})
	queryDeliverySql, args, err := queryDelivery.ToSql()
	if err != nil {
		return FullInfo{}, err
	}

	err = o.db.Select(&order.Delivery, queryDeliverySql, args...)
	if err != nil {
		return FullInfo{}, err
	}

	return order, err
}

func (o *storage) GetAdminOrders(status string, lastOrderCreatedAt string) ([]Order, error) {
	var orders []Order

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	queryOrders := psql.Select(
		"order.id",
		"order.user_lastname",
		"order.user_firstname",
		"order.user_middle_name",
		"order.user_phone_number",
		"order.user_email",
		"order.order_status",
		"order.order_comment",
		"order.order_sum_price",
		"delivery_types.delivery_type_title",
		"payment_types.payment_type_title",
		"order.created_at",
		"order.closed_at",
	).
		From(db.OrdersTable).
		LeftJoin(db.DeliveryTypesTable + " ON order.delivery_type_id=delivery_types.id").
		LeftJoin(db.PaymentTypesTable + " ON order.payment_type_id=payment_types.id")

	if status != "" {
		queryOrders = queryOrders.Where(sq.Eq{"order.order_status": status})
	}

	if lastOrderCreatedAt != "" {
		queryOrders = queryOrders.Where(sq.Lt{"order.created_at": lastOrderCreatedAt})
	}

	queryOrders = queryOrders.OrderBy("order.created_at DESC").Limit(12)

	queryOrdersSql, args, err := queryOrders.ToSql()
	if err != nil {
		return nil, err
	}

	err = o.db.Select(&orders, queryOrdersSql, args...)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *storage) GetDeliveryTypes() (deliveryTypes []server.DeliveryType, err error) {
	queryGetDeliveryTypes := fmt.Sprintf("SELECT * FROM %s", db.DeliveryTypesTable)

	err = o.db.Select(&deliveryTypes, queryGetDeliveryTypes)
	if err != nil {
		return nil, err
	}

	return deliveryTypes, err
}

func (o *storage) GetPaymentTypes() (paymentTypes []server.PaymentType, err error) {
	queryGetPaymentTypes := fmt.Sprintf("SELECT * FROM %s", db.PaymentTypesTable)

	err = o.db.Select(&paymentTypes, queryGetPaymentTypes)
	if err != nil {
		return nil, err
	}

	return paymentTypes, err
}

func (o *storage) ChangeOrderStatus(orderId uuid.UUID, toStatus string) error {
	queryUpdateStatus := fmt.Sprintf("UPDATE %s SET order_status=$1 WHERE id=$2", db.OrdersTable)

	_, err := o.db.Exec(queryUpdateStatus, toStatus, orderId)
	if err != nil {
		return fmt.Errorf("failed to update order status in database due to: %v", err)
	}

	return nil
}