BEGIN;

CREATE TABLE players (
	id uuid PRIMARY KEY,
	username text,
	email text,
	password text
);


COMMIT;
