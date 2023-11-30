package filter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

var testData = []struct {
	testName            string
	filter              Expression
	expectedResultCount int
}{
	{
		"equals",
		Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999"),
		1,
	},
	{
		"equalsAndEquals empty",
		Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999").
			And(Listing.Name.Equals("Test")),
		0,
	},
	{
		"equalsAndEquals found",
		Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999").
			And(Listing.Name.Equals("Horto flat with small garden")),
		1,
	},
	{
		"equalsOrEquals found",
		Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999").
			Or(Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10021707")),
		2,
	},
	{
		"in",
		Listing.ListingUrl.In("https://www.airbnb.com/rooms/10009999"),
		1,
	},
	{
		"in multiple",
		Listing.ListingUrl.In("https://www.airbnb.com/rooms/10009999", "https://www.airbnb.com/rooms/10021707"),
		2,
	},
	{
		"not in",
		Listing.ListingUrl.NotIn("https://www.airbnb.com/rooms/10009999"),
		5554,
	},
	{
		"gt",
		Listing.Bedrooms.Gt(6),
		16,
	},
	{
		"exists",
		Listing.ListingUrl.Exists(),
		5555,
	},
	{
		"not exists",
		Listing.ListingUrl.NotExists(),
		0,
	}}

func TestField_Equals(t *testing.T) {

	for _, datum := range testData {

		t.Run(datum.testName, func(t *testing.T) {

			//given
			f1 := datum.filter

			//when
			ts, err := query[ListingAndReview](f1)
			if err != nil {
				t.Errorf("%v", err)
			}

			if len(ts) != datum.expectedResultCount {
				t.Errorf("expected %d but got: %d", datum.expectedResultCount, len(ts))
			}

		})
	}

}

func query[T any](filter any) ([]T, error) {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongorootuser:mongorootpw@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("airbnb").Collection("listingsAndReviews")

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	count := 0
	defer cursor.Close(ctx)

	var ret = []T{}

	for cursor.Next(ctx) {
		var l T
		err := cursor.Decode(&l)
		if err != nil {
			return nil, err
		}
		ret = append(ret, l)
		count++
	}
	if cursor.Err() != nil {
		return nil, err
	}

	return ret, nil
}