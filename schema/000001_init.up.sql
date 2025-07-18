CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE todo_lists (
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description varchar(255)
);

CREATE TABLE users_lists (
    id serial PRIMARY KEY,
    user_id integer NOT NULL,
    list_id integer NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);
    id serial PRIMARY KEY,

CREATE TABLE todo_items (
    title varchar(255) NOT NULL,
    description varchar(255),
    done boolean NOT NULL DEFAULT false
);

CREATE TABLE lists_items (
    id serial PRIMARY KEY,
    item_id integer NOT NULL,
    list_id integer NOT NULL,
    FOREIGN KEY(item_id) REFERENCES todo_items(id) ON DELETE CASCADE,
    FOREIGN KEY(list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);