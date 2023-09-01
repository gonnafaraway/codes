# codes
![gopher-codes](https://github.com/gonnafaraway/codes/assets/35832930/fb27d3da-8e7d-4242-ab06-7a8871dd0460)


## About
Library with static codes for different cases to use in your code and prevent developers to keep project code base dry and simple.
## List of codes categories
* Linux
* HTTP
* GRPC
* TCP/UDP ports
## Usage
```go
package main

import (
	"github.com/gonnafaraway/codes/linux"
	"os"
)

func main() {
	os.Exit(linux.WithError)
}
```
