-- +goose Up
-- +goose StatementBegin
ALTER TABLE tickets 
ADD COLUMN order_id INT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tickets 
DROP COLUMN order_id;
-- +goose StatementEnd