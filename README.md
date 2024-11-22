# Custom Alerts for Wallet Activities

This project allows users to set up personalized alerts for wallet activities like large token transfers, NFT mints, or specific on-chain events.

## Features

- Personalized alerts for various wallet activities
- Notifications via email, SMS, or push notifications
- Efficient filtering and indexing of on-chain events

## Technical Overview

Built with Go, this microservice connects to the Ethereum network, listens to events, indexes them based on user preferences, and triggers alerts accordingly.

## Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/yourusername/yourrepository.git

2. **Navigate to the project directory**:
    ```sh
    cd yourrepository

3. **Install dependencies**:
    ```sh
    go mod download

## Usage

1. **Configure the Ethereum client**:
    ```sh
    client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/YOUR_INFURA_PROJECT_ID")

2. **Run the application**:
    ```sh
    go run main.go

## Configuration

User Preferences: 
    Configure user alert preferences in the application settings or through the user interface.


