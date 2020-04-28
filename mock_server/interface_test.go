package mock_server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	userId := "666"

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("mock: path: ", r.URL.Path)

			if r.URL.String() == "/user?id="+userId {
				w.Header().Add("Content-Type", "application/json")
				_, err := w.Write([]byte(`{"name": "Maka", "email": "developer@jhonmike.com.br"}`))
				if err != nil {
					log.Fatal("err: ", err)
				}
			}

			if r.URL.Path == "/greet/user/"+userId {
				w.Header().Add("Content-Type", "application/json")
				_, err := w.Write([]byte(`Hello ` + userId))
				if err != nil {
					log.Fatal("err: ", err)
				}
			}
		}),
	)
	defer ts.Close()

	client := New(ts.URL)

	user, err := client.GetUser(userId)

	assert.NoError(t, err)
	assert.Equal(t, "Maka", user.Name)
	assert.Equal(t, "developer@jhonmike.com.br", user.Email)

	msg, err := client.GreetUser(userId)
	assert.NoError(t, err)
	assert.Equal(t, "Hello 66", msg)

}
