package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
)

func TestGenerateClientMethods(t *testing.T) {
	t.Skip("generator")

	files, err := ioutil.ReadDir("../samples")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == "OPERATION.xml" {
			continue
		}

		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		fmt.Printf(`
func (c *Client) %[1]s(params *%[1]sParams) (*%[1]s, error) {
	resp, err := c.get("%[2]s", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %%v", resp.StatusCode)
	}

	op := &%[1]s{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(op)
	if err != nil {
		return nil, err
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
		return op, fmt.Errorf("code: %%s, details: %%s", op.Reply.Code, op.Reply.Detail)
	}
}
`, strings.ToUpper(string(baseName[0]))+baseName[1:], baseName)
	}
}

func TestGenerateModelTest(t *testing.T) {
	t.Skip("generator")

	files, err := ioutil.ReadDir("../samples")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == "OPERATION.xml" {
			continue
		}

		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		fmt.Printf(`
func Test%[1]s(t *testing.T) {
	bytes, err := ioutil.ReadFile("./samples/%[2]s.xml")
	if err != nil {
		t.Fatal(err)
	}

	model := &%[1]s{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprintln(string(indent)) != string(bytes) {
		t.Logf("Got:\n%%s\n\nWant:\n%%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}
`, strings.ToUpper(string(baseName[0]))+baseName[1:], baseName)
	}
}

func TestGenerateClientTest(t *testing.T) {
	t.Skip("generator")

	clientTestsTemplate := `
package namesilo

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

{{range $key, $value := .Names -}}
func TestClient_{{ $value.Upper }}(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/{{ $value.Lower }}", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/{{ $value.Lower }}.xml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		_, err = w.Write(bytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = server.URL

	params := &{{ $value.Upper }}Params{}

	result, err := client.{{ $value.Upper }}(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &{{ $value.Upper }}{}, result)
}
{{end}}
`

	type BaseName struct {
		Lower string
		Upper string
	}

	var baseNames []BaseName

	files, err := ioutil.ReadDir("../samples")
	if err != nil {
		log.Fatal(err)
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

	base := template.New("zz_gen_client_test.go")
	parse, err := base.Parse(clientTestsTemplate)

	b := &bytes.Buffer{}

	data := map[string]interface{}{
		"Names": baseNames,
	}

	err = parse.Execute(b, data)
	if err != nil {
		log.Fatal(err)
	}

	// gofmt
	source, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("../zz_gen_client_test.go", source, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGenerateBaseParams(t *testing.T) {
	t.Skip("generator")

	files, err := ioutil.ReadDir("../samples")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == "OPERATION.xml" {
			continue
		}

		baseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		fmt.Printf(`type %sParams struct {
		
		}
		
		`, strings.ToUpper(string(baseName[0]))+baseName[1:])
	}
}
