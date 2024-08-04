package crud

import (
	"clean_embassy_helper/internal/models"
	"context"
	"fmt"
)

func InsertDocument(client *Client, document models.Embassy) (string, error) {

	collection := client.client.Database(client.mgoDB).Collection(client.mgoCollection)

	insertedDoc, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return "", err
	}

	documentID := fmt.Sprintf("%v", insertedDoc.InsertedID)

	return documentID, nil
}
