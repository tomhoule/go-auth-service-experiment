
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
    id INT PRIMARY KEY,
    email TEXT,
    salt TEXT,
    password TEXT
);

CREATE TABLE capabilities (
    user_id INT REFERENCES user(id),
    is_admin BOOLEAN
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE users;
DROP USER capabilities;
