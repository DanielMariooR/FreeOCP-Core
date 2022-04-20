
CREATE TABLE IF NOT EXISTS User (
    id varchar(255) UNIQUE,
    fullname varchar(255),
    email varchar(255),
    username varchar(255),
    password varchar(255),
    isAdmin boolean
);