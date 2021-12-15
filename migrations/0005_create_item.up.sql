BEGIN;

create table item
(
    id          bigserial primary key,
    category    VARCHAR(45)                 not null,
    description text,
    quantity    int,
    volume      int,
    weight      int,
    date        timestamp                   not null,
    shipment_id bigserial REFERENCES shipment (id),
    user_id     bigserial REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE NOT NULL
);

create table destination
(
    id      bigserial primary key,
    lat     float                           not null,
    lon     float                           not null,
    item_id bigserial REFERENCES item (id)
        ON DELETE CASCADE ON UPDATE CASCADE NOT NULL
);

COMMIT;
