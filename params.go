package namesilo

type AddAccountFundsParams struct {
	Amount    string `url:"amount"`
	PaymentID string `url:"payment_id"`
}

type AddAutoRenewalParams struct {
	Domain string `url:"domain"` // Required
}

type AddPrivacyParams struct {
	Domain string `url:"domain"` // Required
}

type AddRegisteredNameServerParams struct {
	NewHost string `url:"new_host"` // Required
	IP1     string `url:"ip1"`      // Required

	IP2  string `url:"ip2"`  // Optional
	IP3  string `url:"ip3"`  // Optional
	IP4  string `url:"ip4"`  // Optional
	IP5  string `url:"ip5"`  // Optional
	IP6  string `url:"ip6"`  // Optional
	IP7  string `url:"ip7"`  // Optional
	IP8  string `url:"ip8"`  // Optional
	IP9  string `url:"ip9"`  // Optional
	IP10 string `url:"ip10"` // Optional
	IP11 string `url:"ip11"` // Optional
	IP12 string `url:"ip12"` // Optional
	IP13 string `url:"ip13"` // Optional
}

type ChangeNameServersParams struct {
	Domains string `url:"domain"` // Required (A comma-delimited list of up to 200 domains)

	NameServer1 string `url:"ns1"` // Required
	NameServer2 string `url:"ns2"` // Required

	NameServer3  string `url:"ns3"`
	NameServer4  string `url:"ns4"`
	NameServer5  string `url:"ns5"`
	NameServer6  string `url:"ns6"`
	NameServer7  string `url:"ns7"`
	NameServer8  string `url:"ns8"`
	NameServer9  string `url:"ns9"`
	NameServer10 string `url:"ns10"`
	NameServer11 string `url:"ns11"`
	NameServer12 string `url:"ns12"`
	NameServer13 string `url:"ns13"`
}

type CheckRegisterAvailabilityParams struct {
	Domains string `url:"domains"` // Required (A comma-delimited list of domains to check)
}

type CheckTransferAvailabilityParams struct {
	Domains string `url:"domains"` // Required (A comma-delimited list of domains to check)
}

type CheckTransferStatusParams struct {
	Domain string `url:"domain"` // Required
}

type ConfigureEmailForwardParams struct {
	Domain   string `url:"domain"`   // Required
	Email    string `url:"email"`    // Required
	Forward1 string `url:"forward1"` // Required

	Forward2 string `url:"forward12"` // Optional
	Forward3 string `url:"forward13"` // Optional
	Forward4 string `url:"forward14"` // Optional
	Forward5 string `url:"forward15"` // Optional
}

type ContactAddParams struct {
	FirstName                     string `url:"fn"` // Contact Information
	LastName                      string `url:"ln"` // Contact Information
	MailingAddress                string `url:"ad"` // Contact Information
	MailingCity                   string `url:"cy"` // Contact Information
	MailingStateProvinceTerritory string `url:"st"` // Contact Information
	MailingZipPostalCode          string `url:"zp"` // Contact Information
	MailingCountry                string `url:"ct"` // Contact Information
	EmailAddress                  string `url:"em"` // Contact Information
	PhoneNumber                   string `url:"ph"` // Contact Information

	Company         string `url:"cp"`  // Contact Information
	MailingAddress2 string `url:"ad2"` // Contact Information
	Fax             string `url:"fx"`  // Contact Information

	USNexusCategory      string `url:"usnc"` // Contact Information
	USApplicationPurpose string `url:"usap"` // Contact Information

	CIRALegalForm        string `url:"calf"` // CIRA
	CIRALanguage         string `url:"caln"` // CIRA
	CIRAAgreementVersion string `url:"caag"` // CIRA
	CIRAWHOISDisplay     string `url:"cawd"` // CIRA
}

type ContactDeleteParams struct {
	ContactID string `url:"contact_id"`
}

type ContactDomainAssociateParams struct {
	Domain string `url:"domain"` // Required

	Registrant     string `url:"registrant"`     // Optional
	Administrative string `url:"administrative"` // Optional
	Billing        string `url:"billing"`        // Optional
	Technical      string `url:"technical"`      // Optional

	ContactID string `url:"contact_id"` // Contact ID
}

type ContactListParams struct {
	ContactID string `url:"contact_id"` // Optional
}

type ContactUpdateParams struct {
	FirstName                     string `url:"fn"` // Contact Information
	LastName                      string `url:"ln"` // Contact Information
	MailingAddress                string `url:"ad"` // Contact Information
	MailingCity                   string `url:"cy"` // Contact Information
	MailingStateProvinceTerritory string `url:"st"` // Contact Information
	MailingZipPostalCode          string `url:"zp"` // Contact Information
	MailingCountry                string `url:"ct"` // Contact Information
	EmailAddress                  string `url:"em"` // Contact Information
	PhoneNumber                   string `url:"ph"` // Contact Information

	Company         string `url:"cp"`  // Contact Information
	MailingAddress2 string `url:"ad2"` // Contact Information
	Fax             string `url:"fx"`  // Contact Information

	USNexusCategory      string `url:"usnc"` // Contact Information
	USApplicationPurpose string `url:"usap"` // Contact Information

	CIRALegalForm        string `url:"calf"` // CIRA
	CIRALanguage         string `url:"caln"` // CIRA
	CIRAAgreementVersion string `url:"caag"` // CIRA
	CIRAWHOISDisplay     string `url:"cawd"` // CIRA
}

type DeleteEmailForwardParams struct {
	Domain string `url:"domain"` // Required
	Email  string `url:"email"`  // Required
}

type DeleteRegisteredNameServerParams struct {
	CurrentHost string `url:"current_host"` // Required
}

type DnsAddRecordParams struct {
	Domain string `url:"domain"` // Required

	Type     string `url:"rrtype"` // Possible values are "A", "AAAA", "CNAME", "MX" and "TXT"
	Host     string `url:"rrhost"`
	Value    string `url:"rrvalue"`
	Distance int    `url:"rrdistance"`
	TTL      int    `url:"rrttl"`
}

type DnsDeleteRecordParams struct {
	Domain string `url:"domain"` // Required

	ID string `url:"rrid"`
}

type DnsListRecordsParams struct {
	Domain string `url:"domain"` // Required
}

type DnsSecAddRecordParams struct {
	Domain string `url:"domain"` // Required

	Digest     string `url:"digest"`
	KeyTag     string `url:"keyTag"`
	DigestType string `url:"digestType"`
	Alg        string `url:"alg"`
}

type DnsSecDeleteRecordParams struct {
	Domain string `url:"domain"` // Required

	Digest     string `url:"digest"`
	KeyTag     string `url:"keyTag"`
	DigestType string `url:"digestType"`
	Alg        string `url:"alg"`
}

type DnsSecListRecordsParams struct {
	Domain string `url:"domain"` // Required
}

type DnsUpdateRecordParams struct {
	Domain string `url:"domain"` // Required

	ID       string `url:"rrid"`
	Host     string `url:"rrhost"`
	Value    string `url:"rrvalue"`
	Distance int    `url:"rrdistance"`
	TTL      int    `url:"rrttl"`
}

type DomainForwardParams struct {
	Domain   string `url:"domain"`   // Required
	Protocol string `url:"protocol"` // Required
	Address  string `url:"address"`  // Required
	Method   string `url:"method"`   // Required

	MetaTitle       string `url:"meta_title"`       // Optional
	MetaDescription string `url:"meta_description"` // Optional
	MetaKeywords    string `url:"meta_keywords"`    // Optional

}

type DomainForwardSubDomainParams struct {
	Domain    string `url:"domain"`     // Required
	SubDomain string `url:"sub_domain"` // Required
	Protocol  string `url:"protocol"`   // Required
	Address   string `url:"address"`    // Required
	Method    string `url:"method"`     // Required

	MetaTitle       string `url:"meta_title"`       // Optional
	MetaDescription string `url:"meta_description"` // Optional
	MetaKeywords    string `url:"meta_keywords"`    // Optional
}

type DomainForwardSubDomainDeleteParams struct {
	Domain    string `url:"domain"`     // Required
	SubDomain string `url:"sub_domain"` // Required
}

type DomainLockParams struct {
	Domain string `url:"domain"` // Required
}

type DomainUnlockParams struct {
	Domain string `url:"domain"` // Required
}

type EmailVerificationParams struct {
	Email string `url:"email"` // Required
}

type GetAccountBalanceParams struct{}

type GetDomainInfoParams struct {
	Domain string `url:"domain"` // Required
}

type GetPricesParams struct {
	RetailPrices        string `url:"retail_prices"`        // Required
	RegistrationDomains string `url:"registration_domains"` // Required
}

type ListDomainsParams struct {
	Portfolio string `url:"portfolio"` // Optional
}

type ListEmailForwardsParams struct {
	Domain string `url:"domain"` // Required
}

type ListOrdersParams struct{}

type ListRegisteredNameServersParams struct{}

type MarketplaceActiveSalesOverviewParams struct{}

type MarketplaceAddOrModifySaleParams struct {
	Domain   string `url:"domain"`    // Required
	Action   string `url:"action"`    // Required
	SaleType string `url:"sale_type"` // Required

	Reserve                string `url:"reserve"`                   // Optional
	ShowReserve            int32  `url:"show_reserve"`              // Optional
	BuyNow                 string `url:"buy_now"`                   // Optional
	PaymentPlanOffered     int32  `url:"payment_plan_offered"`      // Optional
	PaymentPlanMonths      int32  `url:"payment_plan_months"`       // Optional
	PaymentPlanDownPayment string `url:"payment_plan_down_payment"` // Optional
	EndDate                string `url:"end_date"`                  // Optional
	EndDateUseMaximum      int32  `url:"end_date_use_maximum"`      // Optional
	NotifyBuyers           int32  `url:"notify_buyers"`             // Optional
	Category1              string `url:"category1"`                 // Optional
	Description            string `url:"description"`               // Optional
	UseForSaleLandingPage  int32  `url:"use_for_sale_landing_page"` // Optional
	MpUseOurNameservers    int32  `url:"mp_use_our_nameservers"`    // Optional
	Password               string `url:"password"`                  // Optional
	CancelSale             int32  `url:"cancel_sale"`               // Optional
}

type MarketplaceLandingPageUpdateParams struct {
	Domain string `url:"domain"` // Required

	MpTemplate         int32  `url:"mp_template"`            // Optional
	MpBgcolor          string `url:"mp_bgcolor"`             // Optional
	MpTextcolor        string `url:"mp_textcolor"`           // Optional
	MpShowBuyNow       int32  `url:"mp_show_buy_now"`        // Optional
	MpShowMoreInfo     int32  `url:"mp_show_more_info"`      // Optional
	MpShowRenewalPrice int32  `url:"mp_show_renewal_price"`  // Optional
	MpShowOtherForSale int32  `url:"mp_show_other_for_sale"` // Optional
	MpOtherDomainLinks string `url:"mp_other_domain_links"`  // Optional
	MpMessage          string `url:"mp_message"`             // Optional
}

type ModifyRegisteredNameServerParams struct {
	CurrentHost string `url:"current_host"` // Required
	NewHost     string `url:"new_host"`     // Required
	IP1         string `url:"ip1"`          // Required

	IP2  string `url:"ip2"`  // Optional
	IP3  string `url:"ip3"`  // Optional
	IP4  string `url:"ip4"`  // Optional
	IP5  string `url:"ip5"`  // Optional
	IP6  string `url:"ip6"`  // Optional
	IP7  string `url:"ip7"`  // Optional
	IP8  string `url:"ip8"`  // Optional
	IP9  string `url:"ip9"`  // Optional
	IP10 string `url:"ip10"` // Optional
	IP11 string `url:"ip11"` // Optional
	IP12 string `url:"ip12"` // Optional
	IP13 string `url:"ip13"` // Optional
}

type OrderDetailsParams struct {
	OrderNumber int `url:"order_number"`
}

type PortfolioAddParams struct {
	Portfolio string `url:"portfolio"` // Required
}

type PortfolioDeleteParams struct {
	Portfolio string `url:"portfolio"` // Required
}

type PortfolioDomainAssociateParams struct {
	Portfolio string `url:"portfolio"` // Required
	Domains   string `url:"domains"`   // Required (Comma-delimited list)
}

type PortfolioListParams struct{}

type RegisterDomainParams struct {
	Domain string `url:"domain"` // Required
	Years  int32  `url:"years"`  // Required

	PaymentID string `url:"payment_id"` // Optional
	Private   int32  `url:"private"`    // Optional
	AutoRenew int32  `url:"auto_renew"` // Optional
	Portfolio string `url:"portfolio"`  // Optional
	Coupon    string `url:"coupon"`     // Optional

	NameServer1  string `url:"ns1"`
	NameServer2  string `url:"ns2"`
	NameServer3  string `url:"ns3"`
	NameServer4  string `url:"ns4"`
	NameServer5  string `url:"ns5"`
	NameServer6  string `url:"ns6"`
	NameServer7  string `url:"ns7"`
	NameServer8  string `url:"ns8"`
	NameServer9  string `url:"ns9"`
	NameServer10 string `url:"ns10"`
	NameServer11 string `url:"ns11"`
	NameServer12 string `url:"ns12"`
	NameServer13 string `url:"ns13"`

	FirstName                     string `url:"fn"` // Contact Information
	LastName                      string `url:"ln"` // Contact Information
	MailingAddress                string `url:"ad"` // Contact Information
	MailingCity                   string `url:"cy"` // Contact Information
	MailingStateProvinceTerritory string `url:"st"` // Contact Information
	MailingZipPostalCode          string `url:"zp"` // Contact Information
	MailingCountry                string `url:"ct"` // Contact Information
	EmailAddress                  string `url:"em"` // Contact Information
	PhoneNumber                   string `url:"ph"` // Contact Information

	Company         string `url:"cp"`  // Contact Information
	MailingAddress2 string `url:"ad2"` // Contact Information
	Fax             string `url:"fx"`  // Contact Information

	USNexusCategory      string `url:"usnc"` // Contact Information
	USApplicationPurpose string `url:"usap"` // Contact Information

	ContactID string `url:"contact_id"` // Contact ID
}

type RegisterDomainDropParams struct {
	Domain string `url:"domain"` // Required
	Years  int32  `url:"years"`  // Required

	Private   int32 `url:"private"`    // Optional
	AutoRenew int32 `url:"auto_renew"` // Optional
}

type RegistrantVerificationStatusParams struct{}

type RemoveAutoRenewalParams struct {
	Domain string `url:"domain"` // Required
}

type RemovePrivacyParams struct {
	Domain string `url:"domain"` // Required
}

type RenewDomainParams struct {
	Domain string `url:"domain"` // Required
	Years  int32  `url:"years"`  // Required

	PaymentID string `url:"payment_id"` // Optional
	Coupon    string `url:"coupon"`     // Optional
}

type RetrieveAuthCodeParams struct {
	Domain string `url:"domain"` // Required
}

type TransferDomainParams struct {
	Domain string `url:"domain"` // Required

	PaymentID string `url:"payment_id"` // Optional
	Auth      string `url:"auth"`       // Optional
	Private   int32  `url:"private"`    // Optional
	AutoRenew int32  `url:"auto_renew"` // Optional
	Portfolio string `url:"portfolio"`  // Optional
	Coupon    string `url:"coupon"`     // Optional

	FirstName                     string `url:"fn"` // Contact Information
	LastName                      string `url:"ln"` // Contact Information
	MailingAddress                string `url:"ad"` // Contact Information
	MailingCity                   string `url:"cy"` // Contact Information
	MailingStateProvinceTerritory string `url:"st"` // Contact Information
	MailingZipPostalCode          string `url:"zp"` // Contact Information
	MailingCountry                string `url:"ct"` // Contact Information
	EmailAddress                  string `url:"em"` // Contact Information
	PhoneNumber                   string `url:"ph"` // Contact Information

	Company         string `url:"cp"`  // Contact Information
	MailingAddress2 string `url:"ad2"` // Contact Information
	Fax             string `url:"fx"`  // Contact Information

	USNexusCategory      string `url:"usnc"` // Contact Information
	USApplicationPurpose string `url:"usap"` // Contact Information

	ContactID string `url:"contact_id"` // Contact ID
}

type TransferUpdateChangeEPPCodeParams struct {
	Domain string `url:"domain"` // Required
	Auth   string `url:"auth"`   // Required
}

type TransferUpdateResendAdminEmailParams struct {
	Domain string `url:"domain"` // Required
}

type TransferUpdateResubmitToRegistryParams struct {
	Domain string `url:"domain"` // Required
}
