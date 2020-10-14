
-- +migrate Up
create table faction (
   id int not null auto_increment,
   name varchar(255),
   primary key(id)
);

-- +migrate Down
drop table faction;

