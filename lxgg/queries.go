package main

func createContainer(container Container) (int64, error) {
	//TODO i think we can use named exec here?
	res := LXGGDB.MustExec(`INSERT INTO containers (owner_id, name, tags, ip) VALUES ($1, $2, $3, $4)`,
		container.OwnerID, container.Name, container.Tags, container.IP)

	return res.LastInsertId()
}

func getAllContainers() ([]Container, error) {
	var containers []Container
	err := LXGGDB.Select(&containers, `SELECT * FROM containers`)
	return containers, err
}
