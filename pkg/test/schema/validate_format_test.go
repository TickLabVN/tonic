package schema_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	pkg "github.com/TickLabVN/tonic"
	"github.com/TickLabVN/tonic/parser"
	"github.com/stretchr/testify/assert"
)

var inputMap = map[string]interface{}{
	"base64": struct {
		Base64 string `json:"base64" validate:"format=base64"`
	}{
		Base64: "dGVzdA==",
	},
	"base64url": struct {
		Base64URL string `json:"base64url" validate:"format=base64url"`
	}{
		Base64URL: "dGVzdA==",
	},
	"base64rawurl": struct {
		Base64RawURL string `json:"base64rawurl" validate:"format=base64rawurl"`
	}{
		Base64RawURL: "dGVzdA",
	},
	"bic": struct {
		BIC string `json:"bic" validate:"format=bic"`
	}{
		BIC: "DEUTDEFF",
	},
	"bcp47_language_tag": struct {
		BCP47LanguageTag string `json:"bcp47_language_tag" validate:"format=bcp47_language_tag"`
	}{
		BCP47LanguageTag: "en-US",
	},
	"btc_addr": struct {
		BTCAddr string `json:"btc_addr" validate:"format=btc_addr"`
	}{
		BTCAddr: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
	},
	"btc_addr_bech32": struct {
		BTCAddrBech32 string `json:"btc_addr_bech32" validate:"format=btc_addr_bech32"`
	}{
		BTCAddrBech32: "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf9ls4",
	},
	"credit_card": struct {
		CreditCard string `json:"credit_card" validate:"format=credit_card"`
	}{
		CreditCard: "4111111111111111",
	},
	"mongodb": struct {
		MongoDB string `json:"mongodb" validate:"format=mongodb"`
	}{
		MongoDB: "507f191e810c19729de860ea",
	},
	"mongodb_connection_string": struct {
		MongoDBConnectionString string `json:"mongodb_connection_string" validate:"format=mongodb_connection_string"`
	}{
		MongoDBConnectionString: "mongodb://username:password@localhost:27017/database",
	},
	"cron": struct {
		Cron string `json:"cron" validate:"format=cron"`
	}{
		Cron: "0 0 * * *",
	},
	"spicedb": struct {
		SpiceDb string `json:"spicedb" validate:"format=spicedb"`
	}{
		SpiceDb: "document:example::reader",
	},
	"datetime": struct {
		Datetime string `json:"datetime" validate:"format=datetime"`
	}{
		Datetime: "2024-06-26T15:00:00Z",
	},
	"e164": struct {
		E164 string `json:"e164" validate:"format=e164"`
	}{
		E164: "+14155552671",
	},
	"email": struct {
		Email string `json:"email" validate:"format=email"`
	}{
		Email: "example@example.com",
	},
	"eth_addr": struct {
		EthAddr string `json:"eth_addr" validate:"format=eth_addr"`
	}{
		EthAddr: "0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
	},
	"hexadecimal": struct {
		Hexadecimal string `json:"hexadecimal" validate:"format=hexadecimal"`
	}{
		Hexadecimal: "4d2",
	},
	"hexcolor": struct {
		Hexcolor string `json:"hexcolor" validate:"format=hexcolor"`
	}{
		Hexcolor: "#FFFFFF",
	},
	"hsl": struct {
		HSL string `json:"hsl" validate:"format=hsl"`
	}{
		HSL: "hsl(0, 100%, 50%)",
	},
	"hsla": struct {
		HSLA string `json:"hsla" validate:"format=hsla"`
	}{
		HSLA: "hsla(0, 100%, 50%, 0.5)",
	},
	"html": struct {
		HTML string `json:"html" validate:"format=html"`
	}{
		HTML: "<div>Example</div>",
	},
	"html_encoded": struct {
		HTMLEncoded string `json:"html_encoded" validate:"format=html_encoded"`
	}{
		HTMLEncoded: "&lt;div&gt;Example&lt;/div&gt;",
	},
	"isbn": struct {
		ISBN string `json:"isbn" validate:"format=isbn"`
	}{
		ISBN: "978-3-16-148410-0",
	},
	"isbn10": struct {
		ISBN10 string `json:"isbn10" validate:"format=isbn10"`
	}{
		ISBN10: "0-306-40615-2",
	},
	"isbn13": struct {
		ISBN13 string `json:"isbn13" validate:"format=isbn13"`
	}{
		ISBN13: "978-0-306-40615-7",
	},
	"issn": struct {
		ISSN string `json:"issn" validate:"format=issn"`
	}{
		ISSN: "2049-3630",
	},
	"iso3166_1_alpha2": struct {
		ISO31661Alpha2 string `json:"iso3166_1_alpha2" validate:"format=iso3166_1_alpha2"`
	}{
		ISO31661Alpha2: "US",
	},
	"iso3166_1_alpha3": struct {
		ISO31661Alpha3 string `json:"iso3166_1_alpha3" validate:"format=iso3166_1_alpha3"`
	}{
		ISO31661Alpha3: "USA",
	},
	"iso3166_1_alpha_numeric": struct {
		ISO31661AlphaNumeric string `json:"iso3166_1_alpha_numeric" validate:"format=iso3166_1_alpha_numeric"`
	}{
		ISO31661AlphaNumeric: "840",
	},
	"iso3166_2": struct {
		ISO31662 string `json:"iso3166_2" validate:"format=iso3166_2"`
	}{
		ISO31662: "US-CA",
	},
	"iso4217": struct {
		ISO4217 string `json:"iso4217" validate:"format=iso4217"`
	}{
		ISO4217: "USD",
	},
	"json": struct {
		JSON string `json:"json" validate:"format=json"`
	}{
		JSON: "{\"key\":\"value\"}",
	},
	"jwt": struct {
		JWT string `json:"jwt" validate:"format=jwt"`
	}{
		JWT: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	},
	"latitude": struct {
		Latitude string `json:"latitude" validate:"format=latitude"`
	}{
		Latitude: "37.7749",
	},
	"longitude": struct {
		Longitude string `json:"longitude" validate:"format=longitude"`
	}{
		Longitude: "-122.4194",
	},
	"luhn_checksum": struct {
		LuhnChecksum string `json:"luhn_checksum" validate:"format=luhn_checksum"`
	}{
		LuhnChecksum: "1234567812345670",
	},
	"postcode_iso3166_alpha2": struct {
		PostcodeISO3166Alpha2 string `json:"postcode_iso3166_alpha2" validate:"format=postcode_iso3166_alpha2"`
	}{
		PostcodeISO3166Alpha2: "94103",
	},
	"postcode_iso3166_alpha2_field": struct {
		PostcodeISO3166Alpha2Field string `json:"postcode_iso3166_alpha2_field" validate:"format=postcode_iso3166_alpha2_field"`
	}{
		PostcodeISO3166Alpha2Field: "94103",
	},
	"rgb": struct {
		RGB string `json:"rgb" validate:"format=rgb"`
	}{
		RGB: "rgb(255, 255, 255)",
	},
	"rgba": struct {
		RGBA string `json:"rgba" validate:"format=rgba"`
	}{
		RGBA: "rgba(255, 255, 255, 0.5)",
	},
	"ssn": struct {
		SSN string `json:"ssn" validate:"format=ssn"`
	}{
		SSN: "123-45-6789",
	},
	"timezone": struct {
		Timezone string `json:"timezone" validate:"format=timezone"`
	}{
		Timezone: "America/Los_Angeles",
	},
	"uuid": struct {
		UUID string `json:"uuid" validate:"format=uuid"`
	}{
		UUID: "123e4567-e89b-12d3-a456-426614174000",
	},
	"uuid3": struct {
		UUID3 string `json:"uuid3" validate:"format=uuid3"`
	}{
		UUID3: "123e4567-e89b-12d3-a456-426614174000",
	},
	"uuid3_rfc4122": struct {
		UUID3RFC4122 string `json:"uuid3_rfc4122" validate:"format=uuid3_rfc4122"`
	}{
		UUID3RFC4122: "123e4567-e89b-12d3-a456-426614174000",
	},
	"uuid4": struct {
		UUID4 string `json:"uuid4" validate:"format=uuid4"`
	}{
		UUID4: "123e4567-e89b-12d3-a456-426614174000",
	},
	"uuid4_rfc4122": struct {
		UUID4RFC4122 string `json:"uuid4_rfc4122" validate:"format=uuid4_rfc4122"`
	}{
		UUID4RFC4122: "123e4567-e89b-12d3-a456-426614174000",
	},
	"uuid5": struct {
		UUID5 string `json:"uuid5" validate:"format=uuid5"`
	}{
		UUID5: "123e4567-e89b-12d3-a456-426614174000",
	},
	"uuid5_rfc4122": struct {
		UUID5RFC4122 string `json:"uuid5_rfc4122" validate:"format=uuid5_rfc4122"`
	}{
		UUID5RFC4122: "123e4567-e89b-12d3-a456-426614174000",
	},
	"uuid_rfc4122": struct {
		UUIDRFC4122 string `json:"uuid_rfc4122" validate:"format=uuid_rfc4122"`
	}{
		UUIDRFC4122: "123e4567-e89b-12d3-a456-426614174000",
	},
	"md4": struct {
		MD4 string `json:"md4" validate:"format=md4"`
	}{
		MD4: "81dc9bdb52d04dc20036dbd8313ed055",
	},
	"md5": struct {
		MD5 string `json:"md5" validate:"format=md5"`
	}{
		MD5: "5d41402abc4b2a76b9719d911017c592",
	},
	"sha256": struct {
		SHA256 string `json:"sha256" validate:"format=sha256"`
	}{
		SHA256: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	},
	"sha384": struct {
		SHA384 string `json:"sha384" validate:"format=sha384"`
	}{
		SHA384: "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0ed59e46e2d540a3273d5a0b3b0f14e4dff3e172df1f2a",
	},
	"sha512": struct {
		SHA512 string `json:"sha512" validate:"format=sha512"`
	}{
		SHA512: "cf83e1357eefb8bd... (truncated)",
	},
	"ripemd128": struct {
		RIPEMD128 string `json:"ripemd128" validate:"format=ripemd128"`
	}{
		RIPEMD128: "cdf26213a150dc3ecb610f18f6b38b46",
	},
	"ripemd160": struct {
		RIPEMD160 string `json:"ripemd160" validate:"format=ripemd160"`
	}{
		RIPEMD160: "bcb22960a4a0f0b62d3c3e9f89d06429",
	},
	"tiger128": struct {
		TIGER128 string `json:"tiger128" validate:"format=tiger128"`
	}{
		TIGER128: "3d8986664fd0b82d2007a7096032af30",
	},
	"tiger160": struct {
		TIGER160 string `json:"tiger160" validate:"format=tiger160"`
	}{
		TIGER160: "3d8986664fd0b82d2007a7096032af30b8639586",
	},
	"tiger192": struct {
		TIGER192 string `json:"tiger192" validate:"format=tiger192"`
	}{
		TIGER192: "3d8986664fd0b82d2007a7096032af30b86395863e6ac03a",
	},
	"semver": struct {
		Semver string `json:"semver" validate:"format=semver"`
	}{
		Semver: "1.0.0",
	},
	"ulid": struct {
		ULID string `json:"ulid" validate:"format=ulid"`
	}{
		ULID: "01ARZ3NDEKTSV4RRFFQ69G5FAV",
	},
	"cve": struct {
		CVE string `json:"cve" validate:"format=cve"`
	}{
		CVE: "CVE-2021-44228",
	},
}

type testCase struct {
	name        string
	input       interface{}
	setupOutput func(t *testing.T, st reflect.Type) string
}

func parseFormatTag(t *testing.T, tag string) (string, string) {
	tags := strings.Split(tag, "=")
	if len(tags) < 2 {
		return tags[0], ""
	}

	if tags[0] != "format" || tags[1] == "" {
		t.Error("invalid format tag")
	}
	return tags[0], tags[1]
}

func genOutputStruct(t *testing.T, st reflect.Type) string {
	fmt.Printf("st: %v\n", st)

	if st.Kind() != reflect.Struct {
		t.Errorf("expected struct type, got %v", st.Kind())
	}

	// loop through the fields of the struct
	// and generate the schema for each field
	// and add it to the properties map
	properties := make(map[string]string)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		ft := field.Type
		if ft.Kind() != reflect.String {
			t.Errorf("expected string type, got %v", ft.Kind())
		}

		validateTag := field.Tag.Get("validate")
		jsonTag := field.Tag.Get("json")
		var fieldName string = ""
		if jsonTag != "" {
			fieldName = jsonTag
		}

		if len(fieldName) == 0 {
			fieldName = field.Name
		}

		_, format := parseFormatTag(t, validateTag)
		properties[fieldName] = fmt.Sprintf(`{
			"type": "string",
			"format": "%v"
		}`, format)
	}

	output := `{
		"type": "object",
		"properties": {
		`
	for k, v := range properties {
		output += fmt.Sprintf(`	"%v": %v,`, k, v)
	}

	output = strings.TrimSuffix(output, ",")
	output += "}"

	fmt.Println("output: ", output)

	return output
}

func runTestCases(t *testing.T, testCases []testCase) {
	assert := assert.New(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tcc := tc

			t.Parallel()
			t.Log(tcc.name)

			st := reflect.TypeOf(tcc.input)
			schema, err := parser.ParseStruct(st)
			if err != nil {
				t.Fatal(err)
			}

			b, err := json.Marshal(schema)
			if err != nil {
				t.Fatal(err)
			}

			// test if the output is equal to the expected output
			assert.JSONEq(tcc.setupOutput(t, st), string(b))

			// test if the schema is added to the spec
			spec := pkg.GetSpec()
			schemaName := fmt.Sprintf("#%s/%s", st.PkgPath(), st.Name())
			assert.NotNil(spec.Components.Schemas[schemaName])
		})
	}
}

// go test -v -run=TestValidate_Format_Base64 ./pkg/test/schema/
func TestValidate_Format_Base64(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Base64_WithValidBase64_InStruct",
			input:       inputMap["base64"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

// go test -v -run=TestValidate_Format_Base64URL ./pkg/test/schema/
func TestValidate_Format_Base64URL(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Base64URL_WithValidBase64URL_InStruct",
			input:       inputMap["base64url"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Base64RawURL(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Base64RawURL_WithValidBase64RawURL_InStruct",
			input:       inputMap["base64rawurl"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_BIC(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_BIC_WithValidBIC_InStruct",
			input:       inputMap["bic"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_BCP47LanguageTag(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_BCP47LanguageTag_WithValidBCP47LanguageTag_InStruct",
			input:       inputMap["bcp47_language_tag"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_BTCAddr(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_BTCAddr_WithValidBTCAddr_InStruct",
			input:       inputMap["btc_addr"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_BTCAddrBech32(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_BTCAddrBech32_WithValidBTCAddrBech32_InStruct",
			input:       inputMap["btc_addr_bech32"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_CreditCard(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_CreditCard_WithValidCreditCard_InStruct",
			input:       inputMap["credit_card"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_MongoDB(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_MongoDB_WithValidMongoDB_InStruct",
			input:       inputMap["mongodb"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_MongoDBConnectionString(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_MongoDBConnectionString_WithValidMongoDBConnectionString_InStruct",
			input:       inputMap["mongodb_connection_string"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}
func TestValidate_Format_Cron(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Cron_WithValidCron_InStruct",
			input:       inputMap["cron"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_SpiceDB(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_SpiceDB_WithValidSpiceDB_InStruct",
			input:       inputMap["spicedb"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Datetime(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Datetime_WithValidDatetime_InStruct",
			input:       inputMap["datetime"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_E164(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_E164_WithValidE164_InStruct",
			input:       inputMap["e164"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Email(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Email_WithValidEmail_InStruct",
			input:       inputMap["email"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_EthAddr(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_EthAddr_WithValidEthAddr_InStruct",
			input:       inputMap["eth_addr"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Hexadecimal(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Hexadecimal_WithValidHexadecimal_InStruct",
			input:       inputMap["hexadecimal"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Hexcolor(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Hexcolor_WithValidHexcolor_InStruct",
			input:       inputMap["hexcolor"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_HSL(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_HSL_WithValidHSL_InStruct",
			input:       inputMap["hsl"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_HSLA(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_HSLA_WithValidHSLA_InStruct",
			input:       inputMap["hsla"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_HTML(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_HTML_WithValidHTML_InStruct",
			input:       inputMap["html"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_HTMLEncoded(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_HTMLEncoded_WithValidHTMLEncoded_InStruct",
			input:       inputMap["html_encoded"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISBN(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISBN_WithValidISBN_InStruct",
			input:       inputMap["isbn"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISBN10(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISBN10_WithValidISBN10_InStruct",
			input:       inputMap["isbn10"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISBN13(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISBN13_WithValidISBN13_InStruct",
			input:       inputMap["isbn13"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISSN(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISSN_WithValidISSN_InStruct",
			input:       inputMap["issn"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISO31661Alpha2(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISO31661Alpha2_WithValidISO31661Alpha2_InStruct",
			input:       inputMap["iso3166_1_alpha2"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISO31661Alpha3(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISO31661Alpha3_WithValidISO31661Alpha3_InStruct",
			input:       inputMap["iso3166_1_alpha3"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISO31661AlphaNumeric(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISO31661AlphaNumeric_WithValidISO31661AlphaNumeric_InStruct",
			input:       inputMap["iso3166_1_alpha_numeric"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISO31662(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISO31662_WithValidISO31662_InStruct",
			input:       inputMap["iso3166_2"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ISO4217(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ISO4217_WithValidISO4217_InStruct",
			input:       inputMap["iso4217"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_JSON(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_JSON_WithValidJSON_InStruct",
			input:       inputMap["json"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_JWT(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_JWT_WithValidJWT_InStruct",
			input:       inputMap["jwt"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Latitude(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Latitude_WithValidLatitude_InStruct",
			input:       inputMap["latitude"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Longitude(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Longitude_WithValidLongitude_InStruct",
			input:       inputMap["longitude"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_LuhnChecksum(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_LuhnChecksum_WithValidLuhnChecksum_InStruct",
			input:       inputMap["luhn_checksum"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_PostcodeISO31661Alpha2(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_PostcodeISO31661Alpha2_WithValidPostcodeISO31661Alpha2_InStruct",
			input:       inputMap["postcode_iso3166_alpha2"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_PostcodeISO31661Alpha2Field(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_PostcodeISO31661Alpha2Field_WithValidPostcodeISO31661Alpha2Field_InStruct",
			input:       inputMap["postcode_iso3166_alpha2_field"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_RGB(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_RGB_WithValidRGB_InStruct",
			input:       inputMap["rgb"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_RGBA(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_RGBA_WithValidRGBA_InStruct",
			input:       inputMap["rgba"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_SSN(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_SSN_WithValidSSN_InStruct",
			input:       inputMap["ssn"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Timezone(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Timezone_WithValidTimezone_InStruct",
			input:       inputMap["timezone"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_UUID(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_UUID_WithValidUUID_InStruct",
			input:       inputMap["uuid"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_UUID3(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_UUID3_WithValidUUID3_InStruct",
			input:       inputMap["uuid3"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_UUID3RFC4122(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_UUID3RFC4122_WithValidUUID3RFC4122_InStruct",
			input:       inputMap["uuid3_rfc4122"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_UUID4(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_UUID4_WithValidUUID4_InStruct",
			input:       inputMap["uuid4"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_UUID4RFC4122(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_UUID4RFC4122_WithValidUUID4RFC4122_InStruct",
			input:       inputMap["uuid4_rfc4122"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_UUID5(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_UUID5_WithValidUUID5_InStruct",
			input:       inputMap["uuid5"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_UUID5RFC4122(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_UUID5RFC4122_WithValidUUID5RFC4122_InStruct",
			input:       inputMap["uuid5_rfc4122"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_UUIDRFC4122(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_UUIDRFC4122_WithValidUUIDRFC4122_InStruct",
			input:       inputMap["uuid_rfc4122"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_MD4(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_MD4_WithValidMD4_InStruct",
			input:       inputMap["md4"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_MD5(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_MD5_WithValidMD5_InStruct",
			input:       inputMap["md5"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_SHA256(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_SHA256_WithValidSHA256_InStruct",
			input:       inputMap["sha256"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_SHA384(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_SHA384_WithValidSHA384_InStruct",
			input:       inputMap["sha384"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_SHA512(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_SHA512_WithValidSHA512_InStruct",
			input:       inputMap["sha512"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_RIPEMD128(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_RIPEMD128_WithValidRIPEMD128_InStruct",
			input:       inputMap["ripemd128"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_RIPEMD160(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_RIPEMD160_WithValidRIPEMD160_InStruct",
			input:       inputMap["ripemd160"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_TIGER128(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_TIGER128_WithValidTIGER128_InStruct",
			input:       inputMap["tiger128"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_TIGER160(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_TIGER160_WithValidTIGER160_InStruct",
			input:       inputMap["tiger160"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_TIGER192(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_TIGER192_WithValidTIGER192_InStruct",
			input:       inputMap["tiger192"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_Semver(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_Semver_WithValidSemver_InStruct",
			input:       inputMap["semver"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_ULID(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_ULID_WithValidULID_InStruct",
			input:       inputMap["ulid"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}

func TestValidate_Format_CVE(t *testing.T) {
	testCases := []testCase{
		{
			name:        "TestValidate_Format_CVE_WithValidCVE_InStruct",
			input:       inputMap["cve"],
			setupOutput: genOutputStruct,
		},
	}

	runTestCases(t, testCases)
}
