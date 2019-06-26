package repository

import (
	"app/domain/model"
	"testing"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/stretchr/testify/assert"
)

func initDBPublicAgent(runMigrations bool) *PublicAgentRepository {
	db, err := xorm.NewEngine("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	if runMigrations {
		RunMigrations(db)
	}

	return &PublicAgentRepository{DB: db}
}

func TestPublicAgentRepository_CreateOrUpdatePublicAgent(t *testing.T) {
	uRepo := initDBPublicAgent(true)
	defer uRepo.DB.Close()

	now := time.Now()

	newPublicAgent := model.PublicAgent{
		ID:         "1111",
		Name:       "test",
		Occupation: "Policial",
		Department: "Policia Civil",
		Salary:     5000.2,
		Checked:    now,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	updated, err := uRepo.CreateOrUpdatePublicAgent(&newPublicAgent)
	assert.Nil(t, err)
	assert.True(t, updated)

	publicAgent := model.PublicAgent{}
	ok, err := uRepo.DB.ID(newPublicAgent.ID).Get(&publicAgent)
	assert.Nil(t, err)
	assert.True(t, ok)
	newPublicAgent.Checked = publicAgent.Checked
	newPublicAgent.CreatedAt = publicAgent.CreatedAt
	newPublicAgent.UpdatedAt = publicAgent.UpdatedAt
	assert.Equal(t, publicAgent, newPublicAgent)
}
