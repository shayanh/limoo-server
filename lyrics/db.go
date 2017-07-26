package lyrics

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var _session *mgo.Session

func InitDB(addr string) error {
	if _session != nil {
		return fmt.Errorf("DB has initiated before")
	}

	session, err := mgo.Dial(addr)
	if err != nil {
		return err
	}
	_session = session
	return nil
}

func CloseDB() {
	_session.Close()
}

func newSession() (*mgo.Session, error) {
	if _session == nil {
		return nil, fmt.Errorf("DB must be initiated before")
	}
	session := _session.Copy()
	return session, nil
}