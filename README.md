# go-netdisco

It's a client lib written in go for [netdisco](https://github.com/netdisco/netdisco).

For now only searching device is supported as it was the only thing we wanted for now.

## Usage

import with `go get github.com/orange-cloudfoundry/go-netdisco`

You can now use lib in this way:

```go
package main

import (
	"fmt"
	"github.com/orange-cloudfoundry/go-netdisco"
)

func main() {
	client := netdisco.NewClient("https://netdisco2-demo.herokuapp.com", "guest", "guest", false)
	devices, err := client.SearchDevice(&netdisco.SearchDeviceQuery{
		Layers:   "7",
		Matchall: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(devices)
}
```
