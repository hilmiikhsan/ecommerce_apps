-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_gender AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR(255) NOT NULL,
    "date_of_birth" DATE,
    "phone_number" VARCHAR(255) UNIQUE NOT NULL,
    "gender" user_gender NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "is_active" BOOLEAN NOT NULL DEFAULT true,
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
DROP TABLE IF EXISTS "users";
DROP TYPE IF EXISTS user_gender;
-- +goose StatementEnd
