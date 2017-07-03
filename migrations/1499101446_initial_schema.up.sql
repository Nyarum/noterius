BEGIN;

CREATE TABLE users (
	id uuid PRIMARY KEY,
	username text,
	email text,
	password text
);


COMMIT;
