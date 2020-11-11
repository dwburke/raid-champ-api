
-- +migrate Up
create table api_user_access (
   username varchar(255) not null,
   route varchar(128) not null,
   method varchar(7) not null,
   unique (username, route, method)
);

-- +migrate Down
drop table if exists api_user_access;

