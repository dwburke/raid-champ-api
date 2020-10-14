
-- +migrate Up
create table api_user (
   username varchar(255) not null,
   password varchar(128) not null,
   primary key(username)
);

-- +migrate Down
drop table if exists api_user;

