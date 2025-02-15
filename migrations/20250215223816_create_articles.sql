-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles
(
    id           int           GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    source_id    int           NOT NULL,
    title        text          NOT NULL,
    summary      text          NULL,
    link         text          NOT NULL UNIQUE,
    published_at timestamptz   NOT NULL,
    created_at   timestamptz   NOT NULL DEFAULT now(),
    posted_at    timestamptz   NULL,
    CONSTRAINT fk_articles_source_id
        FOREIGN KEY (source_id)
            REFERENCES sources (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS articles;
-- +goose StatementEnd