package keycloak

// https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims
type User struct {
	Sub                 string      `json:"sub,omitempty"`
	Name                string      `json:"name,omitempty"`
	GivenName           string      `json:"given_name,omitempty"`
	FamilyName          string      `json:"family_name,omitempty"`
	MiddleName          string      `json:"middle_name,omitempty"`
	Nickname            string      `json:"nickname,omitempty"`
	PreferredUsername   string      `json:"preferred_username,omitempty"`
	Profile             string      `json:"profile,omitempty"`
	Picture             string      `json:"picture,omitempty"`
	Website             string      `json:"website,omitempty"`
	Email               string      `json:"email,omitempty"`
	EmailVerified       bool        `json:"email_verified,omitempty"`
	Gender              string      `json:"gender,omitempty"`
	ZoneInfo            string      `json:"zoneinfo,omitempty"`
	Locale              string      `json:"locale,omitempty"`
	PhoneNumber         string      `json:"phone_number,omitempty"`
	PhoneNumberVerified bool        `json:"phone_number_verified,omitempty"`
	Address             UserAddress `json:"address,omitempty"`
	UpdatedAt           int         `json:"updated_at,omitempty"`
}

func (u *User) GetSub() string {
	if u != nil {
		return u.Sub
	}
	return ""
}

func (u *User) GetName() string {
	if u != nil {
		return u.Name
	}
	return ""
}

func (u *User) GetGivenName() string {
	if u != nil {
		return u.GivenName
	}
	return ""
}

func (u *User) GetFamilyName() string {
	if u != nil {
		return u.FamilyName
	}
	return ""
}

func (u *User) GetMiddleName() string {
	if u != nil {
		return u.MiddleName
	}
	return ""
}

func (u *User) GetNickname() string {
	if u != nil {
		return u.Nickname
	}
	return ""
}

func (u *User) GetPreferredUsername() string {
	if u != nil {
		return u.PreferredUsername
	}
	return ""
}

func (u *User) GetProfile() string {
	if u != nil {
		return u.Profile
	}
	return ""
}

func (u *User) GetPicture() string {
	if u != nil {
		return u.Picture
	}
	return ""
}

func (u *User) GetWebsite() string {
	if u != nil {
		return u.Website
	}
	return ""
}

func (u *User) GetEmail() string {
	if u != nil {
		return u.Email
	}
	return ""
}

func (u *User) GetEmailVerified() bool {
	if u != nil {
		return u.EmailVerified
	}
	return false
}

func (u *User) GetGender() string {
	if u != nil {
		return u.Gender
	}
	return ""
}

func (u *User) GetZoneInfo() string {
	if u != nil {
		return u.ZoneInfo
	}
	return ""
}

func (u *User) GetLocale() string {
	if u != nil {
		return u.Locale
	}
	return ""
}

func (u *User) GetPhoneNumber() string {
	if u != nil {
		return u.PhoneNumber
	}
	return ""
}

func (u *User) GetPhoneNumberVerified() bool {
	if u != nil {
		return u.PhoneNumberVerified
	}
	return false
}

func (u *User) GetAddress() UserAddress {
	if u != nil {
		return u.Address
	}
	return UserAddress{}
}

func (u *User) GetUpdatedAt() int {
	if u != nil {
		return u.UpdatedAt
	}
	return 0
}
