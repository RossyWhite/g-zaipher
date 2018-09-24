package zaipher

import (
	"net/http"
	"fmt"
)

const publicPath = "/api/1"

type PublicService struct {
	client *Client
}

type Currency struct {
	Name    string `json:"name"`
	IsToken bool   `json:"is_token"`
}

func (s *PublicService) Currencies(name string) ([]Currency, *http.Response, error) {
	path := fmt.Sprintf("%s/currencies/%s", publicPath, name)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var currencies []Currency
	resp, err := s.client.Do(req, &currencies)
	if err != nil {
		return nil, resp, err
	}

	return currencies, resp, nil
}

type CurrencyPair struct {
	Name            string  `json:"name"`
	Title           string  `json:"title"`
	CurrencyPairStr string  `json:"currency_pair"`
	Description     string  `json:"description"`
	IsToken         bool    `json:"is_token"`
	EventNumber     int     `json:"event_number"`
	Seq             int     `json:"seq"`
	ItemUnitMin     float64 `json:"item_unit_min"`
	ItemUnitStep    float64 `json:"item_unit_step"`
	ItemJapanese    string  `json:"item_japanese"`
	AuxUnitMin      float64 `json:"aux_unit_min"`
	AuxUnitStep     float64 `json:"aux_unit_step"`
	AuxUnitPoint    float64 `json:"aux_unit_point"`
	AuxJapanese     string  `json:"aux_japanese"`
}

func (s *PublicService) CurrencyPairs(name string) ([]CurrencyPair, *http.Response, error) {
	path := fmt.Sprintf("%s/currency_pairs/%s", publicPath, name)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var currencyPairs []CurrencyPair
	resp, err := s.client.Do(req, &currencyPairs)
	if err != nil {
		return nil, resp, err
	}

	return currencyPairs, resp, nil
}

type LastPrice struct {
	Price float64 `json:"last_price"`
}

func (s *PublicService) LastPrice(currencyPair string) (*LastPrice, *http.Response, error) {
	path := fmt.Sprintf("%s/last_price/%s", publicPath, currencyPair)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	lastPrice := &LastPrice{}
	resp, err := s.client.Do(req, lastPrice)
	if err != nil {
		return nil, resp, err
	}

	return lastPrice, resp, nil
}

type Ticker struct {
	Last   float64 `json:"last"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Vwap   float64 `json:"vwap"`
	Volume float64 `json:"volume"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
}

func (s *PublicService) Ticker(currencyPair string) (*Ticker, *http.Response, error) {
	path := fmt.Sprintf("%s/ticker/%s", publicPath, currencyPair)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	ticker := &Ticker{}
	resp, err := s.client.Do(req, ticker)
	if err != nil {
		return nil, resp, err
	}

	return ticker, resp, nil
}

type Trade struct {
	Date         int     `json:"date"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	TID          int     `json:"tid"`
	CurrencyPair string  `json:"currency_pair"`
	TradeType    string  `json:"trade_type"`
}

func (s *PublicService) Trades(currencyPair string) ([]Trade, *http.Response, error) {
	path := fmt.Sprintf("%s/trades/%s", publicPath, currencyPair)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, nil
	}

	var trades []Trade
	resp, err := s.client.Do(req, &trades)
	if err != nil {
		return nil, resp, err
	}

	return trades, resp, nil
}

type Depth struct {
	Bids [][2]float64 `json:"bids"`
	Asks [][2]float64 `json:"asks"`
}

func (s *PublicService) Depth(currencyPair string) (*Depth, *http.Response, error) {
	path := fmt.Sprintf("%s/depth/%s", publicPath, currencyPair)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	depth := &Depth{}
	resp, err := s.client.Do(req, depth)
	if err != nil {
		return nil, resp, err
	}

	return depth, resp, nil
}
