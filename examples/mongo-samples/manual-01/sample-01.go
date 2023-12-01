package main

import (
	"context"
	mq "github.com/sourcefellows/mongo-query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongorootuser:mongorootpw@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("manual").Collection("manual-01")
	err = initDB(collection)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("sample 01")
	err = findWithFilter(ctx, collection, bson.D{{"size.uom", "in"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("--")
	err = findWithFilter(ctx, collection, InventoryFilter.Size.Uom.Equals("in"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("sample 02")
	err = findWithFilter(ctx, collection, bson.D{
		{"size.h", bson.D{
			{"$lt", 15},
		}},
		{"size.uom", "in"},
		{"status", "D"},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("--")
	err = findWithFilter(ctx, collection,
		InventoryFilter.Size.H.Lt(15).
			And(InventoryFilter.Size.Uom.Equals("in"),
				InventoryFilter.Status.Equals("D")))
	if err != nil {
		log.Fatal(err)
	}
}

func findWithFilter(ctx context.Context, collection *mongo.Collection, filter any) error {

	cursor, err := collection.Find(
		ctx,
		filter,
	)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return printOnConsole(ctx, cursor)

}

func printOnConsole(ctx context.Context, cursor *mongo.Cursor) error {

	for cursor.Next(ctx) {
		var m map[string]any
		err := cursor.Decode(&m)

		log.Println(m)
		if err != nil {
			return err
		}
	}
	return cursor.Err()
}

func initDB(coll *mongo.Collection) error {
	docs := []interface{}{
		bson.D{
			{"item", "journal"},
			{"qty", 25},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30},
				{"uom", "cm"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"size", bson.D{
				{"h", 10},
				{"w", 15.25},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
	}
	_, err := coll.InsertMany(context.TODO(), docs)
	return err

}

var InventoryFilter = struct {
	Item mq.Field
	Qty  mq.Field
	Size struct {
		H   mq.Field
		W   mq.Field
		Uom mq.Field
	}
	Status mq.Field
}{
	Item: mq.Field("item"),
	Qty:  mq.Field("qty"),
	Size: struct {
		H   mq.Field
		W   mq.Field
		Uom mq.Field
	}{
		H:   mq.Field("size.h"),
		W:   mq.Field("size.w"),
		Uom: mq.Field("size.uom"),
	},
	Status: mq.Field("status"),
}
