package crud

import (
	"clean_embassy_helper/internal/models"
	"fmt"
)

func InsertDocument(client *Client, document models.Embassy) (string, error) {
	fmt.Printf("Inserting document: %v\n", document)

	client.mgoClientInterface.Collection(client.mgoDB, client.mgoCollection)

	insertedDoc, err := client.mgoClientInterface.InsertOne(document)
	if err != nil {
		return "", err
	}

	return insertedDoc, nil
}
