CREATE TABLE IF NOT EXISTS query (
   id SERIAL PRIMARY KEY,
   client_ip VARCHAR ( 50 ) NOT NULL,
   created_at BIGINT NOT NULL,
   domain VARCHAR ( 50 ) NOT NULL
);

CREATE TABLE IF NOT EXISTS address (
  	id SERIAL PRIMARY KEY,
	query_id INTEGER,
  	ip VARCHAR ( 50 ) NOT NULL,
	FOREIGN KEY (query_id)
      		REFERENCES query (id)
);