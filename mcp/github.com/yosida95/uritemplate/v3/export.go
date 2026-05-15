// export by github.com/goplus/ixgo/cmd/qexp

package uritemplate

import (
	q "github.com/yosida95/uritemplate/v3"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("github.com/yosida95/uritemplate/v3", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "uritemplate",
			Path: "github.com/yosida95/uritemplate/v3",
			Deps: map[string]string{
				"bytes":        "bytes",
				"fmt":          "fmt",
				"log":          "log",
				"regexp":       "regexp",
				"strconv":      "strconv",
				"strings":      "strings",
				"sync":         "sync",
				"unicode":      "unicode",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"CompareFlags": reflect.TypeOf((*q.CompareFlags)(nil)).Elem(),
				"Template":     reflect.TypeOf((*q.Template)(nil)).Elem(),
				"Value":        reflect.TypeOf((*q.Value)(nil)).Elem(),
				"ValueType":    reflect.TypeOf((*q.ValueType)(nil)).Elem(),
				"Values":       reflect.TypeOf((*q.Values)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Equals":  reflect.ValueOf(q.Equals),
				"KV":      reflect.ValueOf(q.KV),
				"List":    reflect.ValueOf(q.List),
				"MustNew": reflect.ValueOf(q.MustNew),
				"New":     reflect.ValueOf(q.New),
				"String":  reflect.ValueOf(q.String),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"CompareVarname": {Typ: reflect.TypeOf(q.CompareVarname), Value: constant.MakeInt64(int64(q.CompareVarname))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"ValueTypeKV":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ValueTypeKV))},
				"ValueTypeList":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ValueTypeList))},
				"ValueTypeString": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ValueTypeString))},
			},
		}
	})
}
