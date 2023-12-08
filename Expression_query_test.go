/**
 * MIT License
 *
 * Copyright (c) 2023 Source Fellows GmbH (https://www.source-fellows.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package filter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

const dbConnectionStringForTesting = "mongodb://mongorootuser:mongorootpw@localhost:27017"
const sampleCollection = "listingsAndReviews"

var queryTestData = []struct {
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
	},
	{
		"array contains all",
		Listing.Amenities.ArrayContainsAll("Wifi", "Kitchen", "Iron"),
		3378,
	},
	{
		"array contains exact",
		Listing.Amenities.ArrayContainsExact("Internet",
			"Wifi",
			"Air conditioning",
			"Kitchen",
			"Buzzer/wireless intercom",
			"Heating",
			"Smoke detector",
			"Carbon monoxide detector",
			"Essentials",
			"Lock on bedroom door"),
		1,
	},
	{
		"array contains with query operator",
		Listing.Amenities.ArrayContainsElement(Equals("Wifi")),
		5303,
	},
	{
		"array contains with query operator and 'and condition'",
		Listing.Bedrooms.Gt(8).And(Listing.Amenities.ArrayContainsElement(Equals("Wifi"))),
		6,
	},
	{
		"array size",
		Listing.Amenities.ArraySize(5),
		31,
	},
	{
		"nested field",
		Listing.Images.PictureUrl.Equals("https://a0.muscache.com/im/pictures/5b408b9e-45da-4808-be65-4edc1f29c453.jpg?aki_policy=large"),
		1,
	},
	{
		"complex query",
		Listing.Amenities.ArraySize(15).
			And(Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999"),
				Listing.Bedrooms.Lt(2),
				Listing.Images.PictureUrl.Equals("https://a0.muscache.com/im/pictures/5b408b9e-45da-4808-be65-4edc1f29c453.jpg?aki_policy=large")),
		1,
	},
	{
		"Specify a Query Condition on a Field Embedded in an Array of Documents",
		Review.ReviewerName.Equals("Milo"),
		2,
	},
	{
		"Specify a Query Condition on a Field Embedded in an Array of Documents",
		Review.ElementNo(50).ReviewerName.Equals("Milo"),
		1,
	},
	{
		"Regex Syntax",
		Review.ReviewerName.Regex("Mi.*"),
		1555,
	}}

func TestField_Equals(t *testing.T) {

	for _, datum := range queryTestData {

		t.Run(datum.testName, func(t *testing.T) {

			//given
			f1 := datum.filter

			//when
			ts, err := query[ListingAndReview](sampleCollection, f1)
			if err != nil {
				t.Errorf("%v", err)
			}

			if len(ts) != datum.expectedResultCount {
				bsonElement := f1.bsonD()
				t.Errorf("expected %d but got: %d. expression used: %v | bson: %v", datum.expectedResultCount, len(ts), f1.bsonD(), bsonElement)
			}

		})
	}

}

func query[T any](collectionName string, filter any) ([]T, error) {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionStringForTesting))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("airbnb").Collection(collectionName)

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
