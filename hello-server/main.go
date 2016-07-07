package main

import (
	"flag"
	"log"

	"github.com/razvanm/vanadium-core-hello"
	"v.io/v23"
	"v.io/v23/security/access"
	"v.io/x/ref/lib/security/securityflag"
	"v.io/x/ref/lib/signals"
	_ "v.io/x/ref/runtime/factories/roaming"
)

var (
	name = flag.String("name", "", "Name to use in mount table")
)

func main() {
	ctx, shutdown := v23.Init()
	defer shutdown()

	perms, err := securityflag.PermissionsFromFlag()
	if err != nil {
		log.Panic("Bad permissions: ", err)
	}
	auth := access.TypicalTagTypePermissionsAuthorizer(perms)

	_, s, err := v23.WithNewServer(ctx, *name, hello.HelloServer(newImpl(perms)), auth)
	if err != nil {
		log.Panic("Error listening: ", err)
	}
	for _, endpoint := range s.Status().Endpoints {
		log.Printf("ENDPOINT=%s\n", endpoint.Name())
	}

	<-signals.ShutdownOnSignals(ctx) // Wait forever.
}
