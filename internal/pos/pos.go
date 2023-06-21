package pos

import (
	"fmt"
	"github.com/spongeling/admin-api/internal/errors"
	"strings"
)

type POS struct {
	Category Category
	Values   map[Field]rune
}

func (x *POS) String() string {
	s := []byte{byte(x.Category.Identifier)}
	for _, field := range x.Category.Fields {
		s = append(s, byte(x.Values[field]))
	}
	return string(s)
}

func (x *POS) Parse(s string) error {
	errInvalidString := errors.New(errors.InvalidArgument, "string is not of valid for the given category")

	if rune(s[0]) != x.Category.Identifier {
		return errInvalidString
	}
	if x.Category.Identifier == Punctuation.Identifier {
		// Special case. Can be two or three characters
		if len(s) < 2 || len(s) > 3 {
			return errInvalidString
		}
	} else if len(s) != 1+len(x.Category.Fields) {
		return errInvalidString
	}

	x.Values = make(map[Field]rune)
	for i, c := range s[1:] {
		x.Values[x.Category.Fields[i]] = c
	}

	return nil
}

func (x *POS) GetConditions() string {
	conditions := []string{fmt.Sprintf(`"category"=%d`, x.Category.Identifier)}
	for field, value := range x.Values {
		conditions = append(conditions, fmt.Sprintf(`"%s"=%d`, field, value))
	}

	return strings.Join(conditions, " AND ")
}

func FromString(s string) (*POS, error) {
	err := errors.New(errors.InvalidArgument, "string is not of valid")
	if s == "" {
		return nil, err
	}

	cat := GetCategory(rune(s[0]))
	if cat.IsEmpty() {
		return nil, err
	}

	if len(s) != 1+len(cat.Fields) {
		return nil, err
	}

	pos := POS{Category: cat}

	err = pos.Parse(s)
	if err != nil {
		return nil, err
	}

	return &pos, nil
}
