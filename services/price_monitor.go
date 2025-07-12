package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type PriceMonitor struct {
	apiKey     string
	httpClient *http.Client
	prices     sync.Map
}

type CoinGeckoPrice struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
	Ethereum struct {
		USD float64 `json:"usd"`
	} `json:"ethereum"`
}

func NewPriceMonitor(apiKey string) *PriceMonitor {
	return &PriceMonitor{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (pm *PriceMonitor) StartMonitoring(checkInterval time.Duration) {
	ticker := time.NewTicker(checkInterval)
	go func() {
		for range ticker.C {
			pm.updatePrices()
		}
	}()
}

func (pm *PriceMonitor) updatePrices() error {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd"
	resp, err := pm.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch prices: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if btc, ok := result["bitcoin"]; ok {
		if price, ok := btc["usd"]; ok {
			pm.prices.Store("BTC", price)
		}
	}

	if eth, ok := result["ethereum"]; ok {
		if price, ok := eth["usd"]; ok {
			pm.prices.Store("ETH", price)
		}
	}

	return nil
}

func (pm *PriceMonitor) GetPrice(cryptoID string) (float64, error) {
	price, ok := pm.prices.Load(cryptoID)
	if !ok {
		return 0, fmt.Errorf("price not available for %s", cryptoID)
	}
	return price.(float64), nil
}
