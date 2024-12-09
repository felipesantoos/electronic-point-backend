CREATE TABLE student (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(200) NOT NULL,
    registration VARCHAR (10) NOT NULL UNIQUE,
    profile_picture TEXT,
    institution VARCHAR(200) NOT NULL,
    course VARCHAR(200) NOT NULL,
    internship_location_name VARCHAR(200) NOT NULL,
    internship_address VARCHAR(200) NOT NULL,
    internship_location VARCHAR(200) NOT NULL,
    total_workload INT,
    CONSTRAINT student_pk PRIMARY KEY (id)
);

CREATE TABLE time_record (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    date DATE NOT NULL,
    entry_time TIMESTAMP NOT NULL,
    exit_time TIMESTAMP NULL,
    location VARCHAR(200) NOT NULL,
    is_off_site BOOLEAN NOT NULL,
    justification TEXT NULL,
    student_id UUID NOT NULL,
    CONSTRAINT time_record_pk PRIMARY KEY (id),
    CONSTRAINT time_record_student_fk FOREIGN KEY (student_id) REFERENCES student (id)
);

COPY student (id, name, registration, profile_picture, institution, course, internship_location_name, internship_address, internship_location, total_workload)
    FROM '/fixtures/000002/student.csv'
    DELIMITER ';' CSV HEADER;

COPY time_record (id, date, entry_time, exit_time, location, is_off_site, justification, student_id)
    FROM '/fixtures/000002/time_record.csv'
    DELIMITER ';' CSV HEADER;

