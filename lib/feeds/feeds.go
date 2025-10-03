package feeds

import(
	"log"
	"context"
	"crypto/sha1"
	"encoding/hex"

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

func NewFeedsInterface() *Feeds {
	return &Feeds{}
}

func (f *Feeds) Startup(ctx context.Context) {
	f.ctx = ctx
}

func (f *Feeds) ListFeeds() []Feed {
	
	db := database.Get()
	feeds := []Feed{}

	rows, err := db.Query("SELECT name, url, icon, hash FROM feeds")
	if err != nil {
		log.Printf("Failed to list feeds: %v", err)
		return feeds
	}
	defer rows.Close()

	for rows.Next() {
		var feed Feed
		err := rows.Scan(&feed.Name, &feed.URL, &feed.Icon, &feed.Hash)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		feeds = append(feeds, feed)
	}

	return feeds

}

func (f *Feeds) AddFeed(feed Feed) error {

	db := database.Get()
	
	hasher := sha1.New()
	hasher.Write([]byte(feed.URL))
	feed.Hash = hex.EncodeToString(hasher.Sum(nil))
	
	query := `INSERT INTO feeds (name, url, icon, hash) VALUES (?, ?, ?, ?)`
	
	_, err := db.Exec(query, feed.Name, feed.URL, feed.Icon, feed.Hash)
	if err != nil {
		log.Printf("Failed to add feed: %v", err)
		return err
	}
	
	return nil

}