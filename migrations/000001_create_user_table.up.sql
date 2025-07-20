create table if not exists user (
	id uuid primary key default get_random_uuid(),
	name varchar(20) not null,
	username varchar(20) unique not null
)