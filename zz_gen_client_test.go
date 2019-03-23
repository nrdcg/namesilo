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
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = server.URL

	params := &AddAccountFundsParams{}

	result, err := client.AddAccountFunds(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddAccountFunds{}, result)
}
func TestClient_AddAutoRenewal(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/addAutoRenewal", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/addAutoRenewal.xml")
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

	params := &AddAutoRenewalParams{}

	result, err := client.AddAutoRenewal(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddAutoRenewal{}, result)
}
func TestClient_AddPrivacy(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/addPrivacy", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/addPrivacy.xml")
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

	params := &AddPrivacyParams{}

	result, err := client.AddPrivacy(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddPrivacy{}, result)
}
func TestClient_AddRegisteredNameServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/addRegisteredNameServer", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/addRegisteredNameServer.xml")
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

	params := &AddRegisteredNameServerParams{}

	result, err := client.AddRegisteredNameServer(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddRegisteredNameServer{}, result)
}
func TestClient_ChangeNameServers(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/changeNameServers", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/changeNameServers.xml")
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

	params := &ChangeNameServersParams{}

	result, err := client.ChangeNameServers(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ChangeNameServers{}, result)
}
func TestClient_CheckRegisterAvailability(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkRegisterAvailability", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/checkRegisterAvailability.xml")
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

	params := &CheckRegisterAvailabilityParams{}

	result, err := client.CheckRegisterAvailability(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckRegisterAvailability{}, result)
}
func TestClient_CheckTransferAvailability(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkTransferAvailability", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/checkTransferAvailability.xml")
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

	params := &CheckTransferAvailabilityParams{}

	result, err := client.CheckTransferAvailability(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckTransferAvailability{}, result)
}
func TestClient_CheckTransferStatus(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkTransferStatus", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/checkTransferStatus.xml")
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

	params := &CheckTransferStatusParams{}

	result, err := client.CheckTransferStatus(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckTransferStatus{}, result)
}
func TestClient_ConfigureEmailForward(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/configureEmailForward", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/configureEmailForward.xml")
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

	params := &ConfigureEmailForwardParams{}

	result, err := client.ConfigureEmailForward(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ConfigureEmailForward{}, result)
}
func TestClient_ContactAdd(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/contactAdd", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/contactAdd.xml")
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

	params := &ContactAddParams{}

	result, err := client.ContactAdd(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactAdd{}, result)
}
func TestClient_ContactDelete(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/contactDelete", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/contactDelete.xml")
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

	params := &ContactDeleteParams{}

	result, err := client.ContactDelete(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactDelete{}, result)
}
func TestClient_ContactDomainAssociate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/contactDomainAssociate", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/contactDomainAssociate.xml")
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

	params := &ContactDomainAssociateParams{}

	result, err := client.ContactDomainAssociate(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactDomainAssociate{}, result)
}
func TestClient_ContactList(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/contactList", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/contactList.xml")
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

	params := &ContactListParams{}

	result, err := client.ContactList(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactList{}, result)
}
func TestClient_ContactUpdate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/contactUpdate", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/contactUpdate.xml")
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

	params := &ContactUpdateParams{}

	result, err := client.ContactUpdate(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactUpdate{}, result)
}
func TestClient_DeleteEmailForward(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/deleteEmailForward", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/deleteEmailForward.xml")
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

	params := &DeleteEmailForwardParams{}

	result, err := client.DeleteEmailForward(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DeleteEmailForward{}, result)
}
func TestClient_DeleteRegisteredNameServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/deleteRegisteredNameServer", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/deleteRegisteredNameServer.xml")
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

	params := &DeleteRegisteredNameServerParams{}

	result, err := client.DeleteRegisteredNameServer(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DeleteRegisteredNameServer{}, result)
}
func TestClient_DnsAddRecord(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/dnsAddRecord", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/dnsAddRecord.xml")
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

	params := &DnsAddRecordParams{}

	result, err := client.DnsAddRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsAddRecord{}, result)
}
func TestClient_DnsDeleteRecord(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/dnsDeleteRecord", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/dnsDeleteRecord.xml")
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

	params := &DnsDeleteRecordParams{}

	result, err := client.DnsDeleteRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsDeleteRecord{}, result)
}
func TestClient_DnsListRecords(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/dnsListRecords", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/dnsListRecords.xml")
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

	params := &DnsListRecordsParams{}

	result, err := client.DnsListRecords(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsListRecords{}, result)
}
func TestClient_DnsSecAddRecord(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/dnsSecAddRecord", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/dnsSecAddRecord.xml")
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

	params := &DnsSecAddRecordParams{}

	result, err := client.DnsSecAddRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecAddRecord{}, result)
}
func TestClient_DnsSecDeleteRecord(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/dnsSecDeleteRecord", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/dnsSecDeleteRecord.xml")
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

	params := &DnsSecDeleteRecordParams{}

	result, err := client.DnsSecDeleteRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecDeleteRecord{}, result)
}
func TestClient_DnsSecListRecords(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/dnsSecListRecords", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/dnsSecListRecords.xml")
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

	params := &DnsSecListRecordsParams{}

	result, err := client.DnsSecListRecords(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecListRecords{}, result)
}
func TestClient_DnsUpdateRecord(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/dnsUpdateRecord", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/dnsUpdateRecord.xml")
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

	params := &DnsUpdateRecordParams{}

	result, err := client.DnsUpdateRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsUpdateRecord{}, result)
}
func TestClient_DomainForward(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/domainForward", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/domainForward.xml")
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

	params := &DomainForwardParams{}

	result, err := client.DomainForward(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForward{}, result)
}
func TestClient_DomainForwardSubDomain(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/domainForwardSubDomain", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/domainForwardSubDomain.xml")
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

	params := &DomainForwardSubDomainParams{}

	result, err := client.DomainForwardSubDomain(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForwardSubDomain{}, result)
}
func TestClient_DomainForwardSubDomainDelete(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/domainForwardSubDomainDelete", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/domainForwardSubDomainDelete.xml")
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

	params := &DomainForwardSubDomainDeleteParams{}

	result, err := client.DomainForwardSubDomainDelete(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForwardSubDomainDelete{}, result)
}
func TestClient_DomainLock(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/domainLock", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/domainLock.xml")
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

	params := &DomainLockParams{}

	result, err := client.DomainLock(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainLock{}, result)
}
func TestClient_DomainUnlock(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/domainUnlock", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/domainUnlock.xml")
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

	params := &DomainUnlockParams{}

	result, err := client.DomainUnlock(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainUnlock{}, result)
}
func TestClient_EmailVerification(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/emailVerification", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/emailVerification.xml")
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

	params := &EmailVerificationParams{}

	result, err := client.EmailVerification(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &EmailVerification{}, result)
}
func TestClient_GetAccountBalance(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/getAccountBalance", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/getAccountBalance.xml")
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

	params := &GetAccountBalanceParams{}

	result, err := client.GetAccountBalance(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetAccountBalance{}, result)
}
func TestClient_GetDomainInfo(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/getDomainInfo", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/getDomainInfo.xml")
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

	params := &GetDomainInfoParams{}

	result, err := client.GetDomainInfo(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetDomainInfo{}, result)
}
func TestClient_GetPrices(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/getPrices", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/getPrices.xml")
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

	params := &GetPricesParams{}

	result, err := client.GetPrices(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetPrices{}, result)
}
func TestClient_ListDomains(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/listDomains", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/listDomains.xml")
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

	params := &ListDomainsParams{}

	result, err := client.ListDomains(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListDomains{}, result)
}
func TestClient_ListEmailForwards(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/listEmailForwards", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/listEmailForwards.xml")
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

	params := &ListEmailForwardsParams{}

	result, err := client.ListEmailForwards(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListEmailForwards{}, result)
}
func TestClient_ListOrders(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/listOrders", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/listOrders.xml")
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

	params := &ListOrdersParams{}

	result, err := client.ListOrders(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListOrders{}, result)
}
func TestClient_ListRegisteredNameServers(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/listRegisteredNameServers", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/listRegisteredNameServers.xml")
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

	params := &ListRegisteredNameServersParams{}

	result, err := client.ListRegisteredNameServers(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListRegisteredNameServers{}, result)
}
func TestClient_MarketplaceActiveSalesOverview(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/marketplaceActiveSalesOverview", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/marketplaceActiveSalesOverview.xml")
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

	params := &MarketplaceActiveSalesOverviewParams{}

	result, err := client.MarketplaceActiveSalesOverview(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceActiveSalesOverview{}, result)
}
func TestClient_MarketplaceAddOrModifySale(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/marketplaceAddOrModifySale", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/marketplaceAddOrModifySale.xml")
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

	params := &MarketplaceAddOrModifySaleParams{}

	result, err := client.MarketplaceAddOrModifySale(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceAddOrModifySale{}, result)
}
func TestClient_MarketplaceLandingPageUpdate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/marketplaceLandingPageUpdate", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/marketplaceLandingPageUpdate.xml")
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

	params := &MarketplaceLandingPageUpdateParams{}

	result, err := client.MarketplaceLandingPageUpdate(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceLandingPageUpdate{}, result)
}
func TestClient_ModifyRegisteredNameServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/modifyRegisteredNameServer", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/modifyRegisteredNameServer.xml")
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

	params := &ModifyRegisteredNameServerParams{}

	result, err := client.ModifyRegisteredNameServer(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ModifyRegisteredNameServer{}, result)
}
func TestClient_OrderDetails(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/orderDetails", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/orderDetails.xml")
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

	params := &OrderDetailsParams{}

	result, err := client.OrderDetails(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &OrderDetails{}, result)
}
func TestClient_PortfolioAdd(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/portfolioAdd", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/portfolioAdd.xml")
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

	params := &PortfolioAddParams{}

	result, err := client.PortfolioAdd(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioAdd{}, result)
}
func TestClient_PortfolioDelete(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/portfolioDelete", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/portfolioDelete.xml")
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

	params := &PortfolioDeleteParams{}

	result, err := client.PortfolioDelete(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioDelete{}, result)
}
func TestClient_PortfolioDomainAssociate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/portfolioDomainAssociate", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/portfolioDomainAssociate.xml")
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

	params := &PortfolioDomainAssociateParams{}

	result, err := client.PortfolioDomainAssociate(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioDomainAssociate{}, result)
}
func TestClient_PortfolioList(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/portfolioList", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/portfolioList.xml")
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

	params := &PortfolioListParams{}

	result, err := client.PortfolioList(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioList{}, result)
}
func TestClient_RegisterDomain(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/registerDomain", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/registerDomain.xml")
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

	params := &RegisterDomainParams{}

	result, err := client.RegisterDomain(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegisterDomain{}, result)
}
func TestClient_RegisterDomainDrop(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/registerDomainDrop", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/registerDomainDrop.xml")
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

	params := &RegisterDomainDropParams{}

	result, err := client.RegisterDomainDrop(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegisterDomainDrop{}, result)
}
func TestClient_RegistrantVerificationStatus(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/registrantVerificationStatus", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/registrantVerificationStatus.xml")
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

	params := &RegistrantVerificationStatusParams{}

	result, err := client.RegistrantVerificationStatus(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegistrantVerificationStatus{}, result)
}
func TestClient_RemoveAutoRenewal(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/removeAutoRenewal", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/removeAutoRenewal.xml")
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

	params := &RemoveAutoRenewalParams{}

	result, err := client.RemoveAutoRenewal(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RemoveAutoRenewal{}, result)
}
func TestClient_RemovePrivacy(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/removePrivacy", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/removePrivacy.xml")
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

	params := &RemovePrivacyParams{}

	result, err := client.RemovePrivacy(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RemovePrivacy{}, result)
}
func TestClient_RenewDomain(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/renewDomain", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/renewDomain.xml")
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

	params := &RenewDomainParams{}

	result, err := client.RenewDomain(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RenewDomain{}, result)
}
func TestClient_RetrieveAuthCode(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/retrieveAuthCode", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/retrieveAuthCode.xml")
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

	params := &RetrieveAuthCodeParams{}

	result, err := client.RetrieveAuthCode(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RetrieveAuthCode{}, result)
}
func TestClient_TransferDomain(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/transferDomain", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/transferDomain.xml")
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

	params := &TransferDomainParams{}

	result, err := client.TransferDomain(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferDomain{}, result)
}
func TestClient_TransferUpdateChangeEPPCode(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/transferUpdateChangeEPPCode", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/transferUpdateChangeEPPCode.xml")
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

	params := &TransferUpdateChangeEPPCodeParams{}

	result, err := client.TransferUpdateChangeEPPCode(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateChangeEPPCode{}, result)
}
func TestClient_TransferUpdateResendAdminEmail(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/transferUpdateResendAdminEmail", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/transferUpdateResendAdminEmail.xml")
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

	params := &TransferUpdateResendAdminEmailParams{}

	result, err := client.TransferUpdateResendAdminEmail(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateResendAdminEmail{}, result)
}
func TestClient_TransferUpdateResubmitToRegistry(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/transferUpdateResubmitToRegistry", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		bytes, err := ioutil.ReadFile("./samples/transferUpdateResubmitToRegistry.xml")
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

	params := &TransferUpdateResubmitToRegistryParams{}

	result, err := client.TransferUpdateResubmitToRegistry(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateResubmitToRegistry{}, result)
}
