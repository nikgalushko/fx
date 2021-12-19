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


## Installation

`go get -u github.com/nikgalushko/fx`

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