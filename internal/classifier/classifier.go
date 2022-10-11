package classifier

import (
	"strings"

	"github.com/Johanx22x/multicore-project/internal/chartm"
	"github.com/Johanx22x/multicore-project/internal/token"
	"github.com/crawlerclub/ce"
)

// A struct to store the websites topics information in a better way
type TopicMap struct {
    topics map[string]float64;
    words []string;
    WebsiteType string;
}

// Naive Bayes classifier
func NaiveBayes(payload map[string]ce.Doc, keywords map[string][]string) {
    // Create a map to store the total of keywords in a topic
    topicsKeywordLength := make(map[string]int)
    totalWords := 0
    for key, value := range keywords {
        // Add the key with the topic name and set the value to the total amount of keywords
        topicsKeywordLength[key] = len(value)
        totalWords += len(value)
    }

    // Create a new map to store the keywords weight
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

    // Iterate over the websites
    for idx, item := range payload {
        if item.Text != "" {
            cont := 0.0
            for idxKey, val := range keywords {
                for _, word := range val {
                    for _, itemWord := range token.Tokenize(item.Text) {
                        str := strings.ToLower(itemWord)
                        // Check if the keywords is a substring
                        if strings.Contains(str, word) {
                            topic := topicsWeight[idx]
                            topic.topics[idxKey]++
                            topic.words = append(topic.words, str)
                        } else { // Increment the OTHER cont
                            cont++
                        } 
                    }
                }
            }
            topic := topicsWeight[idx]
            topic.topics["other"] = cont
        }
    }

    // Create the total topics map to create a new chart
    totalTopics := make(map[string]float64)

    // iterate over the topicsWeight map to get the topic of the webpages
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

            // Naive Bayes classification accordind to some variables
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

        // Create a chart and save the words for the current website
        chartm.CreateChart(chartMap, name, name)
        chartm.SaveWords(val.words, name)
    }

    // Create a chart with all the ocurrences
    chartm.CreateChart(totalTopics, "Keywords ocurrences inside websites", "total-keywords")
}
