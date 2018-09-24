# g-zaipher

go client library for [Zaif API](https://techbureau-api-document.readthedocs.io/ja/latest/#)

## Install

```sh
$ go get github.com/RossyWhite/g-zaipher/zaipher
```

## Usage

```go
func main() {
    // initialize client
    zaif := zaipher.NewClient(&zaipher.Config{
        AuthConfig: &zaipher.AuthConfig{
            APIkey:    "===YOUR API KEY===",
            APISecret: "===YOUR API SECRET==="
        },
    })
    
    // with custom host
    zaif := zaipher.NewClient(&zaipher.Config{
        AuthConfig: &zaipher.AuthConfig{
            APIkey:    "===YOUR API KEY===",
            APISecret: "===YOUR API SECRET==="
        },
        BaseURL: "https://<YOUR_HOST>/",
    })
    
    // public service
    lastPrice, resp, err := zaif.Public.LastPrice("btc_jpy")

    // trade service
    opt := zaipher.TradeOpts{Comment: "Happy Trading!!!"}
    result, resp, err := zaif.Trade.Trade("btc_jpy", "bid", 777777, 77, &opt)
}
```

## License

[MIT License](https://opensource.org/licenses/MIT).
