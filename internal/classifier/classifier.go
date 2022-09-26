package classifier

import (
	"fmt"
	"strings"

	"github.com/Johanx22x/multicore-project/internal/token"
	"github.com/crawlerclub/ce"
)

type topicMap struct {
    topics map[string]float64
}

func NaiveBayes(payload map[string]ce.Doc, keywords map[string][]string) {
    topicsWeight := make(map[string]*topicMap)
    for idx, val := range payload {
        if val.Text != "" {
            for idxKey := range keywords {
                topic := topicsWeight[idx]
                if topic == nil {
                    topic = &topicMap{}
                    topic.topics = make(map[string]float64)
                    topic.topics[idxKey] = 0
                    topicsWeight[idx] = topic
                }
            }
        }
    }

    for idx, item := range payload {
        if item.Text != "" {
            for idxKey, val := range keywords {
                for _, word := range val {
                    for _, itemWord := range token.Tokenize(item.Text) {
                        if word == strings.ToLower(itemWord) {
                            topic := topicsWeight[idx]
                            topic.topics[idxKey]++
                        } 
                    }
                }
            }
        }
    }
    cont := 0
    for key, val := range topicsWeight {
        cont++
        fmt.Println(cont, key, val)
    }
}
