package feedreader

import (
	"fmt"
	aio "github.com/tjololo/io-client-go"
	"net/url"
	"os"
)

func ReadLatestValue(feedName string) string {
	client:=aio.NewClient(os.Getenv("AIO_KEY"))
	client.BaseURL, _ = url.Parse("https://io.adafruit.com")
	feed, _, ferr := client.Feed.Get(feedName)
	if ferr != nil {
		fmt.Printf("unable to load Feed with key %s. %s", feedName, ferr)
		return "N/A"
	}
	client.SetFeed(feed)

	ndp, _, err := client.Data.Last()
	if err != nil {
		fmt.Println("unable to retrieve data")
		return "N/A"
	}
	return ndp.Value
}
