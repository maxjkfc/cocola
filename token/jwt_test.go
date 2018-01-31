package token

import (
	"fmt"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Account string   `json:"account"`
	Name    string   `json:"name"`
	Email   []string `json:"email"`
	jwt.StandardClaims
}

func (u *User) Valid() error {
	if u.Issuer != "max" {
		return fmt.Errorf("Error: Wrong Issuer , %s", u.Issuer)
	}
	return nil
}

type UserMap = jwt.MapClaims

var (
	u1 *User
	u2 UserMap
	t1 string
	t2 string
	r1 *User
	r2 UserMap
)

func init() {
	SetKey("mm")
	SetIssuer("max")
	u1 = &User{
		Account: "a1",
		Name:    "a1Name",
		Email: []string{
			"a1@gmail.com",
			"a1@outlook.com",
		},
	}
	u1.Issuer = "max"
	u1.IssuedAt = time.Now().Unix()

	u2 = make(UserMap)

	u2["iss"] = "max"
	u2["account"] = "a2"
	u2["name"] = "a2name"
	u2["email"] = []string{
		"a2@gmail.com",
		"a2@outlook.com",
	}
	u2["ist"] = time.Now().Unix()
}

func Test_CreateJWTWithClaims(t *testing.T) {
	j := New().Create(u1)
	if j.Error() != nil {
		t.Errorf("Test_CreateJWTWithClaims Error: %v", j.Error())
	} else {
		t1 = j.Get()
		t.Logf("Test_CreateJWTWithClaims Sussess: %v", t1)
	}
}

func Test_CreateJWTWithMap(t *testing.T) {
	u2["ist"] = time.Now().Unix()
	j := New().Create(u2)
	if j.Error() != nil {
		t.Errorf("Test_CreateJWTWithMap Error: %v", j.Error())
	} else {
		t2 = j.Get()
		t.Logf("Test_CreateJWTWithMap Sussess: %v", t2)
	}
}

func Test_ParseJWTWithClaims(t *testing.T) {
	r1 = new(User)
	j := New().Parse(t1, r1)

	if j.Error() != nil {
		t.Errorf("Test_ParseJWTWithClaims Error: %v", j.Error())
	} else {
		t.Logf("Test_ParseJWTWithClaims Sussess: %v", r1)
	}
}

func Test_ParseJWTWithMap(t *testing.T) {
	r2 = make(UserMap)
	j := New().Parse(t2, r2)
	if j.Error() != nil {
		t.Errorf("Test_ParseJWTWithMap Error: %v", j.Error())
	} else {
		t.Logf("Test_ParseJWTWithMap Sussess: %v", r2)
	}
}
