// Command generate creates the entire math package.
// It is intended to be invoked via "go generate" (directive in "gen.go").
package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

const parentPkgName = "go.timothygu.me/math/v2"

const defaultDirPerms os.FileMode = 0777

type unexpandedPkg struct {
	packageName       string
	typeName          string
	unsignedTypeName  string
	camelTypeName     string
	signed            bool
	largerTypes       []string
	largerTypesNonNeg []string
}

type expandedPkg struct {
	ParentPackageName   string
	PackageName         string
	UnsignedPackageName string
	TypeName            string
	CamelTypeName       string
	UnsignedTypeName    string
	Signed              bool
	MaxTestTypes        []string
	MinTestTypes        []string
}

var pkgs = map[string]unexpandedPkg{
	"int64":  {"i64math", "int64", "uint64", "Int64", true, nil, []string{"uint64"}},
	"int32":  {"i32math", "int32", "uint32", "Int32", true, []string{"int", "int64"}, []string{"uint", "uint32"}},
	"int16":  {"i16math", "int16", "uint16", "Int16", true, []string{"int32"}, []string{"uint16"}},
	"int8":   {"i8math", "int8", "uint8", "Int8", true, []string{"int16"}, []string{"uint8"}},
	"int":    {"imath", "int", "uint", "Int", true, []string{"int64"}, []string{"uint", "uint64"}},
	"uint64": {"u64math", "uint64", "", "Uint64", false, nil, nil},
	"uint32": {"u32math", "uint32", "", "Uint32", false, []string{"int64"}, []string{"uint", "uint64"}},
	"uint16": {"u16math", "uint16", "", "Uint16", false, []string{"int", "int32"}, []string{"uint32"}},
	"uint8":  {"u8math", "uint8", "", "Uint8", false, []string{"int16"}, []string{"uint16"}},
	"uint":   {"umath", "uint", "", "Uint", false, nil, []string{"uint64"}},
}
var allTypes = []string{
	"int64", "int32", "int16", "int8", "int",
	"uint64", "uint32", "uint16", "uint8", "uint",
}

var expandedPkgs = map[string]expandedPkg{}

func init() {
	for t, pkg := range pkgs {
		unsignedPackageName := pkg.packageName
		unsignedTypeName := t
		if pkg.signed {
			unsignedPackageName = pkgs[pkg.unsignedTypeName].packageName
			unsignedTypeName = pkg.unsignedTypeName
		}
		expandedPkgs[t] = expandedPkg{
			ParentPackageName:   parentPkgName,
			PackageName:         pkg.packageName,
			UnsignedPackageName: unsignedPackageName,
			TypeName:            t,
			CamelTypeName:       pkg.camelTypeName,
			UnsignedTypeName:    unsignedTypeName,
			Signed:              pkg.signed,
			MaxTestTypes:        maxTestsTypes(t),
			MinTestTypes:        minTestsTypes(t),
		}
	}
}

func maxTestsTypes(typeName string) []string {
	var types []string
	seen := map[string]bool{}
	var dfs func(string)
	dfs = func(typeName string) {
		if seen[typeName] {
			return
		}
		seen[typeName] = true
		types = append(types, typeName)
		for _, t := range pkgs[typeName].largerTypes {
			dfs(t)
		}
		for _, t := range pkgs[typeName].largerTypesNonNeg {
			dfs(t)
		}
	}
	dfs(typeName)
	return types
}

func minTestsTypes(typeName string) []string {
	if !pkgs[typeName].signed {
		return allTypes
	}

	var types []string
	seen := map[string]bool{}
	var dfs func(string)
	dfs = func(typeName string) {
		if seen[typeName] {
			return
		}
		seen[typeName] = true
		types = append(types, typeName)
		for _, t := range pkgs[typeName].largerTypes {
			dfs(t)
		}
	}
	dfs(typeName)
	return types
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Panicln("No caller information")
	}
	dir := filepath.Dir(filename)

	tmpl := template.Must(template.ParseFiles(
		filepath.Join(dir, "const.go.tmpl"),
		filepath.Join(dir, "const_test.go.tmpl"),
		filepath.Join(dir, "math.go.tmpl"),
		filepath.Join(dir, "math_test.go.tmpl"),
		filepath.Join(dir, "multiple.go.tmpl"),
		filepath.Join(dir, "multiple_test.go.tmpl"),
	))

	for _, pkg := range expandedPkgs {
		var err error
		if err = os.RemoveAll(pkg.PackageName); err != nil {
			log.Panicln(err)
		}
		if err = os.MkdirAll(pkg.PackageName, defaultDirPerms); err != nil {
			log.Panicln(err)
		}
		defer func() {
			if err != nil {
				_ = os.RemoveAll(pkg.PackageName)
			}
		}()

		for tmplName, outFile := range map[string]string{
			"const.go.tmpl":         filepath.Join(pkg.PackageName, "const.gen.go"),
			"const_test.go.tmpl":    filepath.Join(pkg.PackageName, "const.gen_test.go"),
			"math.go.tmpl":          filepath.Join(pkg.PackageName, pkg.PackageName+".gen.go"),
			"math_test.go.tmpl":     filepath.Join(pkg.PackageName, pkg.PackageName+".gen_test.go"),
			"multiple.go.tmpl":      filepath.Join(pkg.PackageName, "multiple.gen.go"),
			"multiple_test.go.tmpl": filepath.Join(pkg.PackageName, "multiple.gen_test.go"),
		} {
			log.Println("Generating", outFile)
			var buf bytes.Buffer
			if err = tmpl.ExecuteTemplate(&buf, tmplName, pkg); err != nil {
				log.Panicln(err)
			}
			var formatted []byte
			formatted, err = format.Source(buf.Bytes())
			if err != nil {
				log.Panicln(buf.String(), err)
			}

			log.Println("Writing", outFile)
			var f *os.File
			if f, err = os.Create(outFile); err != nil {
				log.Panicln(err)
			}
			defer func() {
				f.Close()
				if err != nil {
					_ = os.Remove(outFile)
				}
			}()
			if _, err = f.Write(formatted); err != nil {
				log.Panicln(err)
			}
		}
	}
}
