// This file was auto-generated by the vanadium vdl tool.
// Package: hello

package hello

import (
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/services/permissions"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Interface definitions

// HelloClientMethods is the client interface
// containing Hello methods.
type HelloClientMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(perms access.Permissions, version string) error         {Red}
	//    GetPermissions() (perms access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectClientMethods
	Get(*context.T, ...rpc.CallOpt) (greeting string, _ error)
	Set(_ *context.T, greeting string, _ ...rpc.CallOpt) error
}

// HelloClientStub adds universal methods to HelloClientMethods.
type HelloClientStub interface {
	HelloClientMethods
	rpc.UniversalServiceMethods
}

// HelloClient returns a client stub for Hello.
func HelloClient(name string) HelloClientStub {
	return implHelloClientStub{name, permissions.ObjectClient(name)}
}

type implHelloClientStub struct {
	name string

	permissions.ObjectClientStub
}

func (c implHelloClientStub) Get(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Get", nil, []interface{}{&o0}, opts...)
	return
}

func (c implHelloClientStub) Set(ctx *context.T, i0 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Set", []interface{}{i0}, nil, opts...)
	return
}

// HelloServerMethods is the interface a server writer
// implements for Hello.
type HelloServerMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/services/permissions"
	//
	//   type MyObject interface {
	//     permissions.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(perms access.Permissions, version string) error         {Red}
	//    GetPermissions() (perms access.Permissions, version string, err error) {Blue}
	//  }
	permissions.ObjectServerMethods
	Get(*context.T, rpc.ServerCall) (greeting string, _ error)
	Set(_ *context.T, _ rpc.ServerCall, greeting string) error
}

// HelloServerStubMethods is the server interface containing
// Hello methods, as expected by rpc.Server.
// There is no difference between this interface and HelloServerMethods
// since there are no streaming methods.
type HelloServerStubMethods HelloServerMethods

// HelloServerStub adds universal methods to HelloServerStubMethods.
type HelloServerStub interface {
	HelloServerStubMethods
	// Describe the Hello interfaces.
	Describe__() []rpc.InterfaceDesc
}

// HelloServer returns a server stub for Hello.
// It converts an implementation of HelloServerMethods into
// an object that may be used by rpc.Server.
func HelloServer(impl HelloServerMethods) HelloServerStub {
	stub := implHelloServerStub{
		impl:             impl,
		ObjectServerStub: permissions.ObjectServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implHelloServerStub struct {
	impl HelloServerMethods
	permissions.ObjectServerStub
	gs *rpc.GlobState
}

func (s implHelloServerStub) Get(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.Get(ctx, call)
}

func (s implHelloServerStub) Set(ctx *context.T, call rpc.ServerCall, i0 string) error {
	return s.impl.Set(ctx, call, i0)
}

func (s implHelloServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implHelloServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{HelloDesc, permissions.ObjectDesc}
}

// HelloDesc describes the Hello interface.
var HelloDesc rpc.InterfaceDesc = descHello

// descHello hides the desc to keep godoc clean.
var descHello = rpc.InterfaceDesc{
	Name:    "Hello",
	PkgPath: "github.com/razvanm/vanadium-core/hello",
	Embeds: []rpc.EmbedDesc{
		{"Object", "v.io/v23/services/permissions", "// Object provides access control for Vanadium objects.\n//\n// Vanadium services implementing dynamic access control would typically embed\n// this interface and tag additional methods defined by the service with one of\n// Admin, Read, Write, Resolve etc. For example, the VDL definition of the\n// object would be:\n//\n//   package mypackage\n//\n//   import \"v.io/v23/security/access\"\n//   import \"v.io/v23/services/permissions\"\n//\n//   type MyObject interface {\n//     permissions.Object\n//     MyRead() (string, error) {access.Read}\n//     MyWrite(string) error    {access.Write}\n//   }\n//\n// If the set of pre-defined tags is insufficient, services may define their\n// own tag type and annotate all methods with this new type.\n//\n// Instead of embedding this Object interface, define SetPermissions and\n// GetPermissions in their own interface. Authorization policies will typically\n// respect annotations of a single type. For example, the VDL definition of an\n// object would be:\n//\n//  package mypackage\n//\n//  import \"v.io/v23/security/access\"\n//\n//  type MyTag string\n//\n//  const (\n//    Blue = MyTag(\"Blue\")\n//    Red  = MyTag(\"Red\")\n//  )\n//\n//  type MyObject interface {\n//    MyMethod() (string, error) {Blue}\n//\n//    // Allow clients to change access via the access.Object interface:\n//    SetPermissions(perms access.Permissions, version string) error         {Red}\n//    GetPermissions() (perms access.Permissions, version string, err error) {Blue}\n//  }"},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "Get",
			OutArgs: []rpc.ArgDesc{
				{"greeting", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "Set",
			InArgs: []rpc.ArgDesc{
				{"greeting", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
	},
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	return struct{}{}
}