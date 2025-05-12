CREATE TABLE "user" (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(100),
    password VARCHAR(255),
    permission VARCHAR(50),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    token_confirmation VARCHAR,
    token_confirmation_expiration_date TIMESTAMP WITHOUT TIME ZONE
);


CREATE TABLE "order" (
    id BIGSERIAL PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    transaction_value FLOAT NOT NULL
);
