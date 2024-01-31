package catogo

import (
	"encoding/json"
)

type AccountRolesQuery struct {
	AccountRoles struct {
		Items []AccountRole `json:"items,omitempty"`
		Total int `json:"total,omitempty"`
	} `json:"accountRoles,omitempty"`
}

type AccountRole struct {
	ID                   string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Description             string   `json:"description,omitempty"`
	IsPredefined                string   `json:"isPredefined"`
}


func (c *Client) GetAccountRoles(accountID string, accountType string) (*AccountRolesQuery, error) {
	query := graphQLRequest{
		Query: `query accountRoles ($accountID: ID!, $accountType: AccountType!) {
			accountRoles (accountID: $accountID, accountType: $accountType) {
				items {
					id
					name
					description
				}
				total
			}
		}`,
		Variables: map[string]interface{}{
			"accountID": accountID,
			"accountType": accountType,
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
	accountroles := AccountRolesQuery{}
	json.Unmarshal(jsonData, &accountroles)
	return &accountroles, nil

}