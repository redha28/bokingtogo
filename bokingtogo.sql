CREATE TABLE nationality (
  nationality_id SERIAL PRIMARY KEY,
  nationality_name VARCHAR(50) NOT NULL,
  nationality_code CHAR(2) NOT NULL
);

CREATE TABLE customer (
  cst_id SERIAL PRIMARY KEY,
  nationality_id INT NOT NULL,
  cst_name VARCHAR(50) NOT NULL,
  cst_dob DATE NOT NULL,
  cst_phone_num VARCHAR(20) NOT NULL,
  cst_email VARCHAR(50) NOT NULL,
  CONSTRAINT fk_nationality FOREIGN KEY (nationality_id) REFERENCES nationality(nationality_id)
);

CREATE TABLE family (
  fl_id SERIAL PRIMARY KEY,
  cst_id INT NOT NULL,
  fl_relation VARCHAR(50) NOT NULL,
  fl_name VARCHAR(50) NOT NULL,
  fl_dob DATE NOT NULL,
  CONSTRAINT fk_customer FOREIGN KEY (cst_id) REFERENCES customer(cst_id)
);
