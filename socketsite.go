package catogo

import (
	"encoding/json"
)

//move to an other file AccountSnapshot
type AccountSnapshot struct {
	AccountSnapshot struct {
		Sites []SiteSnapshot `json:"sites,omitempty"`
	} `json:"accountSnapshot,omitempty"`
}

type SiteSnapshot struct {
	ID				string   `json:"id,omitempty"`
	Info struct {
		Sockets []SocketInfo `json:"sockets,omitempty"`
	} `json:"info,omitempty"`
}

type SocketInfo struct {
	Serial		string	`json:"serial,omitempty"`
}

type SiteMutation struct {
	Site struct {
		AddSocketSite struct {
			SiteID string `json:"siteId,omitempty"`
		} `json:addSocketSite,omitempty"`
	} `json:"site,omitempty"`
}


type SocketSite struct {
	ID                   string   `json:"id,omitempty"`
	Name               string       `json:"name,omitempty"`
	Serial				string		`json:"serial,omitempty"`
	Description        string       `json:"description,omitempty"`
	SiteType           string       `json:"siteType,omitempty"`
	NativeNetworkRange string       `json:"nativeNetworkRange,omitempty"`
	ConnectionType     string       `json:"connectionType,omitempty"`
	SiteLocation       SiteLocation `json:"siteLocation,omitempty"`
}

type SiteLocation struct {
	CountryCode string `json:"countryCode,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
}


func (c *Client) GetSocketSerial(accountID string, siteID string) (string, error) {
	query := graphQLRequest{
		Query: `query accountSnapshot ($accountID: ID!, $siteIDs: [ID!] ) {
			accountSnapshot (accountID: $accountID) {
				sites (siteIDs: $siteIDs) {
					id
					info {
						sockets {
							serial
						}
					}
					
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountID": accountID,
			"siteIDs": []string{siteID},
		},
	}

	output := validResponse{}

	if err := c.post(query, &output); err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(output.Data)
	if err != nil {
		return "", err
	}
	accountsnapshot := AccountSnapshot{}
	json.Unmarshal(jsonData, &accountsnapshot)
	return accountsnapshot.AccountSnapshot.Sites[0].Info.Sockets[0].Serial, nil
}

func (c *Client) AddSocketSite(accountID string, input *SocketSite) (*SiteMutation, error) {

	query := graphQLRequest{
		Query: `mutation addSocketSite ($accountId: ID!, $input: AddSocketSiteInput!) {
			site(accountId: $accountId) {
				addSocketSite (input: $input) {
					siteId
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": accountID,
			"input":     input,
		},
	}

	output := validResponse{}

	if err := c.post(query, &output); err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(output.Data)
	if err != nil {
		return nil, err
	}

	site := SiteMutation{}
	json.Unmarshal(jsonData, &site)

	return &site, nil
}

