package shield

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shield/domain"
)

var userByname map[string]domain.User

func init() {
	fmt.Println("hello init")
	userByname = make(map[string]domain.User)
	userByname["hello"] = domain.User{"hello", 21}
	userByname["world"] = domain.User{"world", 23}

}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		users := getUser()
		user := domain.User{"liguolin", 21}
		users.Users = append(users.Users, user)
		helloUser := userByname["hello"]
		users.Users = append(users.Users, helloUser)
		bytes, _ := json.Marshal(users)

		writer.Write(bytes)

	})

	http.ListenAndServe(":8888", nil)
}

func getUser() domain.Users {
	var s domain.Users
	s.Users = append(s.Users, domain.User{"smile", 21})
	s.Users = append(s.Users, domain.User{"GG", 33})
	return s
}
