-- +goose Up
-- +goose StatementBegin
create table segments
(
    "segmentID" integer not null
        constraint segments_pk
            primary key,
    slug        text    not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table segments;
-- +goose StatementEnd