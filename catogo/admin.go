package catogo

type Admin struct {
	ID                   string   `json:"id,omitempty"`
	FirstName            string   `json:"firstName,omitempty"`
	LastName             string   `json:"lastName,omitempty"`
	Email                string   `json:"email,omitempty"`
	CreationDate         string   `json:"creationDate,omitempty"`
	PasswordNeverExpires bool     `json:"passwordNeverExpires,omitempty"`
	MfaEnabled           bool     `json:"mfaEnabled,omitempty"`
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

func (c *httpClient) GetAdmins(accountID string) (interface{}, error) {
	query := graphQLRequest{
		Query: `query admins($accountID:ID!) {
			admins(accountID:$accountID) {
				items {
				  id
				  email
				  managedRoles {
					role {
					  name
					}
				  }
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

	return output, nil
}

func (c *httpClient) GetAdmin(accountID string, adminID string) (interface{}, error) {
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

	return output, nil
}

func (c *httpClient) AddAdmin(accountID string, input *Admin) (interface{}, error) {

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

	return output, nil

}

func (c *httpClient) UpdateAdmin(accountID string, adminID string, input *Admin) (interface{}, error) {

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

func (c *httpClient) RemoveAdmin(accountID string, adminID string) (interface{}, error) {

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