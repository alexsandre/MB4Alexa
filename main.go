package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

const baseURL string = "https://www.mercadobitcoin.net/api/%v/%v/"

func main() {
	lambda.Start(Handler)
}

// Handler for lambda
func Handler() (alexa.Response, error) {
	btcTicker := getCoinTicker("BTC")
	message := fmt.Sprintf("O valor do bitcoin é %.2f", btcTicker.Ticker.Last)

	return alexa.NewSimpleResponse("Valor do bitcoin no mercado bitcoin", message), nil
}

func getCoinTicker(coin string) MBResponse {
	var url = fmt.Sprintf(baseURL, coin, "ticker")

	mbresponse := new(MBResponse)
	getJSON(url, mbresponse)

	return *mbresponse
}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// MBResponse is response of ticker endpoint
type MBResponse struct {
	Ticker struct {
		High float32 `json:"high,string"`
		Low  float32 `json:"low,string"`
		Vol  float32 `json:"vol,string"`
		Last float32 `json:"last,string"`
		Buy  float32 `json:"buy,string"`
		Sell float32 `json:"sell,string"`
		Date string  `json:"date"`
	} `json:"ticker"`
}
