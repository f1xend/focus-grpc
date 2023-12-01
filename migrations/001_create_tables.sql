-- +goose Up
create table if not exists rabbit
(
    id int NOT NULL,
    color text,
    PRIMARY KEY(id)
);
-- +goose Down
drop table if exists rabbit;