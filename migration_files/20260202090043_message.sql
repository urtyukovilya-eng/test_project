-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS message (
	id serial8  PRIMARY KEY,
    chat_id serial8 REFERENCES chat(id),
	"text" varchar NOT NULL,
	created_at TIMESTAMP  NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS message;
-- +goose StatementEnd
