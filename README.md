# auth-go

Пакет для работы с keycloak

- GetUserInfo(ctx context.Context, accessToken string, realm string) (User, error)
Получение информации по пользователю из keycloak. Проверка токена на подлинность и валидность.
realm - это область keycloak, в которой управляются объекты, включая пользователей, приложения, роли и группы.

- GetAllUserRoles(accessToken string) (tokenRoles TokenRoles, err error)
Получение всех ролей пользователя через декодирование токена.

- GetUserRolesByClientID(token, client string) (resRoles Roles, err error)
Получение ролей пользователя через декодирование токена, относящихся к конкретному переданному клиенту.

### Пример использования

1. Проверка токена на валидность и получение информации по пользователю:
```
	hostname - путь до keycloak
	realm - нужную область в keycloak
	accessToken - токен

	ctx := context.Background()
	client := keycloak.NewClient(hostname)
	info, err := client.GetUserInfo(ctx, accessToken, realm)
	if err != nil {
		fmt.Println("GetUserInfo failed:" + err.Error())
	}
```

2. Получение ролей пользователя для конкретного клиента из токена
```
	roles, err := keycloak.GetUserRolesByClientID(accessToken, "logistics")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(roles)
```
