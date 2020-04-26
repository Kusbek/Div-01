package user

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/lithammer/shortuuid"
)

type Cookie struct {
	User       *User
	ExpireTime time.Time
}

type Cookies struct {
	Cookies map[string]*Cookie
	mulock  *sync.Mutex
}

var cookies *Cookies

const (
	COOKIEEXPIRETIME = 600
)

func NewCookies() {
	cookies = &Cookies{Cookies: make(map[string]*Cookie), mulock: new(sync.Mutex)}
	go cookies.monitor()
}

func (c *Cookies) monitor() {
	for {
		// fmt.Println(c.Cookies)
		time.Sleep(2 * time.Second)
		for key, value := range c.Cookies {
			if value.ExpireTime.Before(time.Now()) {
				c.mulock.Lock()
				delete(c.Cookies, key)
				c.mulock.Unlock()
			}
		}
	}
}

func GetCookies() *Cookies {
	return cookies
}

func (c *Cookies) Insert(user *User) string {
	c.mulock.Lock()
	defer c.mulock.Unlock()
	uuid := genShortUUID()
	cookie := &Cookie{User: user, ExpireTime: time.Now().Add(COOKIEEXPIRETIME * time.Second)}
	c.Cookies[uuid] = cookie
	return uuid
}

func (c *Cookies) CheckCookie(token string) (*User, error) {
	cookie, ok := c.Cookies[token]
	if !ok {
		return nil, fmt.Errorf("Not Authorized")
	}
	return cookie.User, nil
}

func genShortUUID() string {
	id := shortuuid.New()
	return id
}

func GetCookieFromRequest(req *http.Request) (string, error) {
	c, err := req.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", fmt.Errorf("Cookie is not provided")
		}
		return "", err
	}
	sessionToken := c.Value
	return sessionToken, nil
}

func Authenticate(req *http.Request) (*User, error) {
	c := GetCookies()
	token, err := GetCookieFromRequest(req)
	if err != nil {
		return nil, err
	}
	u, err := c.CheckCookie(token)
	if err != nil {
		return nil, err
	}
	return u, nil
}
