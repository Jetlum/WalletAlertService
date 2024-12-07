package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Jetlum/WalletAlertService/models"
	"github.com/Jetlum/WalletAlertService/repository"
)

type PriceAlertService struct {
	priceMonitor  *PriceMonitor
	alertRepo     *repository.PriceAlertRepository
	emailNotifier EmailNotifier
}

func NewPriceAlertService(
	priceMonitor *PriceMonitor,
	alertRepo *repository.PriceAlertRepository,
	emailNotifier EmailNotifier,
) *PriceAlertService {
	return &PriceAlertService{
		priceMonitor:  priceMonitor,
		alertRepo:     alertRepo,
		emailNotifier: emailNotifier,
	}
}

func (s *PriceAlertService) StartMonitoring() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			s.checkAlerts()
		}
	}()
}

func (s *PriceAlertService) checkAlerts() {
	alerts, err := s.alertRepo.GetActiveAlerts()
	if err != nil {
		log.Printf("Error fetching active alerts: %v", err)
		return
	}

	for _, alert := range alerts {
		price, err := s.priceMonitor.GetPrice(alert.CryptocurrencyID)
		if err != nil {
			log.Printf("Error getting price for %s: %v", alert.CryptocurrencyID, err)
			continue
		}

		threshold, _ := strconv.ParseFloat(alert.ThresholdPrice, 64)
		if s.shouldTriggerAlert(price, threshold, alert.IsUpperBound) {
			s.triggerAlert(&alert, price)
		}
	}
}

func (s *PriceAlertService) shouldTriggerAlert(currentPrice, threshold float64, isUpperBound bool) bool {
	if isUpperBound {
		return currentPrice >= threshold
	}
	return currentPrice <= threshold
}

func (s *PriceAlertService) triggerAlert(alert *models.PriceAlert, currentPrice float64) {
	message := fmt.Sprintf(
		"Price Alert: %s has reached $%.2f (Threshold: $%s)",
		alert.CryptocurrencyID,
		currentPrice,
		alert.ThresholdPrice,
	)

	if alert.EmailNotification {
		// Create event for notification
		event := &models.Event{
			EventType: "PRICE_ALERT",
			Value:     fmt.Sprintf("%.2f", currentPrice),
		}

		userPref := &models.UserPreference{
			UserID:            alert.UserID,
			EmailNotification: true,
		}

		if err := s.emailNotifier.Send(event, userPref); err != nil {
			log.Printf("Failed to send price alert email: %v", err)
		}
	}

	// Deactivate the alert after triggering
	if err := s.alertRepo.DeactivateAlert(alert.ID); err != nil {
		log.Printf("Failed to deactivate alert: %v", err)
	}
}
