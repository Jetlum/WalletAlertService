
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
2.  **Run the application**:
	```sh
	go run main.go