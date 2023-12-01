# monGO Query

![monGO-Query](https://github.com/sourcefellows/mongo-query/actions/workflows/go.yml/badge.svg)

**monGO-Query is a Library which makes building MongoDB Queries easy in Golang**

Formulating requests with the MongoDB API is sometimes very difficult and error-prone. You have to build and nest untyped objects. This will quickly become confusing and therefore difficult to read and maintain. monGO-Query solves this problem with an easier to understand (DSL like) API.

## The difference

The following example shows how the queries differ between the MongoDB API and monGO-Query. First of all, the struct with which the data is mapped in the MongoDB database:

```Golang
type ListingAndReview struct {
	ListingUrl           string               `bson:"listing_url"`
	Name                 string               `bson:"name"`
	Bathrooms            primitive.Decimal128 `bson:"bathrooms"`
	Amenities            []string             `bson:"amenities"`
	Images               struct {
		ThumbnailUrl string `bson:"thumbnail_url"`
		MediumUrl    string `bson:"medium_url"`
		PictureUrl   string `bson:"picture_url"`
		XlPictureUrl string `bson:"xl_picture_url"`
	} `bson:"images"`
}
```

> The data and data structures used in the examples come from a [freely available Airbnb example database](./examples/listingsAndReviews.json). You will find the [complete ListingAndReview struct](./Expression_types_test.go) in this repo (it's only a subset in the sample above). 

If you want to query the MongoDB collection and search for values with a specific `ListingUrl` and `Name`, you have to write the following:

```Golang
filter := bson.D{{"$and", []bson.D{
    {
        {"listing_url", "https://www.airbnb.com/rooms/10009999"},
    },
    {
        {"name", "Horto flat with small garden"},
    },
}}}
```

In this example you have to:

* build the `bson.D` structure right (see the typ, slice, ...)
* know about the MongoDB query operators (`"$and"`)
* know the BSON struct tags for all your fields (`"listing_url"`, `"name"`)
* cross your fingers that no one will change the BSON mapping

In contrast, the same query with monGO-query looks like this:

```Golang
filter := Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999").
	And(Listing.Name.Equals("Horto flat with small garden"))
```

It is:

* readable
* expressive
* easy to write
* independent of the `bson.D` structure


