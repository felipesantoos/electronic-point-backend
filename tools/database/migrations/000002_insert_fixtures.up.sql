COPY account_role(id, name, code)
  FROM '/fixtures/000002/account_role.csv'
  DELIMITER ';' csv header;

COPY person(id, email, name, phone, birth_date, cpf)
  FROM '/fixtures/000002/person.csv'
  DELIMITER ';' csv header;

COPY account(id, email, password, person_id, role_id)
  FROM '/fixtures/000002/account.csv'
  DELIMITER ';' csv header;

COPY knowledge_area(id, name)
  FROM '/fixtures/000002/knowledge_area.csv'
  DELIMITER ';' csv header;

COPY professional(id, person_id)
  FROM '/fixtures/000002/professional.csv'
  DELIMITER ';' csv header;

COPY professional_knowledge_area(professional_id, knowledge_area_id)
  FROM '/fixtures/000002/professional_knowledge_area.csv'
  DELIMITER ';' csv header;
