drop database golang_rpg;

create database golang_rpg;

use golang_rpg;

create table skills
(
	id int auto_increment
		primary key,
	name varchar(45) not null,
	race varchar(45) null,
	required_level int null
)
;

create table bots
(
	id int auto_increment
		primary key,
	name varchar(45) not null,
	race varchar(45) not null,
	level tinyint(3) unsigned default '1' null,
	user_id int not null,
	experience int default '0' null,
	attack int null,
	defense int null,
	fakka int default '100' null,
	skill1_id int NULL,
	skill2_id int NULL,
	ultSkill_id int NULL,
    maxhp int default '1000' not null,
    maxmp int default '500' not null,
	foreign key (skill1_id) references skills (id),
	foreign key (skill2_id) references skills (id),
	foreign key (ultSkill_id) references skills (id),
	UNIQUE (user_id)
)
;

create index user_id_idx
	on bots (user_id)
;

create table enemies
(
	id int auto_increment
		primary key,
	name varchar(45) not null,
	type int null,
	location varchar(45) null,
	attack int null,
	defense int null,
	pp int null,
	agility int null,
	maxhp int null,
	skill1_id int NULL,
	ultSkill_id int NULL,
	fakka int null,
	drop_item int null,
    power int default '0' null,
	constraint name_UNIQUE
		unique (name)
)
;

create index drop_item
	on enemies (drop_item)
;

create table inventory
(
	id int auto_increment
		primary key,
	bot_id int not null,
	item_id int not null,
	quantity int default '1' null,
	constraint inventory_ibfk_1
		foreign key (bot_id) references bots (id)
)
;

create index bot_id
	on inventory (bot_id)
;

create index item_id
	on inventory (item_id)
;

create table items
(
	id int auto_increment
		primary key,
	required_level int default '0' null,
	name varchar(45) null,
	description varchar(200) null,
	race varchar(45) null,
    type int not null,
	price int default '20' null
)
;

alter table enemies
	add constraint drop_item
		foreign key (drop_item) references items (id)
;

alter table inventory
	add constraint inventory_ibfk_2
		foreign key (item_id) references items (id)
;

create table locations
(
	id int auto_increment
		primary key,
	name varchar(45) null,
	type varchar(45) null
)
;

create table shop_items
(
	id int auto_increment
		primary key,
	location_id int null,
	item_id int null,
	price int null,
	constraint location_id
		foreign key (location_id) references locations (id),
	constraint item_id
		foreign key (item_id) references items (id)
)
;

create index item_id_idx
	on shop_items (item_id)
;

create index location_id_idx
	on shop_items (location_id)
;

create table users
(
	id int auto_increment
		primary key,
	username varchar(45) not null,
	password varchar(45) not null,
	name varchar(45) not null,
	age varchar(45) null,
	constraint username_UNIQUE
		unique (username)
)
;

alter table bots
	add constraint user_id
		foreign key (user_id) references users (id)
			on delete cascade
;

INSERT INTO users (username, password, name, age) VALUES ('EngUser', 'EngPass', 'Eng Guy', 21); #ID 1
INSERT INTO users (username, password, name, age) VALUES ('ArtUser', 'ArtPass', 'Art Guy', 21); # ID 2
INSERT INTO users (username, password, name, age) VALUES ('MgtUser', 'MgtPass', 'Mgt Guy', 21); # ID 3
INSERT INTO users (username, password, name, age) VALUES ('LawUser', 'LawPass', 'Law Guy', 21); # ID 4

INSERT INTO skills (name, race, required_level) VALUES ('engSkill1', 'ENG', 1); #ID 1
INSERT INTO skills (name, race, required_level) VALUES ('artSkill1', 'ART', 1); #ID 2
INSERT INTO skills (name, race, required_level) VALUES ('mgtSkill1', 'MGT', 1); #ID 3
INSERT INTO skills (name, race, required_level) VALUES ('lawSkill1', 'LAW', 1); #ID 4
INSERT INTO skills (name, race, required_level) VALUES ('engSkill4', 'ENG', 4); #ID 5

INSERT INTO bots (name, race, user_id, skill1_id) VALUES ('EngBot', 'ENG', 1, 1); # ID 1
INSERT INTO bots (name, race, user_id, skill1_id) VALUES ('ArtBot', 'ART', 2, 2); # ID 2
INSERT INTO bots (name, race, user_id, skill1_id) VALUES ('MgtBot', 'MGT', 3, 3); # ID 3
INSERT INTO bots (name, race, user_id, skill1_id) VALUES ('LawBot', 'LAW', 4, 4); # ID 4

INSERT INTO locations (name, type) VALUES ('moar''s house', 'type1'); #ID 1
INSERT INTO locations (name, type) VALUES ('location1', 'type2'); #ID 2
INSERT INTO locations (name, type) VALUES ('location2', 'type3'); #ID 3

INSERT INTO items (required_level, name, description, race, type, price) VALUES (1, 'Health Potion', 'Restores 15% of your maximum health', '', 1, 100);
INSERT INTO items (required_level, name, description, race, type, price) VALUES (1, 'axe of awesomeness', 'an axe only engs can use because of how complex it is', 'ENG', 2, 10000);

INSERT INTO shop_items (location_id, item_id, price) VALUES (1, 1, 500000);


INSERT INTO enemies (name, type,  attack, defense, pp, agility, maxhp, fakka, power) VALUES ('Heartless', 1, 10, 10, 10, 10, 100, 100, 50); 					 #Normie Enemies
INSERT INTO enemies (name, type,  attack, defense, pp, agility, maxhp, fakka, power) VALUES ('Nobody', 1, 15, 15, 15, 15, 200, 200, 100);					 	 #Normie Enemies
INSERT INTO enemies (name, type, location,  attack, defense, pp, agility, maxhp, fakka, power) VALUES ('Ansem', 2, 'location1', 20, 20, 20, 20, 1000, 1000, 700); 	 #Boss Enemies
INSERT INTO enemies (name, type, location, attack, defense, pp, agility, maxhp, fakka, power) VALUES ('Xehanort', 2, 'location2', 20, 20, 20, 20, 1500, 2000, 1000);  #Boss Enemies