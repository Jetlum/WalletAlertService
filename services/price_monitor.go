package services

import (
	"encoding/json"
	"fmt"
	"log"
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
		USD float64 `json:"ethereum"`
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
			if err := pm.updatePrices(); err != nil {
				log.Printf("Error updating prices: %v", err)
			}
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

	var prices CoinGeckoPrice
	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	pm.prices.Store("BTC", prices.Bitcoin.USD)
	pm.prices.Store("ETH", prices.Ethereum.USD)
	return nil
}

func (pm *PriceMonitor) GetPrice(cryptoID string) (float64, error) {
	price, ok := pm.prices.Load(cryptoID)
	if !ok {
		return 0, fmt.Errorf("price not available for %s", cryptoID)
	}
	return price.(float64), nil
}
