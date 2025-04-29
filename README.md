# Stalcraft: X API Client

Go API client for the [Stalcraft: X API](https://eapi.stalcraft.net/overview.html).

## Installation
`go get -u github.com/kayex/stalcraft`

## Usage
You need to [create an application](https://eapi.stalcraft.net/registration.html#create-application) to access the API.
Authentication is not handled by this package.

The official [demo API](https://eapi.stalcraft.net/overview.html#demo-api) can be used to trial basic functionality. You can
access it by passing either of `stalcraft.DemoAppAccessToken` or `stalcraft.DemoUserAccessToken` along with
`stalcraft.WithDomain(stalcraft.DemoDomain)` to the client constructor.

### Basic example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"github.com/kayex/stalcraft"
)

func main() {
	client, err := stalcraft.NewClient(stalcraft.RegionEU, "your-access-token")
	if err != nil {
		log.Fatal(err)
	}
	
	emission, err := client.EmissionStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}
    
 	fmt.Printf("Current emission start time: %v\n", emission.CurrentStart)
}
```

### Endpoints
All available endpoints are found in [endpoints.go](https://github.com/kayex/stalcraft/tree/main/endpoints.go). Their names match the names in the [API reference](https://eapi.stalcraft.net/reference#/) as closely as possible.

## License
MIT
