# Wallet Alert Service

A Go-based microservice that monitors Ethereum wallet activities and cryptocurrency prices, sending customized alerts based on user preferences.

## Features

- **Real-time Transaction Monitoring**
  - Listens to Ethereum mainnet transactions
  - Detects large token transfers (> 1 ETH)
  - Identifies NFT transfers for popular collections (BAYC, Moonbirds)

- **Cryptocurrency Price Alerts**
  - Real-time price monitoring via CoinGecko API
  - Support for multiple cryptocurrencies (BTC, ETH)
  - Customizable price thresholds
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
  - [`services.EmailNotification`](services/email_notifier.go) for notifications
- **Data Storage**: GORM-based PostgreSQL integration
- **Repository Layer**: 
  - [`EventRepository`](repository/event_repository.go) for event storage
  - [`UserPreferenceRepository`](repository/user_preference.go) for user preferences
  - [`PriceAlertRepository`](repository/price_alert_repository.go) for price alerts
- **Mock Layer**:
  - [`mock`](mock/) package for testing with mock implementations

### Models

- [`Event`](models/event.go): Stores transaction details and event types
- [`UserPreference`](models/models.go): Manages user notification preferences
- [`PriceAlert`](models/models.go): Stores cryptocurrency price alert settings

## Installation

1. **Clone the repository**:

		git clone https://github.com/Jetlum/WalletAlertService.git

2.  **Navigate to the project directory**:
	
		cd WalletAlertService

3.  **Install dependencies**:

		go mod download

## Configuration

Create a `config.yaml` file in the root directory with the following content:

	infura:
		project_id: "YOUR_INFURA_PROJECT_ID"
	database:
		url: "postgresql://username:password@localhost:5432/dbname"
	sendgrid:
		api_key: "YOUR_SENDGRID_API_KEY"
	coingecko:
  		api_key: "YOUR_COINGECKO_API_KEY"
	price_check_interval: 1  # Interval in minutes for price checking

## Usage

1.  **Configure the Ethereum client**:

		go run main.go

2.	**Configure user preferences**

		// Transaction alerts
		userPreference := &models.UserPreference{
			UserID: "user@example.com",
			WalletAddress: "0x...",
			MinEtherValue: "1000000000000000000", // 1 ETH
			TrackNFTs: true,
			EmailNotification: true,
		}

		// Price alerts
		priceAlert := &models.PriceAlert{
			UserID: "user@example.com",
			CryptocurrencyID: "BTC",
			ThresholdPrice: "50000.00",
			IsUpperBound: true,
			EmailNotification: true,
		}
 
## Testing

1.  **Run all tests**:

		go test ./... -v

2.  **Run specific tests**:

		go test -v ./services/... // Test notification services
		go test -v ./repository/... // Test repositories

## Development
Project Structure

	├── config/         # Configuration management
	├── database/       # Database initialization and connection
	├── models/         # Data models and validation
	├── nft/            # NFT detection logic
	├── repository/     # Data access layer
	├── services/       # Business logic and notifications
	├── mock/           # Mock implementations for testing
	└── main.go        # Application entry point

## Key Components

**NFT Detection**:

	Pre-configured list of popular NFT contract addresses
	Extensible for adding new collections

**Transaction Processing**:

	Real-time block monitoring
	Transaction filtering and categorization
	Event creation and storage

**Price Monitoring**:

	Real-time cryptocurrency price tracking
	Configurable check intervals
	Support for multiple cryptocurrencies
	Threshold-based alerts
	Notification System
 
**Notification System**:

	Email notifications via SendGrid
	User preference-based filtering
	Customizable notification templates
	Support for both transaction and price alerts

## Development
Project Structure

	├── config/			# Configuration management
	├── database/			# Database initialization and connection
	├── models/			# Data models
	├── nft/			# NFT detection logic
	├── repository/			# Data access layer
	├── services/			# Business logic and notifications
	└── mock/			# Test mocks

## Testing Environment
The project includes a robust testing setup:

	Unit Tests: Testing individual components
	Mock Implementations: For external services and database
	Integration Tests: Testing component interactions

Set test environment:
	
 	export GO_ENV=test
	
## Dependencies

	go-ethereum: Ethereum client
	sendgrid-go: Email notifications
	gorm: Database ORM
	viper: Configuration management
	coingecko-api: CoinGecko API client
	