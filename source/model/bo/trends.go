package bo

import "go.mongodb.org/mongo-driver/bson/primitive"

type DailyTrends struct {
	Default Default `json:"default"`
}

type Default struct {
	TrendingSearchesDays  TrendingSearchesDays `json:"trendingSearchesDays"`
	EndDateForNextRequest string               `json:"endDateForNextRequest"`
	RssFeedPageUrl        string               `json:"rssFeedPageUrl"`
}

type TrendingSearchesDays struct {
	Date             string  `json:"date"`
	FormattedDate    string  `json:"formattedDate"`
	TrendingSearches []Trend `json:"trendingSearches"`
}

type Trend struct {
	ID               primitive.ObjectID `json:"id"`
	Title            Title              `json:"title"`
	FormattedTraffic string             `json:"formattedTraffic"`
	RelatedQueries   []string           `json:"relatedQueries"`
	Image            Image              `json:"image"`
	Articles         []Article          `json:"articles"`
	ShareUrl         string             `json:"shareUrl"`
}

type Title struct {
	Query       string `json:"query"`
	ExploreLink string `json:"exploreLink"`
}

type Image struct {
	NewsUrl  string `json:"newsUrl"`
	Source   string `json:"source"`
	ImageUrl string `json:"imageUrl"`
}

type Article struct {
	Title   string `json:"title"`
	TimeAgo string `json:"timeAgo"`
	Source  string `json:"source"`
	Image   Image  `json:"image"`
	Url     string `json:"url"`
	Snippet string `json:"snippet"`
}
