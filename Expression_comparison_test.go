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
		bson.D{
			{"listing_url", "https://www.airbnb.com/rooms/10009999"},
			{"name", "Horto flat with small garden"}},
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
		t.Errorf("api and generated results differs %v, %v", len(mongoResult), len(apiResult))
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
		t.Errorf("api and generated results differs %v, %v", len(mongoResult), len(apiResult))
	}

}
