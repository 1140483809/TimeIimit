package models

import "log"

//登录验证
func CheckLogin(username string,password string)string  {
	if _,ok := Usermap[username]; !ok {
		return "用户不存在"
	}
	user := User{
		Username: username,
		Password: password,
	}
	if Usermap[username] != user {
		return "密码错误"
	}
	return "登录成功"
}

//注册
func RegisterUser(username string,password string)string{
	user := User{
		Username: username,
		Password: password,
	}
	if _,ok := Usermap[user.Username]; ok{
		return "用户已存在"
	}
	//写入数据库
	_,err := EngineSql.Insert(user)
	if err != nil {
		log.Fatal(err.Error())
		return "写入失败"
	}
	Usermap[user.Username] = user
	return "注册成功"
}
