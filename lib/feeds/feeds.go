package feeds

import(
	"context"
	"log"

	"onlygood/lib/database"
)

type Feeds struct {
	ctx context.Context
}

type Feed struct {
	Name string `json:"name"`
	URL string `json:"url"`
	Icon string `json:"icon"`
	Hash string `json:"hash"`
}

var feeds = []Feed{
	{
		Name : "BBC News",
		URL : "https://feeds.bbci.co.uk/news/rss.xml",
		Icon : "https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif",
		Hash : "140dffc02f670a97d80a6f99eec3a6db37763c20",
	},
	{
		Name : "The Guardian",
		URL : "https://www.theguardian.com/uk/rss",
		Icon : "https://assets.guim.co.uk/images/guardian-logo-rss.c45beb1bafa34b347ac333af2e6fe23f.png",
		Hash : "657a9846e4d0decff9cccdb29ec3483846906757",
	},
}

func NewFeedsInterface() *Feeds {
	return &Feeds{}
}

func (f *Feeds) ListFeeds(hash string) ([]Feed){

	db := database.Get()
	var feeds []Feed

	rows, err := db.Query("SELECT * FROM feeds")
	if err != nil {
		log.Printf("Failed to list routes: %v", err)
		return feeds
	}
	defer rows.Close()

	for rows.Next() {
		var feed Feed
		err := rows.Scan(&feed)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		feeds = append(feeds, feed)
	}

	return feeds
}