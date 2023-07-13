package token

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/nats-io/nuid"
)

type TokenID struct {
	NUID    string
	Counter int
}

func NewTokenID() *TokenID {
	return &TokenID{NUID: nuid.Next(), Counter: 0}
}

func NewTokenIDFromString(s string) (*TokenID, error) {
	t := NewTokenID()
	err := t.FromString(s)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TokenID) String() string {
	return fmt.Sprintf("%s.%d", t.NUID, t.Counter)
}

func (t *TokenID) FromString(s string) error {
	strs := strings.Split(s, ".")
	if len(strs) != 2 {
		return fmt.Errorf("invalid token format %s", s)
	}
	c, err := strconv.Atoi(strs[1])
	if err != nil {
		return fmt.Errorf("invalid token format %s, %v", s, err)
	}
	t.NUID = strs[0]
	t.Counter = c
	return nil
}

func (t *TokenID) Rotate() *TokenID {
	t.Counter++
	return t
}

func (t *TokenID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return t.FromString(s[1:])
}

func (t *TokenID) MarshalJSON() ([]byte, error) {
	s := "\"" + t.String() + "\""

	return []byte(s), nil
}
