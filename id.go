package pnw

import (
	"strconv"
)

// ID is an integer identifier that the PnW API serialises as a JSON string.
type ID int

func (id *ID) UnmarshalJSON(b []byte) error {
	// Strip surrounding quotes if present, then parse.
	s := string(b)
	if len(s) >= 2 && s[0] == '"' {
		s = s[1 : len(s)-1]
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*id = ID(n)
	return nil
}

func (id ID) Int() int       { return int(id) }
func (id ID) String() string { return strconv.Itoa(int(id)) }
