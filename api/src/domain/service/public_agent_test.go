package service

import (
	"app/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDBPublicAgent struct {
	err error
}

func (r mockDBPublicAgent) CreateOrUpdatePublicAgent(publicAgent *model.PublicAgent) error {
	return r.err
}

func TestPublicAgentService_StartProcess(t *testing.T) {
	mock := mockDBPublicAgent{
		err: nil,
	}
	us := PublicAgentService{mock}
	err := us.StartProcess()
	assert.Nil(t, err)
}
