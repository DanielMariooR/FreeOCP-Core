CREATE TABLE IF NOT EXISTS course_material (
    _id bigint(20) unsigned AUTO_INCREMENT PRIMARY KEY,
    id varchar(255) UNIQUE,
    course_id varchar(255),
    name varchar(255),
    type varchar(255),
    section_id varchar(255),
    content varchar(255),
    content_text TEXT
);
