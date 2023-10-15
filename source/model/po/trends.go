package po

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trend struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	Title            Title              `json:"title" bson:"title"`
	FormattedTraffic string             `json:"formattedTraffic" bson:"formattedTraffic"`
	RelatedQueries   []string           `json:"relatedQueries" bson:"relatedQueries"`
	Image            Image              `json:"image" bson:"image"`
	Articles         []Article          `json:"articles" bson:"articles"`
	ShareUrl         string             `json:"shareUrl" bson:"shareUrl"`
}

type Title struct {
	Query       string `json:"query" bson:"query"`
	ExploreLink string `json:"exploreLink" bson:"exploreLink"`
}

type Image struct {
	NewsUrl  string `json:"newsUrl" bson:"newsUrl"`
	Source   string `json:"source" bson:"source"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl"`
}

type Article struct {
	Title   string `json:"title" bson:"title"`
	TimeAgo string `json:"timeAgo" bson:"timeAgo"`
	Source  string `json:"source" bson:"source"`
	Image   Image  `json:"image" bson:"image"`
	Url     string `json:"url" bson:"url"`
	Snippet string `json:"snippet" bson:"snippet"`
}
