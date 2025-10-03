package main

import (
	"embed"
	"io"
	"net/http"
	"net/url"
	"strings"
	"context"

	"onlygood/lib/app"
	"onlygood/lib/feeds"
	"onlygood/lib/database"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS
func proxyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/api/proxy" {
			handleProxy(w, r)
			return
		}


		next.ServeHTTP(w, r)
	})
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		return
	}

	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	parsedURL, err := url.Parse(targetURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	proxyReq, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Failed to create proxy request", http.StatusInternalServerError)
		return
	}

	for key, values := range r.Header {
		if key != "Host" && key != "Connection" {
			for _, value := range values {
				proxyReq.Header.Add(key, value)
			}
		}
	}

	if proxyReq.Header.Get("User-Agent") == "" {
		proxyReq.Header.Set("User-Agent", "Mozilla/5.0 (compatible; OnlyGood/1.0)")
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to fetch URL: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	excludeHeaders := map[string]bool{
		"X-Frame-Options":           true,
		"Content-Security-Policy":   true,
		"X-Content-Security-Policy": true,
	}

	for key, values := range resp.Header {
		keyLower := strings.ToLower(key)
		if !excludeHeaders[key] && !excludeHeaders[keyLower] {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {

	database.InitDB()
	app := app.NewApp()
	feedsInterface := feeds.NewFeedsInterface()

	err := wails.Run(&options.App{
		Title:  "OnlyGood",
		Width:  1280,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: assetserver.ChainMiddleware(
				proxyMiddleware,
			),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.Startup(ctx)
			feedsInterface.Startup(ctx) // Initialize the context
		},
		Bind: []interface{}{
			app,
			feedsInterface,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 true,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.DefaultAppearance,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "OnlyGood",
				Message: "© Mitchell Technologies Limited",
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}