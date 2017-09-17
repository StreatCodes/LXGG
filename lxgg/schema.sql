CREATE TABLE users(
	id			INTEGER PRIMARY KEY,
	username	TEXT,
	password	TEXT,
	admin		INTEGER
);

CREATE TABLE containers(
	id			INTEGER PRIMARY KEY,
	owner_id	INTEGER,
	name		TEXT,
	tags		TEXT,
	ip			TEXT,
	FOREIGN KEY(owner_id) REFERENCES users(id)
);
