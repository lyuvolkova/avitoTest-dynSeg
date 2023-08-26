-- +goose Up
-- +goose StatementBegin
create table segment_users
(
    "segmentID" integer not null,
    "userID"    integer not null,
    constraint segment_users_pk
            primary key ("segmentID","userID")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table segment_users;
-- +goose StatementEnd