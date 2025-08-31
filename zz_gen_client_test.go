package namesilo

import (
	"context"
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

func setupFakeAPI(dir, operation string) (*http.ServeMux, string, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/"+operation, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		key := query.Get("key")
		if key != "1234" {
			err := xml.NewEncoder(w).Encode(Operation{Reply: Reply{Code: "110", Detail: "Invalid API Key"}})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		f, err := os.Open(filepath.Clean(filepath.Join(".", "samples", dir, operation+".xml")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = io.Copy(w, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return mux, server.URL, func() {
		server.Close()
	}
}

func TestClient_AddAccountFunds(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("account", "addAccountFunds")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &AddAccountFundsParams{}

	result, err := client.AddAccountFunds(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddAccountFunds{}, result)
}

func TestClient_AddAutoRenewal(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "addAutoRenewal")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &AddAutoRenewalParams{}

	result, err := client.AddAutoRenewal(context.TODO(), params)
	require.Error(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddAutoRenewal{}, result)
}

func TestClient_AddPrivacy(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("privacy", "addPrivacy")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &AddPrivacyParams{}

	result, err := client.AddPrivacy(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddPrivacy{}, result)
}

func TestClient_AddRegisteredNameServer(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("nameserver", "addRegisteredNameServer")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &AddRegisteredNameServerParams{}

	result, err := client.AddRegisteredNameServer(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddRegisteredNameServer{}, result)
}

func TestClient_ChangeNameServers(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("nameserver", "changeNameServers")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ChangeNameServersParams{}

	result, err := client.ChangeNameServers(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ChangeNameServers{}, result)
}

func TestClient_CheckRegisterAvailability(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "checkRegisterAvailability")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &CheckRegisterAvailabilityParams{}

	result, err := client.CheckRegisterAvailability(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckRegisterAvailability{}, result)
}

func TestClient_CheckTransferAvailability(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "checkTransferAvailability")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &CheckTransferAvailabilityParams{}

	result, err := client.CheckTransferAvailability(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckTransferAvailability{}, result)
}

func TestClient_CheckTransferStatus(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transfers", "checkTransferStatus")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &CheckTransferStatusParams{}

	result, err := client.CheckTransferStatus(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckTransferStatus{}, result)
}

func TestClient_ConfigureEmailForward(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("email", "configureEmailForward")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ConfigureEmailForwardParams{}

	result, err := client.ConfigureEmailForward(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ConfigureEmailForward{}, result)
}

func TestClient_ContactAdd(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contact", "contactAdd")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactAddParams{}

	result, err := client.ContactAdd(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactAdd{}, result)
}

func TestClient_ContactDelete(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contact", "contactDelete")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactDeleteParams{}

	result, err := client.ContactDelete(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactDelete{}, result)
}

func TestClient_ContactDomainAssociate(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contact", "contactDomainAssociate")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactDomainAssociateParams{}

	result, err := client.ContactDomainAssociate(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactDomainAssociate{}, result)
}

func TestClient_ContactList(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contact", "contactList")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactListParams{}

	result, err := client.ContactList(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactList{}, result)
}

func TestClient_ContactUpdate(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contact", "contactUpdate")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactUpdateParams{}

	result, err := client.ContactUpdate(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactUpdate{}, result)
}

func TestClient_DeleteEmailForward(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("email", "deleteEmailForward")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DeleteEmailForwardParams{}

	result, err := client.DeleteEmailForward(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DeleteEmailForward{}, result)
}

func TestClient_DeleteRegisteredNameServer(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("nameserver", "deleteRegisteredNameServer")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DeleteRegisteredNameServerParams{}

	result, err := client.DeleteRegisteredNameServer(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DeleteRegisteredNameServer{}, result)
}

func TestClient_DnsAddRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dns", "dnsAddRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsAddRecordParams{}

	result, err := client.DnsAddRecord(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsAddRecord{}, result)
}

func TestClient_DnsDeleteRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dns", "dnsDeleteRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsDeleteRecordParams{}

	result, err := client.DnsDeleteRecord(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsDeleteRecord{}, result)
}

func TestClient_DnsListRecords(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dns", "dnsListRecords")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsListRecordsParams{}

	result, err := client.DnsListRecords(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsListRecords{}, result)
}

func TestClient_DnsSecAddRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dns", "dnsSecAddRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsSecAddRecordParams{}

	result, err := client.DnsSecAddRecord(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecAddRecord{}, result)
}

func TestClient_DnsSecDeleteRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dns", "dnsSecDeleteRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsSecDeleteRecordParams{}

	result, err := client.DnsSecDeleteRecord(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecDeleteRecord{}, result)
}

func TestClient_DnsSecListRecords(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dns", "dnsSecListRecords")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsSecListRecordsParams{}

	result, err := client.DnsSecListRecords(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecListRecords{}, result)
}

func TestClient_DnsUpdateRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dns", "dnsUpdateRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsUpdateRecordParams{}

	result, err := client.DnsUpdateRecord(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsUpdateRecord{}, result)
}

func TestClient_DomainForward(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "domainForward")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainForwardParams{}

	result, err := client.DomainForward(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForward{}, result)
}

func TestClient_DomainForwardSubDomain(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "domainForwardSubDomain")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainForwardSubDomainParams{}

	result, err := client.DomainForwardSubDomain(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForwardSubDomain{}, result)
}

func TestClient_DomainForwardSubDomainDelete(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "domainForwardSubDomainDelete")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainForwardSubDomainDeleteParams{}

	result, err := client.DomainForwardSubDomainDelete(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForwardSubDomainDelete{}, result)
}

func TestClient_DomainLock(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "domainLock")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainLockParams{}

	result, err := client.DomainLock(context.TODO(), params)
	require.Error(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainLock{}, result)
}

func TestClient_DomainUnlock(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "domainUnlock")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainUnlockParams{}

	result, err := client.DomainUnlock(context.TODO(), params)
	require.Error(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainUnlock{}, result)
}

func TestClient_EmailVerification(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("email", "emailVerification")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &EmailVerificationParams{}

	result, err := client.EmailVerification(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &EmailVerification{}, result)
}

func TestClient_GetAccountBalance(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("account", "getAccountBalance")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &GetAccountBalanceParams{}

	result, err := client.GetAccountBalance(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetAccountBalance{}, result)
}

func TestClient_GetDomainInfo(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "getDomainInfo")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &GetDomainInfoParams{}

	result, err := client.GetDomainInfo(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetDomainInfo{}, result)
}

func TestClient_GetPrices(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "getPrices")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &GetPricesParams{}

	result, err := client.GetPrices(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetPrices{}, result)
}

func TestClient_ListDomains(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "listDomains")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ListDomainsParams{}

	result, err := client.ListDomains(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListDomains{}, result)
}

func TestClient_ListEmailForwards(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("email", "listEmailForwards")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ListEmailForwardsParams{}

	result, err := client.ListEmailForwards(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListEmailForwards{}, result)
}

func TestClient_ListOrders(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("account", "listOrders")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ListOrdersParams{}

	result, err := client.ListOrders(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListOrders{}, result)
}

func TestClient_ListRegisteredNameServers(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("nameserver", "listRegisteredNameServers")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ListRegisteredNameServersParams{}

	result, err := client.ListRegisteredNameServers(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListRegisteredNameServers{}, result)
}

func TestClient_MarketplaceActiveSalesOverview(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("marketplace", "marketplaceActiveSalesOverview")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &MarketplaceActiveSalesOverviewParams{}

	result, err := client.MarketplaceActiveSalesOverview(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceActiveSalesOverview{}, result)
}

func TestClient_MarketplaceAddOrModifySale(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("marketplace", "marketplaceAddOrModifySale")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &MarketplaceAddOrModifySaleParams{}

	result, err := client.MarketplaceAddOrModifySale(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceAddOrModifySale{}, result)
}

func TestClient_MarketplaceLandingPageUpdate(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("marketplace", "marketplaceLandingPageUpdate")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &MarketplaceLandingPageUpdateParams{}

	result, err := client.MarketplaceLandingPageUpdate(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceLandingPageUpdate{}, result)
}

func TestClient_ModifyRegisteredNameServer(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("nameserver", "modifyRegisteredNameServer")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ModifyRegisteredNameServerParams{}

	result, err := client.ModifyRegisteredNameServer(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ModifyRegisteredNameServer{}, result)
}

func TestClient_OrderDetails(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("account", "orderDetails")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &OrderDetailsParams{}

	result, err := client.OrderDetails(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &OrderDetails{}, result)
}

func TestClient_PortfolioAdd(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("portfolio", "portfolioAdd")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &PortfolioAddParams{}

	result, err := client.PortfolioAdd(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioAdd{}, result)
}

func TestClient_PortfolioDelete(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("portfolio", "portfolioDelete")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &PortfolioDeleteParams{}

	result, err := client.PortfolioDelete(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioDelete{}, result)
}

func TestClient_PortfolioDomainAssociate(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("portfolio", "portfolioDomainAssociate")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &PortfolioDomainAssociateParams{}

	result, err := client.PortfolioDomainAssociate(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioDomainAssociate{}, result)
}

func TestClient_PortfolioList(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("portfolio", "portfolioList")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &PortfolioListParams{}

	result, err := client.PortfolioList(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioList{}, result)
}

func TestClient_RegisterDomain(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "registerDomain")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RegisterDomainParams{}

	result, err := client.RegisterDomain(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegisterDomain{}, result)
}

func TestClient_RegisterDomainDrop(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "registerDomainDrop")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RegisterDomainDropParams{}

	result, err := client.RegisterDomainDrop(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegisterDomainDrop{}, result)
}

func TestClient_RegistrantVerificationStatus(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("email", "registrantVerificationStatus")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RegistrantVerificationStatusParams{}

	result, err := client.RegistrantVerificationStatus(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegistrantVerificationStatus{}, result)
}

func TestClient_RemoveAutoRenewal(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "removeAutoRenewal")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RemoveAutoRenewalParams{}

	result, err := client.RemoveAutoRenewal(context.TODO(), params)
	require.Error(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RemoveAutoRenewal{}, result)
}

func TestClient_RemovePrivacy(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("privacy", "removePrivacy")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RemovePrivacyParams{}

	result, err := client.RemovePrivacy(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RemovePrivacy{}, result)
}

func TestClient_RenewDomain(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domains", "renewDomain")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RenewDomainParams{}

	result, err := client.RenewDomain(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RenewDomain{}, result)
}

func TestClient_RetrieveAuthCode(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transfers", "retrieveAuthCode")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RetrieveAuthCodeParams{}

	result, err := client.RetrieveAuthCode(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RetrieveAuthCode{}, result)
}

func TestClient_TransferDomain(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transfers", "transferDomain")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &TransferDomainParams{}

	result, err := client.TransferDomain(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferDomain{}, result)
}

func TestClient_TransferUpdateChangeEPPCode(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transfers", "transferUpdateChangeEPPCode")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &TransferUpdateChangeEPPCodeParams{}

	result, err := client.TransferUpdateChangeEPPCode(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateChangeEPPCode{}, result)
}

func TestClient_TransferUpdateResendAdminEmail(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transfers", "transferUpdateResendAdminEmail")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &TransferUpdateResendAdminEmailParams{}

	result, err := client.TransferUpdateResendAdminEmail(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateResendAdminEmail{}, result)
}

func TestClient_TransferUpdateResubmitToRegistry(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transfers", "transferUpdateResubmitToRegistry")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &TransferUpdateResubmitToRegistryParams{}

	result, err := client.TransferUpdateResubmitToRegistry(context.TODO(), params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateResubmitToRegistry{}, result)
}
