# Unit

## Table of Content

- [Authors](#authors)
- [Examples](#examples)

## Authors

| Name         | GitHub                                             |
|--------------|----------------------------------------------------|
| Klim Sidorov | [@entrlcom-klim](https://github.com/entrlcom-klim) |

## Examples

```go
package main

import (
	"io"
	"net/http"

	"flida.dev/unit"
)

func main() {
	// Ex: http.MaxBytesReader().
	http.MaxBytesReader(w, r, unit.KiB * 2) // 2 KiB.
	
	// Ex: io.LimitReader().
	io.LimitReader(r, unit.MiB * 4) // 4 MiB.
}

```
