`import github.com/abrander/storm-jsoniter`
===========================================

[jsoniter](https://jsoniter.com/) codec
for [asdine/storm](https://github.com/asdine/storm) for more speed.

Example of Usage
----------------

```go
package main

import (
	"github.com/abrander/storm-jsoniter"
	"github.com/asdine/storm"
)

func main() {
    db, _ := storm.Open(
		"database.db",
        storm.Codec(jsoniter.Codec),
	)

    // Do stuff.

    db.Close()
}
```

Comparison with `encoding/json` and `jsonparser`
------------------------------------------------
<a href="https://jsoniter.com/benchmark.html"><img src="http://jsoniter.com/benchmarks/go-reader.png" alt="benchmark"></a>
