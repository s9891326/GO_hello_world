package note

import "fmt"

func Map() {
	var m1 map[string]string
	fmt.Println("m1 == nil?:", m1 == nil)
	m1 = make(map[string]string, 2)
	m1["1"] = "1111"
	m1["2"] = "222"
	m1["3"] = "33333"
	fmt.Println("m1=", m1)

	m2 := map[string]string{
		"1": "11111",
		"2": "22222",
	}
	fmt.Println("m2=", m2, ",m2[1]=", m2["1"])

	v, ok := m2["23"]
	if ok {
		fmt.Println("v=", v)
	} else {
		fmt.Println("不存在")
	}

	delete(m1, "1")
	fmt.Println("m1=", m1)

	// 清除整個map
	//m1 = nil
	m2 = make(map[string]string)
	fmt.Println("m2=", m2)

	for key, value := range m1 {
		fmt.Printf("m1[%v]=%v\n", key, value)
	}
}

func TypeDefinitionAndTypeAlias() {
	fmt.Println("4.4.1 自定義數據")
	type mesType uint16
	var u100 uint16 = 100
	var textMessage mesType = mesType(u100) // mesType(100)
	fmt.Printf("textMessage=%v, type=%T\n", textMessage, textMessage)

	fmt.Println("4.4.2 類型別名")
	type myUnit16 = uint16
	var myu16 myUnit16 = u100
	fmt.Printf("myu16=%v, type=%T\n", myu16, myu16)
}

// 4.5 結構體
type User struct {
	Name string `json:"name"`
	Id   uint32
}
type Account struct {
	User
	password string
}
type Contact struct {
	*User
	Remark string
}

func Struct() {
	var u1 User = User{
		Name: "eddy",
	}
	u1.Id = 1000
	fmt.Println("u1=", u1)
	var u2 *User = &User{
		Name: "eddy2",
		Id:   222,
	}
	fmt.Println("u2=", u2)

	a1 := Account{
		User: User{
			Name: u1.Name,
		},
		password: "666",
	}

	var c1 *Contact = &Contact{
		User:   &User{},
		Remark: "no limit",
	}
	c1.Name = "5555" // = c1.User.Name = "5555"

	fmt.Println("a1=", a1)
	fmt.Printf("c1=%v,type=%T\n", c1, c1)
	fmt.Println("c1.User=", c1.User)
	fmt.Println("c1.User=", *((*c1).User))

}

// 5.1 方法
func (u User) printName() {
	fmt.Println("u.Name=", u.Name)
}

func (u *User) setId() {
	fmt.Println("u.setid=", u)
	(*u).Id = 1000 // = u.Id = 1000
}

func Method() {
	u := User{
		Name: "eddy",
	}
	u.printName()
	u.setId()
	fmt.Println("Method u=", u)
}

func main20() {
	fmt.Println("123")
	//Map()
	//TypeDefinitionAndTypeAlias()
	//Struct()
	Method()
}
