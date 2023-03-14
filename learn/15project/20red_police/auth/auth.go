package auth

type authList struct {
	IsNeedLogin bool
}

const (
	server_01 = "Server."
)

const (
	REGISTER    = "Register"
	LOGIN       = "Login"
	CREATE_PMAP = "CreatePMap"
)

var AuthListMapping map[string]*authList = map[string]*authList{
	server_01 + REGISTER:    &authList{IsNeedLogin: false},
	server_01 + LOGIN:       &authList{IsNeedLogin: false},
	server_01 + CREATE_PMAP: &authList{IsNeedLogin: true},
}
