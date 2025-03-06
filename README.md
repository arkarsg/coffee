# Coffee Bot

## Backend

### Pre-requisites
- Env vars
    ```
    TELEGRAM_TOKEN=
    DB_URI=mongodb://mongo:mongo@localhost:27017/
    ```
- `Go 1.24.1`
- Docker

1. Run MongoDB DB
    ```
    docker compose up
    ```

2. Start Telegram Bot
    ```
    make app
    ```
or
    ```
    go run ./cmd/*.go
    ```
