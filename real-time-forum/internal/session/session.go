package session

import (
	"DIV-01/real-time-forum/internal/model"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/lithammer/shortuuid"
)

//Cookie ...
type Cookie interface {
	Insert(*model.User) string
	Check(string) (*model.User, error)
	Delete(string)
}

type session struct {
	user       *model.User
	expireTime time.Time
}

type cookie struct {
	expiration time.Duration
	sessions   map[string]*session
	mu         *sync.Mutex
}

//New ...
func New() Cookie {
	c := &cookie{
		expiration: 10,
		sessions:   make(map[string]*session),
		mu:         &sync.Mutex{},
	}
	go c.monitor()
	return c
}

func (c *cookie) Insert(u *model.User) string {
	uuid := shortuuid.New()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sessions[uuid] = &session{user: u, expireTime: time.Now().Add(c.expiration * time.Hour)}
	return uuid
}

func (c *cookie) Check(uuid string) (*model.User, error) {
	session, ok := c.sessions[uuid]
	if !ok {
		return nil, errors.New("Unauthorized")
	}
	return session.user, nil
}

func (c *cookie) Delete(uuid string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.sessions, uuid)
}

func (c *cookie) monitor() {
	for {
		// fmt.Println("cookie monitoring!!!")
		time.Sleep(1 * time.Second)
		fmt.Println(c.sessions)
		for key, value := range c.sessions {

			if value.expireTime.Before(time.Now()) {
				c.mu.Lock()
				delete(c.sessions, key)
				c.mu.Unlock()
			}
		}
	}
}
