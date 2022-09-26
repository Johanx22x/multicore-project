package token;

import "strings"

func tokenize(text string) []string {
    // FIX: Find a better way to tokenize and delete unnecesary words
    return strings.Fields(text);
}
