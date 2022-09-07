package keycloak

type TokenRoles struct {
	RealmRoles    Roles
	ResourceRoles ResourceRolesList
}

type ResourceRoles struct {
	resource string
	roles    Roles
}

type ResourceRolesList []ResourceRoles

type Roles map[string]struct{}
