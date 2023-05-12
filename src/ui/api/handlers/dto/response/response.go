package response

import "github.com/google/uuid"

type ID struct {
	ID string `json:"id,omitempty"`
}

type idBuilder struct{}

func IDBuilder() *idBuilder {
	return &idBuilder{}
}

func (*idBuilder) FromID(id string) *ID {
	return &ID{id}
}

func (*idBuilder) FromUUID(id uuid.UUID) *ID {
	return &ID{id.String()}
}
