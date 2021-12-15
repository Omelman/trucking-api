BEGIN;

CREATE TABLE "users"
(
    "id"         bigserial primary key,
    "first_name" character varying,
    "last_name"  character varying,
    "role"       varchar NOT NULL,
    "email"      character varying NOT NULL,
    "password"   text NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "language"   varchar NOT NULL
);

COMMIT;
