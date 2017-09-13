CREATE TABLE users(
	id			INTEGER PRIMARY KEY,
	username	TEXT,
	password	TEXT,
	admin		INTEGER
);

CREATE TABLE folders(
	id			INTEGER PRIMARY KEY,
	name		TEXT
);

CREATE TABLE containers(
	id			INTEGER PRIMARY KEY,
	owner		INTEGER,
	folder		INTEGER,
	name		TEXT,
	FOREIGN KEY(owner) REFERENCES users(id),
	FOREIGN KEY(folder) REFERENCES folders(id)
);
