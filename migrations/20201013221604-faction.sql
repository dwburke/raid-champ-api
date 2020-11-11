
-- +migrate Up
create table faction (
   id int not null,
   name varchar(255),
   primary key(id)
);

insert into faction values (1,  'Banner Lords');
insert into faction values (2,  'High Elves');
insert into faction values (3,  'The Sacred Order');
insert into faction values (4,  'Barbarians');
insert into faction values (5,  'Ogryn Tribes');
insert into faction values (6,  'Lizardmen');
insert into faction values (7,  'Skinwalkers');
insert into faction values (8,  'Orcs');
insert into faction values (9,  'Demonspawn');
insert into faction values (10, 'Undead Hordes');
insert into faction values (11, 'Dark Elves');
insert into faction values (12, 'Knight Revenant');
insert into faction values (13, 'Dwarves');

-- +migrate Down
drop table faction;

