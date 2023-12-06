package generator

import "go.mongodb.org/mongo-driver/bson/primitive"

//go:generate mongo-query-gen -in Types.go -outDir .

type ListingAndReview struct {
	Id                   string               `bson:"_id"`
	ListingUrl           string               `bson:"listing_url"`
	Name                 string               `bson:"name"`
	Summary              string               `bson:"summary"`
	Space                string               `bson:"space"`
	Description          string               `bson:"description"`
	NeighborhoodOverview string               `bson:"neighborhood_overview"`
	Notes                string               `bson:"notes"`
	Transit              string               `bson:"transit"`
	Access               string               `bson:"access"`
	Interaction          string               `bson:"interaction"`
	HouseRules           string               `bson:"house_rules"`
	PropertyType         string               `bson:"property_type"`
	RoomType             string               `bson:"room_type"`
	BedType              string               `bson:"bed_type"`
	MinimumNights        string               `bson:"minimum_nights"`
	MaximumNights        string               `bson:"maximum_nights"`
	CancellationPolicy   string               `bson:"cancellation_policy"`
	LastScraped          primitive.DateTime   `bson:"last_scraped"`
	CalendarLastScraped  primitive.DateTime   `bson:"calendar_last_scraped"`
	Accommodates         int                  `bson:"accommodates"`
	Bedrooms             int                  `bson:"bedrooms"`
	Beds                 int                  `bson:"beds"`
	NumberOfReviews      int                  `bson:"number_of_reviews"`
	Bathrooms            primitive.Decimal128 `bson:"bathrooms"`
	Amenities            []string             `bson:"amenities"`
	Price                primitive.Decimal128 `bson:"price"`
	WeeklyPrice          primitive.Decimal128 `bson:"weekly_price"`
	MonthlyPrice         primitive.Decimal128 `bson:"monthly_price"`
	CleaningFee          primitive.Decimal128 `bson:"cleaning_fee"`
	ExtraPeople          primitive.Decimal128 `bson:"extra_people"`
	GuestsIncluded       primitive.Decimal128 `bson:"guests_included"`
	Images               struct {
		ThumbnailUrl string `bson:"thumbnail_url"`
		MediumUrl    string `bson:"medium_url"`
		PictureUrl   string `bson:"picture_url"`
		XlPictureUrl string `bson:"xl_picture_url"`
	} `bson:"images"`
	Host struct {
		HostId                 string   `bson:"host_id"`
		HostUrl                string   `bson:"host_url"`
		HostName               string   `bson:"host_name"`
		HostLocation           string   `bson:"host_location"`
		HostAbout              string   `bson:"host_about"`
		HostThumbnailUrl       string   `bson:"host_thumbnail_url"`
		HostPictureUrl         string   `bson:"host_picture_url"`
		HostNeighbourhood      string   `bson:"host_neighbourhood"`
		HostIsSuperhost        bool     `bson:"host_is_superhost"`
		HostHasProfilePic      bool     `bson:"host_has_profile_pic"`
		HostIdentityVerified   bool     `bson:"host_identity_verified"`
		HostListingsCount      int      `bson:"host_listings_count"`
		HostTotalListingsCount int      `bson:"host_total_listings_count"`
		HostVerifications      []string `bson:"host_verifications"`
	} `bson:"host"`
	Address struct {
		Street         string `bson:"street"`
		Suburb         string `bson:"suburb"`
		GovernmentArea string `bson:"government_area"`
		Market         string `bson:"market"`
		Country        string `bson:"country"`
		CountryCode    string `bson:"country_code"`
		Location       struct {
			Type            string    `bson:"type"`
			Coordinates     []float64 `bson:"coordinates"`
			IsLocationExact bool      `bson:"is_location_exact"`
		} `bson:"location"`
	} `bson:"address"`
	Availability struct {
		Availability30  int `bson:"availability_30"`
		Availability60  int `bson:"availability_60"`
		Availability90  int `bson:"availability_90"`
		Availability365 int `bson:"availability_365"`
	} `bson:"availability"`
	ReviewScores struct {
		ReviewScoresAccuracy      int `bson:"review_scores_accuracy"`
		ReviewScoresCleanliness   int `bson:"review_scores_cleanliness"`
		ReviewScoresCheckin       int `bson:"review_scores_checkin"`
		ReviewScoresCommunication int `bson:"review_scores_communication"`
		ReviewScoresLocation      int `bson:"review_scores_location"`
		ReviewScoresValue         int `bson:"review_scores_value"`
		ReviewScoresRating        int `bson:"review_scores_rating"`
	} `bson:"review_scores"`
	Reviews []struct {
		Id           string             `bson:"_id"`
		Date         primitive.DateTime `bson:"date"`
		ListingId    string             `bson:"listing_id"`
		ReviewerId   string             `bson:"reviewer_id"`
		ReviewerName string             `bson:"reviewer_name"`
		Comments     string             `bson:"comments"`
	} `bson:"reviews"`
}
