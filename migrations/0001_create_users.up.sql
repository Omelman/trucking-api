BEGIN;

CREATE TABLE "users"
(
    "id"         bigserial primary key,
    "first_name" character varying,
    "last_name"  character varying,
    "email"      character varying NOT NULL,
    "password"   text NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "language"   varchar NOT NULL
);

CREATE TABLE "roles"
(
    "id"         bigserial primary key,
    "name"       character varying,
    "editable"   boolean,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);

ALTER TABLE "users" ADD "role_id" bigint;
CREATE  INDEX  "index_users_on_role_id" ON "users"  ("role_id");
ALTER TABLE "users" ADD CONSTRAINT "fk_role_users"
    FOREIGN KEY ("role_id")
        REFERENCES "roles" ("id");

COMMIT;
