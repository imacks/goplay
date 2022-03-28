Goplay
======
This is a command-line client for https://play.golang.org. It uses https://github.com/haya14busa/goplay API client under 
the hood.

Key differences from https://github.com/haya14busa/goplay/tree/v1.0.0/cmd/goplay:

1. added the `-endpoint` flag option
2. the `-share` flag now defaults to `false`
3. the `-openbrowser` flag now defaults to `false`

Example:

```bash
echo 'package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
' | goplay
```

**Hint**: China users can try `goplay -endpoint https://golang.google.cn/_`.


Build from source
-----------------
Trim path and strip symbols for a minimal build:

```bash
GO111MODULE=on go install -v -x -trimpath -ldflags="-s -w" github.com/imacks/goplay/cmd/goplay@v1.0.0
```
