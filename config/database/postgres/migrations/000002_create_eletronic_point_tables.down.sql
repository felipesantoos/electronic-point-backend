-- Remove data inserted through COPY statements
DELETE FROM time_record_status;
DELETE FROM student_linked_to_teacher;
DELETE FROM time_record;
DELETE FROM internship;
DELETE FROM student;
DELETE FROM course;
DELETE FROM campus;
DELETE FROM institution;
DELETE FROM internship_location;

-- Drop tables in reverse order of their creation to avoid dependency issues
DROP TABLE IF EXISTS time_record_status;
DROP TABLE IF EXISTS student_linked_to_teacher;
DROP TABLE IF EXISTS time_record;
DROP TABLE IF EXISTS internship;
DROP TABLE IF EXISTS student;
DROP TABLE IF EXISTS course;
DROP TABLE IF EXISTS campus;
DROP TABLE IF EXISTS institution;
DROP TABLE IF EXISTS internship_location;
