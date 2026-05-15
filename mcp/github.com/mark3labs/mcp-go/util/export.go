// export by github.com/goplus/ixgo/cmd/qexp

package util

import (
	q "github.com/mark3labs/mcp-go/util"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("github.com/mark3labs/mcp-go/util", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "util",
			Path: "github.com/mark3labs/mcp-go/util",
			Deps: map[string]string{
				"log": "log",
			},
			Interfaces: map[string]reflect.Type{
				"Logger": reflect.TypeOf((*q.Logger)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"DefaultLogger": reflect.ValueOf(q.DefaultLogger),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
