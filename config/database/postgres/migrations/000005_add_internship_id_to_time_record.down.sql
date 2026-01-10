ALTER TABLE time_record DROP CONSTRAINT IF EXISTS time_record_internship_fk;
ALTER TABLE time_record DROP COLUMN IF EXISTS internship_id;
