# Gator CLI

A command-line RSS feed aggregator and reader built with Go.

## Prerequisites

To use the Gator CLI, you'll need the following installed on your system:

- **PostgreSQL**: Used to store and retrieve data
- **Go** (version 1.18 or later): Required to build and install the CLI tool

Make sure both are available in your system's `PATH`.

## Installation

1. Install the Gator CLI using Go:
    ```bash
    go install github.com/yourusername/gatorcli@latest
    ```

2. Ensure your `GOPATH/bin` is in your `PATH` to run the CLI directly from the terminal.

3. Create a configuration file:
    - Create `.gatorconfig.json` in your home directory with the following content:
    ```json
    {
        "url": "postgres://username:password@localhost:5432/yourdb"
    }
    ```
    > **Note**: Replace `username`, `password`, and `yourdb` with your actual PostgreSQL credentials and database name.

## Usage

Run the CLI using:
```bash
gator <command>
```

### Available Commands

| Command | Description |
|---------|-------------|
| `gator register <username>` | Registers and logs in a new user |
| `gator login <username>` | Logs in a user |
| `gator addfeed <url>` | Adds a URL to fetch from RSS |
| `gator agg <interval>` | Begins fetching the oldest feeds at specified interval |
| `gator browse <limit>` | Prints posts the logged-in user is following (up to limit) |

