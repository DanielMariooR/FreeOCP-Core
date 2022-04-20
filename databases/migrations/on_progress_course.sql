CREATE TABLE IF NOT EXISTS on_progress_course (
    user_id varchar(255),
    course_id varchar(255),
    start_date DATETIME DEFAULT CURRENT_TIMESTAMP
);