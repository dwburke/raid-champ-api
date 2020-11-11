
-- +migrate Up
create table champ (
   id UUID NOT NULL DEFAULT gen_random_uuid(),
   name varchar(255),
   rarity int default 0,
   affinity_id int,
   faction_id int,
   unique (name)
);

-- +migrate Down
drop table champ;

