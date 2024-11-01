package basic

import (
	"fmt"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	user := &basic.User{Id: 1, UserName: "Imran Ali", IsActive: true}
	user.Password = []byte("mySecret")
	user.Gender = basic.Gender_MALE
	user.Emails = []string{"imran@ali", "imran@ali2"}
	fmt.Println(user.String())
	fmt.Println(user.String())
	fmt.Println(user.Gender.String())
}

func ProtoToJson() {
	user := &basic.User{Id: 1, UserName: "Imran Ali", IsActive: true}
	user.Password = []byte("mySecret")
	user.Gender = basic.Gender_MALE
	user.Emails = []string{"imran@ali", "imran@ali2"}
	json, _ := protojson.Marshal(user)
	fmt.Println(string(json))
}
