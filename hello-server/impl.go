package main

import (
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/verror"
	"github.com/razvanm/vanadium-core-hello"
)

type impl struct {
	greeting string
	perms    access.Permissions
}

func newImpl(perms access.Permissions) hello.HelloServerMethods {
	return &impl{
		greeting: "Hello World!",
		perms:    perms,
	}
}

func (f *impl) Get(_ *context.T, _ rpc.ServerCall) (greeting string, err error) {
	return f.greeting, nil
}

func (f *impl) Set(_ *context.T, _ rpc.ServerCall, greeting string) error {
	f.greeting = greeting
	return nil
}

func (f *impl) SetPermissions(ctx *context.T, _ rpc.ServerCall, _ access.Permissions, _ string) error {
	return verror.NewErrNotImplemented(ctx)
}

func (f *impl) GetPermissions(_ *context.T, _ rpc.ServerCall) (access.Permissions, string, error) {
	return f.perms, "", nil
}
