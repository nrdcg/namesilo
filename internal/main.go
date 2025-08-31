package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"
)

//go:embed client-methods.go.tmpl
var clientMethodsTemplate string

//go:embed client-tests.go.tmpl
var clientTestsTemplate string

//go:embed model-tests.go.tmpl
var modelTestsTemplate string

func main() {
	err := generate(clientMethodsTemplate, "zz_gen_client.go")
	if err != nil {
		log.Fatal(err)
	}

	// IMPORTANT: several exceptions.
	err = generate(clientTestsTemplate, "zz_gen_client_test.go")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("zz_gen_client_test.go must be reviewed manually because of several exceptions.")

	err = generate(modelTestsTemplate, "zz_gen_types_test.go")
	if err != nil {
		log.Fatal(err)
	}
}

//nolint:unused // only to initiate the structures.
func generateBaseParams() error {
	paramsTemplate := `package namesilo
{{range $key, $value := .Names }}
// {{ $value.Upper }}Params Parameters for operation {{ $value.Lower }}.
type {{ $value.Upper }}Params struct {
	// TODO
}
{{end}}
`

	return generate(paramsTemplate, "zz_gen_params.go")
}

type BaseName struct {
	Lower string
	Upper string
	Dir   string
}

func generate(tmpl, filename string) error {
	var baseNames []BaseName

	err := filepath.WalkDir(filepath.FromSlash("samples"), func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if d.Name() == "OPERATION.xml" {
			return nil
		}

		if strings.Contains(d.Name(), "-") {
			return nil
		}

		if strings.Contains(d.Name(), "_") {
			return nil
		}

		baseName := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		baseNames = append(baseNames, BaseName{
			Lower: baseName,
			Upper: strings.ToUpper(string(baseName[0])) + baseName[1:],
			Dir:   filepath.Base(filepath.Dir(path)),
		})

		return nil
	})
	if err != nil {
		return err
	}

	slices.SortFunc(baseNames, func(a, b BaseName) int {
		return strings.Compare(a.Lower, b.Lower)
	})

	data := map[string]any{
		"Names": baseNames,
	}

	base := template.New(filename)

	parse, err := base.Parse(tmpl)
	if err != nil {
		return err
	}

	b := &bytes.Buffer{}

	err = parse.Execute(b, data)
	if err != nil {
		return err
	}

	// gofmt
	source, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	return os.WriteFile(filename, source, 0o666)
}
