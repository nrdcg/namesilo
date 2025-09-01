package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// https://github.com/miku/zek
// for i in $(ls -1); do zek -c -n ${i%.xml} "./$i" >> "bbb.go"; done

const baseDir = "fixtures"

type NamesiloXML struct {
	XMLName xml.Name `xml:"namesilo"`
	Text    string   `xml:",chardata"`
	Request struct {
		Text      string `xml:",chardata"`
		Operation string `xml:"operation"`
		Ip        string `xml:"ip"`
	} `xml:"request"`
	Reply struct {
		Text   string `xml:",chardata"`
		Code   string `xml:"code"`
		Detail string `xml:"detail"`
	} `xml:"reply"`
}

type NamesiloJSON struct {
	Request struct {
		Operation string `json:"operation"`
		Ip        string `json:"ip"`
	} `json:"request"`
	Reply any `json:"reply"`
}

func main() {
	// Extracted with:
	// ```
	// document.querySelectorAll('.j-link').forEach(e => console.log(e.href))
	// ```

	pagesUID := []string{
		"domains/register-domain",
		"domains/register-domain-drop",
		"domains/register-domain-claims",
		"domains/renew-domain",
		"domains/transfer-domain",
		"domains/list-domains",
		"domains/get-domain-info",
		"account/get-prices",
		"domains/domain-forward",
		"domains/domain-forward-sub-domain",
		"domains/domain-forward-sub-domain-delete",
		"domains/add-auto-renewal",
		"domains/remove-auto-renewal",
		"domains/domain-lock",
		"domains/domain-unlock",
		"domains/domain-push",
		"domains/check-register-availability",
		"domains/check-transfer-availability",
		"domains/whois",
		"transfers/transfer-domain",
		"transfers/check-transfer-status",
		"transfers/transfer-update-change-epp-code",
		"transfers/transfer-update-resend-admin-email",
		"transfers/transfer-update-resubmit-registry",
		"transfers/retrieve-auth-code",
		"contact/contact-list",
		"contact/contact-add",
		"contact/contact-update",
		"contact/contact-delete",
		"contact/contact-domain-associate",
		"nameserver/change-nameserver",
		"nameserver/list-registered-nameservers",
		"nameserver/add-registered-nameserver",
		"nameserver/modify-registered-nameserver",
		"nameserver/delete-registered-nameserver",
		"dns/dns-list-records",
		"dns/dns-add-record",
		"dns/dns-update-record",
		"dns/dns-delete-record",
		"dns/dns-seclist-records",
		"dns/dns-secadd-records",
		"dns/dns-secdelete-record",
		"portfolio/portfolio-list",
		"portfolio/portfolio-add",
		"portfolio/portfolio-delete",
		"portfolio/portfolio-domain-associate",
		"privacy/add-privacy",
		"privacy/remove-privacy",
		"email/list-email-forwards",
		"email/configure-email-forward",
		"email/delete-email-forward",
		"email/registrant-verification-status",
		"email/email-verification",
		"email/transfer-update",
		"marketplace/marketplace-active-sales-overview",
		"marketplace/marketplace-add-modify-sale",
		"marketplace/marketplace-landing-page-update",
		"forwarding/forward-domain",
		"forwarding/forward-domain-sub-domain",
		"forwarding/domain-forward-subdomain-delete",
		"forwarding/list-email-forwards",
		"forwarding/configure-email-forwards",
		"forwarding/delete-email-forwards",
		"account/get-account-balance",
		"account/add-account-funds",
		"account/list-orders",
		"account/order-details",
		"account/list-expiring-domains",
		"account/count-expiring-domains",
		"auctions/list-auctions",
		"auctions/view-auction",
		"auctions/view-auctions",
		"auctions/watch-auction",
		"auctions/bid-auction",
		"auctions/bulk-bid-auction",
		"auctions/buy-now-auction",
		"auctions/view-auction-history",
	}

	slices.SortFunc(pagesUID, func(a, b string) int {
		return strings.Compare(path.Base(a), path.Base(b))
	})

	err := scrape(pagesUID)
	if err != nil {
		log.Fatal(err)
	}
}

func scrape(pagesUID []string) error {
	operations := map[string]string{}

	for _, uid := range pagesUID {
		doc, err := getPage(uid)
		if err != nil {
			return fmt.Errorf("%s: %w", uid, err)
		}

		var operation, baseOperation string

		doc.Find("input.apiRequestExamples").Each(func(i int, s *goquery.Selection) {
			val, exists := s.Attr("value")
			if exists {
				uri, err := url.Parse(val)
				if err != nil {
					log.Printf("%s: %v", uid, err)
					return
				}

				operation = path.Base(uri.Path)
				baseOperation = path.Base(uri.Path)
			}
		})

		if prev, find := operations[operation]; find {
			if prev == uid || path.Base(prev) == path.Base(uid) {
				log.Printf("SKIP: Multiple operations with the same name: %s (prev: %s, curr: %s)", operation, prev, uid)
				continue
			}

			log.Printf("Multiple operations with the same name: %s (prev: %s, curr: %s)", operation, prev, uid)

			operation = operation + extractPrefix(uid, prev)
		} else {
			operations[operation] = uid
		}

		// XML

		dstXMLDir := filepath.Join(baseDir, "xml")
		err = os.MkdirAll(dstXMLDir, 0755)
		if err != nil {
			return fmt.Errorf("%s (%s): %w", operation, uid, err)
		}

		doc.Find(".xml-json-switch .xml-code .html").Each(func(i int, s *goquery.Selection) {
			namesilo := NamesiloXML{}
			err := xml.Unmarshal([]byte(s.Text()), &namesilo)
			if err != nil {
				log.Printf("XML: %s (%s): %v", baseOperation, uid, err)
			} else if namesilo.Request.Operation != baseOperation {
				log.Printf("XML: operation inside the XML is inconsistent: got %q, want %q (%s)", namesilo.Request.Operation, baseOperation, uid)
			}

			file, err := os.Create(filepath.Join(dstXMLDir, operation+".xml"))
			if err != nil {
				log.Fatalf("XML: %s (%s): %v", baseOperation, uid, err)
			}

			defer file.Close()

			_, err = file.WriteString(s.Text() + "\n")
			if err != nil {
				log.Fatalf("XML: %s (%s): %v", baseOperation, uid, err)
			}
		})

		// JSON

		dstJSONDir := filepath.Join(baseDir, "json")
		err = os.MkdirAll(dstJSONDir, 0755)
		if err != nil {
			return fmt.Errorf("XML: %s (%s): %v", baseOperation, uid, err)
		}

		doc.Find(".xml-json-switch .json-code .json").Each(func(i int, s *goquery.Selection) {
			namesilo := NamesiloJSON{}
			err := json.Unmarshal([]byte(s.Text()), &namesilo)
			if err != nil {
				log.Printf("JSON: %s (%s): %v", baseOperation, uid, err)
			} else if namesilo.Request.Operation != baseOperation {
				log.Printf("JSON: operation inside the JSON is inconsistent: got %q, want %q (%s)", namesilo.Request.Operation, baseOperation, uid)
			}

			file, err := os.Create(filepath.Join(dstJSONDir, operation+".json"))
			if err != nil {
				log.Fatalf("JSON: %s (%s): %v", baseOperation, uid, err)
			}

			defer file.Close()

			_, err = file.WriteString(s.Text() + "\n")
			if err != nil {
				log.Fatalf("JSON: %s (%s): %v", baseOperation, uid, err)
			}
		})
	}

	return nil
}

func getPage(uid string) (*goquery.Document, error) {
	req, err := newRequest(uid)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)
}

func newRequest(uid string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.namesilo.com/api-reference/pages", nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Set("uid", uid)
	req.URL.RawQuery = query.Encode()

	setFakeHeaders(req)

	return req, nil
}

func setFakeHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:143.0) Gecko/20100101 Firefox/143.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.7,fr;q=0.3")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.namesilo.com/api-reference")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
}

func extractPrefix(a, b string) string {
	aBase := path.Base(a)
	bBase := path.Base(b)

	prefix := strings.TrimPrefix(aBase, bBase)
	if aBase < bBase {
		prefix = strings.TrimPrefix(bBase, aBase)
	}

	return "_" + strings.TrimPrefix(prefix, "-")
}
