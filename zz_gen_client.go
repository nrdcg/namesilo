package namesilo

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// AddAccountFunds Execute operation addAccountFunds.
func (c *Client) AddAccountFunds(ctx context.Context, params *AddAccountFundsParams) (*AddAccountFunds, error) {
	resp, err := c.get(ctx, "addAccountFunds", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &AddAccountFunds{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// AddAutoRenewal Execute operation addAutoRenewal.
func (c *Client) AddAutoRenewal(ctx context.Context, params *AddAutoRenewalParams) (*AddAutoRenewal, error) {
	resp, err := c.get(ctx, "addAutoRenewal", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &AddAutoRenewal{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// AddPrivacy Execute operation addPrivacy.
func (c *Client) AddPrivacy(ctx context.Context, params *AddPrivacyParams) (*AddPrivacy, error) {
	resp, err := c.get(ctx, "addPrivacy", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &AddPrivacy{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// AddRegisteredNameServer Execute operation addRegisteredNameServer.
func (c *Client) AddRegisteredNameServer(ctx context.Context, params *AddRegisteredNameServerParams) (*AddRegisteredNameServer, error) {
	resp, err := c.get(ctx, "addRegisteredNameServer", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &AddRegisteredNameServer{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ChangeNameServers Execute operation changeNameServers.
func (c *Client) ChangeNameServers(ctx context.Context, params *ChangeNameServersParams) (*ChangeNameServers, error) {
	resp, err := c.get(ctx, "changeNameServers", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ChangeNameServers{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// CheckRegisterAvailability Execute operation checkRegisterAvailability.
func (c *Client) CheckRegisterAvailability(
	ctx context.Context,
	params *CheckRegisterAvailabilityParams,
) (*CheckRegisterAvailability, error) {
	resp, err := c.get(ctx, "checkRegisterAvailability", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &CheckRegisterAvailability{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// CheckTransferAvailability Execute operation checkTransferAvailability.
func (c *Client) CheckTransferAvailability(
	ctx context.Context,
	params *CheckTransferAvailabilityParams,
) (*CheckTransferAvailability, error) {
	resp, err := c.get(ctx, "checkTransferAvailability", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &CheckTransferAvailability{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// CheckTransferStatus Execute operation checkTransferStatus.
func (c *Client) CheckTransferStatus(ctx context.Context, params *CheckTransferStatusParams) (*CheckTransferStatus, error) {
	resp, err := c.get(ctx, "checkTransferStatus", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &CheckTransferStatus{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ConfigureEmailForward Execute operation configureEmailForward.
func (c *Client) ConfigureEmailForward(ctx context.Context, params *ConfigureEmailForwardParams) (*ConfigureEmailForward, error) {
	resp, err := c.get(ctx, "configureEmailForward", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ConfigureEmailForward{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ContactAdd Execute operation contactAdd.
func (c *Client) ContactAdd(ctx context.Context, params *ContactAddParams) (*ContactAdd, error) {
	resp, err := c.get(ctx, "contactAdd", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ContactAdd{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ContactDelete Execute operation contactDelete.
func (c *Client) ContactDelete(ctx context.Context, params *ContactDeleteParams) (*ContactDelete, error) {
	resp, err := c.get(ctx, "contactDelete", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ContactDelete{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ContactDomainAssociate Execute operation contactDomainAssociate.
func (c *Client) ContactDomainAssociate(ctx context.Context, params *ContactDomainAssociateParams) (*ContactDomainAssociate, error) {
	resp, err := c.get(ctx, "contactDomainAssociate", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ContactDomainAssociate{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ContactList Execute operation contactList.
func (c *Client) ContactList(ctx context.Context, params *ContactListParams) (*ContactList, error) {
	resp, err := c.get(ctx, "contactList", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ContactList{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ContactUpdate Execute operation contactUpdate.
func (c *Client) ContactUpdate(ctx context.Context, params *ContactUpdateParams) (*ContactUpdate, error) {
	resp, err := c.get(ctx, "contactUpdate", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ContactUpdate{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DeleteEmailForward Execute operation deleteEmailForward.
func (c *Client) DeleteEmailForward(ctx context.Context, params *DeleteEmailForwardParams) (*DeleteEmailForward, error) {
	resp, err := c.get(ctx, "deleteEmailForward", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DeleteEmailForward{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DeleteRegisteredNameServer Execute operation deleteRegisteredNameServer.
func (c *Client) DeleteRegisteredNameServer(
	ctx context.Context,
	params *DeleteRegisteredNameServerParams,
) (*DeleteRegisteredNameServer, error) {
	resp, err := c.get(ctx, "deleteRegisteredNameServer", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DeleteRegisteredNameServer{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DnsAddRecord Execute operation dnsAddRecord.
func (c *Client) DnsAddRecord(ctx context.Context, params *DnsAddRecordParams) (*DnsAddRecord, error) {
	resp, err := c.get(ctx, "dnsAddRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DnsAddRecord{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DnsDeleteRecord Execute operation dnsDeleteRecord.
func (c *Client) DnsDeleteRecord(ctx context.Context, params *DnsDeleteRecordParams) (*DnsDeleteRecord, error) {
	resp, err := c.get(ctx, "dnsDeleteRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DnsDeleteRecord{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DnsListRecords Execute operation dnsListRecords.
func (c *Client) DnsListRecords(ctx context.Context, params *DnsListRecordsParams) (*DnsListRecords, error) {
	resp, err := c.get(ctx, "dnsListRecords", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DnsListRecords{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DnsSecAddRecord Execute operation dnsSecAddRecord.
func (c *Client) DnsSecAddRecord(ctx context.Context, params *DnsSecAddRecordParams) (*DnsSecAddRecord, error) {
	resp, err := c.get(ctx, "dnsSecAddRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DnsSecAddRecord{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DnsSecDeleteRecord Execute operation dnsSecDeleteRecord.
func (c *Client) DnsSecDeleteRecord(ctx context.Context, params *DnsSecDeleteRecordParams) (*DnsSecDeleteRecord, error) {
	resp, err := c.get(ctx, "dnsSecDeleteRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DnsSecDeleteRecord{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DnsSecListRecords Execute operation dnsSecListRecords.
func (c *Client) DnsSecListRecords(ctx context.Context, params *DnsSecListRecordsParams) (*DnsSecListRecords, error) {
	resp, err := c.get(ctx, "dnsSecListRecords", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DnsSecListRecords{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DnsUpdateRecord Execute operation dnsUpdateRecord.
func (c *Client) DnsUpdateRecord(ctx context.Context, params *DnsUpdateRecordParams) (*DnsUpdateRecord, error) {
	resp, err := c.get(ctx, "dnsUpdateRecord", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DnsUpdateRecord{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DomainForward Execute operation domainForward.
func (c *Client) DomainForward(ctx context.Context, params *DomainForwardParams) (*DomainForward, error) {
	resp, err := c.get(ctx, "domainForward", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DomainForward{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DomainForwardSubDomain Execute operation domainForwardSubDomain.
func (c *Client) DomainForwardSubDomain(ctx context.Context, params *DomainForwardSubDomainParams) (*DomainForwardSubDomain, error) {
	resp, err := c.get(ctx, "domainForwardSubDomain", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DomainForwardSubDomain{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DomainForwardSubDomainDelete Execute operation domainForwardSubDomainDelete.
func (c *Client) DomainForwardSubDomainDelete(
	ctx context.Context,
	params *DomainForwardSubDomainDeleteParams,
) (*DomainForwardSubDomainDelete, error) {
	resp, err := c.get(ctx, "domainForwardSubDomainDelete", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DomainForwardSubDomainDelete{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DomainLock Execute operation domainLock.
func (c *Client) DomainLock(ctx context.Context, params *DomainLockParams) (*DomainLock, error) {
	resp, err := c.get(ctx, "domainLock", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DomainLock{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// DomainUnlock Execute operation domainUnlock.
func (c *Client) DomainUnlock(ctx context.Context, params *DomainUnlockParams) (*DomainUnlock, error) {
	resp, err := c.get(ctx, "domainUnlock", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &DomainUnlock{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// EmailVerification Execute operation emailVerification.
func (c *Client) EmailVerification(ctx context.Context, params *EmailVerificationParams) (*EmailVerification, error) {
	resp, err := c.get(ctx, "emailVerification", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &EmailVerification{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// GetAccountBalance Execute operation getAccountBalance.
func (c *Client) GetAccountBalance(ctx context.Context, params *GetAccountBalanceParams) (*GetAccountBalance, error) {
	resp, err := c.get(ctx, "getAccountBalance", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &GetAccountBalance{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// GetDomainInfo Execute operation getDomainInfo.
func (c *Client) GetDomainInfo(ctx context.Context, params *GetDomainInfoParams) (*GetDomainInfo, error) {
	resp, err := c.get(ctx, "getDomainInfo", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &GetDomainInfo{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// GetPrices Execute operation getPrices.
func (c *Client) GetPrices(ctx context.Context, params *GetPricesParams) (*GetPrices, error) {
	resp, err := c.get(ctx, "getPrices", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &GetPrices{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ListDomains Execute operation listDomains.
func (c *Client) ListDomains(ctx context.Context, params *ListDomainsParams) (*ListDomains, error) {
	resp, err := c.get(ctx, "listDomains", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ListDomains{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ListEmailForwards Execute operation listEmailForwards.
func (c *Client) ListEmailForwards(ctx context.Context, params *ListEmailForwardsParams) (*ListEmailForwards, error) {
	resp, err := c.get(ctx, "listEmailForwards", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ListEmailForwards{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ListOrders Execute operation listOrders.
func (c *Client) ListOrders(ctx context.Context, params *ListOrdersParams) (*ListOrders, error) {
	resp, err := c.get(ctx, "listOrders", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ListOrders{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ListRegisteredNameServers Execute operation listRegisteredNameServers.
func (c *Client) ListRegisteredNameServers(
	ctx context.Context,
	params *ListRegisteredNameServersParams,
) (*ListRegisteredNameServers, error) {
	resp, err := c.get(ctx, "listRegisteredNameServers", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ListRegisteredNameServers{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// MarketplaceActiveSalesOverview Execute operation marketplaceActiveSalesOverview.
func (c *Client) MarketplaceActiveSalesOverview(
	ctx context.Context,
	params *MarketplaceActiveSalesOverviewParams,
) (*MarketplaceActiveSalesOverview, error) {
	resp, err := c.get(ctx, "marketplaceActiveSalesOverview", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &MarketplaceActiveSalesOverview{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// MarketplaceAddOrModifySale Execute operation marketplaceAddOrModifySale.
func (c *Client) MarketplaceAddOrModifySale(
	ctx context.Context,
	params *MarketplaceAddOrModifySaleParams,
) (*MarketplaceAddOrModifySale, error) {
	resp, err := c.get(ctx, "marketplaceAddOrModifySale", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &MarketplaceAddOrModifySale{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// MarketplaceLandingPageUpdate Execute operation marketplaceLandingPageUpdate.
func (c *Client) MarketplaceLandingPageUpdate(
	ctx context.Context,
	params *MarketplaceLandingPageUpdateParams,
) (*MarketplaceLandingPageUpdate, error) {
	resp, err := c.get(ctx, "marketplaceLandingPageUpdate", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &MarketplaceLandingPageUpdate{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// ModifyRegisteredNameServer Execute operation modifyRegisteredNameServer.
func (c *Client) ModifyRegisteredNameServer(
	ctx context.Context,
	params *ModifyRegisteredNameServerParams,
) (*ModifyRegisteredNameServer, error) {
	resp, err := c.get(ctx, "modifyRegisteredNameServer", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &ModifyRegisteredNameServer{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// OrderDetails Execute operation orderDetails.
func (c *Client) OrderDetails(ctx context.Context, params *OrderDetailsParams) (*OrderDetails, error) {
	resp, err := c.get(ctx, "orderDetails", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &OrderDetails{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// PortfolioAdd Execute operation portfolioAdd.
func (c *Client) PortfolioAdd(ctx context.Context, params *PortfolioAddParams) (*PortfolioAdd, error) {
	resp, err := c.get(ctx, "portfolioAdd", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &PortfolioAdd{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// PortfolioDelete Execute operation portfolioDelete.
func (c *Client) PortfolioDelete(ctx context.Context, params *PortfolioDeleteParams) (*PortfolioDelete, error) {
	resp, err := c.get(ctx, "portfolioDelete", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &PortfolioDelete{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// PortfolioDomainAssociate Execute operation portfolioDomainAssociate.
func (c *Client) PortfolioDomainAssociate(
	ctx context.Context,
	params *PortfolioDomainAssociateParams,
) (*PortfolioDomainAssociate, error) {
	resp, err := c.get(ctx, "portfolioDomainAssociate", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &PortfolioDomainAssociate{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// PortfolioList Execute operation portfolioList.
func (c *Client) PortfolioList(ctx context.Context, params *PortfolioListParams) (*PortfolioList, error) {
	resp, err := c.get(ctx, "portfolioList", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &PortfolioList{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// RegisterDomain Execute operation registerDomain.
func (c *Client) RegisterDomain(ctx context.Context, params *RegisterDomainParams) (*RegisterDomain, error) {
	resp, err := c.get(ctx, "registerDomain", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &RegisterDomain{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// RegisterDomainDrop Execute operation registerDomainDrop.
func (c *Client) RegisterDomainDrop(ctx context.Context, params *RegisterDomainDropParams) (*RegisterDomainDrop, error) {
	resp, err := c.get(ctx, "registerDomainDrop", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &RegisterDomainDrop{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// RegistrantVerificationStatus Execute operation registrantVerificationStatus.
func (c *Client) RegistrantVerificationStatus(
	ctx context.Context,
	params *RegistrantVerificationStatusParams,
) (*RegistrantVerificationStatus, error) {
	resp, err := c.get(ctx, "registrantVerificationStatus", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &RegistrantVerificationStatus{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// RemoveAutoRenewal Execute operation removeAutoRenewal.
func (c *Client) RemoveAutoRenewal(ctx context.Context, params *RemoveAutoRenewalParams) (*RemoveAutoRenewal, error) {
	resp, err := c.get(ctx, "removeAutoRenewal", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &RemoveAutoRenewal{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// RemovePrivacy Execute operation removePrivacy.
func (c *Client) RemovePrivacy(ctx context.Context, params *RemovePrivacyParams) (*RemovePrivacy, error) {
	resp, err := c.get(ctx, "removePrivacy", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &RemovePrivacy{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// RenewDomain Execute operation renewDomain.
func (c *Client) RenewDomain(ctx context.Context, params *RenewDomainParams) (*RenewDomain, error) {
	resp, err := c.get(ctx, "renewDomain", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &RenewDomain{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// RetrieveAuthCode Execute operation retrieveAuthCode.
func (c *Client) RetrieveAuthCode(ctx context.Context, params *RetrieveAuthCodeParams) (*RetrieveAuthCode, error) {
	resp, err := c.get(ctx, "retrieveAuthCode", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &RetrieveAuthCode{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// TransferDomain Execute operation transferDomain.
func (c *Client) TransferDomain(ctx context.Context, params *TransferDomainParams) (*TransferDomain, error) {
	resp, err := c.get(ctx, "transferDomain", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &TransferDomain{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// TransferUpdateChangeEPPCode Execute operation transferUpdateChangeEPPCode.
func (c *Client) TransferUpdateChangeEPPCode(
	ctx context.Context,
	params *TransferUpdateChangeEPPCodeParams,
) (*TransferUpdateChangeEPPCode, error) {
	resp, err := c.get(ctx, "transferUpdateChangeEPPCode", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &TransferUpdateChangeEPPCode{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// TransferUpdateResendAdminEmail Execute operation transferUpdateResendAdminEmail.
func (c *Client) TransferUpdateResendAdminEmail(
	ctx context.Context,
	params *TransferUpdateResendAdminEmailParams,
) (*TransferUpdateResendAdminEmail, error) {
	resp, err := c.get(ctx, "transferUpdateResendAdminEmail", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &TransferUpdateResendAdminEmail{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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

// TransferUpdateResubmitToRegistry Execute operation transferUpdateResubmitToRegistry.
func (c *Client) TransferUpdateResubmitToRegistry(
	ctx context.Context,
	params *TransferUpdateResubmitToRegistryParams,
) (*TransferUpdateResubmitToRegistry, error) {
	resp, err := c.get(ctx, "transferUpdateResubmitToRegistry", params)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP status code %v", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	op := &TransferUpdateResubmitToRegistry{}
	err = xml.Unmarshal(bytes, op)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w: %s", err, bytes)
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
