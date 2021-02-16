package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"stocksync/pkg/client/contracts"
	"stocksync/pkg/client/internal"
	"stocksync/pkg/stockinfo/model"
	"strings"
)

const (
	priceDetailsPath = "/data/pricemultifull"
)

type StockClient struct {
	client  HTTPClient
	baseURL string
}

func (sc *StockClient) GetPriceData(ctx context.Context, fsym string, tsym string) (*model.StockInfo, error) {
	params := map[string]string{
		"fsyms": fsym,
		"tsyms": tsym,
	}

	u, err := internal.BuildURL(sc.baseURL, priceDetailsPath, params)
	if err != nil {
		return nil, fmt.Errorf("priceDetails.BuildURL. error: %v", err)
	}
	fmt.Println("URL: " + u.String())
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("priceDetails.NewRequest. error: %v", err)
	}

	resp, err := sc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("priceDetails.client.Do error: %v", err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("priceDetails.Resp.Reading error: %v", err)
	}


	fmt.Println("----------------")
	fmt.Println(string(data))
	fmt.Println("----------------")
	jsonData := gjson.Get(string(data), strings.Join([]string{"RAW", fsym, tsym}, "."))

	stockInfoFormat := &contracts.StockPrice{}

	fmt.Println("++++++++++++++++")
	fmt.Println(jsonData.String())
	fmt.Println("++++++++++++++++")
	err = json.Unmarshal([]byte(jsonData.String()), stockInfoFormat)
	if err != nil {
		return nil, fmt.Errorf("priceDetails.Resp.Parsing error: %v", err)
	}

	return stockInfoFormat.ToStockInfo(fsym, tsym), nil
}

func NewStockClient(client HTTPClient, baseURL string) *StockClient {
	return &StockClient{
		client:  client,
		baseURL: baseURL,
	}
}
