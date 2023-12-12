-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "categories" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "created_by" UUID NOT NULL,
    "updated_by" UUID NULL,
    "deleted_by" UUID NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL,
    FOREIGN KEY ("created_by") REFERENCES "auth" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("updated_by") REFERENCES "auth" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("deleted_by") REFERENCES "auth" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "categories";
-- +goose StatementEnd
