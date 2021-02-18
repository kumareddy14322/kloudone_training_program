package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Transfer struct {
	Id		   			primitive.ObjectID	`bson:"_id"`
	OriginAccount		string 				`json:"originAccount"`
	DestinationAccount  string			 	`json:"destAccount"`
	Amount    			float64 			`json:"amount"`
	Created_at 			time.Time			`json:"createdAt"`
}

func (transfer *Transfer) Save() {
	transfer.Id = primitive.NewObjectID()
	transfer.Created_at = time.Now()
}

func GetTransfers(db *mongo.Client) []*Transfer{
	var transfers []*Transfer
	cur, err := GetCollection(db, "transfers").Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var transfer Transfer
		err := cur.Decode(&transfer)

		if err != nil {
			log.Fatal(err)
		}

		transfers = append(transfers, &transfer)
	}

	cur.Close(context.TODO())
	return transfers
}

func (transfer *Transfer) MakeTransfer(db *mongo.Client, originAccount *Account, destinationAccount *Account) (*mongo.InsertOneResult, string) {
	if transfer.Amount > originAccount.Balance {
		return nil, "Cannot make transfer with value that surpasses origin account's balance."
	}

	transfer.MakeTransaction(db, originAccount, destinationAccount)
	insertResult, err := GetCollection(db, "transfers").InsertOne(context.TODO(), transfer)

	if err != nil {
		return insertResult, err.Error()
	} else {
		return insertResult, ""
	}
}

func (transfer *Transfer) MakeTransaction(db *mongo.Client, originAccount *Account, destinationAccount *Account) {
	newOriginBalance := originAccount.Balance - transfer.Amount
	newDestinationBalance := destinationAccount.Balance + transfer.Amount

	log.Print(newOriginBalance)
	log.Print(newDestinationBalance)

	originAccount.UpdateBalance(db, newOriginBalance)
	destinationAccount.UpdateBalance(db, newDestinationBalance)
}