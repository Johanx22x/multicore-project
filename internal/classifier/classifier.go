package classifier

import (
	"fmt"
	"strings"

	"github.com/Johanx22x/multicore-project/internal/chartm"
	"github.com/Johanx22x/multicore-project/internal/token"
	"github.com/crawlerclub/ce"
)

type TopicMap struct {
    topics map[string]float64;
    words []string;
    WebsiteType string;
}

func findMajor(topicMap map[string]*TopicMap) string {
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

    topicsWeight := make(map[string]*TopicMap)
    for idx, val := range payload {
        if (val.Text == "") {
            continue
        }
        topic := topicsWeight[idx]
        topic = &TopicMap{}
        topic.topics = make(map[string]float64)
        for idxKey := range keywords {
            topic.topics[idxKey] = 0
        }
        topic.WebsiteType = "Unknow"
        topicsWeight[idx] = topic
    }

    for idx, item := range payload {
        if item.Text != "" {
            cont := 0.0
            for idxKey, val := range keywords {
                for _, word := range val {
                    for _, itemWord := range token.Tokenize(item.Text) {
                        str := strings.ToLower(itemWord)
                        if strings.Contains(str, word) {
                            topic := topicsWeight[idx]
                            topic.topics[idxKey]++
                            topic.words = append(topic.words, str)
                        } else {
                            cont++
                        } 
                    }
                }
            }
            topic := topicsWeight[idx]
            topic.topics["other"] = cont
        }
    }

    totalTopics := make(map[string]float64)
    for websiteKey, val := range topicsWeight {
        chartMap := make(map[string]float64)
        majorName := "Unknow"
        majorValue := 0.0
        for key, valfloat := range val.topics {
            if strings.ToLower(key) == "other" {
                continue
            }
            chartMap[key] = valfloat
            wordPerTopic := topicsKeywordLength[key]

            result := ((float64(wordPerTopic) / float64(totalWords)) * (valfloat / float64(wordPerTopic)))

            if majorValue < result  {
                majorValue = result
                majorName = key
            }
            totalTopics[key] += valfloat       
        }
        topicsWeight[websiteKey].WebsiteType = majorName 
        name := strings.Trim(websiteKey, "https://")
        name += " [" + val.WebsiteType + "]"

        chartm.CreateChart(chartMap, name, name)
        chartm.SaveWords(val.words, name)
    }

    chartm.CreateChart(totalTopics, "Keywords ocurrences inside websites", "total-keywords")
}
