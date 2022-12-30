package roles

const (
	RoleUsersManager    = "ROLE_USERS_MNG"
	RoleMachinesManager = "ROLE_MACHINES_MNG"
)

type Roles struct {
	list []string `json:"list" bson:"list"`
}

func (u *Roles) Has(role string) bool {
	for _, v := range u.list {
		if v == role {
			return true
		}
	}
	return false
}

func (u *Roles) HasAll(roles ...string) bool {
	for _, r := range roles {
		if !u.Has(r) {
			return false
		}
	}
	return true
}

func (u *Roles) Set(roles ...string) {
	u.list = roles
}

func (u *Roles) Add(roles ...string) {
	u.list = append(u.list, roles...)
}

func (u *Roles) Delete(role string) {
	roles := []string{}
	for _, value := range u.list {
		if value != role {
			roles = append(roles, value)
		}
	}
	u.list = roles
}

func (u *Roles) Reset() {
	u.list = []string{}
}

func (u *Roles) Length() int {
	return len(u.list)
}

func (r *Roles) StringsSlice() (res []string) {
	res = make([]string, len(r.list))
	copy(res, r.list)
	return
}
