package mock_server

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// Client interface
type Client interface {
	GetUser(id string) (*User, error)
	GreetUser(id string) (string, error)
}

type client struct {
	http.Client
	baseURL string
}

type User struct {
	Name  string
	Email string
}

// New client with defaults
func New(url string) Client {
	return &client{
		http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
		url,
	}
}

// GetUser returns a user
func (c *client) GetUser(id string) (*User, error) {
	req, err := http.NewRequest("GET", c.baseURL+"/user?id="+id, nil)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal("error closing response body")
		}
	}()

	var user *User
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

// GreetUser adds a greet message for user
func (c *client) GreetUser(id string) (string, error) {
	req, err := http.NewRequest("GET", c.baseURL+"/greet/user/"+id, nil)

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal("error closing response body")
		}
	}()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		log.Fatal("err: ", err)

	}
	return buf.String(), nil
}
