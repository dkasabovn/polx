package alpaca

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"polx/app/domain/bo"
	"polx/app/domain/definition"
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

func (a *alpacaRepo) GetBars(ticker, startDate, endDate string ) (*bo.AlpacaResponse, error) {
	url := fmt.Sprintf(baseUrl, ticker, startDate, endDate)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}


	request.Header.Set(		"APCA-API-KEY-ID" , "PKU0STVX1YQ9PG1OUKV3")
	request.Header.Set(		"APCA-API-SECRET-KEY" , "pAbAmL0OxCIEpkmRDWVg2f0G0gzrWeXwpwscNBX3")

	resp, err := a.client.Do(request)
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var trades bo.AlpacaResponse
	json.Unmarshal(body, &trades);
	fmt.Print(trades.Bars)
	fmt.Print("END")

	return &trades, nil
}
