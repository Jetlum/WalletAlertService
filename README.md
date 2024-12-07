# Wallet Alert Service

A Go-based microservice that monitors Ethereum wallet activities and cryptocurrency prices, sending customized alerts based on user preferences.

## Features

- **Real-time Transaction Monitoring**
  - Listens to Ethereum mainnet transactions
  - Detects large token transfers (> 1 ETH)
  - Identifies NFT transfers for popular collections (BAYC, Moonbirds)

- **Cryptocurrency Price Alerts**
  - Real-time price monitoring via CoinGecko API
  - Customizable price thresholds
  - Support for multiple cryptocurrencies (BTC, ETH)
  - Upper and lower bound price alerts

- **Customizable User Alerts**
  - Email notifications via SendGrid
  - Configurable alert thresholds
  - Per-wallet notification preferences
  - Price alert preferences

## Architecture

### Core Components

- **Event Detection**: [`nfts.NFTDetector`](nft/nftdetector.go) for NFT transfers
- **Price Monitoring**: [`services.PriceMonitor`](services/price_monitor.go) for cryptocurrency prices
- **Alert Services**: 
  - [`services.PriceAlertService`](services/price_alert.go) for price alerts
  - [`services.EmailNotification`](services/notification.go) for notifications
- **Data Storage**: GORM-based PostgreSQL integration
- **Repository Layer**: 
  - [`EventRepository`](repository/event_repository.go) for event storage
  - [`UserPreferenceRepository`](repository/user_preference.go) for user preferences
  - [`PriceAlertRepository`](repository/price_alert_repository.go) for price alerts

### Models

- [`Event`](models/event.go): Stores transaction details and event types
- [`UserPreference`](models/models.go): Manages user notification preferences
- [`PriceAlert`](models/models.go): Stores cryptocurrency price alert settings

### Technical Overview

Built with Go, this microservice connects to the Ethereum network, listens to events, indexes them based on user preferences, and triggers alerts accordingly.

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/Jetlum/WalletAlertService.git
    ```

2. **Navigate to the project directory**:
    ```sh
    cd WalletAlertService
    ```

3. **Install dependencies**:
    ```sh
    go mod download
    ```

## Configuration

Create a `config.yaml` file in the root directory:

	```yaml
	infura:
	  project_id: "YOUR_INFURA_PROJECT_ID"
	database:
	  url: "postgresql://username:password@localhost:5432/dbname"
	sendgrid:
	  api_key: "YOUR_SENDGRID_API_KEY"

## Usage

1.  **Configure the Ethereum client**:

	```sh
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	
2.	**Configure user preferences**

		// Transaction alerts
		userPreference := &models.UserPreference{
			UserID: "user@example.com",
			WalletAddress: "0x...",
			MinEtherValue: "1000000000000000000", // 1 ETH
			TrackNFTs: true,
			EmailNotification: true
		}

		// Price alerts
		priceAlert := &models.PriceAlert{
			UserID: "user@example.com",
			CryptocurrencyID: "BTC",
			ThresholdPrice: "50000.00",
			IsUpperBound: true,
			EmailNotification: true
		}
 
3.  **Run the application**:
	```sh
	go run main.go

## Testing

1.  **Run all tests**:

	```sh
	go test ./... -v

2.  **Run specific tests**:

	```sh
	go test -v ./services/... // Test notification services
	go test -v ./repository/... // Test repositories

## Key Components

**NFT Detection**:

	Pre-configured list of popular NFT contract addresses
	Extensible for adding new collections

**Transaction Processing**:

	Real-time block monitoring
	Transaction filtering and categorization
	Event creation and storage
 
**Notification System**:

	Email notifications via SendGrid
	User preference-based filtering
	Customizable notification templates

**Price Monitoring**:

	Real-time price monitoring via CoinGecko API
	Customizable price thresholds
	Support for multiple cryptocurrencies (BTC, ETH)
	Upper and lower bound price alerts

## Development
Project Structure

	├── config/			# Configuration management
	├── database/			# Database initialization and connection
	├── models/			# Data models
	├── nft/			# NFT detection logic
	├── repository/			# Data access layer
	├── services/			# Business logic and notifications

## Testing Environment
The project includes a robust testing setup:

	Unit Tests: Testing individual components
	Integration Tests: Testing component interactions

Set test environment:
	
 	export GO_ENV=test
	
## Dependencies

	go-ethereum: Ethereum client
	sendgrid-go: Email notifications
	gorm: Database ORM
	viper: Configuration management
	coingecko-api: CoinGecko API client