package alpaca

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"polx/app/domain/bo"
	"polx/app/domain/definition"
	"polx/app/system/environment"
	"sync"
)

var (
	alpacaOnce sync.Once
	alpacaInst *alpacaRepo
)

const (
	baseUrl = "https://data.alpaca.markets/v2/stocks/%s/bars?start=%s&end=%s&timeframe=1Day"
)

type alpacaRepo struct {
	client *http.Client
}

func GetAlpacaRepo() definition.AlpacaRepo {
	alpacaOnce.Do(func() {
		alpacaInst = &alpacaRepo{
			client: &http.Client{},
		}
	})
	return alpacaInst
}

func (a *alpacaRepo) GetBars(ticker, startDate, endDate string) (*bo.AlpacaResponse, error) {
	url := fmt.Sprintf(baseUrl, ticker, startDate, endDate)
	fmt.Println(startDate)
	fmt.Println(endDate)
	fmt.Println(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("APCA-API-KEY-ID", environment.ALPACA_KEY)
	request.Header.Set("APCA-API-SECRET-KEY", environment.ALPACA_SECRET)

	resp, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	var trades bo.AlpacaResponse
	json.Unmarshal(body, &trades)
	fmt.Print(trades)
	fmt.Print("END")

	return &trades, nil
}
