CREATE TABLE IF NOT EXISTS Course_Enrollment(
  user_id VARCHAR(255),
  course_id VARCHAR(255),
  is_enrolled BIT,
  PRIMARY KEY (user_id, course_id),
  FOREIGN KEY (user_id) REFERENCES User.id,
  FOREIGN KEY (course_id) REFERENCES Course.id
);