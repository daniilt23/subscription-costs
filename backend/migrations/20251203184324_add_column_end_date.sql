-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE subscriptions
ADD end_date DATE
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
