Unofficial and Incomplete Plivo HTTP API Helper Library in Go

<div style="margin: 25px;">
<a href="https://rapidapi.com/package/Plivo/functions?utm_source=PlivoGitHub-Go&utm_medium=button&utm_content=Vendor_GitHub" style="
    all: initial;
    background-color: #498FE1;
    border-width: 0;
    border-radius: 5px;
    padding: 10px 20px;
    color: white;
    font-family: 'Helvetica';
    font-size: 12pt;
    background-image: url(https://scdn.rapidapi.com/logo-small.png);
    background-size: 25px;
    background-repeat: no-repeat;
    background-position-y: center;
    background-position-x: 10px;
    padding-left: 44px;
    cursor: pointer;">
  Run now on <b>RapidAPI</b>
</a>
</div>

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
