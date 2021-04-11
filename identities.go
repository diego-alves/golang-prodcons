package main

import "encoding/json"

type ExternalIdentities struct {
	TotalCount int
	PageInfo   struct {
		HasNextPage bool
		EndCursor   string
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

const query = `query($org: String!, $after: String) {
	organization(login: $org) {
	  samlIdentityProvider {
		externalIdentities(first: 100, after: $after) {
		  totalCount
		  pageInfo {
			startCursor
			hasNextPage
			endCursor
		  }
		  edges {
			cursor
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

func GetExternalIdentities(org string, after string) ExternalIdentities {
	var data map[string]map[string]map[string]map[string]ExternalIdentities
	vars := map[string]interface{}{
		"org": org,
	}

	if after != "" {
		vars["after"] = after
	}

	res := RequestE(query, vars)

	json.Unmarshal(res, &data)
	return data["data"]["organization"]["samlIdentityProvider"]["externalIdentities"]
}
