package multikmodel

type User struct {
	Id       int
	NickName string

	*multik.Controller
}

func (u *User) Id() {

}
