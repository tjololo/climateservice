package temperature

import (
	"fmt"
	"github.com/tjololo/climateservice/internal/pkg/feedreader"
	"net/http"
	"os"
	"strconv"
)

func ReturnTemperature(w http.ResponseWriter, r *http.Request) {
	feedName := os.Getenv("TEMPERATURE_FEED_NAME")
	fmt.Printf("Reading feed: %s\n", feedName)
	temp := feedreader.ReadLatestValue(feedName)
	f, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		fmt.Printf("unable to parse value to float. %s", err)
	}
	fmt.Fprintf(w, "Current temperature is: %.1f", f)
}
