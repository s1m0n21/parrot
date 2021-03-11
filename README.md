# 📝 go-log

 ## Installation
```shell
go get github.com/s1m0n21/go-log
```

## Quick start
```go
package main

import (
	logger "github.com/s1m0n21/go-log"
)

func main() {
	// log to console
	log := logger.New("a")
	log.Infof("test")
	
	// set log level dynamic
	_ = logger.SetLevel("*", "error")
	
	// log to file
	// log file max size: 100MiB
	// log file max backups: 5
	// log file max age: 7days
	log = logger.New("b", logger.OptSetLogFile("b.log", 100, 5, 7))
	log.Infof("test")
}
```