package feeds

import(
	"io"
	"log"
	"context"
	"strings"
	"net/http"
	"net/url"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"

	"onlygood/lib/database"

	readability "github.com/go-shiori/go-readability"
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

type Article struct {
	Title string `json:"title"`
	Content string `json:"content"`
	SentimentGroup string `json:"sentimentGroup"`
	SentimentScore float64 `json:"sentimentScore"`
	URL string `json:"url"`
	AlreadyRead bool `json:"alreadyRead"`
	Hash string `json:"hash"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

func NewFeedsInterface() *Feeds {
	return &Feeds{}
}

func hashValue(valueToHash string) (string){
	
	hasher := sha1.New()
	hasher.Write([]byte(valueToHash))
	result := hex.EncodeToString(hasher.Sum(nil))

	return result

}

func ProcessArticle(articleURL string) {

	resp, err := http.Get(articleURL)
	if err != nil {
		log.Printf("failed to download %s: %v\n", articleURL, err)
		resp.Body.Close()
		return
	}
	defer resp.Body.Close()

	parsedURL, err := url.Parse(articleURL)
	if err != nil {
		log.Fatalf("error parsing url")
	}

	article, err := readability.FromReader(resp.Body, parsedURL)
	if err != nil {
		log.Fatalf("failed to parse %s: %v\n", articleURL, err)
	}

	log.Printf("Content for %s:\n%s\n", articleURL, article.TextContent)

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

	feed.Hash = hashValue(feed.URL)
	
	query := `INSERT INTO feeds (name, url, icon, hash) VALUES (?, ?, ?, ?)`
	
	_, err := db.Exec(query, feed.Name, feed.URL, feed.Icon, feed.Hash)
	if err != nil {
		log.Printf("Failed to add feed: %v", err)
		return err
	}
	
	return nil

}

func (f *Feeds) GetArticlesForFeed(feedHash string) []Article {

	var articles = []Article{}

	db := database.Get()

	query := `SELECT url FROM feeds WHERE hash = ?`

	var feedURL string
	err := db.QueryRow(query, feedHash).Scan(&feedURL)
	if err != nil {
		log.Printf("Failed to get feed for hash: %v", err)
		return articles
	}

	// Fetch RSS feed
	resp, err := http.Get(feedURL)
	if err != nil {
		log.Printf("Failed to fetch RSS feed: %v", err)
		return articles
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read RSS feed body: %v", err)
		return articles
	}

	// Parse RSS XML
	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		log.Printf("Failed to parse RSS XML: %v", err)
		return articles
	}

	// Convert RSS items to Article structs
	for _, item := range rss.Channel.Items {

		cleanURL := item.Link
		if idx := strings.Index(item.Link, "?"); idx != -1 {
			cleanURL = item.Link[:idx]
		}

		go ProcessArticle(cleanURL)

		article := Article{
			Title: item.Title,
			Content: item.Description,
			URL: cleanURL,
			SentimentGroup: "unknown",
			SentimentScore: -1.0,
			AlreadyRead: false,
			Hash: hashValue(cleanURL),
		}

		articles = append(articles, article)

	}

	log.Printf("Fetched %d articles from %s\n", len(articles), feedURL)

	return articles

}