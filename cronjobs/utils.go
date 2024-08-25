package cronjobs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Arinji2/news-backend/api"
	"github.com/Arinji2/news-backend/types"
)

func parseAndCleanupNewsItems(results map[string]interface{}) []types.NewsItem {
	resultsJSON, err := json.Marshal(results["results"])
	if err != nil {
		log.Fatalf("Error marshalling results to JSON: %v", err)
	}

	var newsItems []types.NewsItem
	err = json.Unmarshal(resultsJSON, &newsItems)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return newsItems
}

func verifyAllFields(newsItem types.NewsItem) bool {
	return newsItem.Title != "" && newsItem.Description != "" && newsItem.PublishedAt != "" && newsItem.Author != "" && newsItem.URL != "" && newsItem.URLToImage != ""
}

func checkIfNewsExists(newsItem types.NewsItem, client *api.ApiClient, table string) bool {
	results, err := client.SendRequestWithQuery("GET", fmt.Sprintf("/api/collections/%s/records", table), map[string]string{
		"page":    "1",
		"perPage": "1",
		"filter":  fmt.Sprintf(`"title=%s" || url=%s`, newsItem.Title, newsItem.URL)}, nil)

	if err != nil {
		return false
	}

	totalItems, ok := results["totalItems"].(float64)
	if !ok {
		log.Fatalf("Error converting totalItems to float64: %v", err)
	}
	return totalItems > 0
}
