CREATE TABLE IF NOT EXISTS orders (
    id uuid PRIMARY KEY,
    code VARCHAR(128),
    note TEXT,
    total_price INTEGER CHECK(total_price>0),
    status VARCHAR(128),    -- new -> preparation -> ready -> in-delivery -> delivered -> closed
    restaurant_id UUID NOT NULL,
    customer_id UUID,
    courier_id VARCHAR(56),
    address TEXT,
    lat DECIMAL(7,5) NOT NULL,
    long DECIMAL(7,5) NOT NULL,
    ready_at TIMESTAMP,
    delivery_started_at TIMESTAMP,
    delivered_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS items (
    id uuid PRIMARY KEY,
    item_id uuid NOT NULL,
    order_id UUID REFERENCES orders(id),
    price INTEGER CHECK(price>0),
    qty INTEGER CHECK(qty>0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);