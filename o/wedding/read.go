package wedding

func GetWeddings() ([]*Wedding, error) {
	var result []*Wedding
	var err = weddingTable.FindAll(&result)
	return result, err
}
