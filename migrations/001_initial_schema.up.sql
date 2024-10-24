CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE base_table (
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
) INHERITS (base_table);

CREATE TABLE items (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    title VARCHAR(255) NOT NULL,
    notes TEXT,
    seller_id uuid,
    price DOUBLE PRECISION,
    FOREIGN KEY (seller_id) REFERENCES users (id) ON DELETE CASCADE
) INHERITS (base_table);

CREATE TABLE purchases (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    buyer_id uuid,
    item_id uuid,
    price DOUBLE PRECISION,
    title VARCHAR(255) NOT NULL,
    FOREIGN KEY (buyer_id) REFERENCES users (id),
    FOREIGN KEY (item_id) REFERENCES items (id)
) INHERITS (base_table)

