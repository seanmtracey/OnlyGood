package app

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

type Article struct {
	Title string `json:"title"`
	Content string `json:"content"`
	SentimentGroup string `json:"sentimentGroup"`
	SentimentScore float64 `json:"sentimentScore"`
	URL string `json:"url"`
	AlreadyRead bool `json:"alreadyRead"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Echo(value string) string {
	return fmt.Sprintf("%s", value)
}

func (a *App)GetArticles() []Article{

	/*var articles = []Article{
		{
			Title: "Eyewitnesses describe Manchester synagogue attack",
			Content: "Some content",
			SentimentGroup: "negative",
			SentimentScore: 0.25,
			URL: "https://www.bbc.co.uk/news/articles/cn0rp05ykx7o",
			AlreadyRead: false,
		},
		{
			Title: "Unmasked: Secret BBC filming exposes hidden culture of misogyny and racism inside Met Police",
			Content: "More content",
			SentimentGroup: "middling",
			SentimentScore: 0.55,
			URL: "https://www.bbc.co.uk/news/articles/cvgq06d44jyo",
			AlreadyRead: true,
		},
		{
			Title: "Unmasked: Secret BBC filming exposes hidden culture of misogyny and racism inside Met Police",
			Content: "More content",
			SentimentGroup: "positive",
			SentimentScore: 0.75,
			URL: "https://www.bbc.co.uk/news/articles/cvgq06d44jyo",
			AlreadyRead: true,
		},
	}*/

	return []Article{}
}