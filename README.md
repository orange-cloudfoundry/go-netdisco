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

**Important**:
For now user/password login produce an `api key`, a user can't have multiple on netdisco which can block usage with the
same login on multiple clients which are not aware of the api key.

Until netdisco allow multiple api key you can:

1. set `api_token_lifetime: 1576800000` in netdisco configuration, this will set expiration time to 50 years for a token
2. Produce an api key on https://my.netdisco.com/swagger-ui
3. use in this lib `netdisco.NewClientWithApiKey("https://netdisco2-demo.herokuapp.com", "api key", false)` everywhere
   you need it

This is a weak solution as at any login or logout made on api, will make api key change.
