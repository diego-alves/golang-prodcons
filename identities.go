package main

import "encoding/json"

type ExternalIdentities struct {
	TotalCount int
	PageInfo   struct {
		HasNextPage bool
	}
	Edges []struct {
		Node struct {
			SamlIdentity struct {
				NameId string
			}
			User struct {
				Login string
			}
		}
	}
}

const query = `query($org: String!) {
    organization(login: $org) {
        samlIdentityProvider {
            externalIdentities(first: 100) {
                totalCount
                pageInfo {
                    hasNextPage
                }
                edges {
                    node {
                        samlIdentity {
                            nameId
                        }
                        user {
                            login
                        }
                    }
                }
            }
        }
    }
}`

func GetExternalIdentities(org string) ExternalIdentities {
	vars := map[string]interface{}{
		"org": org,
	}
	res := RequestE(query, vars)
	var data map[string]map[string]map[string]map[string]ExternalIdentities
	json.Unmarshal(res, &data)
	return data["data"]["organization"]["samlIdentityProvider"]["externalIdentities"]

}
