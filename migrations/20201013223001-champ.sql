
-- +migrate Up
create table champ (
   id int not null auto_increment,
   name varchar(255),
   affinity_id int,
   faction_id int,
   primary key(id)
);

-- +migrate Down
drop table champ;

