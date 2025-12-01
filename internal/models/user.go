package models

type User struct {
	CommonName        string   `json:"commonName"`
	DistinguishedName string   `json:"distinguishedName"`
	DisplayName       string   `json:"displayName"`
	GivenName         string   `json:"givenName"`
	Surname           string   `json:"surname"`
	Email             string   `json:"email"`
	UserPrincipalname string   `json:"userPrincipalName"`
	Groups            []string `json:"groups"`

	AdditionalAttributes []map[string]string `json:"additionalAttributes"`
}
