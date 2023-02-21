--- Create tables
CREATE TABLE environment (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(20) NOT NULL
);

CREATE TABLE project (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(60) NOT NULL,
  description TEXT NULL,
  parent_id uuid NULL,
  creation_date DATE DEFAULT NOW(),
  due_date DATE DEFAULT NOW(),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  CONSTRAINT project_parent_fk FOREIGN KEY(parent_id) REFERENCES project(id)
);

CREATE TABLE project_environment (
  project_id uuid NOT NULL,
  environment_id uuid NOT NULL,
  CONSTRAINT project_environment_project_fk FOREIGN KEY (project_id) REFERENCES project(id),
  CONSTRAINT project_environment_environment_fk FOREIGN KEY (environment_id) REFERENCES environment(id)
);

CREATE TABLE project_knowledge_area (
  knowledge_area_id uuid NOT NULL,
  project_id uuid NOT NULL,
  CONSTRAINT project_knowledge_area_karea_fk FOREIGN KEY (knowledge_area_id) REFERENCES knowledge_area(id),
  CONSTRAINT project_knowledge_area_project_fk FOREIGN KEY (project_id) REFERENCES project(id)
);

CREATE TABLE project_professional (
  project_id uuid NOT NULL,
  professional_id uuid NOT NULL,
  CONSTRAINT project_professional_project_fk FOREIGN KEY (project_id) REFERENCES project(id),
  CONSTRAINT project_professional_professional_fk FOREIGN KEY (professional_id) REFERENCES professional(id)
);
--- ######################

--- Creating Triggers
CREATE TRIGGER update_entry_updated_at AFTER UPDATE ON project FOR EACH ROW EXECUTE PROCEDURE update_updated_at_prop();
--- ######################

--- Inserting fixtures
COPY environment(id, name)
  FROM '/fixtures/000002/environment.csv'
  DELIMITER ';' csv header;
--- ######################
