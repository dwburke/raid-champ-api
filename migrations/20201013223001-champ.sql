
-- +migrate Up
create table champ (
   id serial primary key,
   name varchar(255),
   rarity int default 0,
   affinity_id int,
   faction_id int,
   unique (name)
);

-- +migrate Down
drop table champ;

