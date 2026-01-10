-- 1. Add column allowing NULL initially
ALTER TABLE time_record ADD COLUMN internship_id UUID;

-- 2. Update existing records to link to the most appropriate internship
-- First try: matching by date range
UPDATE time_record tr
SET internship_id = (
    SELECT i.id FROM internship i
    WHERE i.student_id = tr.student_id
    AND i.started_in <= tr.date
    AND (i.ended_in IS NULL OR i.ended_in >= tr.date)
    ORDER BY i.created_at DESC
    LIMIT 1
);

-- Second try: for any remaining NULLs, pick the most recent internship for that student
UPDATE time_record tr
SET internship_id = (
    SELECT i.id FROM internship i
    WHERE i.student_id = tr.student_id
    ORDER BY i.created_at DESC
    LIMIT 1
)
WHERE tr.internship_id IS NULL;

-- 3. Delete any remaining records that couldn't be linked to an internship
-- This ensures the NOT NULL constraint below doesn't fail.
DELETE FROM time_record WHERE internship_id IS NULL;

-- 4. Apply constraints
ALTER TABLE time_record ALTER COLUMN internship_id SET NOT NULL;
ALTER TABLE time_record ADD CONSTRAINT time_record_internship_fk FOREIGN KEY (internship_id) REFERENCES internship (id);
