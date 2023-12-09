-- +goose Up
-- +goose StatementBegin
CREATE TYPE payment_status AS ENUM ('UNPAID', 'PAID', 'EXPIRED');

CREATE TABLE IF NOT EXISTS "orders" (
    "id" VARCHAR(255) NOT NULL,
    "user_id" UUID NOT NULL,
    "trx_id" VARCHAR(255) NOT NULL,
    "total_price" INTEGER NOT NULL DEFAULT 0,
    "status" payment_status NOT NULL,
    "invoice_url" VARCHAR(255) NOT NULL,
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
DROP TABLE IF EXISTS "orders";
DROP TYPE IF EXISTS payment_status;
-- +goose StatementEnd
