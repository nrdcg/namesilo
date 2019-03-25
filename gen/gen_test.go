package gen

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

func TestGenerateClientMethods(t *testing.T) {
	t.Skip("generator")

	clientMethodsTemplate := `package namesilo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)
{{range $key, $value := .Names }}
// {{ $value.Upper }} Execute operation {{ $value.Lower }}. 
func (c *Client) {{ $value.Upper }}(params *{{ $value.Upper }}Params) (*{{ $value.Upper }}, error) {
	resp, err := c.get("{{ $value.Lower }}", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &{{ $value.Upper }}{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %v: %s", err, bytes)
	}

	switch op.Reply.Code {
	case SuccessfulAPIOperation:
		// Successful API operation
		return op, nil
	case SuccessfulRegistration:
		// Successful registration, but not all provided hosts were valid resulting in our nameservers being used
		return op, nil
	case SuccessfulOrder:
		// Successful order, but there was an error with the contact information provided so your account default contact profile was used
		return op, nil
	default:
		// error
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}
{{end}}
`

	err := generate(clientMethodsTemplate, "zz_gen_client.go")
	require.NoError(t, err)
}

func TestGenerateClientTest(t *testing.T) {
	t.Skip("generator - several exception")

	clientTestsTemplate := `package namesilo

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupFakeAPI(operation string) (*http.ServeMux, string, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/"+operation, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		f, err := os.Open(filepath.Join(".", "samples", operation+".xml"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		_, err = io.Copy(w, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return mux, server.URL, func() {
		server.Close()
	}
}
{{range $key, $value := .Names }}
func TestClient_{{ $value.Upper }}(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("{{ $value.Lower }}")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &{{ $value.Upper }}Params{}

	result, err := client.{{ $value.Upper }}(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &{{ $value.Upper }}{}, result)
}
{{end}}
`

	err := generate(clientTestsTemplate, "zz_gen_client_test.go")
	require.NoError(t, err)
}

func TestGenerateModelTest(t *testing.T) {
	t.Skip("generator - one exception")

	modelTestsTemplate := `package namesilo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)
{{range $key, $value := .Names }}
func Test{{ $value.Upper }}(t *testing.T) {
	bytes, err := ioutil.ReadFile("./samples/{{ $value.Lower }}.xml")
	if err != nil {
		t.Fatal(err)
	}

	model := &{{ $value.Upper }}{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprintln(string(indent)) != string(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}
{{end}}
`

	err := generate(modelTestsTemplate, "model_test.go")
	require.NoError(t, err)
}

func TestGenerateBaseParams(t *testing.T) {
	t.Skip("generator - only to initiate the structures")

	paramsTemplate := `package namesilo
{{range $key, $value := .Names }}
// {{ $value.Upper }}Params Parameters for operation {{ $value.Lower }}.
type {{ $value.Upper }}Params struct {
	// TODO
}
{{end}}
`

	err := generate(paramsTemplate, "zz_gen_params.go")
	require.NoError(t, err)
}

func generate(tmpl string, filename string) error {
	type BaseName struct {
		Lower string
		Upper string
	}

	var baseNames []BaseName

	files, err := ioutil.ReadDir("../samples")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == "OPERATION.xml" {
			continue
		}

		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		baseNames = append(baseNames, BaseName{
			Lower: baseName,
			Upper: strings.ToUpper(string(baseName[0])) + baseName[1:],
		})
	}

	data := map[string]interface{}{
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

	return ioutil.WriteFile(filepath.Join("..", filename), source, 0666)
}
