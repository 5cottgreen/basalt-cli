# basalt-cli
basalt-cli is a package created as a test exercise for "Базальт СПО".
# Get started
go get github.com/5cottgreen/basalt-cli
# Example
```golang
package main
    
import (
  "fmt"
  "github.com/5cottgreen/basalt-cli"
)

func main() {
  basaltcli.HandleCommand()
}
```
# Usage
Commands
1. get [arg] [arg] -s simple fetch
2. get [arg] [arg] -c comparative fetch
3. get [arg] [arg] -cr comparative fetch in reverse way
3. get [arg] [arg] -cr comparative fetch by version


