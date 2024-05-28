# go-rate-limiter

A simple rate limiter for Go.

## Installation

```bash
go get github.com/iwasherd/ratelimiter
```

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/iwasherd/ratelimiter"
)

func main() {
	// Create a new timestamps storage
	storage := ratelimiter.NewMemoryTimeStorage()
    // Create a new rate limiter that allows 10 requests per second
    limiter := ratelimiter.New(10, time.Second, storage)

    // Check if the rate limiter allows a request
    if limiter.Allow() {
        fmt.Println("Request allowed")
    } else {
        fmt.Println("Request denied")
    }
}
```

## License
```
MIT
```


## Local Development

### Prerequisites

- [Go](https://golang.org/dl/) (v1.22)
- [Make](https://www.gnu.org/software/make/) (v4.3)

### Setup

1. Clone the repository

```bash
git clone
```

2. Change the working directory

```bash
cd go-rate-limiter
```

3. Run the tests

```bash

make test
```

