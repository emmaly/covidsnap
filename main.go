package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	selector := `div[data-hveid="CBQQPQ"]`
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(`https://www.google.com/search?q=covid-19+stats+united+states`),
		chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.Screenshot(selector, &buf, chromedp.NodeVisible, chromedp.ByQuery),
	}); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("COVID-19 US Snapshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}
