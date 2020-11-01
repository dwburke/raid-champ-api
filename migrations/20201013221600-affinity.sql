
-- +migrate Up
create table affinity (
   id int not null,
   name varchar(255),
   primary key(id)
);

insert into affinity values (1, "Magic");
insert into affinity values (2, "Spirit");
insert into affinity values (3, "Force");
insert into affinity values (4, "Void");

-- +migrate Down
drop table affinity;

