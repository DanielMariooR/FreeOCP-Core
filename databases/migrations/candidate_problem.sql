CREATE TABLE IF NOT EXISTS Candidate_Problem (
	id varchar(255) PRIMARY KEY,
    creator varchar(255) DEFAULT NULL,
    title varchar(255) DEFAULT NULL,
    type varchar(255) DEFAULT NULL,
    topic varchar(255) DEFAULT NULL,
    difficulty varchar(255) DEFAULT NULL,
    status varchar(255) DEFAULT NULL
)
