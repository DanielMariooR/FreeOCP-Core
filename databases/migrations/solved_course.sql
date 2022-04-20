CREATE TABLE IF NOT EXISTS solved_course (
    user_id varchar(255),
    course_id varchar(255),
    start_date DATETIME,
    finish_date DATETIME,
    final_score int,
);