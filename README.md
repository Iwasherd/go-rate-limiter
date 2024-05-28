# go-rate-limiter

A simple rate limiter for Go.

## Installation

```bash
go get github.com/rohanprasad/go-rate-limiter
```

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/rohanprasad/go-rate-limiter"
)

func main() {
    // Create a new rate limiter that allows 10 requests per second
    limiter := rate.NewLimiter(10, time.Second)

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

