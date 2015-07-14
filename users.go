package gogitlab

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const (
	users_url        = "/users"     // Get users list
	user_url         = "/users/:id" // Get a single user.
	current_user_url = "/user"      // Get current user
)

type User struct {
	Id            int    `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	Name          string `json:"name,omitempty"`
	State         string `json:"state,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	Bio           string `json:"bio,omitempty"`
	Skype         string `json:"skype,omitempty"`
	LinkedIn      string `json:"linkedin,omitempty"`
	Twitter       string `json:"twitter,omitempty"`
	ExternUid     string `json:"extern_uid,omitempty"`
	Provider      string `json:"provider,omitempty"`
	ThemeId       int    `json:"theme_id,omitempty"`
	ColorSchemeId int    `json:"color_scheme_id,color_scheme_id"`
}

func (g *Gitlab) Users() ([]*User, error) {

	url := g.ResourceUrl(users_url, nil)

	var users []*User

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &users)
	}

	return users, err
}

/* Get/Search for user by username */
func (g *Gitlab) UserByUsername(username string) (*User, error) {
	resourceUrl := g.ResourceUrl(users_url, nil)

	body := url.Values{"search": []string{username}}
	resourceUrl += "&" + body.Encode()

	contents, err := g.buildAndExecRequest("GET", resourceUrl, nil)
	if err != nil {
		return nil, err
	}

	var user []User
	if err = json.Unmarshal(contents, &user); err != nil {
		return nil, err
	}

	if len(user) < 1 {
		return nil, fmt.Errorf("User not found: %s", username)
	} else if len(user) > 1 {
		return nil, fmt.Errorf("Multiple users found: %s", contents)
	}
	return &user[0], err
}

/*
Get a single user.

    GET /users/:id

Parameters:

    id The ID of a user

Usage:

	user, err := gitlab.User("your_user_id")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", user)
*/
func (g *Gitlab) User(id string) (*User, error) {

	url := g.ResourceUrl(user_url, map[string]string{":id": id})

	user := new(User)

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, err
}

func (g *Gitlab) DeleteUser(id string) error {
	url := g.ResourceUrl(user_url, map[string]string{":id": id})
	var err error
	_, err = g.buildAndExecRequest("DELETE", url, nil)
	return err
}

func (g *Gitlab) CurrentUser() (User, error) {
	url := g.ResourceUrl(current_user_url, nil)
	var user User

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, err
}
