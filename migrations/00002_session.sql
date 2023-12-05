-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE sessions (
  id SERIAL PRIMARY KEY,
  user_id INT UNIQUE REFERENCES users (id) ON DELETE CASCADE,
  token_hash TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE sessions;
-- +goose StatementEnd
