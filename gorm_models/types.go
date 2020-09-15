package gorm_models

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

// ID
func MarshalID(id uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

func UnmarshalID(v interface{}) (uint, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ids must be strings")
	}
	i, e := strconv.Atoi(id)
	return uint(i), e
}
