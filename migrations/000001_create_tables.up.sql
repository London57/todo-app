create table if not exists "user"(
	id uuid primary key default gen_random_uuid(),
	name varchar(20) not null,
	username varchar(20) unique,
	email varchar(255) unique not null,
	password text
);


create table todo_lists(
    id uuid primary key default gen_random_uuid(),
    title varchar(255) not null,
    description varchar(255)
);

create table users_lists(
    id serial primary key,
    foreign key(user_id) references "user" (id) on delete cascade,
    foreign key(list_id) references todo_lists (id) on delete cascade
);

create table todo_items(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);


create table lists_items(
    id serial primary key,
    foreign key(item_id) references todo_items (id) on delete cascade,
    foreign key(list_id) references todo_lists (id) on delete cascade
);