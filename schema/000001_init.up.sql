CREATE TABLE users
(
    user_id       serial not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE categories
(
    category_id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE transactions
(
    transaction_id serial not null unique,
    type varchar(255) not null,
    date timestamptz not null,
    amount int not null,
    category_id int references categories(category_id) on delete cascade,
    user_id int references users(user_id) on delete cascade
);

CREATE TABLE transaction_category
(
    id      serial                                           not null unique,
    transaction_id int references transactions(transaction_id) on delete cascade,
    category_id int references categories(category_id) on delete cascade
);