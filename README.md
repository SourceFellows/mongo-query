# monGO Query

![monGO-Query](https://github.com/sourcefellows/mongo-query/actions/workflows/go.yml/badge.svg)

<center><img src="https://www.source-fellows.com/mongo-query.png"></center>

----

**monGO-Query is a Library which makes building MongoDB Queries easy in Golang**

Formulating requests with the MongoDB API is sometimes very difficult and error-prone. You have to build and nest untyped objects. This will quickly become confusing and therefore difficult to read and maintain. monGO-Query solves this problem with an easier to understand (DSL like) API.

## MongoDB API vs monGO-Query

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

If you want to query the MongoDB collection and search for values with a specific `ListingUrl` and `Name` with the **MongoDB-API**, you have to write the following:

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

In contrast, the *same* query with **monGO-query** looks like this:

```Golang
filter := Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999").
	And(Listing.Name.Equals("Horto flat with small garden"))
```

It is:

* readable
* expressive
* easy to write
* independent of the `bson.D` structure

## How to use monGO-Query

monGO-Query uses its own simple API with `Expression`s, `Field`s and `Operator`s.

You can simply define a filter type for each struct you want to query. Instances of this filter types can than be used as parameter to the MongoDB API. They will automatically be marshalled as MongoDB `bson.D` objects. 

> **You can generate this filter types with a generator which is also part of this project**! [See below](#generating-filter-types).

```Golang
type ListingFilter struct {
	ListingUrl Field
	Name       Field
	Bedrooms   Field
	Amenities  ArrayField
	Images     ImagesFilter
}

type ImagesFilter struct {
	ThumbnailUrl Field
	MediumUrl    Field
	PictureUrl   Field
	XlPictureUrl Field
}

var Listing = ListingFilter{
    ListingUrl: Field("listing_url"),
    Name:       Field("name"),
    Bedrooms:   Field("bedrooms"),
    Amenities:  ArrayField("amenities"),
    Images: ImagesFilter{
        ThumbnailUrl: Field("images.thumbnail_url"),
        MediumUrl:    Field("images.medium_url"),
        PictureUrl:   Field("images.picture_url"),
        XlPictureUrl: Field("images.xl_picture_url"),
    },
}
```

And then you can use the filter and query data via the MongoDB API:

```Golang
...
filter := Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999")
collection := client.Database("airbnb").Collection("listingsAndReviews")
cursor, err := collection.Find(ctx, filter)
...
```

> You will find a complete sample within the [unit tests](./Expression_query_test.go) of this project.

## Generating filter types

Defining filter types is easy. Just use the generator which is also included in the project. Install it via `go install` and use it (see an [example here](./examples/generator)):

```bash
go install github.com/sourcefellows/mongo-query/cmd/mongo-query-gen@latest

mongo-query-gen -in Types.go -outDir .
```

## Samples from MongoDB manual

* Query nested equals ([see here](https://www.mongodb.com/docs/manual/tutorial/query-embedded-documents/) or [local impl](./examples/mongo-samples/manual-01))

```Golang
//MongoDB API
err = findwithFilter(ctx, collection, bson.D{{"size.uom", "in"}})
//monGO-Query
err = findwithFilter(ctx, collection, InventoryFilter.Size.Uom.Equals("in"))
```
