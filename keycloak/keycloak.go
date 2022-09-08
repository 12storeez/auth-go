package keycloak

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"strings"
)

type keycloak struct {
	basePath    string
	restyClient *resty.Client
}

func NewClient(hostname string) *keycloak {
	return &keycloak{
		basePath:    hostname,
		restyClient: resty.New(),
	}
}

func (k *keycloak) GetUserInfo(ctx context.Context, accessToken string, realm string) (*User, error) {
	userInfoUrl := strings.Join([]string{k.basePath, "auth", "realms", realm, "protocol", "openid-connect", "userinfo"}, "/")

	resp, err := k.restyClient.R().
		SetContext(ctx).
		SetAuthToken(accessToken).
		SetHeader("Content-Type", "application/json").
		Get(userInfoUrl)
	if err != nil {
		return nil, APIError{
			Code:    resp.StatusCode(),
			Message: err.Error(),
		}
	}
	defer resp.RawBody().Close()

	if resp.Body() == nil {
		return nil, APIError{
			Message: "empty response",
		}
	}

	if resp.IsError() {
		return nil, APIError{
			Code:    resp.StatusCode(),
			Message: resp.String(),
		}
	}

	var result User
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func GetAllUserRoles(accessToken string) (TokenRoles, error) {
	var hmacSampleSecret []byte
	var tokenRoles TokenRoles

	token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	// не обрабатываем err, потому err=invalid token, т.к. у нас нет ключа, чтобы его провалидировать
	if token == nil {
		return tokenRoles, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return tokenRoles, fmt.Errorf("invalid token")
	}

	realmRoleList := make(Roles)
	if realmAccess, ok := claims["realm_access"]; ok {
		switch realmAccess.(type) {
		case map[string]interface{}:
			realmRoleList = getRolesFromMap(realmAccess.(map[string]interface{}))
		}
	}

	var resourceRoleList ResourceRolesList
	if resourceAccess, ok := claims["resource_access"]; ok {
		switch resourceAccess.(type) {
		case map[string]interface{}:
			for resourceName, accessList := range resourceAccess.(map[string]interface{}) {
				switch accessList.(type) {
				case map[string]interface{}:
					roles := getRolesFromMap(accessList.(map[string]interface{}))
					resourceRole := ResourceRoles{
						resource: resourceName,
						roles:    roles,
					}
					resourceRoleList = append(resourceRoleList, &resourceRole)
				}
			}
		}
	}

	tokenRoles.RealmRoles = realmRoleList
	tokenRoles.ResourceRoles = resourceRoleList

	return tokenRoles, nil
}

func getRolesFromMap(accessMap map[string]interface{}) Roles {
	roleList := make(Roles)
	roles, ok := accessMap["roles"]
	if ok {
		switch roles.(type) {
		case []interface{}:
			for _, role := range roles.([]interface{}) {
				switch role.(type) {
				case string:
					roleList[role.(string)] = struct{}{}
				}
			}
		}
	}
	return roleList
}

func GetUserRolesByClientID(token, client string) (resRoles Roles, err error) {
	roles, err := GetAllUserRoles(token)
	if err != nil {
		return nil, err
	}

	for _, resourceRoles := range roles.ResourceRoles {
		if resourceRoles.resource == client {
			resRoles = resourceRoles.roles
			return
		}
	}

	return
}
