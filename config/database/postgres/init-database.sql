-- Grant permissions only on the specific database
GRANT CONNECT ON DATABASE eletronic_point TO eletronic_point;
GRANT USAGE, SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO eletronic_point;
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA public TO eletronic_point;

-- Allow the user to have permissions on future tables and sequences
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO eletronic_point;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO eletronic_point;
