-- Drop triggers in reverse order of their creation to avoid dependency issues
DROP TRIGGER IF EXISTS update_entry_updated_at ON account;
DROP TRIGGER IF EXISTS update_entry_updated_at ON person;

-- Drop tables in reverse order of their creation to avoid dependency issues
DROP TABLE IF EXISTS professional;
DROP TABLE IF EXISTS password_reset;
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS person;
DROP TABLE IF EXISTS account_role;

-- Drop functions
DROP FUNCTION IF EXISTS update_updated_at_prop();

-- Drop extensions
DROP EXTENSION IF EXISTS "uuid-ossp";
