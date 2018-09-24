package zaipher

import (
	"net/url"
	"net/http"
	"encoding/json"
	"strings"
	"reflect"
	"strconv"
	"unicode"
)

const (
	defaultBaseURL = "https://api.zaif.jp/"
	userAgent      = "g-zaipher"
)

type Client struct {
	BaseURL    *url.URL
	AuthConfig *AuthConfig
	httpClient *http.Client
	UserAgent  string

	// services
	Public *PublicService
	Trade  *TradeService
}

type Config struct {
	AuthConfig *AuthConfig
	BaseURL    string
}

type ResponseEnvelop struct {
	Return interface{} `json:"return"`
}

func NewClient(opts *Config) *Client {
	httpClient := http.DefaultClient
	c := &Client{
		httpClient: httpClient,
		BaseURL:    parseHost(opts.BaseURL),
		AuthConfig: opts.AuthConfig,
		UserAgent:  userAgent,
	}

	c.Public = &PublicService{client: c}
	c.Trade = &TradeService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, params url.Values) (*http.Request, error) {
	ref, _ := url.Parse(urlStr)
	u := c.BaseURL.ResolveReference(ref)

	req, err := http.NewRequest(method, u.String(), strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) NewRequestWithAuth(method, urlStr string, params url.Values) (*http.Request, error) {
	if c.AuthConfig == nil || *c.AuthConfig == (AuthConfig{}) {
		return nil, newAPIError("API keys are not set")
	}

	ref, _ := url.Parse(urlStr)
	u := c.BaseURL.ResolveReference(ref)

	params.Add("nonce", GetNonce())
	encParams := params.Encode()
	req, err := http.NewRequest(method, u.String(), strings.NewReader(encParams))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Key", c.AuthConfig.APIkey)
	req.Header.Set("Sign", MakeHMAC(encParams, c.AuthConfig.APISecret))

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := checkApiError(json.NewDecoder(resp.Body))
	if err != nil {
		return resp, err
	}

	if err = json.Unmarshal(*b, v); err != nil {
		return resp, err
	}

	return resp, nil
}

func AddOptions(params interface{}) url.Values {
	values := url.Values{}

	if params == nil || reflect.ValueOf(params).IsNil() {
		return values
	}

	iVal := reflect.ValueOf(params).Elem()
	typ := iVal.Type()

	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)

		if isZeroValue(f.Interface()) {
			continue
		}

		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(toSnakeCase(typ.Field(i).Name), v)
	}

	return values
}

func isZeroValue(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func parseHost(urlString string) *url.URL {
	var host *url.URL
	if len(urlString) != 0 {
		host, _ = url.Parse(urlString)
		return host
	}

	host, _ = url.Parse(defaultBaseURL)
	return host
}

func checkApiError(d *json.Decoder) (*json.RawMessage, error) {
	var raw json.RawMessage
	if err := d.Decode(&raw); err != nil {
		return nil, err
	}

	var e struct{ Msg string `json:"error"` }
	if err := json.Unmarshal(raw, &e); err == nil && len(e.Msg) > 0 {
		return nil, newAPIError(e.Msg)
	}
	return &raw, nil
}

func toSnakeCase(in string) string {
	runes := []rune(in)

	var out []rune
	for i := 0; i < len(runes); i++ {
		if i > 0 && (unicode.IsUpper(runes[i]) || unicode.IsNumber(runes[i])) && ((i+1 < len(runes) && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
