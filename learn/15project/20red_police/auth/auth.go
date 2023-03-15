package auth

type authList struct {
	IsNeedLogin bool
}

const (
	server_01 = "Server."
)

const (
	REGISTER   = "Register"
	LOGIN      = "Login"
	CREATEPMAP = "CreatePMap"
	USERLIST   = "UserList"
	LOGINOUT   = "LoginOut"
	DELETEROOM = "DeleteRoom"
)

var AuthListMapping map[string]*authList = map[string]*authList{
	server_01 + REGISTER:   &authList{IsNeedLogin: false},
	server_01 + LOGIN:      &authList{IsNeedLogin: false},
	server_01 + LOGINOUT:   &authList{IsNeedLogin: false},
	server_01 + CREATEPMAP: &authList{IsNeedLogin: true},
	server_01 + USERLIST:   &authList{IsNeedLogin: true},
	server_01 + DELETEROOM: &authList{IsNeedLogin: true},
}
