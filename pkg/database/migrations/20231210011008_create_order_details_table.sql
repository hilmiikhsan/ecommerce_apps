-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "order_details" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "order_id" VARCHAR(255) NOT NULL,
    "total_price_product" INTEGER NOT NULL DEFAULT 0,
    "product_id" INTEGER NOT NULL,
    "quantity" INTEGER NOT NULL DEFAULT 0,
    "created_by" UUID NOT NULL,
    "updated_by" UUID NOT NULL,
    "deleted_by" UUID NOT NULL,
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
DROP TABLE IF EXISTS "order_details";
-- +goose StatementEnd
