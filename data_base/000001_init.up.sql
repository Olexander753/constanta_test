CREATE TABLE users
{
    id serial primary key,
    email varchar(255) not null unique
};

CREATE TABLE transactions
{
    id serial primary key,
    user_id int not null,
    sum int not null,
    currensy varchar(255) not null,
    date_create timestamp,
    date_last_update timestamp,
    status varchar(255)
};

INSERT INTO users(email) values('1@ya.ru');
INSERT INTO users(email) values('2@ya.ru');
INSERT INTO users(email) values('3@ya.ru');
INSERT INTO users(email) values('4@ya.ru');
