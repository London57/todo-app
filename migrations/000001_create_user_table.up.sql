create table if not exists "user"(
	id uuid primary key default gen_random_uuid(),
	name varchar(20) not null,
	username varchar(20) unique,
	email varchar(255) unique not null,
	password text
)