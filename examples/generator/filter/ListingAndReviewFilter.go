// Code generated by monGo-Query. DO NOT EDIT.
// see https://github.com/SourceFellows/mongo-query or
// https://www.source-fellows.com
package filter

import (
	mq "github.com/sourcefellows/mongo-query"
)

type ListingAndReviewFilter struct {
	Id                   mq.Field
	ListingUrl           mq.Field
	Name                 mq.Field
	Summary              mq.Field
	Space                mq.Field
	Description          mq.Field
	NeighborhoodOverview mq.Field
	Notes                mq.Field
	Transit              mq.Field
	Access               mq.Field
	Interaction          mq.Field
	HouseRules           mq.Field
	PropertyType         mq.Field
	RoomType             mq.Field
	BedType              mq.Field
	MinimumNights        mq.Field
	MaximumNights        mq.Field
	CancellationPolicy   mq.Field
	LastScraped          mq.Field
	CalendarLastScraped  mq.Field
	Accommodates         mq.Field
	Bedrooms             mq.Field
	Beds                 mq.Field
	NumberOfReviews      mq.Field
	Bathrooms            mq.Field
	Amenities            mq.ArrayField
	Price                mq.Field
	WeeklyPrice          mq.Field
	MonthlyPrice         mq.Field
	CleaningFee          mq.Field
	ExtraPeople          mq.Field
	GuestsIncluded       mq.Field
	Reviews              mq.ArrayField
	//nested struct Images
	Images struct {
		ThumbnailUrl mq.Field
		MediumUrl    mq.Field
		PictureUrl   mq.Field
		XlPictureUrl mq.Field
	}
	//nested struct Host
	Host struct {
		HostId                 mq.Field
		HostUrl                mq.Field
		HostName               mq.Field
		HostLocation           mq.Field
		HostAbout              mq.Field
		HostThumbnailUrl       mq.Field
		HostPictureUrl         mq.Field
		HostNeighbourhood      mq.Field
		HostIsSuperhost        mq.Field
		HostHasProfilePic      mq.Field
		HostIdentityVerified   mq.Field
		HostListingsCount      mq.Field
		HostTotalListingsCount mq.Field
		HostVerifications      mq.ArrayField
	}
	//nested struct Address
	Address struct {
		Street         mq.Field
		Suburb         mq.Field
		GovernmentArea mq.Field
		Market         mq.Field
		Country        mq.Field
		CountryCode    mq.Field
		//nested struct Location
		Location struct {
			Type            mq.Field
			Coordinates     mq.ArrayField
			IsLocationExact mq.Field
		}
	}
	//nested struct Availability
	Availability struct {
		Availability30  mq.Field
		Availability60  mq.Field
		Availability90  mq.Field
		Availability365 mq.Field
	}
	//nested struct ReviewScores
	ReviewScores struct {
		ReviewScoresAccuracy      mq.Field
		ReviewScoresCleanliness   mq.Field
		ReviewScoresCheckin       mq.Field
		ReviewScoresCommunication mq.Field
		ReviewScoresLocation      mq.Field
		ReviewScoresValue         mq.Field
		ReviewScoresRating        mq.Field
	}
}

var ListingAndReview = ListingAndReviewFilter{

	Id:                   mq.Field("_id"),
	ListingUrl:           mq.Field("listing_url"),
	Name:                 mq.Field("name"),
	Summary:              mq.Field("summary"),
	Space:                mq.Field("space"),
	Description:          mq.Field("description"),
	NeighborhoodOverview: mq.Field("neighborhood_overview"),
	Notes:                mq.Field("notes"),
	Transit:              mq.Field("transit"),
	Access:               mq.Field("access"),
	Interaction:          mq.Field("interaction"),
	HouseRules:           mq.Field("house_rules"),
	PropertyType:         mq.Field("property_type"),
	RoomType:             mq.Field("room_type"),
	BedType:              mq.Field("bed_type"),
	MinimumNights:        mq.Field("minimum_nights"),
	MaximumNights:        mq.Field("maximum_nights"),
	CancellationPolicy:   mq.Field("cancellation_policy"),
	LastScraped:          mq.Field("last_scraped"),
	CalendarLastScraped:  mq.Field("calendar_last_scraped"),
	Accommodates:         mq.Field("accommodates"),
	Bedrooms:             mq.Field("bedrooms"),
	Beds:                 mq.Field("beds"),
	NumberOfReviews:      mq.Field("number_of_reviews"),
	Bathrooms:            mq.Field("bathrooms"),
	Amenities:            mq.ArrayField("amenities"),
	Price:                mq.Field("price"),
	WeeklyPrice:          mq.Field("weekly_price"),
	MonthlyPrice:         mq.Field("monthly_price"),
	CleaningFee:          mq.Field("cleaning_fee"),
	ExtraPeople:          mq.Field("extra_people"),
	GuestsIncluded:       mq.Field("guests_included"),
	Reviews:              mq.ArrayField("reviews"),
	/* Images */
	Images: struct {
		ThumbnailUrl mq.Field
		MediumUrl    mq.Field
		PictureUrl   mq.Field
		XlPictureUrl mq.Field
	}{
		ThumbnailUrl: mq.Field("images.thumbnail_url"),
		MediumUrl:    mq.Field("images.medium_url"),
		PictureUrl:   mq.Field("images.picture_url"),
		XlPictureUrl: mq.Field("images.xl_picture_url"),
	},
	/* Host */
	Host: struct {
		HostId                 mq.Field
		HostUrl                mq.Field
		HostName               mq.Field
		HostLocation           mq.Field
		HostAbout              mq.Field
		HostThumbnailUrl       mq.Field
		HostPictureUrl         mq.Field
		HostNeighbourhood      mq.Field
		HostIsSuperhost        mq.Field
		HostHasProfilePic      mq.Field
		HostIdentityVerified   mq.Field
		HostListingsCount      mq.Field
		HostTotalListingsCount mq.Field
		HostVerifications      mq.ArrayField
	}{
		HostId:                 mq.Field("host.host_id"),
		HostUrl:                mq.Field("host.host_url"),
		HostName:               mq.Field("host.host_name"),
		HostLocation:           mq.Field("host.host_location"),
		HostAbout:              mq.Field("host.host_about"),
		HostThumbnailUrl:       mq.Field("host.host_thumbnail_url"),
		HostPictureUrl:         mq.Field("host.host_picture_url"),
		HostNeighbourhood:      mq.Field("host.host_neighbourhood"),
		HostIsSuperhost:        mq.Field("host.host_is_superhost"),
		HostHasProfilePic:      mq.Field("host.host_has_profile_pic"),
		HostIdentityVerified:   mq.Field("host.host_identity_verified"),
		HostListingsCount:      mq.Field("host.host_listings_count"),
		HostTotalListingsCount: mq.Field("host.host_total_listings_count"),
		HostVerifications:      mq.ArrayField("host.host_verifications"),
	},
	/* Address */
	Address: struct {
		Street         mq.Field
		Suburb         mq.Field
		GovernmentArea mq.Field
		Market         mq.Field
		Country        mq.Field
		CountryCode    mq.Field
		//nested struct Location
		Location struct {
			Type            mq.Field
			Coordinates     mq.ArrayField
			IsLocationExact mq.Field
		}
	}{
		Street:         mq.Field("address.street"),
		Suburb:         mq.Field("address.suburb"),
		GovernmentArea: mq.Field("address.government_area"),
		Market:         mq.Field("address.market"),
		Country:        mq.Field("address.country"),
		CountryCode:    mq.Field("address.country_code"),
		/* Location */
		Location: struct {
			Type            mq.Field
			Coordinates     mq.ArrayField
			IsLocationExact mq.Field
		}{
			Type:            mq.Field("address.location.type"),
			Coordinates:     mq.ArrayField("address.location.coordinates"),
			IsLocationExact: mq.Field("address.location.is_location_exact"),
		},
	},
	/* Availability */
	Availability: struct {
		Availability30  mq.Field
		Availability60  mq.Field
		Availability90  mq.Field
		Availability365 mq.Field
	}{
		Availability30:  mq.Field("availability.availability_30"),
		Availability60:  mq.Field("availability.availability_60"),
		Availability90:  mq.Field("availability.availability_90"),
		Availability365: mq.Field("availability.availability_365"),
	},
	/* ReviewScores */
	ReviewScores: struct {
		ReviewScoresAccuracy      mq.Field
		ReviewScoresCleanliness   mq.Field
		ReviewScoresCheckin       mq.Field
		ReviewScoresCommunication mq.Field
		ReviewScoresLocation      mq.Field
		ReviewScoresValue         mq.Field
		ReviewScoresRating        mq.Field
	}{
		ReviewScoresAccuracy:      mq.Field("review_scores.review_scores_accuracy"),
		ReviewScoresCleanliness:   mq.Field("review_scores.review_scores_cleanliness"),
		ReviewScoresCheckin:       mq.Field("review_scores.review_scores_checkin"),
		ReviewScoresCommunication: mq.Field("review_scores.review_scores_communication"),
		ReviewScoresLocation:      mq.Field("review_scores.review_scores_location"),
		ReviewScoresValue:         mq.Field("review_scores.review_scores_value"),
		ReviewScoresRating:        mq.Field("review_scores.review_scores_rating"),
	},
}
