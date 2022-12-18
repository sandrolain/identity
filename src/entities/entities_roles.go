package entities

const (
	ROLE_USERS_MANAGER    = "ROLE_USERS_MNG"
	ROLE_SESSIONS_MANAGER = "ROLE_SESSIONS_MNG"
	ROLE_KEYS_MANAGER     = "ROLE_KEYS_MNG"
)

func (u *Entity) GetRoles() EntityRoles {
	return u.Roles
}

func (u *Entity) getRolesMap() map[EntityRole]bool {
	r := make(map[EntityRole]bool)
	for _, value := range u.Roles {
		r[value] = true
	}
	return r
}

func (u *Entity) HasRole(role EntityRole) bool {
	m := u.getRolesMap()
	_, ok := m[role]
	return ok
}

func (u *Entity) HasRoles(roles EntityRoles) bool {
	m := u.getRolesMap()
	for _, r := range roles {
		_, ok := m[r]
		if !ok {
			return false
		}
	}
	return true
}

func (u *Entity) SetRoles(roles EntityRoles) {
	u.Roles = roles
}

func (u *Entity) AddRoles(roles EntityRoles) {
	u.Roles = append(u.Roles, roles...)
}

func (u *Entity) DeleteRole(role EntityRole) {
	roles := make(EntityRoles, 0)
	for _, value := range u.Roles {
		if value != role {
			roles = append(roles, value)
		}
	}
	u.Roles = roles
}

func (u *Entity) ResetRoles() {
	u.Roles = make(EntityRoles, 0)
}
