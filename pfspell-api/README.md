# PF-Spell API

Requires GO installed to run, more info can be found [here](https://go.dev/)

Required Meilisearch, can be any of the installations described [here](https://docs.meilisearch.com/learn/getting_started/quick_start.html)


## Setup

```bash
go install
```


## Run the server

```bash
MEILI_HOST=http://127.0.0.1:7700 MEILI_KEY=MASTER_KEY go run cmd/server/main.go
```

