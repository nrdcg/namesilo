package namesilo

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func toCleanString(data []byte) string {
	return strings.TrimSuffix(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func TestAddAccountFunds(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/account/addAccountFunds.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &AddAccountFunds{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestAddAutoRenewal(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/addAutoRenewal.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &AddAutoRenewal{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestAddPrivacy(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/privacy/addPrivacy.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &AddPrivacy{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestAddRegisteredNameServer(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/nameserver/addRegisteredNameServer.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &AddRegisteredNameServer{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestChangeNameServers(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/nameserver/changeNameServers.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ChangeNameServers{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestCheckRegisterAvailability(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/checkRegisterAvailability.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &CheckRegisterAvailability{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestCheckTransferAvailability(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/checkTransferAvailability.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &CheckTransferAvailability{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestCheckTransferStatus(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/transfers/checkTransferStatus.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &CheckTransferStatus{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestConfigureEmailForward(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/email/configureEmailForward.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ConfigureEmailForward{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestContactAdd(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/contact/contactAdd.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ContactAdd{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestContactDelete(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/contact/contactDelete.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ContactDelete{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestContactDomainAssociate(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/contact/contactDomainAssociate.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ContactDomainAssociate{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestContactList(t *testing.T) {
	t.Skip("because <usnc/> become <usnc></usnc>")

	bytes, err := os.ReadFile(filepath.FromSlash("./samples/contact/contactList.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ContactList{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestContactUpdate(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/contact/contactUpdate.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ContactUpdate{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDeleteEmailForward(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/email/deleteEmailForward.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DeleteEmailForward{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDeleteRegisteredNameServer(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/nameserver/deleteRegisteredNameServer.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DeleteRegisteredNameServer{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDnsAddRecord(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/dns/dnsAddRecord.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DnsAddRecord{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDnsDeleteRecord(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/dns/dnsDeleteRecord.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DnsDeleteRecord{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDnsListRecords(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/dns/dnsListRecords.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DnsListRecords{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDnsSecAddRecord(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/dns/dnsSecAddRecord.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DnsSecAddRecord{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDnsSecDeleteRecord(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/dns/dnsSecDeleteRecord.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DnsSecDeleteRecord{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDnsSecListRecords(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/dns/dnsSecListRecords.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DnsSecListRecords{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDnsUpdateRecord(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/dns/dnsUpdateRecord.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DnsUpdateRecord{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDomainForward(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/domainForward.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DomainForward{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDomainForwardSubDomain(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/domainForwardSubDomain.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DomainForwardSubDomain{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDomainForwardSubDomainDelete(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/domainForwardSubDomainDelete.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DomainForwardSubDomainDelete{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDomainLock(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/domainLock.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DomainLock{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestDomainUnlock(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/domainUnlock.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &DomainUnlock{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestEmailVerification(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/email/emailVerification.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &EmailVerification{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestGetAccountBalance(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/account/getAccountBalance.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &GetAccountBalance{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestGetDomainInfo(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/getDomainInfo.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &GetDomainInfo{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestGetPrices(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/getPrices.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &GetPrices{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestListDomains(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/listDomains.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ListDomains{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestListEmailForwards(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/email/listEmailForwards.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ListEmailForwards{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestListOrders(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/account/listOrders.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ListOrders{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestListRegisteredNameServers(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/nameserver/listRegisteredNameServers.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ListRegisteredNameServers{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestMarketplaceActiveSalesOverview(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/marketplace/marketplaceActiveSalesOverview.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &MarketplaceActiveSalesOverview{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestMarketplaceAddOrModifySale(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/marketplace/marketplaceAddOrModifySale.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &MarketplaceAddOrModifySale{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestMarketplaceLandingPageUpdate(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/marketplace/marketplaceLandingPageUpdate.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &MarketplaceLandingPageUpdate{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestModifyRegisteredNameServer(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/nameserver/modifyRegisteredNameServer.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &ModifyRegisteredNameServer{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestOrderDetails(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/account/orderDetails.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &OrderDetails{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestPortfolioAdd(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/portfolio/portfolioAdd.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &PortfolioAdd{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestPortfolioDelete(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/portfolio/portfolioDelete.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &PortfolioDelete{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestPortfolioDomainAssociate(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/portfolio/portfolioDomainAssociate.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &PortfolioDomainAssociate{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestPortfolioList(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/portfolio/portfolioList.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &PortfolioList{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestRegisterDomain(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/registerDomain.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &RegisterDomain{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestRegisterDomainDrop(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/registerDomainDrop.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &RegisterDomainDrop{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestRegistrantVerificationStatus(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/email/registrantVerificationStatus.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &RegistrantVerificationStatus{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestRemoveAutoRenewal(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/removeAutoRenewal.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &RemoveAutoRenewal{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestRemovePrivacy(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/privacy/removePrivacy.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &RemovePrivacy{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestRenewDomain(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/domains/renewDomain.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &RenewDomain{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestRetrieveAuthCode(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/transfers/retrieveAuthCode.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &RetrieveAuthCode{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestTransferDomain(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/transfers/transferDomain.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &TransferDomain{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestTransferUpdateChangeEPPCode(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/transfers/transferUpdateChangeEPPCode.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &TransferUpdateChangeEPPCode{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestTransferUpdateResendAdminEmail(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/transfers/transferUpdateResendAdminEmail.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &TransferUpdateResendAdminEmail{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}

func TestTransferUpdateResubmitToRegistry(t *testing.T) {
	bytes, err := os.ReadFile(filepath.FromSlash("./samples/transfers/transferUpdateResubmitToRegistry.xml"))
	if err != nil {
		t.Fatal(err)
	}

	model := &TransferUpdateResubmitToRegistry{}
	err = xml.Unmarshal(bytes, model)
	if err != nil {
		t.Fatal(err)
	}

	indent, err := xml.MarshalIndent(model, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	if toCleanString(indent) != toCleanString(bytes) {
		t.Logf("Got:\n%s\n\nWant:\n%s\n", string(indent), string(bytes))
		t.Error("Errors")
	}
}
