# Go Soccer MCP Server

## Overview

Go Soccer is a Go application that provides soccer-related data for teams through both a RESTful API and an MCP (Model Context Protocol) server. It delivers information such as latest and upcoming matches, player rosters, transfer history, and trophies. The API scrapes public soccer data sources and exposes endpoints for easy integration with other applications or services, while the MCP server enables integration with MCP-compatible platforms (such as VS Code Copilot) via the MCP protocol.

## Features

- **Latest Matches**: Retrieve the most recent matches for a given team.
- **Next Matches**: Get upcoming matches for a team.
- **Players**: List all players in a team, including position, age, and country.
- **Transfers**: View recent player transfers for a team.
- **Trophies**: List all trophies won by a team.

## Getting Started

### Prerequisites

- Go 1.20+ installed
- (Optional) `make` for build/run convenience

## MCP Server

In addition to the RESTful API, this project also implements an MCP (Model Context Protocol) server for integration with MCP-compatible platforms, such as VS Code Copilot.

### Running the MCP Server

You can run the MCP server directly with Go:

```sh
$ go run cmd/soccer/mcp/main.go
```

> **Note:** The MCP server uses stdio transport and does not expose HTTP endpoints. It is intended for programmatic integrations, such as MCP agents or plugins that support the protocol.

The available MCP tools are:

- get_latest_matches
- get_next_matches
- get_players
- get_transfers
- get_trophies

Each tool accepts a `team` parameter and returns data similar to the REST endpoints.

### Installation

Clone the repository:

```sh
$ git clone https://github.com/Sup3r-Us3r/go-soccer.git
$ cd go-soccer
```

Build the application:

```sh
$ make build
```

### Running the Application

You can run the server directly with Go:

```sh
$ make run
```

Or run the compiled binary:

```sh
$ make start
```

The server will start on `http://localhost:8080`.

## API Endpoints

All endpoints require a `teamName` query parameter.

### 1. Get Latest Matches

- **Endpoint:** `GET /api/soccer/matches/latest?teamName=<team>`
- **Response:**

```json
{
  "matches": [
    {
      "title": "Team A vs Team B",
      "home": "Team A",
      "homeLogo": "https://...",
      "away": "Team B",
      "awayLogo": "https://...",
      "date": "2025-08-01",
      "scoreBoard": "2-1"
    }
  ]
}
```

### 2. Get Next Matches

- **Endpoint:** `GET /api/soccer/matches/next?teamName=<team>`
- **Response:**

```json
{
  "matches": [
    {
      "title": "Team A vs Team C",
      "home": "Team A",
      "homeLogo": "https://...",
      "away": "Team C",
      "awayLogo": "https://...",
      "date": "2025-08-15"
    }
  ]
}
```

### 3. Get Players

- **Endpoint:** `GET /api/soccer/players?teamName=<team>`
- **Response:**

```json
{
  "players": [
    {
      "position": "Forward",
      "player": "John Doe",
      "age": "28",
      "country": "Brazil"
    }
  ]
}
```

### 4. Get Transfers

- **Endpoint:** `GET /api/soccer/transfers?teamName=<team>`
- **Response:**

```json
{
  "transfers": [
    {
      "date": "2025-07-10",
      "type": "In",
      "player": "Jane Smith",
      "team": {
        "name": "Team D",
        "url": "https://..."
      }
    }
  ]
}
```

### 5. Get Trophies

- **Endpoint:** `GET /api/soccer/trophies?teamName=<team>`
- **Response:**

```json
{
  "trophies": [
    {
      "year": "2024",
      "championship": {
        "name": "Champions League",
        "url": "https://..."
      }
    }
  ]
}
```

## Error Handling

If the `teamName` parameter is missing or invalid, the API will return an error message with an appropriate HTTP status code.

## License

This project is licensed under the MIT License.
