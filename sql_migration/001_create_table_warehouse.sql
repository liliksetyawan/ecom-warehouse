CREATE TABLE warehouses (
    id BIGSERIAL PRIMARY KEY,
    shop_id BIGINT,
    product_id BIGINT,
    code VARCHAR(16),
    name VARCHAR(256),
    stock BIGINT,
    location VARCHAR(255) NOT NULL,
    status VARCHAR(10),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT uq_wrcode_shop_product UNIQUE (shop_id, product_id, code)
);