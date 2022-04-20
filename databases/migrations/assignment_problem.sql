CREATE TABLE IF NOT EXISTS assignment_problem (
  assignment_id VARCHAR(255),
  problem_id VARCHAR(255),
  PRIMARY KEY (assignment_id, problem_id),
  FOREIGN KEY (assignment_id) REFERENCES assignment(id),
  FOREIGN KEY (problem_id) REFERENCES Candidate_Problem(id)
);