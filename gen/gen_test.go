package gen

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"testing"
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
		return nil, fmt.Errorf("code: %%s, details: %%s", op.Reply.Code, op.Reply.Detail)
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
