package bo

type DailyTrends struct {
	Default *Default `json:"default"`
}

type Default struct {
	TrendingSearchesDays  []*TrendingSearchesDay `json:"trendingSearchesDays"`
	EndDateForNextRequest string                 `json:"endDateForNextRequest"`
	RssFeedPageUrl        string                 `json:"rssFeedPageUrl"`
}

type TrendingSearchesDay struct {
	Date             string   `json:"date"`
	FormattedDate    string   `json:"formattedDate"`
	TrendingSearches []*Trend `json:"trendingSearches"`
}

type Trend struct {
	Title            *Title     `json:"title"`
	FormattedTraffic string     `json:"formattedTraffic"`
	RelatedQueries   []*Title   `json:"relatedQueries"`
	Image            *Image     `json:"image"`
	Articles         []*Article `json:"articles"`
	ShareUrl         string     `json:"shareUrl"`
}

type Title struct {
	Query       string `json:"query"`
	ExploreLink string `json:"exploreLink"`
}

type Image struct {
	NewsUrl  string `json:"newsUrl"`
	Source   string `json:"internal"`
	ImageUrl string `json:"imageUrl"`
}

type Article struct {
	Title   string `json:"title"`
	TimeAgo string `json:"timeAgo"`
	Source  string `json:"internal"`
	Image   *Image `json:"image"`
	Url     string `json:"url"`
	Snippet string `json:"snippet"`
}
