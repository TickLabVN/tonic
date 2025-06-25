package schema_test

import (
	"testing"

	"github.com/TickLabVN/tonic/core/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidate_Format(t *testing.T) {
	type Test struct {
		Base64                     string `json:"base64" validate:"base64"`
		Base64URL                  string `json:"base64url" validate:"base64url"`
		Base64RawURL               string `json:"base64rawurl" validate:"base64rawurl"`
		BIC                        string `json:"bic" validate:"bic"`
		BCP47LanguageTag           string `json:"bcp47LanguageTag" validate:"bcp47_language_tag"`
		BTCAddr                    string `json:"btcAddr" validate:"btc_addr"`
		BTCAddrBech32              string `json:"btcAddrBech32" validate:"btc_addr_bech32"`
		CreditCard                 string `json:"creditCard" validate:"credit_card"`
		MongoDB                    string `json:"mongoDB" validate:"mongodb"`
		MongoDBConnectionString    string `json:"mongoDBConnectionString" validate:"mongodb_connection_string"`
		Cron                       string `json:"cron" validate:"cron"`
		SpiceDB                    string `json:"spiceDB" validate:"spicedb"`
		Datetime                   string `json:"datetime" validate:"datetime"`
		E164                       string `json:"e164" validate:"e164"`
		Email                      string `json:"email" validate:"email"`
		EthAddr                    string `json:"ethAddr" validate:"eth_addr"`
		Hexadecimal                string `json:"hexadecimal" validate:"hexadecimal"`
		Hexcolor                   string `json:"hexcolor" validate:"hexcolor"`
		HSL                        string `json:"hsl" validate:"hsl"`
		HSLA                       string `json:"hsla" validate:"hsla"`
		HTML                       string `json:"html" validate:"html"`
		HTMLEncoded                string `json:"htmlEncoded" validate:"html_encoded"`
		ISBN                       string `json:"isbn" validate:"isbn"`
		ISBN10                     string `json:"isbn10" validate:"isbn10"`
		ISBN13                     string `json:"isbn13" validate:"isbn13"`
		ISSN                       string `json:"issn" validate:"issn"`
		ISO31661Alpha2             string `json:"iso31661Alpha2" validate:"iso3166_1_alpha2"`
		ISO31661Alpha3             string `json:"iso31661Alpha3" validate:"iso3166_1_alpha3"`
		ISO31661AlphaNumeric       string `json:"iso31661AlphaNumeric" validate:"iso3166_1_alpha_numeric"`
		ISO31662                   string `json:"iso31662" validate:"iso3166_2"`
		ISO4217                    string `json:"iso4217" validate:"iso4217"`
		JSON                       string `json:"json" validate:"json"`
		JWT                        string `json:"jwt" validate:"jwt"`
		Latitude                   string `json:"latitude" validate:"latitude"`
		Longitude                  string `json:"longitude" validate:"longitude"`
		LuhnChecksum               string `json:"luhnChecksum" validate:"luhn_checksum"`
		PostcodeISO3166Alpha2      string `json:"postcodeISO3166Alpha2" validate:"postcode_iso3166_alpha2"`
		PostcodeISO3166Alpha2Field string `json:"postcodeISO3166Alpha2Field" validate:"postcode_iso3166_alpha2_field"`
		RGB                        string `json:"rgb" validate:"rgb"`
		RGBA                       string `json:"rgba" validate:"rgba"`
		SSN                        string `json:"ssn" validate:"ssn"`
		Timezone                   string `json:"timezone" validate:"timezone"`
		UUID                       string `json:"uuid" validate:"uuid"`
		UUID3                      string `json:"uuid3" validate:"uuid3"`
		UUID3RFC4122               string `json:"uuid3RFC4122" validate:"uuid3_rfc4122"`
		UUID4                      string `json:"uuid4" validate:"uuid4"`
		UUID4RFC4122               string `json:"uuid4RFC4122" validate:"uuid4_rfc4122"`
		UUID5                      string `json:"uuid5" validate:"uuid5"`
		UUID5RFC4122               string `json:"uuid5RFC4122" validate:"uuid5_rfc4122"`
		UUIDRFC4122                string `json:"uuidRFC4122" validate:"uuid_rfc4122"`
		MD4                        string `json:"md4" validate:"md4"`
		MD5                        string `json:"md5" validate:"md5"`
		SHA256                     string `json:"sha256" validate:"sha256"`
		SHA384                     string `json:"sha384" validate:"sha384"`
		SHA512                     string `json:"sha512" validate:"sha512"`
		RIPEMD128                  string `json:"ripemd128" validate:"ripemd128"`
		RIPEMD160                  string `json:"ripemd160" validate:"ripemd160"`
		TIGER128                   string `json:"tiger128" validate:"tiger128"`
		TIGER160                   string `json:"tiger160" validate:"tiger160"`
		TIGER192                   string `json:"tiger192" validate:"tiger192"`
		Semver                     string `json:"semver" validate:"semver"`
		ULID                       string `json:"ulid" validate:"ulid"`
		CVE                        string `json:"cve" validate:"cve"`
	}
	assert := assert.New(t)
	schema, err := utils.AssertParse(assert, Test{})
	assert.Nil(err)

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"base64": { "type": "string", "format": "base64", "description": "Base64 String" },
			"base64url": { "type": "string", "format": "base64url", "description": "Base64URL String" },
			"base64rawurl": { "type": "string", "format": "base64rawurl", "description": "Base64RawURL String" },
			"bic": { "type": "string", "format": "bic", "description": "Business Identifier Code (ISO 9362)" },
			"bcp47LanguageTag": { "type": "string", "format": "bcp47_language_tag", "description": "Language tag (BCP 47)" },
			"btcAddr": { "type": "string", "format": "btc_addr", "description": "Bitcoin Address" },
			"btcAddrBech32": { "type": "string", "format": "btc_addr_bech32", "description": "Bitcoin Bech32 Address (segwit)" },
			"creditCard": { "type": "string", "format": "credit_card", "description": "Credit Card Number" },
			"mongoDB": { "type": "string", "format": "mongodb", "description": "MongoDB ObjectID" },
			"mongoDBConnectionString": { "type": "string", "format": "mongodb_connection_string", "description": "MongoDB Connection String" },
			"cron": { "type": "string", "format": "cron", "description": "Cron" },
			"spiceDB": { "type": "string", "format": "spicedb", "description": "SpiceDb ObjectID/Permission/Type" },
			"datetime": { "type": "string", "format": "datetime", "description": "Datetime" },
			"e164": { "type": "string", "format": "e164", "description": "e164 formatted phone number" },
			"email": { "type": "string", "format": "email", "description": "E-mail String" },
			"ethAddr": { "type": "string", "format": "eth_addr", "description": "Ethereum Address" },
			"hexadecimal": { "type": "string", "format": "hexadecimal", "description": "Hexadecimal String" },
			"hexcolor": { "type": "string", "format": "hexcolor", "description": "Hexcolor String" },
			"hsl": { "type": "string", "format": "hsl", "description": "HSL String" },
			"hsla": { "type": "string", "format": "hsla", "description": "HSLA String" },
			"html": { "type": "string", "format": "html", "description": "HTML Tags" },
			"htmlEncoded": { "type": "string", "format": "html_encoded", "description": "HTML Encoded" },
			"isbn": { "type": "string", "format": "isbn", "description": "International Standard Book Number" },
			"isbn10": { "type": "string", "format": "isbn10", "description": "International Standard Book Number 10" },
			"isbn13": { "type": "string", "format": "isbn13", "description": "International Standard Book Number 13" },
			"issn": { "type": "string", "format": "issn", "description": "International Standard Serial Number" },
			"iso31661Alpha2": { "type": "string", "format": "iso3166_1_alpha2", "description": "Two-letter country code (ISO 3166-1 alpha-2)" },
			"iso31661Alpha3": { "type": "string", "format": "iso3166_1_alpha3", "description": "Three-letter country code (ISO 3166-1 alpha-3)" },
			"iso31661AlphaNumeric": { "type": "string", "format": "iso3166_1_alpha_numeric", "description": "Numeric country code (ISO 3166-1 numeric)" },
			"iso31662": { "type": "string", "format": "iso3166_2", "description": "Country subdivision code (ISO 3166-2)" },
			"iso4217": { "type": "string", "format": "iso4217", "description": "Currency code (ISO 4217)" },
			"json": { "type": "string", "format": "json", "description": "JSON" },
			"jwt": { "type": "string", "format": "jwt", "description": "JSON Web Token (JWT)" },
			"latitude": { "type": "string", "format": "latitude", "description": "Latitude" },
			"longitude": { "type": "string", "format": "longitude", "description": "Longitude" },
			"luhnChecksum": { "type": "string", "format": "luhn_checksum", "description": "Luhn Algorithm Checksum (for strings and (u)int)" },
			"postcodeISO3166Alpha2": { "type": "string", "format": "postcode_iso3166_alpha2", "description": "Postcode" },
			"postcodeISO3166Alpha2Field": { "type": "string", "format": "postcode_iso3166_alpha2_field", "description": "Postcode" },
			"rgb": { "type": "string", "format": "rgb", "description": "RGB String" },
			"rgba": { "type": "string", "format": "rgba", "description": "RGBA String" },
			"ssn": { "type": "string", "format": "ssn", "description": "Social Security Number SSN" },
			"timezone": { "type": "string", "format": "timezone", "description": "Timezone" },
			"uuid": { "type": "string", "format": "uuid", "description": "Universally Unique Identifier UUID" },
			"uuid3": { "type": "string", "format": "uuid3", "description": "Universally Unique Identifier UUID v3" },
			"uuid3RFC4122": { "type": "string", "format": "uuid3_rfc4122", "description": "Universally Unique Identifier UUID v3 RFC4122" },
			"uuid4": { "type": "string", "format": "uuid4", "description": "Universally Unique Identifier UUID v4" },
			"uuid4RFC4122": { "type": "string", "format": "uuid4_rfc4122", "description": "Universally Unique Identifier UUID v4 RFC4122" },
			"uuid5": { "type": "string", "format": "uuid5", "description": "Universally Unique Identifier UUID v5" },
			"uuid5RFC4122": { "type": "string", "format": "uuid5_rfc4122", "description": "Universally Unique Identifier UUID v5 RFC4122" },
			"uuidRFC4122": { "type": "string", "format": "uuid_rfc4122", "description": "Universally Unique Identifier UUID RFC4122" },
			"md4": { "type": "string", "format": "md4", "description": "MD4 hash" },
			"md5": { "type": "string", "format": "md5", "description": "MD5 hash" },
			"sha256": { "type": "string", "format": "sha256", "description": "SHA256 hash" },
			"sha384": { "type": "string", "format": "sha384", "description": "SHA384 hash" },
			"sha512": { "type": "string", "format": "sha512", "description": "SHA512 hash" },
			"ripemd128": { "type": "string", "format": "ripemd128", "description": "RIPEMD-128 hash" },
			"ripemd160": { "type": "string", "format": "ripemd160", "description": "RIPEMD-160 hash" },
			"tiger128": { "type": "string", "format": "tiger128", "description": "TIGER128 hash" },
			"tiger160": { "type": "string", "format": "tiger160", "description": "TIGER160 hash" },
			"tiger192": { "type": "string", "format": "tiger192", "description": "TIGER192 hash" },
			"semver": { "type": "string", "format": "semver", "description": "Semantic Versioning 2.0.0" },
			"ulid": { "type": "string", "format": "ulid", "description": "Universally Unique Lexicographically Sortable Identifier ULID" },
			"cve": { "type": "string", "format": "cve", "description": "Common Vulnerabilities and Exposures Identifier (CVE id)" }
		}
	}`, schema)
}
