# codes
![gopher-codes](https://github.com/gonnafaraway/codes/assets/35832930/d945c23f-7c2d-4623-a5b3-47f09185db0c)

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
