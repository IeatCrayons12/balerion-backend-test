# Balerion Backend — Take Home #8

Convert decimal values to Thai baht text.

**Examples:**
- `1234` → `หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน`
- `33333.75` → `สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์`

---

## Project Structure

```
balerion-backend-test/
├── cmd/
│   └── main.go                  # Entry point — runs example conversions
├── internal/
│   └── thaibaht/
│       ├── converter.go         # Core conversion logic
│       └── converter_test.go    # Unit tests
├── go.mod
├── go.sum
└── README.md
```

---

## Prerequisites

Install Go on your machine if you haven't already.

**macOS (Homebrew):**
```bash
brew install go
```

**Linux:**
```bash
sudo apt-get update && sudo apt-get install -y golang
```

**Windows:**

Download and run the installer from https://go.dev/dl/ then restart your terminal.

Verify installation:
```bash
go version
# should print: go version go1.21.x ...
```

---

## Setup & Run (from a fresh machine)

```bash
# 1. Clone the repo
git clone https://github.com/IeatCrayons12/balerion-backend-test.git
cd balerion-backend-test

# 2. Install dependencies
go mod tidy

# 3. Run the program
go run ./cmd/main.go
```

**Expected output:**
```
1234
หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน

33333.75
สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์
```

---

## Run Tests

```bash
go test ./internal/thaibaht/... -v
```

---

## Design Notes

### Package structure
The converter lives in `internal/thaibaht` — keeping it internal means it can be imported by any service within this module (HTTP handler, gRPC service, worker) without being exposed as a public API. This is idiomatic Go for shared business logic.

### How the algorithm works
1. Split the decimal into integer part and fractional part (×100 for satang/สตางค์)
2. Recursively group by millions, then resolve positions (สิบ / ร้อย / พัน / หมื่น / แสน) within each group
3. Apply Thai language special rules:
   - `1` in tens place → `สิบ` (not `หนึ่งสิบ`)
   - `2` in tens place → `ยี่สิบ` (not `สองสิบ`)
   - `1` in ones place when preceded by tens → `เอ็ด` (not `หนึ่ง`)
4. Append `ถ้วน` if no fractional part, else convert fractional part and append `สตางค์`

### Integrating into a service

```go
import (
    "github.com/shopspring/decimal"
    "github.com/IeatCrayons12/balerion-backend-test/internal/thaibaht"
)

amount := decimal.NewFromFloat(1234.50)
text := thaibaht.Convert(amount)
// → "หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบสตางค์"
```

---

## Dependencies

| Package | Version | Purpose |
|---|---|---|
| github.com/shopspring/decimal | v1.4.0 | Arbitrary-precision decimal arithmetic |