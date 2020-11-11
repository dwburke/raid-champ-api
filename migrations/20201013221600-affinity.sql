
-- +migrate Up
create table affinity (
   id int not null,
   name varchar(255),
   primary key(id)
);

insert into affinity (id,name) values (1, 'Magic');
insert into affinity (id,name) values (2, 'Spirit');
insert into affinity (id,name) values (3, 'Force');
insert into affinity (id,name) values (4, 'Void');

-- +migrate Down
drop table affinity;

