-- +goose Up

create table rental_offer
(
    id           serial primary key,
    latitude     numeric      not null default 0,
    longitude    numeric      not null default 0,
    full_address varchar(128) not null default '',
    link         varchar(128) not null default '',
    created      timestamp             default CURRENT_TIMESTAMP
);

create index if not exists rental_offer_coords_idx on rental_offer (latitude, longitude);
create index if not exists rental_offer_full_address_gin_idx on rental_offer using gin (to_tsvector('russian', full_address));
create index rental_offer_created_idx on rental_offer (created);

-- +goose Down

drop table if exists rental_offer;