`import github.com/abrander/storm-jsoniter`
===========================================

[jsoniter](https://github.com/json-iterator/go/) codec
for [asdine/storm](https://github.com/asdine/storm).

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
