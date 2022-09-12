# 🦜 Parrot
A logging tool based on zap

 ## Installation
```shell
go get github.com/s1m0n21/parrot
```

## Quick start
```go
package main

import "github.com/s1m0n21/parrot"

func main() {
	// log to console
	log := parrot.New("a")
	log.Infof("test")
	
	// change log level
	_ = parrot.SetLevel("*", "error")
	
	// log to file
	// log file max size: 100MiB
	// log file max backups: 5
	// log file max age: 7days
	log = parrot.New("b", parrot.OptSetLogFile("b.log", 100, 5, 7))
	log.Infof("test")
}
```
