package feeds

import(
	"context"
)

type Feeds struct {
	ctx context.Context
}

type Feed struct {
	Name string `json:"name"`
	URL string `json:"url"`
	Icon string `json:"icon"`
}

func NewFeedsInterface() *Feeds {
	return &Feeds{}
}

func (f *Feeds) ListFeeds() []Feed{

	var feeds = []Feed{
		{
			Name : "BBC News",
			URL : "https://feeds.bbci.co.uk/news/rss.xml",
			Icon : "https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif",
		},
		{
			Name : "The Guardian",
			URL : "https://www.theguardian.com/uk/rss",
			Icon : "https://assets.guim.co.uk/images/guardian-logo-rss.c45beb1bafa34b347ac333af2e6fe23f.png",
		},
	}

	return feeds

}