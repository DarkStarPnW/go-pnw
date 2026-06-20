package pnw

import "strings"

// APIError is returned when the GraphQL response contains one or more errors.
type APIError []graphQLError

func (e APIError) Error() string {
	msgs := make([]string, len(e))
	for i, err := range e {
		msgs[i] = err.Message
	}
	return "pnw: api error: " + strings.Join(msgs, "; ")
}
