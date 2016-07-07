package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/razvanm/vanadium-core-hello"
	"v.io/v23"
	"v.io/v23/context"
	_ "v.io/x/ref/runtime/factories/generic"
)

var (
	server = flag.String("server", "", "Name of the server to connect to")
	setValue = flag.String("set", "", "If not empty, issue a Set RPC with the indicated string")
)

func main() {
	ctx, shutdown := v23.Init()
	defer shutdown()

	f := hello.HelloClient(*server)
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	if *setValue != "" {
		f.Set(ctx, *setValue)
	} else {
		hello, _ := f.Get(ctx)
		fmt.Println(hello)
	}
}