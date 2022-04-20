CREATE TABLE IF NOT EXISTS Course (
    id varchar(255) UNIQUE,
    course_name varchar(255),
    description varchar(255),
    thumbnail varchar(255),
    creator varchar(255),
);