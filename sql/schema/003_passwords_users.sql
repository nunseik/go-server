-- +goose Up
alter TABLE users
ADD COLUMN hashed_password TEXT NOT NULL
DEFAULT 'unset';

-- +goose Down
alter TABLE users
DROP COLUMN hashed_password;