package uuid

import (
	"github.com/google/uuid"
	"github.com/mkaiho/go-graphql-sample/usecase/gateway"
)

var _ gateway.IDManager

type UUIDManager struct{}

func (m *UUIDManager) Generate() string {
	return uuid.New().String()
}
