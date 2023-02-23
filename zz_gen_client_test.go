package namesilo

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
				return
			}
		}

		f, err := os.Open(filepath.Clean(filepath.Join(".", "samples", operation+".xml")))
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
	_, serverURL, teardown := setupFakeAPI("addAccountFunds")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &AddAccountFundsParams{}

	result, err := client.AddAccountFunds(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddAccountFunds{}, result)
}

func TestClient_AddAutoRenewal(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("addAutoRenewal")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &AddAutoRenewalParams{}

	result, err := client.AddAutoRenewal(params)
	require.Error(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddAutoRenewal{}, result)
}

func TestClient_AddPrivacy(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("addPrivacy")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &AddPrivacyParams{}

	result, err := client.AddPrivacy(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddPrivacy{}, result)
}

func TestClient_AddRegisteredNameServer(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("addRegisteredNameServer")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &AddRegisteredNameServerParams{}

	result, err := client.AddRegisteredNameServer(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &AddRegisteredNameServer{}, result)
}

func TestClient_ChangeNameServers(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("changeNameServers")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ChangeNameServersParams{}

	result, err := client.ChangeNameServers(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ChangeNameServers{}, result)
}

func TestClient_CheckRegisterAvailability(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("checkRegisterAvailability")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &CheckRegisterAvailabilityParams{}

	result, err := client.CheckRegisterAvailability(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckRegisterAvailability{}, result)
}

func TestClient_CheckTransferAvailability(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("checkTransferAvailability")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &CheckTransferAvailabilityParams{}

	result, err := client.CheckTransferAvailability(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckTransferAvailability{}, result)
}

func TestClient_CheckTransferStatus(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("checkTransferStatus")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &CheckTransferStatusParams{}

	result, err := client.CheckTransferStatus(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &CheckTransferStatus{}, result)
}

func TestClient_ConfigureEmailForward(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("configureEmailForward")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ConfigureEmailForwardParams{}

	result, err := client.ConfigureEmailForward(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ConfigureEmailForward{}, result)
}

func TestClient_ContactAdd(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contactAdd")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactAddParams{}

	result, err := client.ContactAdd(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactAdd{}, result)
}

func TestClient_ContactDelete(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contactDelete")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactDeleteParams{}

	result, err := client.ContactDelete(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactDelete{}, result)
}

func TestClient_ContactDomainAssociate(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contactDomainAssociate")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactDomainAssociateParams{}

	result, err := client.ContactDomainAssociate(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactDomainAssociate{}, result)
}

func TestClient_ContactList(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contactList")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactListParams{}

	result, err := client.ContactList(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactList{}, result)
}

func TestClient_ContactUpdate(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("contactUpdate")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ContactUpdateParams{}

	result, err := client.ContactUpdate(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ContactUpdate{}, result)
}

func TestClient_DeleteEmailForward(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("deleteEmailForward")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DeleteEmailForwardParams{}

	result, err := client.DeleteEmailForward(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DeleteEmailForward{}, result)
}

func TestClient_DeleteRegisteredNameServer(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("deleteRegisteredNameServer")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DeleteRegisteredNameServerParams{}

	result, err := client.DeleteRegisteredNameServer(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DeleteRegisteredNameServer{}, result)
}

func TestClient_DnsAddRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dnsAddRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsAddRecordParams{}

	result, err := client.DnsAddRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsAddRecord{}, result)
}

func TestClient_DnsDeleteRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dnsDeleteRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsDeleteRecordParams{}

	result, err := client.DnsDeleteRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsDeleteRecord{}, result)
}

func TestClient_DnsListRecords(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dnsListRecords")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsListRecordsParams{}

	result, err := client.DnsListRecords(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsListRecords{}, result)
}

func TestClient_DnsSecAddRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dnsSecAddRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsSecAddRecordParams{}

	result, err := client.DnsSecAddRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecAddRecord{}, result)
}

func TestClient_DnsSecDeleteRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dnsSecDeleteRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsSecDeleteRecordParams{}

	result, err := client.DnsSecDeleteRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecDeleteRecord{}, result)
}

func TestClient_DnsSecListRecords(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dnsSecListRecords")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsSecListRecordsParams{}

	result, err := client.DnsSecListRecords(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsSecListRecords{}, result)
}

func TestClient_DnsUpdateRecord(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("dnsUpdateRecord")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DnsUpdateRecordParams{}

	result, err := client.DnsUpdateRecord(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DnsUpdateRecord{}, result)
}

func TestClient_DomainForward(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domainForward")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainForwardParams{}

	result, err := client.DomainForward(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForward{}, result)
}

func TestClient_DomainForwardSubDomain(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domainForwardSubDomain")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainForwardSubDomainParams{}

	result, err := client.DomainForwardSubDomain(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForwardSubDomain{}, result)
}

func TestClient_DomainForwardSubDomainDelete(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domainForwardSubDomainDelete")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainForwardSubDomainDeleteParams{}

	result, err := client.DomainForwardSubDomainDelete(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainForwardSubDomainDelete{}, result)
}

func TestClient_DomainLock(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domainLock")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainLockParams{}

	result, err := client.DomainLock(params)
	require.Error(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainLock{}, result)
}

func TestClient_DomainUnlock(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("domainUnlock")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &DomainUnlockParams{}

	result, err := client.DomainUnlock(params)
	require.Error(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &DomainUnlock{}, result)
}

func TestClient_EmailVerification(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("emailVerification")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &EmailVerificationParams{}

	result, err := client.EmailVerification(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &EmailVerification{}, result)
}

func TestClient_GetAccountBalance(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("getAccountBalance")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &GetAccountBalanceParams{}

	result, err := client.GetAccountBalance(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetAccountBalance{}, result)
}

func TestClient_GetDomainInfo(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("getDomainInfo")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &GetDomainInfoParams{}

	result, err := client.GetDomainInfo(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetDomainInfo{}, result)
}

func TestClient_GetPrices(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("getPrices")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &GetPricesParams{}

	result, err := client.GetPrices(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &GetPrices{}, result)
}

func TestClient_ListDomains(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("listDomains")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ListDomainsParams{}

	result, err := client.ListDomains(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListDomains{}, result)
}

func TestClient_ListEmailForwards(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("listEmailForwards")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ListEmailForwardsParams{}

	result, err := client.ListEmailForwards(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListEmailForwards{}, result)
}

func TestClient_ListOrders(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("listOrders")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ListOrdersParams{}

	result, err := client.ListOrders(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListOrders{}, result)
}

func TestClient_ListRegisteredNameServers(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("listRegisteredNameServers")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ListRegisteredNameServersParams{}

	result, err := client.ListRegisteredNameServers(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ListRegisteredNameServers{}, result)
}

func TestClient_MarketplaceActiveSalesOverview(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("marketplaceActiveSalesOverview")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &MarketplaceActiveSalesOverviewParams{}

	result, err := client.MarketplaceActiveSalesOverview(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceActiveSalesOverview{}, result)
}

func TestClient_MarketplaceAddOrModifySale(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("marketplaceAddOrModifySale")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &MarketplaceAddOrModifySaleParams{}

	result, err := client.MarketplaceAddOrModifySale(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceAddOrModifySale{}, result)
}

func TestClient_MarketplaceLandingPageUpdate(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("marketplaceLandingPageUpdate")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &MarketplaceLandingPageUpdateParams{}

	result, err := client.MarketplaceLandingPageUpdate(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &MarketplaceLandingPageUpdate{}, result)
}

func TestClient_ModifyRegisteredNameServer(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("modifyRegisteredNameServer")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &ModifyRegisteredNameServerParams{}

	result, err := client.ModifyRegisteredNameServer(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &ModifyRegisteredNameServer{}, result)
}

func TestClient_OrderDetails(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("orderDetails")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &OrderDetailsParams{}

	result, err := client.OrderDetails(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &OrderDetails{}, result)
}

func TestClient_PortfolioAdd(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("portfolioAdd")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &PortfolioAddParams{}

	result, err := client.PortfolioAdd(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioAdd{}, result)
}

func TestClient_PortfolioDelete(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("portfolioDelete")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &PortfolioDeleteParams{}

	result, err := client.PortfolioDelete(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioDelete{}, result)
}

func TestClient_PortfolioDomainAssociate(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("portfolioDomainAssociate")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &PortfolioDomainAssociateParams{}

	result, err := client.PortfolioDomainAssociate(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioDomainAssociate{}, result)
}

func TestClient_PortfolioList(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("portfolioList")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &PortfolioListParams{}

	result, err := client.PortfolioList(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &PortfolioList{}, result)
}

func TestClient_RegisterDomain(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("registerDomain")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RegisterDomainParams{}

	result, err := client.RegisterDomain(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegisterDomain{}, result)
}

func TestClient_RegisterDomainDrop(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("registerDomainDrop")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RegisterDomainDropParams{}

	result, err := client.RegisterDomainDrop(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegisterDomainDrop{}, result)
}

func TestClient_RegistrantVerificationStatus(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("registrantVerificationStatus")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RegistrantVerificationStatusParams{}

	result, err := client.RegistrantVerificationStatus(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RegistrantVerificationStatus{}, result)
}

func TestClient_RemoveAutoRenewal(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("removeAutoRenewal")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RemoveAutoRenewalParams{}

	result, err := client.RemoveAutoRenewal(params)
	require.Error(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RemoveAutoRenewal{}, result)
}

func TestClient_RemovePrivacy(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("removePrivacy")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RemovePrivacyParams{}

	result, err := client.RemovePrivacy(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RemovePrivacy{}, result)
}

func TestClient_RenewDomain(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("renewDomain")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RenewDomainParams{}

	result, err := client.RenewDomain(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RenewDomain{}, result)
}

func TestClient_RetrieveAuthCode(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("retrieveAuthCode")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &RetrieveAuthCodeParams{}

	result, err := client.RetrieveAuthCode(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &RetrieveAuthCode{}, result)
}

func TestClient_TransferDomain(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transferDomain")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &TransferDomainParams{}

	result, err := client.TransferDomain(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferDomain{}, result)
}

func TestClient_TransferUpdateChangeEPPCode(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transferUpdateChangeEPPCode")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &TransferUpdateChangeEPPCodeParams{}

	result, err := client.TransferUpdateChangeEPPCode(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateChangeEPPCode{}, result)
}

func TestClient_TransferUpdateResendAdminEmail(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transferUpdateResendAdminEmail")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &TransferUpdateResendAdminEmailParams{}

	result, err := client.TransferUpdateResendAdminEmail(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateResendAdminEmail{}, result)
}

func TestClient_TransferUpdateResubmitToRegistry(t *testing.T) {
	_, serverURL, teardown := setupFakeAPI("transferUpdateResubmitToRegistry")
	defer teardown()

	transport, err := NewTokenTransport("1234")
	require.NoError(t, err)

	client := NewClient(transport.Client())
	client.Endpoint = serverURL

	params := &TransferUpdateResubmitToRegistryParams{}

	result, err := client.TransferUpdateResubmitToRegistry(params)
	require.NoError(t, err)

	require.NotNil(t, result)

	assert.IsType(t, &TransferUpdateResubmitToRegistry{}, result)
}
