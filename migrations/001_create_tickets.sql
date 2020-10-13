-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tickets (
	ticket_id SERIAL PRIMARY KEY,

    state TEXT,

	created_at TIMESTAMP,
    accepted_at TIMESTAMP,
    prepared_at TIMESTAMP,
    ready_to_pickup_at TIMESTAMP,
    cancelled_at TIMESTAMP,

    CHECK(state IN (
        'CREATED',
        'ACCEPTED',
        'PREPARED',
        'READY_FOR_PICKUP',
        'CANCELLED'
    ))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tickets;
-- +goose StatementEnd