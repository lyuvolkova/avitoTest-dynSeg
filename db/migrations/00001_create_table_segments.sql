-- +goose Up
-- +goose StatementBegin
create table segments
(
    "segmentID" serial not null
        constraint segments_pk
            primary key,
    slug        text    not null
        constraint segments_slug_uniq_idx
            unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table segments;
-- +goose StatementEnd