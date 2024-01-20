package catogo

func (c *Client) GetEntities(accountID string, entityType string) (interface{}, error) {

	query := graphQLRequest{
		Query: `query entityLookup ($accountID: ID!, $type: EntityType!) {
			entityLookup (accountID: $accountID, type: $type) {
				items {
					entity {
						id
						name
						type
					}
					description
					helperFields
				}
				total
			}
		}`,
		Variables: map[string]interface{}{
			"accountID": accountID,
			"type":      entityType,
		},
	}

	output := validResponse{}

	if err := c.post(query, &output); err != nil {
		return nil, err
	}

	return output, nil

}

func (c *Client) GetEntityByID(accountID string, entityType string, entityIDs string) (interface{}, error) {

	query := graphQLRequest{
		Query: `query entityLookup ($accountID: ID!, $type: EntityType!, $entityIDs: [ID!]) {
			entityLookup (accountID: $accountID, type: $type, entityIDs: $entityIDs) {
				items {
					entity {
						id
						name
						type
					}
					description
					helperFields
				}
				total
			}
		}`,
		Variables: map[string]interface{}{
			"accountID": accountID,
			"type":      entityType,
			"entityIDs": entityIDs,
		},
	}

	output := validResponse{}

	if err := c.post(query, &output); err != nil {
		return nil, err
	}

	return output, nil

}
