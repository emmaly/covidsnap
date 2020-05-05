package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

func main() {
	allocOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), allocOpts...)
	defer allocCancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var buf []byte
	var x string
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Emulate(device.Pixel2),
		chromedp.Navigate(`https://www.google.com/search?q=covid-19+stats+united+states`),
		chromedp.EvaluateAsDevTools("window.scroll(0, 155); '1';", &x),
		chromedp.CaptureScreenshot(&buf),
	}); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("COVID-19 US Snapshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}
