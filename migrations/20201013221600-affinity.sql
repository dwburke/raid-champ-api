
-- +migrate Up
create table affinity (
   id int not null auto_increment,
   name varchar(255),
   primary key(id)
);

-- +migrate Down
drop table affinity;

