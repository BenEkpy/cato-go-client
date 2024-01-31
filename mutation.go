package catogo

type SSocketSite struct {
	Name               string       `json:"name"`
	Description        string       `json:"description"`
	SiteType           string       `json:"siteType"`
	NativeNetworkRange string       `json:"nativeNetworkRange"`
	ConnectionType     string       `json:"connectionType"`
	SiteLocation       SiteLocation `json:"siteLocation"`
}

type SSiteLocation struct {
	CountryCode string `json:"countryCode"`
	Timezone    string `json:"timezone"`
}

func (c *Client) AAddSocketSite(accountID string, input *SocketSite) (interface{}, error) {

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

	return output, nil

}
