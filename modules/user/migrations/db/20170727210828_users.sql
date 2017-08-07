
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table users
(
	id int auto_increment
		primary key,
	email varchar(50) not null,
	password varchar(60) not null,
	access_token varchar(60) not null,
	avatar varchar(255) null,
	user_type enum('personal', 'corporation') not null,
	status enum('registered', 'blocked') default 'registered' not null,
	created_at timestamp not null,
	updated_at timestamp not null,
	constraint users_email_uindex
<<<<<<< fa3abc7160a3f63575262da5d189787bdb93a99c
		unique (email),
	constraint users_token_uindex
	unique (access_token)
=======
		unique (email)
>>>>>>> fix migration script in make file
)
;

create table user_personal
(
	id int auto_increment
		primary key,
	user_id int not null,
	first_name varchar(40) null,
	last_name varchar(40) null,
	gender enum('male', 'female') null,
	cellphone varchar(20) null,
	phone varchar(20) null,
	address varchar(255) null,
	city_id int null,
	created_at timestamp not null,
	updated_at timestamp not null,
	constraint user_personal_user_id_uindex
		unique (user_id),
	constraint user_personal_users_id_fk
		foreign key (user_id) references users (id)
)
;

create table user_corporation
(
	id int auto_increment
		primary key,
	user_id int not null,
	name varchar(50) null,
	cellphone varchar(20) null,
	phone varchar(20) null,
	address varchar(255) null,
	economic_code varchar(40) null,
	register_code varchar(40) null,
	city_id int null,
	created_at timestamp not null,
	updated_at timestamp not null,
	constraint user_corporation_user_id_uindex
	unique (user_id),
	constraint user_corporation_users_id_fk
	foreign key (user_id) references users (id)
)
;

create table roles
(
	id int auto_increment
		primary key,
	name varchar(40) not null,
	description varchar(255) null,
	domain_id int not null,
	created_at timestamp not null,
	updated_at timestamp not null,
	constraint roles_name_uindex
	unique (name)
)
;

create index roles_domains_id_fk
	on roles (domain_id)
;



create table role_user
(
	user_id int not null,
	role_id int not null,
	created_at timestamp not null,
	primary key (user_id, role_id),
	constraint role_user_users_id_fk
	foreign key (user_id) references users (id),
	constraint role_user_roles_id_fk
	foreign key (role_id) references roles (id)
)
;

create index role_user_roles_id_fk
	on role_user (role_id)
;

create table role_permission
(
	id int auto_increment
		primary key,
	role_id int not null,
	perm varchar(60) not null,
	scope enum('self', 'parent', 'global') not null,
	created_at timestamp not null,
	updated_at timestamp not null,
	constraint role_permission_roles_id_fk
	foreign key (role_id) references roles (id)
)
;

create index role_permission_roles_id_fk
	on role_permission (role_id)
;





-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE role_permission;
DROP TABLE roles;
DROP TABLE domain_user;
DROP TABLE domains;
DROP TABLE user_corporation;
DROP TABLE user_personal;
DROP TABLE users;
