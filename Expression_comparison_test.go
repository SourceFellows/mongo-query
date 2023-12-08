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
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func Test_Compare_EqualsAndEquals(t *testing.T) {

	//given
	f1 := Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999").
		And(Listing.Name.Equals("Horto flat with small garden"))
	mongoFilter := f1.bsonD()
	apiFilter := bson.D{{"$and", []bson.D{
		{
			{"listing_url", "https://www.airbnb.com/rooms/10009999"},
		},
		{
			{"name", "Horto flat with small garden"},
		},
	}}}

	//when

	//then
	if !reflect.DeepEqual(mongoFilter, apiFilter) {
		t.Error("api and generated value differs")
	}

}

func Test_Compare_In(t *testing.T) {

	//given
	f1 := Listing.ListingUrl.In("https://www.airbnb.com/rooms/10009999")
	mongoFilter := f1.bsonD()
	apiFilter := bson.D{
		{"listing_url", bson.D{{"$in", []string{"https://www.airbnb.com/rooms/10009999"}}}},
	}

	//when
	//check if api works
	apiResult, err := query[ListingAndReview](apiFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}
	mongoResult, err := query[ListingAndReview](mongoFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}

	//then
	if !reflect.DeepEqual(mongoResult, apiResult) {
		t.Errorf("api and generated results differs lib: %v, api: %v", len(mongoResult), len(apiResult))
	}

}

func Test_Compare_Gt(t *testing.T) {

	//given
	f1 := Listing.Bedrooms.Gt(6)
	mongoFilter := f1.bsonD()
	apiFilter := bson.D{
		{"bedrooms", bson.D{{"$gt", 6}}},
	}

	//when
	//check if api works
	apiResult, err := query[ListingAndReview](apiFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}
	mongoResult, err := query[ListingAndReview](mongoFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}

	//then

	if !reflect.DeepEqual(mongoFilter, apiFilter) {
		t.Errorf("api and generated value differs: %v %v", mongoFilter, apiFilter)
	}
	if !reflect.DeepEqual(mongoResult, apiResult) {
		t.Errorf("api and generated results differs lib: %v, api: %v", len(mongoResult), len(apiResult))
	}

}

func Test_Compare_EqualsAndSize(t *testing.T) {

	//given
	f1 := Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999").And(Listing.Amenities.ArraySize(15))
	mongoFilter := f1.bsonD()
	apiFilter := bson.D{{"$and", []bson.D{
		{
			{"listing_url", "https://www.airbnb.com/rooms/10009999"},
		},
		{
			{"amenities", bson.D{{"$size", 15}}},
		},
	}}}

	//when
	apiResult, err := query[ListingAndReview](apiFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}
	mongoResult, err := query[ListingAndReview](mongoFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}

	//then
	if !reflect.DeepEqual(mongoResult, apiResult) {
		t.Errorf("api and generated results differs lib: %v, api: %v", len(mongoResult), len(apiResult))
	}
	if !reflect.DeepEqual(mongoFilter, apiFilter) {
		t.Errorf("mongoquery and api value differs lib: %v, api: %v", mongoFilter, apiFilter)
	}

}

func Test_Compare_ArrayContainsQueryOperator(t *testing.T) {

	//given
	f1 := Listing.Amenities.ArrayContainsElement(Equals("Wifi"))
	mongoFilter := f1.bsonD()
	apiFilter := bson.D{{"amenities", bson.D{
		{"$eq", "Wifi"},
	}}}

	//when
	apiResult, err := query[ListingAndReview](apiFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}
	mongoResult, err := query[ListingAndReview](mongoFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}

	//then
	if !reflect.DeepEqual(mongoResult, apiResult) {
		t.Errorf("api and generated results differs lib: %v, api: %v", len(mongoResult), len(apiResult))
	}
	if !reflect.DeepEqual(mongoFilter, apiFilter) {
		t.Errorf("mongoquery and api value differs lib: %v, api: %v", mongoFilter, apiFilter)
	}

}

func Test_Compare_ArrayContainsQueryOperatorWithAndCondition(t *testing.T) {

	//given
	f1 := Listing.Bedrooms.Gt(8).And(Listing.Amenities.ArrayContainsElement(Equals("Wifi")))
	mongoFilter := f1.bsonD()
	apiFilter := bson.D{{Key: "$and", Value: []bson.D{
		{{Key: "bedrooms", Value: bson.D{{"$gt", 8}}}},
		{{Key: "amenities", Value: bson.D{{"$eq", "Wifi"}}}},
	}}}

	//when
	apiResult, err := query[ListingAndReview](apiFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}
	mongoResult, err := query[ListingAndReview](mongoFilter)
	if err != nil {
		t.Errorf("could not execute query %v", err)
	}

	//then
	if !reflect.DeepEqual(mongoResult, apiResult) {
		t.Errorf("api and generated results differs lib: %v, api: %v", len(mongoResult), len(apiResult))
	}
	if !reflect.DeepEqual(mongoFilter, apiFilter) {
		t.Errorf("mongoquery and api value differs lib: %v, api: %v", mongoFilter, apiFilter)
	}

}
