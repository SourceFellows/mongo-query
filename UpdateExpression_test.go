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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

var updateCollectionName = "clonedForUpdate"

var updateTestData = []struct {
	testName                        string
	filterBeforeUpdate              Expression
	expectedResultCountBeforeUpdate int
	updateFilter                    Expression
	update                          UpdateExpression
	expectedResultCountAfterUpdate  int
}{
	{"simple updateOne",
		Listing.ListingUrl.Equals("http://www.source-fellows.com"),
		0,
		Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999"),
		Listing.ListingUrl.Set("http://www.source-fellows.com"),
		1,
	},
	{"simple updateOne with and condition",
		Listing.ListingUrl.Equals("http://www.source-fellows.com").And(Listing.Name.Equals("Horst")),
		0,
		Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999"),
		Listing.ListingUrl.Set("http://www.source-fellows.com").And(Listing.Name.Set("Horst")),
		1,
	},
}

func TestUpdateExpressions(t *testing.T) {

	for _, datum := range updateTestData {

		t.Run(datum.testName, func(t *testing.T) {

			//given
			err := cloneCollection(updateCollectionName)
			if err != nil {
				t.Errorf("could not clone collection for testing %v", err)
				return
			}
			defer removeCollection(updateCollectionName)

			f1 := datum.filterBeforeUpdate

			//when
			ts, err := query[ListingAndReview](updateCollectionName, f1)
			if err != nil {
				t.Errorf("%v", err)
				return
			}

			if len(ts) != datum.expectedResultCountBeforeUpdate {
				t.Errorf("[BEFORE UPDATE] expected %d but got: %d. expression used: %v", datum.expectedResultCountBeforeUpdate, len(ts), f1)
			}

			_, err = updateOne(datum.updateFilter, datum.update)
			if err != nil {
				t.Errorf("[UPDATE] error while updating %v", err)
			}

			ts, err = query[ListingAndReview](updateCollectionName, f1)
			if err != nil {
				t.Errorf("%v", err)
			}

			if len(ts) != datum.expectedResultCountAfterUpdate {
				t.Errorf("[AFTER UPDATE] expected %d but got: %d. expression used: %v", datum.expectedResultCountAfterUpdate, len(ts), f1)
			}

		})
	}

}

func updateOne(filter any, update any) (int64, error) {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionStringForTesting))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("airbnb").Collection(updateCollectionName)

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return -1, err
	}

	return result.MatchedCount, nil
}

func cloneCollection(name string) error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionStringForTesting))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	command := bson.A{}
	command = append(command, bson.D{{"$match", bson.D{}}})
	command = append(command, bson.D{{"$out", name}})

	_, err = client.Database("airbnb").Collection(sampleCollection).Aggregate(ctx, command)
	return err
}

func removeCollection(name string) error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionStringForTesting))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("airbnb").Collection(name)
	return collection.Drop(ctx)

}
