// TODO: Add documentation to this module
package classifier

import (
	"fmt"
	"strings"

	"github.com/Johanx22x/multicore-project/internal/chartm"
	"github.com/Johanx22x/multicore-project/internal/token"
	"github.com/crawlerclub/ce"
)

type topicMap struct {
    topics map[string]float64;
    WebsiteType string;
}

func findMajor(topicMap map[string]*topicMap) string {
    for key, value := range topicMap {
        fmt.Println(key, value)
    }
    return "Know"
}

// TODO: Apply classification
func NaiveBayes(payload map[string]ce.Doc, keywords map[string][]string) {
    // Create a map to store the total of keywords in a topic
    topicsKeywordLength := make(map[string]int)
    totalWords := 0
    for key, value := range keywords {
        // Add the key with the topic name and set the value to the total amount of keywords
        topicsKeywordLength[key] = len(value)
        totalWords += len(value)
    }

    topicsWeight := make(map[string]*topicMap)
    for idx, val := range payload {
        if (val.Text == "") {
            continue
        }
        topic := topicsWeight[idx]
        topic = &topicMap{}
        topic.topics = make(map[string]float64)
        for idxKey := range keywords {
            topic.topics[idxKey] = 0
        }
        topic.WebsiteType = "Unknow"
        topicsWeight[idx] = topic
    }

    for idx, item := range payload {
        if item.Text != "" {
            cont :=  float64(len(token.Tokenize(item.Text)))
            for idxKey, val := range keywords {
                for _, word := range val {
                    for _, itemWord := range token.Tokenize(item.Text) {
                        if word == strings.ToLower(itemWord) {
                            topic := topicsWeight[idx]
                            topic.topics[idxKey]++
                            cont--
                        } 
                    }
                }
            }
            topic := topicsWeight[idx]
            topic.topics["other"] = cont
        }
    }

    majorValue := 0.0
    totalTopics := make(map[string]float64)
    for websiteKey, val := range topicsWeight {
        for key, valfloat := range val.topics {
            if strings.ToLower(key) == "other" {
                continue
            }
            wordPerTopic := topicsKeywordLength[key]

            // FORMULA = (wordPerTopic / totalWords) * (Incidences / wordPerTopic)
            result := ((float64(wordPerTopic) / float64(totalWords)) * (valfloat / float64(wordPerTopic)))
            // FIXME: Skip results when are equal
            if majorValue < result  {
                topicsWeight[websiteKey].WebsiteType = key
            } 
            totalTopics[key] += valfloat       
        }
        fmt.Println(websiteKey, val)
    }

    chartm.CreateChart(totalTopics, "Keywords ocurrences inside websites", "total-keywords")
}
