
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
	email_confirmed enum('yes', 'no') default 'no' not null,
	mobile_confirmed enum('yes', 'no') default 'no' not null,
	status enum('registered', 'blocked') default 'registered' not null,
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null,
	constraint users_email_uindex
		unique (email),
	constraint users_users_id_fk
		foreign key (parent_id) references users (id)
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
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null,
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
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null,
	constraint user_corporation_user_id_uindex
	unique (user_id),
	constraint user_corporation_users_id_fk
	foreign key (user_id) references users (id)
)
;

create table domains
(
	id int auto_increment
		primary key,
	name varchar(40) not null,
	description varchar(255) null,
	active enum('yes', 'no') default 'yes' not null,
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null,
	constraint domains_name_uindex
	unique (name)
)
;

create table domain_user
(
	domain_id int not null,
	user_id int not null,
	primary key (domain_id, user_id),
	constraint domain_user_domains_id_fk
	foreign key (domain_id) references domains (id),
	constraint domain_user_users_id_fk
	foreign key (user_id) references users (id)
)
;

create index domain_user_users_id_fk
	on domain_user (user_id)
;

create table roles
(
	id int auto_increment
		primary key,
	name varchar(40) not null,
	description varchar(255) null,
	domain_id int not null,
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null,
	constraint roles_name_uindex
	unique (name),
	constraint roles_domains_id_fk
	foreign key (domain_id) references domains (id)
)
;

create index roles_domains_id_fk
	on roles (domain_id)
;



create table role_user
(
	user_id int not null,
	role_id int not null,
	created_at timestamp default CURRENT_TIMESTAMP not null,
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
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null,
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
