package schema_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TickLabVN/tonic/core"
	"github.com/TickLabVN/tonic/core/utils"
	"github.com/TickLabVN/tonic/core/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidate_Networks(t *testing.T) {
	type Test struct {
		Cidr            string `json:"cidr" validate:"cidr"`
		Cidrv4          string `json:"cidrv4" validate:"cidrv4"`
		Cidrv6          string `json:"cidrv6" validate:"cidrv6"`
		DataUri         string `json:"dataUri" validate:"datauri"`
		Fqdn            string `json:"fqdn" validate:"fqdn"`
		Hostname        string `json:"hostname" validate:"hostname"`
		HostnamePort    string `json:"hostnamePort" validate:"hostname_port"`
		HostnameRfc1123 string `json:"hostnameRfc1123" validate:"hostname_rfc1123"`
		Ip              string `json:"ip" validate:"ip"`
		Ip4Addr         string `json:"ip4Addr" validate:"ip4_addr"`
		Ip6Addr         string `json:"ip6Addr" validate:"ip6_addr"`
		IpAddr          string `json:"ipAddr" validate:"ip_addr"`
		Ipv4            string `json:"ipv4" validate:"ipv4"`
		Ipv6            string `json:"ipv6" validate:"ipv6"`
		Mac             string `json:"mac" validate:"mac"`
		Tcp4Addr        string `json:"tcp4Addr" validate:"tcp4_addr"`
		Tcp6Addr        string `json:"tcp6Addr" validate:"tcp6_addr"`
		TcpAddr         string `json:"tcpAddr" validate:"tcp_addr"`
		Udp4Addr        string `json:"udp4Addr" validate:"udp4_addr"`
		Udp6Addr        string `json:"udp6Addr" validate:"udp6_addr"`
		UdpAddr         string `json:"udpAddr" validate:"udp_addr"`
		UnixAddr        string `json:"unixAddr" validate:"unix_addr"`
		Uri             string `json:"uri" validate:"uri"`
		Url             string `json:"url" validate:"url"`
		HttpUrl         string `json:"httpUrl" validate:"http_url"`
		UrlEncoded      string `json:"urlEncoded" validate:"url_encoded"`
		UrnRfc2141      string `json:"urnRfc2141" validate:"urn_rfc2141"`
	}

	assert := assert.New(t)
	dt := reflect.TypeOf(Test{})

	err := validator.ParseStruct(dt)
	if err != nil {
		t.Fatal(err)
	}

	spec := core.Init()
	schemaName := utils.GetSchemaPath(dt)

	schema, ok := spec.Components.Schemas[schemaName]
	assert.True(ok)
	assert.NotNil(schema)

	b, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"cidr": { "type": "string", "format": "cidr", "description": "Classless Inter-Domain Routing CIDR"},
			"cidrv4": { "type": "string", "format": "cidrv4", "description": "Classless Inter-Domain Routing CIDRv4"},
			"cidrv6": { "type": "string", "format": "cidrv6", "description": "Classless Inter-Domain Routing CIDRv6"},
			"dataUri": { "type": "string", "format": "data-uri", "description": "Data URI"},
			"fqdn": { "type": "string", "format": "fqdn", "description": "Fully Qualified Domain Name (FQDN)" },
			"hostname": { "type": "string", "format": "hostname", "description": "Hostname RFC 952" },
			"hostnamePort": { "type": "string", "format": "hostname-port", "description": "HostPort" },
			"hostnameRfc1123": { "type": "string", "format": "hostname-rfc1123", "description": "Hostname RFC 1123" },
			"ip": { "type": "string", "format": "ip", "description": "Internet Protocol Address IP" },
			"ip4Addr": { "type": "string", "format": "ipv4", "description": "Internet Protocol Address IPv4" },
			"ip6Addr": { "type": "string", "format": "ipv6", "description": "Internet Protocol Address IPv6" },
			"ipAddr": { "type": "string", "format": "ip", "description": "Internet Protocol Address IP" },
			"ipv4": { "type": "string", "format": "ipv4", "description": "Internet Protocol Address IPv4" },
			"ipv6": { "type": "string", "format": "ipv6", "description": "Internet Protocol Address IPv6" },
			"mac": { "type": "string", "format": "mac", "description": "Media Access Control Address MAC" },
			"tcp4Addr": { "type": "string", "format": "tcp4-addr", "description": "Transmission Control Protocol Address TCPv4" },
			"tcp6Addr": { "type": "string", "format": "tcp6-addr", "description": "Transmission Control Protocol Address TCPv6" },
			"tcpAddr": { "type": "string", "format": "tcp-addr", "description": "Transmission Control Protocol Address TCP" },
			"udp4Addr": { "type": "string", "format": "udp4-addr", "description": "User Datagram Protocol Address UDPv4" },
			"udp6Addr": { "type": "string", "format": "udp6-addr", "description": "User Datagram Protocol Address UDPv6" },
			"udpAddr": { "type": "string", "format": "udp-addr", "description": "User Datagram Protocol Address UDP" },
			"unixAddr": { "type": "string", "format": "unix-addr", "description": "Unix domain socket end point Address" },
			"uri": { "type": "string", "format": "uri", "description": "URI String" },
			"url": { "type": "string", "format": "url", "description": "URL String" },
			"httpUrl": { "type": "string", "format": "http-url", "description": "HTTP URL String" },
			"urlEncoded": { "type": "string", "format": "url-encoded", "description": "URL Encoded" },
			"urnRfc2141": { "type": "string", "format": "urn-rfc2141", "description": "Urn RFC 2141 String" }
		}
	}`, string(b))
}
