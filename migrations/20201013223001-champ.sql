
-- +migrate Up
create table champ (
   id int not null auto_increment,
   name varchar(255),
   rarity int default 0,
   affinity_id int,
   faction_id int,
   primary key(id),
   unique key(name)
);

-- +migrate Down
drop table champ;

