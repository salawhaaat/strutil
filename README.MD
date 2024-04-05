# strutil

## Functions

### Reverse

```go
func Reverse(s string) string
```

Reverse returns its argument string reversed rune-wise left to right.

### SymbolCount

```go
func SymbolCount(s string) int
```

SymbolCount returns the number of symbols in the given string.

## Example

```go
package main

import (
    "fmt"
    "github.com/salawhaaat/strutil"
)

func main() {
    // Example for Reverse
    fmt.Println(strutil.Reverse("hello"))
    // Output: olleh

    // Example for SymbolCount
    fmt.Println(strutil.SymbolCount("Hello, 世界"))
    // Output: 9
}
```

## Testing

- **Table-Driven Tests**: For easy addition of new cases.
- **Anonymous Structs**: Used in test cases for potential reuse.
- **Unordered Testing with Maps**: Ensures code doesn't rely on specific execution order.
- **Subtests**: Continues testing even if a test fails.
- **Google's `cmp` Function**: Used for better, more informative error messages.

## Tests

To run tests, use the following command:

```bash
go test -v
```
