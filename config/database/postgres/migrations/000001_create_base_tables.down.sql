--- Dropping Triggers
DROP TRIGGER IF EXISTS update_entry_updated_at ON account;
DROP TRIGGER IF EXISTS update_entry_updated_at ON person;

--- Dropping Tables
DROP TABLE IF EXISTS professional;
DROP TABLE IF EXISTS password_reset;
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS person;
DROP TABLE IF EXISTS account_role;

--- Dropping Function
DROP FUNCTION IF EXISTS update_updated_at_prop;

--- Dropping Extension
DROP EXTENSION IF EXISTS "uuid-ossp";
