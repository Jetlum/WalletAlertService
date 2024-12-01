# Wallet Alert Service

A Go-based microservice that monitors Ethereum wallet activities and sends customized alerts based on user preferences.

## Features

- **Real-time Transaction Monitoring**
  - Listens to Ethereum mainnet transactions
  - Detects large token transfers (> 1 ETH)
  - Identifies NFT transfers for popular collections (BAYC, Moonbirds)

- **Customizable User Alerts**
  - Email notifications via SendGrid
  - Configurable alert thresholds
  - Per-wallet notification preferences

- **Event Types**
  - Large transfers (> 1 ETH)
  - NFT transfers
  - Custom threshold alerts

## Architecture

### Core Components

- **Event Detection**: [`nfts.NFTDetector`](nft/nftdetector.go) for NFT transfers
- **Notification Service**: [`services.EmailNotification`](services/email_notifier.go) for sending alerts
- **Data Storage**: GORM-based PostgreSQL integration
- **Repository Layer**: 
  - [`EventRepository`](repository/event_repository.go) for event storage
  - [`UserPreferenceRepository`](repository/user_preference.go) for user preferences

### Models

- [`Event`](models/event.go): Stores transaction details and event types
- [`UserPreference`](models/models.go): Manages user notification preferences

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

## Usage

1.  **Configure the Ethereum client**:

	go run main.go

2.	**Configure user preferences**

	userPreference := &models.UserPreference{
		UserID: "user@example.com",
		WalletAddress: "0x...",
		MinEtherValue: "1000000000000000000", // 1 ETH
		TrackNFTs: true,
		EmailNotification: true
	}
 
## Testing

1.  **Run all tests**:

	go test ./... -v

2.  **Run specific tests**:

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