# fx
Fx is a useful functional programming helpers. 
Support **only** Go 1.18+.

## Features
- Slice
  - Each
  - Collect
  - Reduce
  - Find
  - Filter
  - Every
  - Some
  - Contains
  - Max
  - Min
  - GroupBy
  - Sample
  - SampleN
  - Union
  - Intersection
  - Uniq
  - IndexOf
  - LastIndexOf
  - Reverse
- Map
  - Keys
  - Values
  - Each
  - Filter
- Channel
  - Merge

## Documentation

Documentation with examples can be found here: https://nikgalushko.github.io/fx/

## Installation

slice helpers `go get github.com/nikgalushko/fx/slice`

map helpers `go get github.com/nikgalushko/fx/kv`

channel helpers `go get github.com/nikgalushko/fx/ch`

## Usage

```go
import (
  "fmt"

  "github.com/nikgalushko/fx/kv"
  "github.com/nikgalushko/fx/slice"
)

type (
  ID string

  Attribute struct {
    Value string
  }
)

func main() {
  m := map[ID]Attribute{
    ID("1"): {Value: "blah"},
    ID("1861"): {Value: "!"},
    ID("1234"): {Value: "yeah"},
  }

  fmt.Println("ids", kv.Keys(m), "contains special", slice.Contains(kv.Values(m), Attribute{Value: "!"}))
}
```