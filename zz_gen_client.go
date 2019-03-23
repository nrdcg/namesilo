package namesilo

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func (c *Client) AddAccountFunds(params *AddAccountFundsParams) (*AddAccountFunds, error) {
	resp, err := c.get("addAccountFunds", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &AddAccountFunds{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) AddAutoRenewal(params *AddAutoRenewalParams) (*AddAutoRenewal, error) {
	resp, err := c.get("addAutoRenewal", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &AddAutoRenewal{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) AddPrivacy(params *AddPrivacyParams) (*AddPrivacy, error) {
	resp, err := c.get("addPrivacy", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &AddPrivacy{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) AddRegisteredNameServer(params *AddRegisteredNameServerParams) (*AddRegisteredNameServer, error) {
	resp, err := c.get("addRegisteredNameServer", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &AddRegisteredNameServer{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ChangeNameServers(params *ChangeNameServersParams) (*ChangeNameServers, error) {
	resp, err := c.get("changeNameServers", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ChangeNameServers{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) CheckRegisterAvailability(params *CheckRegisterAvailabilityParams) (*CheckRegisterAvailability, error) {
	resp, err := c.get("checkRegisterAvailability", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &CheckRegisterAvailability{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) CheckTransferAvailability(params *CheckTransferAvailabilityParams) (*CheckTransferAvailability, error) {
	resp, err := c.get("checkTransferAvailability", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &CheckTransferAvailability{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) CheckTransferStatus(params *CheckTransferStatusParams) (*CheckTransferStatus, error) {
	resp, err := c.get("checkTransferStatus", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &CheckTransferStatus{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ConfigureEmailForward(params *ConfigureEmailForwardParams) (*ConfigureEmailForward, error) {
	resp, err := c.get("configureEmailForward", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ConfigureEmailForward{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ContactAdd(params *ContactAddParams) (*ContactAdd, error) {
	resp, err := c.get("contactAdd", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ContactAdd{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ContactDelete(params *ContactDeleteParams) (*ContactDelete, error) {
	resp, err := c.get("contactDelete", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ContactDelete{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ContactDomainAssociate(params *ContactDomainAssociateParams) (*ContactDomainAssociate, error) {
	resp, err := c.get("contactDomainAssociate", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ContactDomainAssociate{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ContactList(params *ContactListParams) (*ContactList, error) {
	resp, err := c.get("contactList", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ContactList{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ContactUpdate(params *ContactUpdateParams) (*ContactUpdate, error) {
	resp, err := c.get("contactUpdate", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ContactUpdate{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DeleteEmailForward(params *DeleteEmailForwardParams) (*DeleteEmailForward, error) {
	resp, err := c.get("deleteEmailForward", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DeleteEmailForward{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DeleteRegisteredNameServer(params *DeleteRegisteredNameServerParams) (*DeleteRegisteredNameServer, error) {
	resp, err := c.get("deleteRegisteredNameServer", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DeleteRegisteredNameServer{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DnsAddRecord(params *DnsAddRecordParams) (*DnsAddRecord, error) {
	resp, err := c.get("dnsAddRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DnsAddRecord{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DnsDeleteRecord(params *DnsDeleteRecordParams) (*DnsDeleteRecord, error) {
	resp, err := c.get("dnsDeleteRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DnsDeleteRecord{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DnsListRecords(params *DnsListRecordsParams) (*DnsListRecords, error) {
	resp, err := c.get("dnsListRecords", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DnsListRecords{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DnsSecAddRecord(params *DnsSecAddRecordParams) (*DnsSecAddRecord, error) {
	resp, err := c.get("dnsSecAddRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DnsSecAddRecord{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DnsSecDeleteRecord(params *DnsSecDeleteRecordParams) (*DnsSecDeleteRecord, error) {
	resp, err := c.get("dnsSecDeleteRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DnsSecDeleteRecord{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DnsSecListRecords(params *DnsSecListRecordsParams) (*DnsSecListRecords, error) {
	resp, err := c.get("dnsSecListRecords", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DnsSecListRecords{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DnsUpdateRecord(params *DnsUpdateRecordParams) (*DnsUpdateRecord, error) {
	resp, err := c.get("dnsUpdateRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DnsUpdateRecord{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DomainForward(params *DomainForwardParams) (*DomainForward, error) {
	resp, err := c.get("domainForward", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DomainForward{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DomainForwardSubDomain(params *DomainForwardSubDomainParams) (*DomainForwardSubDomain, error) {
	resp, err := c.get("domainForwardSubDomain", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DomainForwardSubDomain{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DomainForwardSubDomainDelete(params *DomainForwardSubDomainDeleteParams) (*DomainForwardSubDomainDelete, error) {
	resp, err := c.get("domainForwardSubDomainDelete", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DomainForwardSubDomainDelete{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DomainLock(params *DomainLockParams) (*DomainLock, error) {
	resp, err := c.get("domainLock", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DomainLock{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) DomainUnlock(params *DomainUnlockParams) (*DomainUnlock, error) {
	resp, err := c.get("domainUnlock", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &DomainUnlock{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) EmailVerification(params *EmailVerificationParams) (*EmailVerification, error) {
	resp, err := c.get("emailVerification", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &EmailVerification{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) GetAccountBalance(params *GetAccountBalanceParams) (*GetAccountBalance, error) {
	resp, err := c.get("getAccountBalance", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &GetAccountBalance{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) GetDomainInfo(params *GetDomainInfoParams) (*GetDomainInfo, error) {
	resp, err := c.get("getDomainInfo", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &GetDomainInfo{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) GetPrices(params *GetPricesParams) (*GetPrices, error) {
	resp, err := c.get("getPrices", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &GetPrices{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ListDomains(params *ListDomainsParams) (*ListDomains, error) {
	resp, err := c.get("listDomains", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ListDomains{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ListEmailForwards(params *ListEmailForwardsParams) (*ListEmailForwards, error) {
	resp, err := c.get("listEmailForwards", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ListEmailForwards{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ListOrders(params *ListOrdersParams) (*ListOrders, error) {
	resp, err := c.get("listOrders", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ListOrders{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ListRegisteredNameServers(params *ListRegisteredNameServersParams) (*ListRegisteredNameServers, error) {
	resp, err := c.get("listRegisteredNameServers", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ListRegisteredNameServers{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) MarketplaceActiveSalesOverview(params *MarketplaceActiveSalesOverviewParams) (*MarketplaceActiveSalesOverview, error) {
	resp, err := c.get("marketplaceActiveSalesOverview", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &MarketplaceActiveSalesOverview{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) MarketplaceAddOrModifySale(params *MarketplaceAddOrModifySaleParams) (*MarketplaceAddOrModifySale, error) {
	resp, err := c.get("marketplaceAddOrModifySale", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &MarketplaceAddOrModifySale{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) MarketplaceLandingPageUpdate(params *MarketplaceLandingPageUpdateParams) (*MarketplaceLandingPageUpdate, error) {
	resp, err := c.get("marketplaceLandingPageUpdate", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &MarketplaceLandingPageUpdate{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) ModifyRegisteredNameServer(params *ModifyRegisteredNameServerParams) (*ModifyRegisteredNameServer, error) {
	resp, err := c.get("modifyRegisteredNameServer", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &ModifyRegisteredNameServer{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) OrderDetails(params *OrderDetailsParams) (*OrderDetails, error) {
	resp, err := c.get("orderDetails", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &OrderDetails{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) PortfolioAdd(params *PortfolioAddParams) (*PortfolioAdd, error) {
	resp, err := c.get("portfolioAdd", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &PortfolioAdd{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) PortfolioDelete(params *PortfolioDeleteParams) (*PortfolioDelete, error) {
	resp, err := c.get("portfolioDelete", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &PortfolioDelete{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) PortfolioDomainAssociate(params *PortfolioDomainAssociateParams) (*PortfolioDomainAssociate, error) {
	resp, err := c.get("portfolioDomainAssociate", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &PortfolioDomainAssociate{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) PortfolioList(params *PortfolioListParams) (*PortfolioList, error) {
	resp, err := c.get("portfolioList", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &PortfolioList{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) RegisterDomain(params *RegisterDomainParams) (*RegisterDomain, error) {
	resp, err := c.get("registerDomain", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &RegisterDomain{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) RegisterDomainDrop(params *RegisterDomainDropParams) (*RegisterDomainDrop, error) {
	resp, err := c.get("registerDomainDrop", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &RegisterDomainDrop{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) RegistrantVerificationStatus(params *RegistrantVerificationStatusParams) (*RegistrantVerificationStatus, error) {
	resp, err := c.get("registrantVerificationStatus", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &RegistrantVerificationStatus{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) RemoveAutoRenewal(params *RemoveAutoRenewalParams) (*RemoveAutoRenewal, error) {
	resp, err := c.get("removeAutoRenewal", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &RemoveAutoRenewal{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) RemovePrivacy(params *RemovePrivacyParams) (*RemovePrivacy, error) {
	resp, err := c.get("removePrivacy", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &RemovePrivacy{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) RenewDomain(params *RenewDomainParams) (*RenewDomain, error) {
	resp, err := c.get("renewDomain", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &RenewDomain{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) RetrieveAuthCode(params *RetrieveAuthCodeParams) (*RetrieveAuthCode, error) {
	resp, err := c.get("retrieveAuthCode", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &RetrieveAuthCode{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) TransferDomain(params *TransferDomainParams) (*TransferDomain, error) {
	resp, err := c.get("transferDomain", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &TransferDomain{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) TransferUpdateChangeEPPCode(params *TransferUpdateChangeEPPCodeParams) (*TransferUpdateChangeEPPCode, error) {
	resp, err := c.get("transferUpdateChangeEPPCode", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &TransferUpdateChangeEPPCode{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) TransferUpdateResendAdminEmail(params *TransferUpdateResendAdminEmailParams) (*TransferUpdateResendAdminEmail, error) {
	resp, err := c.get("transferUpdateResendAdminEmail", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &TransferUpdateResendAdminEmail{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}

func (c *Client) TransferUpdateResubmitToRegistry(params *TransferUpdateResubmitToRegistryParams) (*TransferUpdateResubmitToRegistry, error) {
	resp, err := c.get("transferUpdateResubmitToRegistry", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	op := &TransferUpdateResubmitToRegistry{}

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
		return op, fmt.Errorf("code: %s, details: %s", op.Reply.Code, op.Reply.Detail)
	}
}
