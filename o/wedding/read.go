package wedding

func GetWeddings() ([]*Wedding, error) {
	var result []*Wedding
	var err = weddingTable.FindAll(&result)
	return result, err
}

func GetWedding(id string) (*Wedding, error) {
	var result *Wedding
	var err = weddingTable.FindID(id, &result)
	return result, err
}

func DeleteWeddingByID(id string) error {
	return weddingTable.RemoveId(id)
}
