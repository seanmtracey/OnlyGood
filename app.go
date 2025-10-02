package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

type Article struct {
	Title string
	Content string
	Sentiment float64
	URL string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Echo(value string) string {
	return fmt.Sprintf("%s", value)
}

func (a *App)GetArticles() []Article{

	var articles = []Article{
		{
			Title: "Eyewitnesses describe Manchester synagogue attack",
			Content: "Some content",
			Sentiment: 0.25,
			URL: "https://www.bbc.co.uk/news/articles/cn0rp05ykx7o",
		},
		{
			Title: "Unmasked: Secret BBC filming exposes hidden culture of misogyny and racism inside Met Police",
			Content: "More content",
			Sentiment: 0.35,
			URL: "https://www.bbc.co.uk/news/articles/cvgq06d44jyo",
		},
	}

	return articles
}