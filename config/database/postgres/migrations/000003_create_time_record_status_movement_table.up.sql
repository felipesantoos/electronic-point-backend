CREATE TABLE IF NOT EXISTS time_record_status_movement (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    time_record_id UUID NOT NULL,
    status_id UUID NOT NULL,
    person_id UUID NOT NULL,
    comments TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    terminated_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT time_record_status_movement_time_record_fk FOREIGN KEY (time_record_id) REFERENCES time_record (id),
    CONSTRAINT time_record_status_movement_status_fk FOREIGN KEY (status_id) REFERENCES time_record_status (id),
    CONSTRAINT time_record_status_movement_person_fk FOREIGN KEY (person_id) REFERENCES person (id)
);

COPY time_record_status_movement (
    id,
    time_record_id,
    status_id,
    person_id,
    comments,
    created_at,
    terminated_at,
    updated_at,
    deleted_at
)
FROM '/fixtures/000003/time_record_status_movement.csv'
DELIMITER ';' CSV HEADER;

