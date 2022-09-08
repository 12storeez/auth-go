package keycloak

// https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
type UserAddress struct {
	Formatted     string `json:"formatted,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
	Locality      string `json:"locality,omitempty"`
	Region        string `json:"region,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Country       string `json:"country,omitempty"`
}

func (a *UserAddress) GetFormatted() string {
	if a != nil {
		return a.Formatted
	}
	return ""
}

func (a *UserAddress) GetStreetAddress() string {
	if a != nil {
		return a.StreetAddress
	}
	return ""
}

func (a *UserAddress) GetLocality() string {
	if a != nil {
		return a.Locality
	}
	return ""
}

func (a *UserAddress) GetRegion() string {
	if a != nil {
		return a.Region
	}
	return ""
}

func (a *UserAddress) GetPostalCode() string {
	if a != nil {
		return a.PostalCode
	}
	return ""
}

func (a *UserAddress) GetCountry() string {
	if a != nil {
		return a.Country
	}
	return ""
}
