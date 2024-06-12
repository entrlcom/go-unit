# Unit

## Table of Content

- [Examples](#examples)
- [License](#license)

## Examples

```go
package main

import (
	"io"
	"net/http"

	"entrlcom.dev/unit"
)

func main() {
	// Ex: http.MaxBytesReader().
	http.MaxBytesReader(w, r, unit.KiB * 2) // 2 KiB.
	
	// Ex: io.LimitReader().
	io.LimitReader(r, unit.MiB * 4) // 4 MiB.
}

```

## License

[MIT](https://choosealicense.com/licenses/mit/)
