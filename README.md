# semaphore
POSIX semaphore library for Golang


# Usage

```go
package main

import (
	"log"
	"time"

	"github.com/zzhaolei/semaphore"
)

func main() {
	name := "test"
	mode := 0o644
	value := 2

	s := semaphore.New()
	if err := s.Open(name, mode, value); err != nil {
		log.Fatal(err)
	}

	log.Println("acquire...")
	s.Acquire()
	// if err := s.TryAcquire(); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("running...")
	time.Sleep(time.Second * 5)

	s.Release()
	log.Println("release")

	s.Unlink("test")
	log.Println("close and unlink")
}
```
