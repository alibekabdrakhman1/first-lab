create table users (
    id varchar primary key,
    name varchar(50) not null,
    surname varchar(50) not null,
    login varchar(50) not null unique,
    password varchar(50) not null
)