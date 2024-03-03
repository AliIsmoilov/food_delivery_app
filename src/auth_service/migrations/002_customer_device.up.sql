CREATE TABLE IF NOT EXISTS customers_device (
    id uuid PRIMARY KEY NOT NULL,
    customer_id uuid REFERENCES customers(id) ON DELETE CASCADE,
    device_id VARCHAR(255),
    device_name VARCHAR(64),
    notification_id VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX indx_customer_device ON customer_device(device_id);
