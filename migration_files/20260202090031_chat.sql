-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chat (
	id serial8 PRIMARY KEY,
	title varchar NOT NULL,
	created_at TIMESTAMP  NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chat;
-- +goose StatementEnd
