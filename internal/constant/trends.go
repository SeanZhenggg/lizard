package constant

const (
	GoogleDailyTrendApiDomain = "https://trends.google.com/trends/api/dailytrends"
)

const (
	MIN_ID_VALUE = 1111111111
)

const (
	DefaultGoogleDailyTrendUrl = "https://trends.google.com/trends/trendingsearches/daily?geo=TW&hl=zh-TW"
)

var (
	DailyTrendApiReqParams = map[string]string{
		"hl":  "zh-TW",
		"tz":  "-480",
		"geo": "TW",
		"ns":  "15",
	}
)
