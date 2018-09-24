package zaipher

import (
	"net/http"
	"net/url"
	"strconv"
)

const tradePath = "/tapi"

type TradeService struct {
	client *Client
}

type Info struct {
	Info2
	TradeCount int `json:"trade_count"`
}

func (s *TradeService) GetInfo() (*Info, *http.Response, error) {
	path := tradePath
	params := url.Values{}
	params.Add("method", "get_info")

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	info := &Info{}
	resp, err := s.client.Do(req, &ResponseEnvelop{Return: info})
	if err != nil {
		return nil, resp, err
	}
	return info, resp, nil
}

type Info2 struct {
	Funds      map[string]float64 `json:"funds"`
	Deposit    map[string]float64 `json:"deposit"`
	Rights     map[string]int     `json:"rights"`
	OpenOrders int                `json:"open_orders"`
	ServerTime int                `json:"server_time"`
}

func (s *TradeService) GetInfo2() (*Info2, *http.Response, error) {
	path := tradePath
	params := url.Values{}
	params.Add("method", "get_info2")

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	info := &Info2{}
	resp, err := s.client.Do(req, &ResponseEnvelop{Return: info})
	if err != nil {
		return nil, resp, err
	}
	return info, resp, nil
}

type PersonalInfo struct {
	RankingNickname string `json:"ranking_nickname"`
	IconPath        string `json:"icon_path"`
}

func (s *TradeService) GetPersonalInfo() (*PersonalInfo, *http.Response, error) {
	path := tradePath
	params := url.Values{}
	params.Add("method", "get_personal_info")

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	info := &PersonalInfo{}
	resp, err := s.client.Do(req, &ResponseEnvelop{Return: info})
	if err != nil {
		return nil, resp, err
	}
	return info, resp, nil
}

type IDInfo struct {
	User struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		Name      string `json:"name"`
		Kana      string `json:"kana"`
		Certified bool   `json:"certified"`
	}
}

func (s *TradeService) GetIDInfo() (*IDInfo, *http.Response, error) {
	path := tradePath
	params := url.Values{}
	params.Add("method", "get_id_info")

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	info := &IDInfo{}
	resp, err := s.client.Do(req, &ResponseEnvelop{Return: info})
	if err != nil {
		return nil, resp, err
	}

	return info, resp, nil
}

type TradeHistory struct {
	CurrencyPair string  `json:"currency_pair"`
	Action       string  `json:"action"`
	Amount       float64 `json:"amount"`
	Price        float64 `json:"price"`
	Fee          float64 `json:"fee"`
	YourAction   string  `json:"your_action"`
	Bonus        float64 `json:"bonus"`
	Timestamp    string  `json:"timestamp"`
	Comment      string  `json:"comment"`
}

type HistoryOpts struct {
	From   int
	Count  int
	FromID int
	EndID  int
	Order  string
	Since  int
	End    int
}

type TradeHistoryOpts struct {
	HistoryOpts
	CurrencyPair string
}

func (s *TradeService) TradeHistory(opts *TradeHistoryOpts) (map[string]TradeHistory, *http.Response, error) {
	path := tradePath
	params := AddOptions(opts)
	params.Add("method", "trade_history")

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	var info map[string]TradeHistory
	data := &ResponseEnvelop{Return: &info}
	resp, err := s.client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return info, resp, nil

}

type ActiveOrder struct {
	CurrencyPair string `json:"currency_pair"`
	Action       string `json:"action"`
	Amount       string `json:"amount"`
	Price        int    `json:"price"`
	Timestamp    string `json:"timestamp"`
	Comment      string `json:"comment"`
}

func (s *TradeService) ActiveOrders(currencyPair string) (map[string]ActiveOrder, *http.Response, error) {
	path := tradePath
	params := url.Values{}
	params.Add("method", "active_orders")
	params.Add("currency_pair", currencyPair)

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	var orders map[string]ActiveOrder
	data := &ResponseEnvelop{Return: &orders}
	resp, err := s.client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return orders, resp, nil
}

type TradeResult struct {
	Received float64            `json:"received"`
	Remains  float64            `json:"remains"`
	OrderID  int                `json:"order_id"`
	Funds    map[string]float64 `json:"funds"`
}

type TradeOpts struct {
	Limit   float64
	Comment string
}

func (s *TradeService) Trade(currencyPair, action string, price, amount float64,
	opts *TradeOpts) (*TradeResult, *http.Response, error) {

	path := tradePath
	params := AddOptions(opts)
	params.Add("method", "trade")
	params.Add("currency_pair", currencyPair)
	params.Add("action", action)
	params.Add("price", strconv.FormatFloat(price, 'f', 4, 64))
	params.Add("amount", strconv.FormatFloat(amount, 'f', 4, 64))

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	result := &TradeResult{}
	data := &ResponseEnvelop{Return: result}
	resp, err := s.client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

type WithdrawResult struct {
	ID    int                `json:"id"`
	TxID  string             `json:"txid"`
	Fee   float64            `json:"fee"`
	Funds map[string]float64 `json:"funds"`
}

type WithDrawOpts struct {
	Message string
	OptFee  float64
}

func (s *TradeService) WithDraw(currencyPair, address string, amount float64,
	opts *WithDrawOpts) (*WithdrawResult, *http.Response, error) {

	path := tradePath
	params := AddOptions(opts)
	params.Add("method", "withdraw")
	params.Add("currency_pair", currencyPair)
	params.Add("address", address)
	params.Add("amount", strconv.FormatFloat(amount, 'f', 4, 64))

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	result := &WithdrawResult{}
	data := &ResponseEnvelop{Return: result}
	resp, err := s.client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

type depositAndWithdrawHistory struct {
	Timestamp int     `json:"timestamp"`
	Address   string  `json:"address"`
	Amount    float64 `json:"amount"`
	TxID      string  `json:"txid"`
}

type DepositHistory struct {
	depositAndWithdrawHistory
}

func (s *TradeService) DepositHistory(currency string, opts *HistoryOpts) (map[string]DepositHistory, *http.Response, error) {
	path := tradePath
	params := AddOptions(opts)
	params.Add("method", "deposit_history")
	params.Add("currency", currency)

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	var histories map[string]DepositHistory
	data := &ResponseEnvelop{Return: &histories}
	resp, err := s.client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return histories, resp, nil

}

type WithdrawHistory struct {
	depositAndWithdrawHistory
}

func (s *TradeService) WithdrawHistory(currency string, opts *HistoryOpts) (map[string]WithdrawHistory, *http.Response, error) {
	path := tradePath
	params := AddOptions(opts)
	params.Add("method", "withdraw_history")
	params.Add("currency", currency)

	req, err := s.client.NewRequestWithAuth("POST", path, params)
	if err != nil {
		return nil, nil, err
	}

	var histories map[string]WithdrawHistory
	data := &ResponseEnvelop{Return: &histories}
	resp, err := s.client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return histories, resp, nil

}
