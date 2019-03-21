package namesilo

import "encoding/xml"

type Request struct {
	Operation string `xml:"operation"`
	IP        string `xml:"ip"`
}

type Reply struct {
	Code   string `xml:"code"`
	Detail string `xml:"detail"`
}

// Operation was generated 2019-03-20 19:35:05.
type Operation struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// AddAccountFunds was generated 2019-03-20 19:35:05.
type AddAccountFunds struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code       string `xml:"code"`
		Detail     string `xml:"detail"`
		NewBalance string `xml:"new_balance"`
	} `xml:"reply"`
}

// AddAutoRenewal was generated 2019-03-20 19:35:05.
type AddAutoRenewal struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// AddPrivacy was generated 2019-03-20 19:35:05.
type AddPrivacy struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// AddRegisteredNameServer was generated 2019-03-20 19:35:05.
type AddRegisteredNameServer struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// ChangeNameServers was generated 2019-03-20 19:35:05.
type ChangeNameServers struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// CheckRegisterAvailability was generated 2019-03-20 19:35:05.
type CheckRegisterAvailability struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code      string `xml:"code"`
		Detail    string `xml:"detail"`
		Available struct {
			Domain []string `xml:"domain"`
		} `xml:"available"`
		Unavailable struct {
			Domain string `xml:"domain"`
		} `xml:"unavailable"`
		Invalid struct {
			Domain string `xml:"domain"`
		} `xml:"invalid"`
	} `xml:"reply"`
}

// CheckTransferAvailability was generated 2019-03-20 19:35:05.
type CheckTransferAvailability struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code      string `xml:"code"`
		Detail    string `xml:"detail"`
		Available struct {
			Domain []string `xml:"domain"`
		} `xml:"available"`
		Unavailable struct {
			Domain []struct {
				Reason string `xml:"reason,attr"`
			} `xml:"domain"`
		} `xml:"unavailable"`
	} `xml:"reply"`
}

// CheckTransferStatus was generated 2019-03-20 19:35:05.
type CheckTransferStatus struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Date    string `xml:"date"`
		Status  string `xml:"status"`
		Message string `xml:"message"`
	} `xml:"reply"`
}

// ConfigureEmailForward was generated 2019-03-20 19:35:05.
type ConfigureEmailForward struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Message string `xml:"message"`
	} `xml:"reply"`
}

// ContactAdd was generated 2019-03-20 19:35:05.
type ContactAdd struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code      string `xml:"code"`
		Detail    string `xml:"detail"`
		ContactID string `xml:"contact_id"`
	} `xml:"reply"`
}

// ContactDomainAssociate was generated 2019-03-20 19:35:05.
type ContactDomainAssociate struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// ContactList was generated 2019-03-20 19:35:05.
type ContactList struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Contact []struct {
			ContactID      string `xml:"contact_id"`
			DefaultProfile string `xml:"default_profile"`
			Nickname       string `xml:"nickname"`
			Company        string `xml:"company"`
			FirstName      string `xml:"first_name"`
			LastName       string `xml:"last_name"`
			Address        string `xml:"address"`
			Address2       string `xml:"address2"`
			City           string `xml:"city"`
			State          string `xml:"state"`
			Zip            string `xml:"zip"`
			Country        string `xml:"country"`
			Email          string `xml:"email"`
			Phone          string `xml:"phone"`
			Fax            string `xml:"fax"`
			Usnc           string `xml:"usnc"`
			Usap           string `xml:"usap"`
			Calf           string `xml:"calf"`
			Caln           string `xml:"caln"`
			Caag           string `xml:"caag"`
			Cawd           string `xml:"cawd"`
		} `xml:"contact"`
	} `xml:"reply"`
}

// ContactUpdate was generated 2019-03-20 19:35:05.
type ContactUpdate struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// ContactDelete was generated 2019-03-20 19:35:05.
type ContactDelete struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// DeleteEmailForward was generated 2019-03-20 19:35:05.
type DeleteEmailForward struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Message string `xml:"message"`
	} `xml:"reply"`
}

// DeleteRegisteredNameServer was generated 2019-03-20 19:35:05.
type DeleteRegisteredNameServer struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// DnsAddRecord was generated 2019-03-20 19:35:05.
type DnsAddRecord struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code     string `xml:"code"`
		Detail   string `xml:"detail"`
		RecordID string `xml:"record_id"`
	} `xml:"reply"`
}

// DnsDeleteRecord was generated 2019-03-20 19:35:05.
type DnsDeleteRecord struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// DnsListRecords was generated 2019-03-20 19:35:05.
type DnsListRecords struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code           string `xml:"code"`
		Detail         string `xml:"detail"`
		ResourceRecord []struct {
			RecordID string `xml:"record_id"`
			Type     string `xml:"type"`
			Host     string `xml:"host"`
			Value    string `xml:"value"`
			TTL      string `xml:"ttl"`
			Distance string `xml:"distance"`
		} `xml:"resource_record"`
	} `xml:"reply"`
}

// DnsSecAddRecord was generated 2019-03-20 19:35:05.
type DnsSecAddRecord struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// DnsSecDeleteRecord was generated 2019-03-20 19:35:05.
type DnsSecDeleteRecord struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// DnsSecListRecords was generated 2019-03-20 19:35:05.
type DnsSecListRecords struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code     string `xml:"code"`
		Detail   string `xml:"detail"`
		DsRecord []struct {
			Digest     string `xml:"digest"`
			DigestType string `xml:"digest_type"`
			Algorithm  string `xml:"algorithm"`
			KeyTag     string `xml:"key_tag"`
		} `xml:"ds_record"`
	} `xml:"reply"`
}

// DnsUpdateRecord was generated 2019-03-20 19:35:05.
type DnsUpdateRecord struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code     string `xml:"code"`
		Detail   string `xml:"detail"`
		RecordID string `xml:"record_id"`
	} `xml:"reply"`
}

// DomainForwardSubDomainDelete was generated 2019-03-20 19:35:05.
type DomainForwardSubDomainDelete struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// DomainForwardSubDomain was generated 2019-03-20 19:35:05.
type DomainForwardSubDomain struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Message string `xml:"message"`
	} `xml:"reply"`
}

// DomainForward was generated 2019-03-20 19:35:05.
type DomainForward struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// DomainLock was generated 2019-03-20 19:35:05.
type DomainLock struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// DomainUnlock was generated 2019-03-20 19:35:05.
type DomainUnlock struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// EmailVerification was generated 2019-03-20 19:35:05.
type EmailVerification struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Message string `xml:"message"`
	} `xml:"reply"`
}

// GetAccountBalance was generated 2019-03-20 19:35:05.
type GetAccountBalance struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Balance string `xml:"balance"`
	} `xml:"reply"`
}

// GetDomainInfo was generated 2019-03-20 19:35:05.
type GetDomainInfo struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code                      string `xml:"code"`
		Detail                    string `xml:"detail"`
		Created                   string `xml:"created"`
		Expires                   string `xml:"expires"`
		Status                    string `xml:"status"`
		Locked                    string `xml:"locked"`
		Private                   string `xml:"private"`
		AutoRenew                 string `xml:"auto_renew"`
		TrafficType               string `xml:"traffic_type"`
		EmailVerificationRequired string `xml:"email_verification_required"`
		Portfolio                 string `xml:"portfolio"`
		ForwardURL                string `xml:"forward_url"`
		ForwardType               string `xml:"forward_type"`
		Nameservers               struct {
			Nameserver []struct {
				Position string `xml:"position,attr"`
			} `xml:"nameserver"`
		} `xml:"nameservers"`
		ContactIds struct {
			Registrant     string `xml:"registrant"`
			Administrative string `xml:"administrative"`
			Technical      string `xml:"technical"`
			Billing        string `xml:"billing"`
		} `xml:"contact_ids"`
	} `xml:"reply"`
}

// GetPrices was generated 2019-03-20 19:35:05.
type GetPrices struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code   string `xml:"code"`
		Detail string `xml:"detail"`
		Com    struct {
			Registration string `xml:"registration"`
			Transfer     string `xml:"transfer"`
			Renew        string `xml:"renew"`
		} `xml:"com"`
		Net struct {
			Registration string `xml:"registration"`
			Transfer     string `xml:"transfer"`
			Renew        string `xml:"renew"`
		} `xml:"net"`
	} `xml:"reply"`
}

// ListDomains was generated 2019-03-20 19:35:05.
type ListDomains struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Domains struct {
			Domain []string `xml:"domain"`
		} `xml:"domains"`
	} `xml:"reply"`
}

// ListEmailForwards was generated 2019-03-20 19:35:05.
type ListEmailForwards struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code      string `xml:"code"`
		Detail    string `xml:"detail"`
		Addresses []struct {
			Email      string   `xml:"email"`
			ForwardsTo []string `xml:"forwards_to"`
		} `xml:"addresses"`
	} `xml:"reply"`
}

// ListOrders was generated 2019-03-20 19:35:05.
type ListOrders struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code   string `xml:"code"`
		Detail string `xml:"detail"`
		Order  []struct {
			OrderNumber string `xml:"order_number"`
			OrderDate   string `xml:"order_date"`
			Method      string `xml:"method"`
			Total       string `xml:"total"`
		} `xml:"order"`
	} `xml:"reply"`
}

// ListRegisteredNameServers was generated 2019-03-20 19:35:05.
type ListRegisteredNameServers struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code   string `xml:"code"`
		Detail string `xml:"detail"`
		Hosts  []struct {
			Host string   `xml:"host"`
			IP   []string `xml:"ip"`
		} `xml:"hosts"`
	} `xml:"reply"`
}

// MarketplaceActiveSalesOverview was generated 2019-03-20 19:35:05.
type MarketplaceActiveSalesOverview struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code        string `xml:"code"`
		Detail      string `xml:"detail"`
		SaleDetails []struct {
			Domain           string `xml:"domain"`
			Status           string `xml:"status"`
			Reserve          string `xml:"reserve"`
			BuyNow           string `xml:"buy_now"`
			Portfolio        string `xml:"portfolio"`
			SaleType         string `xml:"sale_type"`
			PayPlanOffered   string `xml:"pay_plan_offered"`
			EndDate          string `xml:"end_date"`
			AutoExtendDays   string `xml:"auto_extend_days"`
			TimeRemaining    string `xml:"time_remaining"`
			Private          string `xml:"private"`
			ActiveBidOrOffer string `xml:"active_bid_or_offer"`
		} `xml:"sale_details"`
	} `xml:"reply"`
}

// MarketplaceAddOrModifySale was generated 2019-03-20 19:35:05.
type MarketplaceAddOrModifySale struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Message string `xml:"message"`
	} `xml:"reply"`
}

// MarketplaceLandingPageUpdate was generated 2019-03-20 19:35:05.
type MarketplaceLandingPageUpdate struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// ModifyRegisteredNameServer was generated 2019-03-20 19:35:05.
type ModifyRegisteredNameServer struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// OrderDetails was generated 2019-03-20 19:35:05.
type OrderDetails struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code         string `xml:"code"`
		Detail       string `xml:"detail"`
		OrderDate    string `xml:"order_date"`
		Method       string `xml:"method"`
		Total        string `xml:"total"`
		OrderDetails []struct {
			Description    string `xml:"description"`
			YearsQty       string `xml:"years_qty"`
			Price          string `xml:"price"`
			Subtotal       string `xml:"subtotal"`
			Status         string `xml:"status"`
			CreditedDate   string `xml:"credited_date"`
			CreditedAmount string `xml:"credited_amount"`
		} `xml:"order_details"`
	} `xml:"reply"`
}

// PortfolioAdd was generated 2019-03-20 19:35:05.
type PortfolioAdd struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// PortfolioDelete was generated 2019-03-20 19:35:05.
type PortfolioDelete struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// PortfolioDomainAssociate was generated 2019-03-20 19:35:05.
type PortfolioDomainAssociate struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code    string `xml:"code"`
		Detail  string `xml:"detail"`
		Message string `xml:"message"`
	} `xml:"reply"`
}

// PortfolioList was generated 2019-03-20 19:35:05.
type PortfolioList struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code       string `xml:"code"`
		Detail     string `xml:"detail"`
		Portfolios struct {
			Name string `xml:"name"`
			Code string `xml:"code"`
		} `xml:"portfolios"`
	} `xml:"reply"`
}

// RegisterDomainDrop was generated 2019-03-20 19:35:05.
type RegisterDomainDrop struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code        string `xml:"code"`
		Detail      string `xml:"detail"`
		Message     string `xml:"message"`
		Domain      string `xml:"domain"`
		OrderAmount string `xml:"order_amount"`
	} `xml:"reply"`
}

// RegisterDomain was generated 2019-03-20 19:35:05.
type RegisterDomain struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code        string `xml:"code"`
		Detail      string `xml:"detail"`
		Message     string `xml:"message"`
		Domain      string `xml:"domain"`
		OrderAmount string `xml:"order_amount"`
	} `xml:"reply"`
}

// RegistrantVerificationStatus was generated 2019-03-20 19:35:05.
type RegistrantVerificationStatus struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code   string `xml:"code"`
		Detail string `xml:"detail"`
		Email  []struct {
			EmailAddress string `xml:"email_address"`
			Domains      string `xml:"domains"`
			Verified     string `xml:"verified"`
		} `xml:"email"`
	} `xml:"reply"`
}

// RemoveAutoRenewal was generated 2019-03-20 19:35:05.
type RemoveAutoRenewal struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// RemovePrivacy was generated 2019-03-20 19:35:05.
type RemovePrivacy struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// RenewDomain was generated 2019-03-20 19:35:05.
type RenewDomain struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code        string `xml:"code"`
		Detail      string `xml:"detail"`
		Message     string `xml:"message"`
		Domain      string `xml:"domain"`
		OrderAmount string `xml:"order_amount"`
	} `xml:"reply"`
}

// RetrieveAuthCode was generated 2019-03-20 19:35:05.
type RetrieveAuthCode struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// TransferDomain was generated 2019-03-20 19:35:05.
type TransferDomain struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   struct {
		Code        string `xml:"code"`
		Detail      string `xml:"detail"`
		Message     string `xml:"message"`
		Domain      string `xml:"domain"`
		OrderAmount string `xml:"order_amount"`
	} `xml:"reply"`
}

// TransferUpdateChangeEPPCode was generated 2019-03-20 19:35:05.
type TransferUpdateChangeEPPCode struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// TransferUpdateResendAdminEmail was generated 2019-03-20 19:35:05.
type TransferUpdateResendAdminEmail struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}

// TransferUpdateResubmitToRegistry was generated 2019-03-20 19:35:05.
type TransferUpdateResubmitToRegistry struct {
	XMLName xml.Name `xml:"namesilo"`
	Request Request  `xml:"request"`
	Reply   Reply    `xml:"reply"`
}
