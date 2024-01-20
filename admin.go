package catogo

import (
	"encoding/json"
)

type AdminsQuery struct {
	Admins struct {
		Items []Admin `json:"items,omitempty"`
		Total int `json:"total,omitempty"`
	} `json:"admins,omitempty"`
}

type AdminQuery struct {
	Admin Admin `json:"admin,omitempty"`
}

type AdminMutation struct {
	Admin struct {
		AddAdmin struct {
			AdminID string `json:"adminID,omitempty"`
		} `json:"addAdmin,omitempty"`
	} `json:"admin,omitempty"`
}

type Admin struct {
	ID                   string   `json:"id,omitempty"`
	FirstName            string   `json:"firstName,omitempty"`
	LastName             string   `json:"lastName,omitempty"`
	Email                string   `json:"email,omitempty"`
	CreationDate         string   `json:"creationDate,omitempty"`
	PasswordNeverExpires bool     `json:"passwordNeverExpires"`
	MfaEnabled           bool     `json:"mfaEnabled"`
	ManagedRoles         []Roles `json:"managedRoles,omitempty"`
	ResellerRoles        []Roles `json:"resellerRoles,omitempty"`
}

type AdminUpdate struct {
	FirstName            string   `json:"firstName,omitempty"`
	LastName             string   `json:"lastName,omitempty"`
	PasswordNeverExpires bool     `json:"passwordNeverExpires"`
	MfaEnabled           bool     `json:"mfaEnabled"`
	ManagedRoles         []Roles `json:"managedRoles,omitempty"`
	ResellerRoles        []Roles `json:"resellerRoles,omitempty"`
}

type Roles struct {
	Role Role `json:"role"`
}

type Role struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) GetAdmins(accountID string) (*AdminsQuery, error) {
	query := graphQLRequest{
		Query: `query admins($accountID:ID!) {
			admins(accountID:$accountID) {
				items {
				  id
				  lastName
				  firstName
				  email
				  passwordNeverExpires
				  mfaEnabled
				}
				total
			}
		}`,
		Variables: map[string]interface{}{
			"accountID": accountID,
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
	admins := AdminsQuery{}
	json.Unmarshal(jsonData, &admins)
	return &admins, nil

}

func (c *Client) GetAdmin(accountID string, adminID string) (*AdminQuery, error) {
	query := graphQLRequest{
		Query: `query admin($accountId: ID!, $adminID: ID!) {
			admin(accountId:$accountId, adminID:$adminID) {
				id
				firstName
				lastName
				email
				creationDate
				mfaEnabled
				managedRoles {
					role {
						name
					}
				}
				resellerRoles {
					role {
						name
					}
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": accountID,
			"adminID":	adminID,
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
	admin := AdminQuery{}
	json.Unmarshal(jsonData, &admin)

	return &admin, nil
}

func (c *Client) AddAdmin(accountID string, input *Admin) (*AdminMutation, error) {

	query := graphQLRequest{
		Query: `mutation addAdmin($accountId: ID!, $input: AddAdminInput!) {
			admin(accountId: $accountId) {
				addAdmin (input: $input) {
					adminID
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
	admin := AdminMutation{}
	json.Unmarshal(jsonData, &admin)

	return &admin, nil

}

func (c *Client) UpdateAdmin(accountID string, adminID string, input *AdminUpdate) (interface{}, error) {

	// if _, ok := input["Email"]; ok {
	// 	delete(input, "Email")
	// }

	query := graphQLRequest{
		Query: `mutation updateAdmin($accountId:ID!, $adminID:ID!, $input: UpdateAdminInput!) {
			admin(accountId: $accountId) {
				updateAdmin (adminID:$adminID,input:$input) {
					adminID
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": accountID,
			"adminID":     adminID,
			"input":     input,
		},
	}

	output := validResponse{}

	if err := c.post(query, &output); err != nil {
		return nil, err
	}

	return output, nil

}

func (c *Client) RemoveAdmin(accountID string, adminID string) (interface{}, error) {

	query := graphQLRequest{
		Query: `mutation removeAdmin($accountId: ID!, $adminID: ID!) {
			admin(accountId: $accountId) {
				removeAdmin (adminID: $adminID) {
					adminID
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": accountID,
			"adminID":     adminID,
		},
	}

	output := validResponse{}

	if err := c.post(query, &output); err != nil {
		return nil, err
	}

	return output, nil

}