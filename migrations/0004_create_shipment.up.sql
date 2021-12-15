BEGIN;

create table shipment
(
    id         bigserial primary key,
    status     VARCHAR(10)                  not null,
    vehicle_id bigserial REFERENCES vehicle (id)
        ON DELETE CASCADE ON UPDATE CASCADE NOT NULL
);

COMMIT;
