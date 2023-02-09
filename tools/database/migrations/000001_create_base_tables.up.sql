CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_updated_at_prop()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- TODO: add trigger for updating updated_at attr when any attr of an entity is changed

CREATE TABLE account_role (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(30) NOT NULL,
  code VARCHAR(30) NOT NULL
);

CREATE TABLE person (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  email VARCHAR(50) UNIQUE,
  name VARCHAR(60) NOT NULL,
  phone VARCHAR(11),
  birth_date DATE NOT NULL,
  cpf VARCHAR(11) UNIQUE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE account (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  email VARCHAR(50) UNIQUE,
  password VARCHAR(512) NOT NULL,
  person_id uuid UNIQUE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  role_id uuid,
  CONSTRAINT person_fk FOREIGN KEY (person_id) REFERENCES person(id),
  CONSTRAINT account_role_fk FOREIGN KEY (role_id) REFERENCES account_role(id)
);

CREATE TABLE password_reset (
  account_id uuid NOT NULL,
  token TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  CONSTRAINT account_fk FOREIGN KEY (account_id) REFERENCES account(id)
);

CREATE TABLE knowledge_area (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name varchar(30)
);

CREATE TABLE professional (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  person_id uuid UNIQUE,
  CONSTRAINT person_fk FOREIGN KEY (person_id) REFERENCES person(id)
);

CREATE TABLE professional_knowledge_areas (
  professional_id uuid NOT NULL,
  knowledge_area_id uuid NOT NULL,
  CONSTRAINT professional_knowledge_areas_prof_fk FOREIGN KEY (professional_id) REFERENCES professional(id),
  CONSTRAINT professional_knowledge_areas_knowledge_areas_fk FOREIGN KEY (knowledge_area_id) REFERENCES knowledge_area(id)
);

CREATE TRIGGER update_entry_updated_at AFTER UPDATE ON account FOR EACH ROW EXECUTE PROCEDURE update_updated_at_prop();
CREATE TRIGGER update_entry_updated_at AFTER UPDATE ON person FOR EACH ROW EXECUTE PROCEDURE update_updated_at_prop();
