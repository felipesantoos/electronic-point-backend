CREATE TABLE internship_location (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(200) NOT NULL,
    address VARCHAR(200) NOT NULL,
    city VARCHAR(100) NOT NULL,
    lat NUMERIC NULL,
    long NUMERIC NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT internship_location_pk PRIMARY KEY (id)
);

CREATE TABLE student (
    id UUID NOT NULL,
    registration VARCHAR(10) NOT NULL UNIQUE,
    profile_picture TEXT NULL,
    institution VARCHAR(200) NOT NULL,
    course VARCHAR(200) NOT NULL,
    total_workload INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT student_pk PRIMARY KEY (id),
    CONSTRAINT student_person_fk FOREIGN KEY (id) REFERENCES person (id),
    CONSTRAINT student_registration_uk UNIQUE (registration)
);

CREATE TABLE internship (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    student_id UUID NOT NULL,
    internship_location_id UUID NOT NULL,
    started_in DATE NOT NULL,
    ended_in DATE NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT internship_pk PRIMARY KEY (id),
    CONSTRAINT internship_student_fk FOREIGN KEY (student_id) REFERENCES student (id),
    CONSTRAINT internship_internship_location_fk FOREIGN KEY (internship_location_id) REFERENCES internship_location (id)
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
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT time_record_pk PRIMARY KEY (id),
    CONSTRAINT time_record_student_fk FOREIGN KEY (student_id) REFERENCES student (id)
);

COPY internship_location (id, name, address, city, lat, long, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/internship_location.csv'
    DELIMITER ';' CSV HEADER;

COPY student (id, registration, profile_picture, institution, course, total_workload, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/student.csv'
    DELIMITER ';' CSV HEADER;

COPY internship (id, student_id, internship_location_id, started_in, ended_in, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/internship.csv'
    DELIMITER ';' CSV HEADER;

COPY time_record (id, date, entry_time, exit_time, location, is_off_site, justification, student_id, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/time_record.csv'
    DELIMITER ';' CSV HEADER;

