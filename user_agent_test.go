package useragents

import (
	r "github.com/stretchr/testify/require"
	"testing"
)

const UserAgentsJson = "user_agents.json"

func TestNewUserAgents(t *testing.T) {
	u, err := New(UserAgentsJson)

	r.Nil(t, err)
	r.True(t, !u.IsEmpty())

	_, err = New("asd")
	r.NotNil(t, err)

	_, err = New("user_agents_bad.json")
	r.NotNil(t, err)
}

func TestUserAgents_GetRandom(t *testing.T) {
	u, err := New(UserAgentsJson)
	r.Nil(t, err)
	r.True(t, !u.IsEmpty())

	ua1 := u.GetRandom() // получим случайный useragent
	ua2 := u.GetRandom() // получим еще один случайный useragent
	ua3 := u.GetRandom() // получим еще один случайный useragent
	r.True(t, ua1 != ua2 || ua1 != ua3 || ua2 != ua3)
}
