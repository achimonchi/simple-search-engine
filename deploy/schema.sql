CREATE TABLE "products" (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    description text NOT NULL,
    price int(11) NOT NULL,
    stock int(11) NOT NULL,
    created_at timestamptz DEFAULT NOW()
);