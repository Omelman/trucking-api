BEGIN;

create table vehicle
(
    id                bigserial primary key,
    type              VARCHAR(45)           not null,
    connection_string text                  not null,
    carrying_capacity int                   not null,
    useful_volume     int                   not null,
    created_at        timestamp             not null,
    length            int,
    height            int,
    width             int,
    owner_id          bigserial REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE NOT NULL
);

COMMIT;
