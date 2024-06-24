package res

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"time"
	"wowpow/internal/proto"
)

//go:embed quotes.json
var quotesSource []byte

var quotes []*proto.Quote

func LoadQuotes() error {
	rand.Seed(time.Now().UnixNano())
	return json.Unmarshal(quotesSource, &quotes)
}

func GetRandomQuote() *proto.Quote {
	return quotes[rand.Intn(len(quotes))]
}

// func ConvertQuotes() error {
// 	lines := strings.Split(quotesSource, "\r\n\r\n")
// 	quotes := []*proto.Quote{}
// 	for _, line := range lines {
// 		parts := strings.Split(line, "\r\n--")
// 		if len(parts) != 2 {
// 			panic("SSSS")
// 		}
// 		quotes = append(quotes, &proto.Quote{
// 			Text:   strings.TrimSpace(parts[0]),
// 			Author: strings.TrimSpace(parts[1]),
// 		})
// 	}

// 	b, err := json.MarshalIndent(quotes, "", "  ")
// 	if err != nil {
// 		panic(err)
// 	}

// 	os.WriteFile("quotes.json", b, 0644)

// 	return nil
// }
