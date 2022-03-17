-- +goose Up

create table address
(
    id           serial primary key,
    latitude     numeric      not null default 0,
    longitude    numeric      not null default 0,
    full_address varchar(128) not null default ''
);

create index if not exists coords_idx on address (latitude, longitude);

-- +goose Down

drop table if exists address;