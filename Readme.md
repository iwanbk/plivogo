Unofficial and Incomplete Plivo HTTP API Helper Library in Go
## INSTALL
    go get github.com/iwanbk/plivogo

## Implemented API
    Account
        get Account Detail
    Call
        make call

## Example

```go
package main

import (
    "github.com/iwanbk/plivogo"
)

func main() {
	authId := "AAAA"
	authToken := "BBBBBB"

	c, _ := plivogo.NewClient(AUTH_ID, AUTH_TOKEN)

    //get Account detail
	c.Account.Get()
	
    //make a call
	p, _ := plivogo.NewParams()
	p.Set("to", "sip:ch1303@phone.plivo.com")
	p.Set("caller_name", "Somenone")
	p.Set("from", "12345678901")
	p.Set("answer_url", "http://my.web.com/123.xml")
	p.Set("answer_method", "GET")
	c.Call.Call(p.Dumps())
}
```
