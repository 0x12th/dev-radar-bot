-- +goose Up
-- +goose StatementBegin
CREATE TABLE sources
(
    id           int           GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name         text          NOT NULL,
    feed_url     text          NOT NULL,
    priority     int           NOT NULL,
    created_at   timestamptz   NOT NULL DEFAULT now(),
    updated_at   timestamptz   NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sources;
-- +goose StatementEnd
