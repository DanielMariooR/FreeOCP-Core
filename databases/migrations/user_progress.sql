create table user_progress (
    _id bigint(20) unsigned AUTO_INCREMENT PRIMARY KEY,
    user_id varchar(255),
    course_id varchar(255),
    material_id varchar(255),
    score int DEFAULT 100
);