# mongo Query

![mongo Query](https://github.com/sourcefellows/mongo-query/actions/workflows/go.yml/badge.svg)

<center><img width="100%" alt="mongo Query" src="./assets/mongo-query-logo-black.png"></center>

----

**mongo Query is a library that makes it easy to build MongoDB queries in Golang**

Formulating requests with the MongoDB API is sometimes very difficult and error-prone. You have to build and nest untyped objects (`bson.M, bson.D, ...`). This will quickly become confusing and therefore difficult to read and maintain. mongo Query solves this problem with an easier to understand (DSL like) API.

For example, the following 'mongo Query' expression finds all documents in a MongoDB database where the array of 'amenities' has a size of 15 and the 'ListingUrl' equals '<value>' and the 'pictureUrl' of the 'image' equals '<value>':

```Golang
Listing.Amenities.ArraySize(15).
  And(Listing.ListingUrl.Equals("<value>"),
      Listing.Images.PictureUrl.Equals("<value>"))
```

> The (DSL-) queries are created with 'Filter structs' which can be generated using a commandline tool, which is also part of this project (see [Generating filter types](#generating-filter-types)).

The [API documentation for the filter structs can be found here](https://pkg.go.dev/github.com/sourcefellows/mongo-query#section-documentation). It is based on the 'native MongoDB query API'.

## MongoDB API vs mongo Query

The following example shows how the queries differ between the MongoDB API and mongo Query. First of all, the struct with which the data is mapped in the MongoDB database:

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

In contrast, the *same* query with **mongo Query** looks like this:

```Golang
filter := Listing.ListingUrl.Equals("https://www.airbnb.com/rooms/10009999").
	And(Listing.Name.Equals("Horto flat with small garden"))
```

It is:

* readable
* expressive
* easy to write
* independent of the `bson.D` structure
* resistant to renaming BSON field names

## How to use mongo Query

mongo Query uses its own simple API with `Expression`, `Field` and `Operator` types.

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

or with embedded documents...

```Golang
...
filter := Listing.Amenities.ArraySize(15).
    And(Listing.ListingUrl.Equals("<value>"),
        Listing.Images.PictureUrl.Equals("<value>"))

collection := client.Database("airbnb").Collection("listingsAndReviews")
cursor, err := collection.Find(ctx, filter)
...
```

> You will find a complete sample within the [unit tests](./Expression_query_test.go) of this project.

## Generating filter types

Defining filter types is easy. Just use the generator which is also included in the project. Install it via `go install` and use it (see an [example here](./examples/generator)):

```bash
go install github.com/sourcefellows/mongo-query/cmd/mongo-query-gen@latest
```

```bash
kkoehler@project$ mongo-query-gen 
no input file given
  -in string
        path to file with Golang structs
  -only string
        list of struct names - only given struct names will be used for code generation
  -outDir string
        path to output directory - a subdirectory "filter" will be generated automatically
```

For example:

```bash
mongo-query-gen -in Types.go -outDir .
```

## Samples from MongoDB manual

* Query embedded documents (Specify Equality Match on a Nested Field) ([see here](https://www.mongodb.com/docs/manual/tutorial/query-embedded-documents/) or [local impl](./examples/mongo-samples/manual-01))

```Golang
//MongoDB API
err = findWithFilter(ctx, collection, bson.D{{"size.uom", "in"}})
//mongo Query
err = findWithFilter(ctx, collection, InventoryFilter.Size.Uom.Equals("in"))
```

* Query embedded documents (Specify AND Condition) ([see here](https://www.mongodb.com/docs/manual/tutorial/query-embedded-documents/) or [local impl](./examples/mongo-samples/manual-01))

```Golang
//MongoDB API
err = findWithFilter(ctx, collection, bson.D{
    {"size.h", bson.D{
        {"$lt", 15},
    }},
    {"size.uom", "in"},
    {"status", "D"},
})
//mongo Query
err = findWithFilter(ctx, collection,
    InventoryFilter.Size.H.Lt(15).
        And(InventoryFilter.Size.Uom.Equals("in"),
            InventoryFilter.Status.Equals("D")))
```
* More samples can be found in the [unit tests of the project](./Expression_query_test.go).