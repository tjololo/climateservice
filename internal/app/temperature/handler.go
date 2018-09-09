package temperature

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tjololo/climateservice/internal/pkg/feedreader"
	"net/http"
	"os"
	"strconv"
)

func ReturnTemperature(w http.ResponseWriter, r *http.Request) {
	feedName := os.Getenv("TEMPERATURE_FEED_NAME")
	fmt.Printf("Reading feed: %s\n", feedName)
	temp, err := feedreader.ReadLatestValue(feedName)
	if err != nil {
		log.Errorf("Failed to fetch temperature. %s", err)
		http.Error(w, "Could not fetch temperature", 500)
		return
	}
	f, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		log.Errorf("unable to parse value to float. %s", err)
		http.Error(w, "Could not fetch temperature", 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SensorData{
		Unit:  "°C",
		Value: fmt.Sprintf("%.1f", f),
	})
}

func ReturnHumidity(w http.ResponseWriter, r *http.Request) {
	feedName := os.Getenv("HUMIDITY_FEED_NAME")
	fmt.Printf("Reading feed: %s\n", feedName)
	temp, err := feedreader.ReadLatestValue(feedName)
	if err != nil {
		log.Errorf("unable to fetch value. %s", temp, err)
		http.Error(w, "Failed to parse temperature", 500)
		return
	}
	f, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		log.Errorf("unable to parse value: \"%s\" to float. %s", temp, err)
		http.Error(w, "Failed to parse temperature", 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SensorData{
		Unit:  "%",
		Value: fmt.Sprintf("%.1f", f),
	})
}
