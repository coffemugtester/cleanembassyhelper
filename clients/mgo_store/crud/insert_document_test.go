package crud

//	//TODO: use mock db for testing
//func TestInsertDocument(t *testing.T) {
//
//	CRUDExpectations := func(m *MockClient) {
//		m.On("InsertDocument", models.Embassy{}).Return(&mongo.InsertOneResult{}, nil)
//	}
//
//	mockedClient := NewMockCRUDClient(CRUDExpectations)
//
//	test, err := InsertDocument(mockedClient, models.Embassy{})
//
//	//mgoConfig := conf.MgoConfig{
//	//	MongoUri:        "",
//	//	MongoCollection: "embassy",
//	//	MongoDb:         "embassy",
//	//}
//
//
//	// client, err := NewMongoClient(CRUDClientExpectations)
//	//client := NewCRUDClient(mgoConfig)
//
//	//embassy := models.Embassy{}
//	//if insertedDocID, err := InsertDocument(client, embassy); err != nil {
//	//	t.Error("Expected true, got false")
//	//} else {
//	//	fmt.Printf("INSERTED DOC: %v\n", insertedDocID)
//	//	assert.NotNil(t, embassy)
//	//}
//}
