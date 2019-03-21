package namesilo

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_AddAccountFunds(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/addAccountFunds", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/addAccountFunds.xml")
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
	if err != nil {
		t.Fatal(err)
	}

	client := NewClient(transport.Client())
	client.Endpoint = server.URL

	params := &AddAccountFundsParams{
		Amount:    "1000000",
		PaymentID: "acbd",
	}

	funds, err := client.AddAccountFunds(params)
	if err != nil {
		t.Fatal(err)
	}

	if funds == nil {
		t.Error("empty response")
	}
}
