CREATE TABLE internship_location (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(200) NOT NULL,
    number VARCHAR(5) NOT NULL,
    street VARCHAR(100) NOT NULL,
    neighborhood VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    zip_code VARCHAR(100) NOT NULL,
    lat NUMERIC NOT NULL,
    long NUMERIC NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT internship_location_pk PRIMARY KEY (id)
);

CREATE TABLE institution (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT institution_pk PRIMARY KEY (id)
);

CREATE TABLE campus (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(200) NOT NULL,
    institution_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT campus_pk PRIMARY KEY (id),
    CONSTRAINT campus_institution_fk FOREIGN KEY (institution_id) REFERENCES institution (id)
);

CREATE TABLE course (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT course_pk PRIMARY KEY (id)
);

CREATE TABLE student (
    id UUID NOT NULL,
    registration VARCHAR(10) NOT NULL UNIQUE,
    profile_picture TEXT NULL,
    campus_id UUID NOT NULL,
    course_id UUID NOT NULL,
    total_workload INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT student_pk PRIMARY KEY (id),
    CONSTRAINT student_person_fk FOREIGN KEY (id) REFERENCES person (id),
    CONSTRAINT student_campus_fk FOREIGN KEY (campus_id) REFERENCES campus (id),
    CONSTRAINT student_course_fk FOREIGN KEY (course_id) REFERENCES course (id),
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
    CONSTRAINT time_record_student_fk FOREIGN KEY (student_id) REFERENCES person (id)
);

CREATE TABLE student_linked_to_teacher (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    student_id UUID NOT NULL,
    teacher_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT student_linked_to_teacher_pk PRIMARY KEY (id),
    CONSTRAINT student_linked_to_teacher_student_fk FOREIGN KEY (student_id) REFERENCES person (id),
    CONSTRAINT student_linked_to_teacher_teacher_fk FOREIGN KEY (teacher_id) REFERENCES person (id)
);

COPY internship_location (id, name, number, street, neighborhood, city, zip_code, lat, long, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/internship_location.csv'
    DELIMITER ';' CSV HEADER;

COPY institution (id, name, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/institution.csv'
    DELIMITER ';' CSV HEADER;

COPY campus (id, name, institution_id, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/campus.csv'
    DELIMITER ';' CSV HEADER;

COPY course (id, name, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/course.csv'
    DELIMITER ';' CSV HEADER;

COPY student (id, registration, profile_picture, campus_id, course_id, total_workload, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/student.csv'
    DELIMITER ';' CSV HEADER;

COPY internship (id, student_id, internship_location_id, started_in, ended_in, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/internship.csv'
    DELIMITER ';' CSV HEADER;

COPY time_record (id, date, entry_time, exit_time, location, is_off_site, justification, student_id, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/time_record.csv'
    DELIMITER ';' CSV HEADER;

COPY student_linked_to_teacher (id, student_id, teacher_id, created_at, updated_at, deleted_at)
    FROM '/fixtures/000002/student_linked_to_teacher.csv'
    DELIMITER ';' CSV HEADER;
