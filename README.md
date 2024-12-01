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

## Technical Overview

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

Create a `config.yaml` file in the root directory with the following content:

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

		userPreference := &models.UserPreference{
			UserID: "user@example.com",
			WalletAddress: "0x...",
			MinEtherValue: "1000000000000000000", // 1 ETH
			TrackNFTs: true,
			EmailNotification: true,
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
	Mock Implementations: For external services and database
	Integration Tests: Testing component interactions

Set test environment:
	
 	export GO_ENV=test
	
## Dependencies

	go-ethereum: Ethereum client
	sendgrid-go: Email notifications
	gorm: Database ORM
	viper: Configuration management