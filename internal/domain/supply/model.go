package supply

import "time"

type InfoDTO struct {
	Id         int        `json:"-" db:"id"`
	Supplier   string     `json:"supplier" db:"supplier" binding:"required"`
	SupplyTime *time.Time `json:"supplyTime" db:"supply_time"`
	Comment    string     `json:"comment" db:"comment"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at"`
}

type PaymentDTO struct {
	PaymentAccount string     `json:"paymentType" db:"payment_account" binding:"required"`
	PaymentTime    *time.Time `json:"paymentTime" db:"payment_time"`
	PaymentSum     float64    `json:"paymentSum" db:"payment_sum" binding:"required"`
}

type ProductDTO struct {
	SupplyId      int     `json:"-" db:"supply_id"`
	ProductId     int     `json:"productId" db:"product_id"`
	Packaging     string  `json:"packaging" db:"packaging"`
	Amount        float64 `json:"amount" db:"amount"`
	PriceForUnit  float64 `json:"priceForUnit" db:"price_for_unit"`
	SumWithoutTax float64 `json:"sumWithoutTax" db:"sum_without_tax"`
	Tax           float64 `json:"tax" db:"tax"`
	TotalSum      float64 `json:"totalSum" db:"total_sum"`
}

type Supply struct {
	Info     InfoDTO      `json:"info"`
	Payment  []PaymentDTO `json:"payment"`
	Products []ProductDTO `json:"products"`
}