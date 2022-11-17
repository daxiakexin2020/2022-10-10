package service

import "path/filepath"

func FilePathJoin() string {
	//filepath.Join  将路径组合起来   c/  b/  a/ =>   c/b/a
	return filepath.Join("./")
}

func TmpTest() string {
	var tokenString []string
	if len(tokenString) > 0 && tokenString[0] != "" {
		return tokenString[0]
	}
	return "not exists"
}
