

# Custom Alerts for Wallet Activities

  

This project allows users to set up personalized alerts for wallet activities like large token transfers, NFT mints, or specific on-chain events.

  

## Features

  

- Personalized alerts for various wallet activities

- Notifications via email, SMS, or push notifications

- Efficient filtering and indexing of on-chain events

- NFT transfer detection

  

## Technical Overview

  

Built with Go, this microservice connects to the Ethereum network, listens to events, indexes them based on user preferences, and triggers alerts accordingly.

  

## Installation

  

1.  **Clone the repository**:

	```sh
	git clone https://github.com/Jetlum/WalletAlertServicee.git
2.  **Navigate to the project directory**:
	```sh
	cd WalletAlertService
3.  **Install dependencies**:
	```sh
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

	```sh
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/YOUR_INFURA_PROJECT_ID")
2.	**Configure user preferences**

		userPreference := &models.UserPreference{
			UserID: "user@example.com",
			WalletAddress: "0x...",
		    MinEtherValue: "1000000000000000000", // 1 ETH
		    TrackNFTs: true,
		    EmailNotification: true
		})
 
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
	Transaction Processing:

**Real-time block monitoring**:

	Transaction filtering and categorization
	Event creation and storage
 
**Notification System**:

	Email notifications via SendGrid
	User preference-based filtering
	Customizable notification templates

## Development
Project Structure

	├── config/         		# Configuration management
	├── database/       	# Database initialization and connection
	├── models/         	# Data models
	├── nft/           			# NFT detection logic
	├── repository/    	# Data access layer
	├── services/      		# Business logic and notifications
	└── mock/          		# Test mocks