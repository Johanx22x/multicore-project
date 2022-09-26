package token;

import "strings"

func Tokenize(text string) []string {
    // FIX: Find a better way to tokenize and delete unnecesary words
    return strings.Fields(text);
}
