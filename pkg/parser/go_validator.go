// eq	Equals
// eq_ignore_case	Equals ignoring case
// gt	Greater than
// gte	Greater than or equal
// lt	Less Than
// lte	Less Than or Equal
// ne	Not Equal
// ne_ignore_case	Not Equal ignoring case
package parser

import (
	"encoding/json"
	"errors"
	"strings"
)

type ValidateTag struct {
	// Fields
	EqCsField     bool `json:"eqcsfield,omitempty"`
	EqField       bool `json:"eqfield,omitempty"`
	FieldContains bool `json:"fieldcontains,omitempty"`
	GtCsField     bool `json:"gtcsfield,omitempty"`
	GtField       bool `json:"gtfield,omitempty"`
	GteCsField    bool `json:"gtecsfield,omitempty"`
	GteField      bool `json:"gtefield,omitempty"`
	LtCsField     bool `json:"ltcsfield,omitempty"`
	LtField       bool `json:"ltfield,omitempty"`
	LteCsField    bool `json:"ltecsfield,omitempty"`
	LteField      bool `json:"ltefield,omitempty"`
	NeCsField     bool `json:"necsfield,omitempty"`
	NeField       bool `json:"nefield,omitempty"`

	// Network
	Cidr            bool `json:"cidr"`
	CidrV4          bool `json:"cidrv4"`
	CidrV6          bool `json:"cidrv6"`
	DataURI         bool `json:"datauri"`
	Fqdn            bool `json:"fqdn"`
	Hostname        bool `json:"hostname"`
	HostnamePort    bool `json:"hostname_port"`
	HostnameRfc1123 bool `json:"hostname_rfc1123"`
	IP              bool `json:"ip"`
	IP4Addr         bool `json:"ip4_addr"`
	IP6Addr         bool `json:"ip6_addr"`
	IPAddr          bool `json:"ip_addr"`
	IPv4            bool `json:"ipv4"`
	IPv6            bool `json:"ipv6"`
	MAC             bool `json:"mac"`
	TCP4Addr        bool `json:"tcp4_addr"`
	TCP6Addr        bool `json:"tcp6_addr"`
	TCPAddr         bool `json:"tcp_addr"`
	UDP4Addr        bool `json:"udp4_addr"`
	UDP6Addr        bool `json:"udp6_addr"`
	UDPAddr         bool `json:"udp_addr"`
	UnixAddr        bool `json:"unix_addr"`
	URI             bool `json:"uri"`
	URL             bool `json:"url"`
	HTTPURL         bool `json:"http_url"`
	URLEncoded      bool `json:"url_encoded"`
	URNRfc2141      bool `json:"urn_rfc2141"`

	// String
	Alpha           bool `json:"alpha"`
	Alphanum        bool `json:"alphanum"`
	AlphanumUnicode bool `json:"alphanumunicode"`
	AlphaUnicode    bool `json:"alphaunicode"`
	ASCII           bool `json:"ascii"`
	Boolean         bool `json:"boolean"`
	Contains        bool `json:"contains"`
	ContainsAny     bool `json:"containsany"`
	ContainsRune    bool `json:"containsrune"`
	EndsNotWith     bool `json:"endsnotwith"`
	EndsWith        bool `json:"endswith"`
	Excludes        bool `json:"excludes"`
	ExcludesAll     bool `json:"excludesall"`
	ExcludesRune    bool `json:"excludesrune"`
	Lowercase       bool `json:"lowercase"`
	Multibyte       bool `json:"multibyte"`
	Number          bool `json:"number"`
	Numeric         bool `json:"numeric"`
	PrintASCII      bool `json:"printascii"`
	StartsNotWith   bool `json:"startsnotwith"`
	StartsWith      bool `json:"startswith"`
	Uppercase       bool `json:"uppercase"`

	// Formats
	Base64                     bool `json:"base64"`
	Base64URL                  bool `json:"base64url"`
	Base64RawURL               bool `json:"base64rawurl"`
	BIC                        bool `json:"bic"`
	BCP47Language              bool `json:"bcp47_language_tag"`
	BTCAddress                 bool `json:"btc_addr"`
	BTCAddressBech32           bool `json:"btc_addr_bech32"`
	CreditCard                 bool `json:"credit_card"`
	MongoDB                    bool `json:"mongodb"`
	MongoDBConnectionString    bool `json:"mongodb_connection_string"`
	Cron                       bool `json:"cron"`
	SpiceDBObject              bool `json:"spicedb"`
	Datetime                   bool `json:"datetime"`
	E164                       bool `json:"e164"`
	Email                      bool `json:"email"`
	EthereumAddress            bool `json:"eth_addr"`
	Hexadecimal                bool `json:"hexadecimal"`
	Hexcolor                   bool `json:"hexcolor"`
	HSL                        bool `json:"hsl"`
	HSLA                       bool `json:"hsla"`
	HTML                       bool `json:"html"`
	HTMLEncoded                bool `json:"html_encoded"`
	ISBN                       bool `json:"isbn"`
	ISBN10                     bool `json:"isbn10"`
	ISBN13                     bool `json:"isbn13"`
	ISSN                       bool `json:"issn"`
	ISO31661Alpha2             bool `json:"iso3166_1_alpha2"`
	ISO31661Alpha3             bool `json:"iso3166_1_alpha3"`
	ISO31661AlphaNumeric       bool `json:"iso3166_1_alpha_numeric"`
	ISO31662                   bool `json:"iso3166_2"`
	ISO4217                    bool `json:"iso4217"`
	JSON                       bool `json:"json"`
	JWT                        bool `json:"jwt"`
	Latitude                   bool `json:"latitude"`
	Longitude                  bool `json:"longitude"`
	LuhnChecksum               bool `json:"luhn_checksum"`
	PostcodeISO3166Alpha2      bool `json:"postcode_iso3166_alpha2"`
	PostcodeISO3166Alpha2Field bool `json:"postcode_iso3166_alpha2_field"`
	RGB                        bool `json:"rgb"`
	RGBA                       bool `json:"rgba"`
	SSN                        bool `json:"ssn"`
	Timezone                   bool `json:"timezone"`
	UUID                       bool `json:"uuid"`
	UUID3                      bool `json:"uuid3"`
	UUID3RFC4122               bool `json:"uuid3_rfc4122"`
	UUID4                      bool `json:"uuid4"`
	UUID4RFC4122               bool `json:"uuid4_rfc4122"`
	UUID5                      bool `json:"uuid5"`
	UUID5RFC4122               bool `json:"uuid5_rfc4122"`
	UUIDRFC4122                bool `json:"uuid_rfc4122"`
	MD4                        bool `json:"md4"`
	MD5                        bool `json:"md5"`
	SHA256                     bool `json:"sha256"`
	SHA384                     bool `json:"sha384"`
	SHA512                     bool `json:"sha512"`
	RIPEMD128                  bool `json:"ripemd128"`
	RIPEMD160                  bool `json:"ripemd160"`
	TIGER128                   bool `json:"tiger128"`
	TIGER160                   bool `json:"tiger160"`
	TIGER192                   bool `json:"tiger192"`
	Semver                     bool `json:"semver"`
	ULID                       bool `json:"ulid"`
	CVE                        bool `json:"cve"`

	// Comparison
	Eq           string `json:"eq,omitempty"`
	EqIgnoreCase string `json:"eq_ignore_case,omitempty"`
	Ne           string `json:"ne,omitempty"`
	NeIgnoreCase string `json:"ne_ignore_case,omitempty"`
	Gt           string `json:"gt,omitempty"`
	Gte          string `json:"gte,omitempty"`
	Lt           string `json:"lt,omitempty"`
	Lte          string `json:"lte,omitempty"`

	// Others
	Dir                bool            `json:"dir"`
	DirPath            bool            `json:"dirpath"`
	File               bool            `json:"file"`
	FilePath           bool            `json:"filepath"`
	Image              bool            `json:"image"`
	IsDefault          bool            `json:"isdefault"`
	Len                string          `json:"len"`
	Max                string          `json:"max"`
	Min                string          `json:"min"`
	OneOf              []string        `json:"oneof,omitempty"`
	Required           bool            `json:"required"`
	RequiredIf         *TagComparision `json:"required_if,omitempty"`
	RequiredUnless     *TagComparision `json:"required_unless,omitempty"`
	RequiredWith       string          `json:"required_with,omitempty"`
	RequiredWithAll    []string        `json:"required_with_all,omitempty"`
	RequiredWithout    string          `json:"required_without,omitempty"`
	RequiredWithoutAll []string        `json:"required_without_all,omitempty"`
	ExcludedIf         *TagComparision `json:"excluded_if,omitempty"`
	ExcludedUnless     *TagComparision `json:"excluded_unless,omitempty"`
	ExcludedWith       string          `json:"excluded_with,omitempty"`
	ExcludedWithAll    []string        `json:"excluded_with_all,omitempty"`
	ExcludedWithout    string          `json:"excluded_without,omitempty"`
	ExcludedWithoutAll []string        `json:"excluded_without_all,omitempty"`
	Unique             bool            `json:"unique"`
}

type TagComparision struct {
	FieldName string   `json:"fieldname"`
	Values    []string `json:"values"`
}

func ParseTag(tag string) (*ValidateTag, error) {
	tag = strings.TrimSpace(tag)
	if len(tag) == 0 {
		return nil, errors.New("empty tag")
	}
	parts := strings.Split(tag, ",")
	rawTag := make(map[string]interface{})

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
		kv := strings.Split(parts[i], "=")
		var key = kv[0]
		if len(kv) > 1 {
			values := strings.Split(kv[1], " ")
			if len(values) > 1 {
				rawTag[key] = values
			} else {
				rawTag[key] = kv[1]
			}
		} else {
			rawTag[key] = true
		}
	}

	validateTag := &ValidateTag{}
	b, err := json.Marshal(rawTag)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, validateTag)
	if err != nil {
		return nil, err
	}
	return validateTag, nil
}
