package feedreader

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	aio "github.com/tjololo/io-client-go"
	"net/url"
	"os"
)

func ReadLatestValue(feedName string) (string, error) {
	client := aio.NewClient(os.Getenv("AIO_KEY"))
	client.BaseURL, _ = url.Parse("https://io.adafruit.com")
	feed, _, err := client.Feed.Get(feedName)
	if err != nil {
		fmt.Printf("unable to load Feed with key %s. %s", feedName, err)
		return "N/A", err
	}
	client.SetFeed(feed)

	ndp, _, err := client.Data.Last()
	if err != nil {
		log.Errorf("unable to retrieve data. %s", err)
		return "N/A", err
	}
	return ndp.Value, nil
}
