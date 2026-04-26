# ifsc-search

A self-contained HTTP service for locating Indian bank branches by bank plus a
fuzzy free-text query over branch name, address, and city. Built on a Bleve
index generated from `IFSC.csv` shipped with each
[`razorpay/ifsc` release](https://github.com/razorpay/ifsc/releases).

## Build the index

```bash
make build-index                      # downloads latest release CSV
make build-index-from CSV=./IFSC.csv  # uses a local CSV
```

The index lands in `ifsc-api/index/` (gitignored) along with `version.json`.

## Run the server

```bash
make run-search
# IFSC_SEARCH_PORT and IFSC_SEARCH_INDEX_PATH override defaults
```

## API

### `GET /search`

Query params:

| Name     | Required        | Notes                                          |
| -------- | --------------- | ---------------------------------------------- |
| `bank`   | one of bank/q   | 4-char IFSC bank code or fuzzy bank name       |
| `q`      | one of bank/q   | free-text over branch, city, address           |
| `limit`  | no              | default 20, max 100                            |
| `offset` | no              | default 0                                      |

Example:

```bash
curl 'http://localhost:8080/search?bank=HDFC&q=andheri+west&limit=5'
```

### `GET /healthz`

Returns the index version metadata and document count.

## Tests

```bash
make ifsc-api-test
```
