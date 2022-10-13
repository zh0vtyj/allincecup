ALTER TABLE categories ADD COLUMN category_description text;

CREATE TABLE categories_filtration
(
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE DEFAULT NULL,
    img_url TEXT,
    info_description TEXT NOT NULL,
    filtration_title TEXT NOT NULL,
    filtration_description TEXT,
    filtration_list_id INT REFERENCES categories_filtration(id) DEFAULT NULL
);

