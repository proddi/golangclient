package provisionclient

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Define a custom type that implements the json.Unmarshaler interface.
type PVID string

func (n *PVID) UnmarshalJSON(b []byte) error {
	// Convert the number to a string and store it in the Number value.
	if b[0] == '"' {
		*n = PVID(strings.Trim(string(b), "\""))
	} else if string(b) == "null" {
		*n = ""
	} else {
		*n = PVID(string(b))
	}
	return nil
}

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Auth       AuthStruct
	DNS        DNSMethods
	Resources  ResourceMethods
	IPAM       IPAMMethods
	DHCP       DHCPMethods
	Umbrella   UmbrellaMethods
	Contacts   ContactsMethods
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewClient -
func NewClient(host, username, password string, skipTLSVerify bool) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: skipTLSVerify,
				},
			}},
		HostURL: strings.Trim(host, "/"),
	}

	c.Auth = AuthStruct{
		Username: username,
		Password: password,
	}

	c.DNS.Client = &c
	c.Resources.Client = &c
	c.IPAM.Client = &c
	c.DHCP.Client = &c
	c.Umbrella.Client = &c
	c.Contacts.Client = &c

	return &c, nil
}

func (c *Client) doRequest(method, relative_url string, input_body io.Reader) ([]byte, error) {
	url := c.HostURL + "/api/v2/" + strings.Trim(relative_url, "/")
	req, err := http.NewRequest(method, url, input_body)
	if err != nil {
		return nil, err
	}

	auth := c.Auth.Username + ":" + c.Auth.Password
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	req.Header.Add("Accept", "application/json")
	if method != "GET" {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
