Unofficial and Incomplete Plivo HTTP API Helper Library in Go
## INSTALL
    go get github.com/iwanbk/plivogo

## Implemented API
    Account
    Endpoint
    Application

## Example

```go
package main

import (
    "github.com/iwanbk/plivogo"
)

func main() {
	authId := "AAAA"
	authToken := "BBBBBB"

	p, _ := plivogo.NewPlivo(AUTH_ID, AUTH_TOKEN)

    //get endpoint details
	endpoint, _ := p.Endpoint.Get("your_endpoint_id")
	
}
```

You can also take a look at test files
