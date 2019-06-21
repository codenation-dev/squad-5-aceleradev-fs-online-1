package repository

import (
	"app/domain/model"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/stretchr/testify/assert"
)

func initDBAlert(runMigrations bool) *AlertRepository {
	db, err := xorm.NewEngine("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	if runMigrations {
		RunMigrations(db)
	}

	return &AlertRepository{DB: db}
}

func TestAlertRepository_GetAlert(t *testing.T) {
	uRepo := initDBAlert(true)
	defer uRepo.DB.Close()

	mockAlerts(uRepo.DB)

	newAlert, err := uRepo.GetAlert("1111")
	assert.Nil(t, err)
	assert.NotNil(t, newAlert)
	assert.Equal(t, "1111", newAlert.ID)
	assert.Equal(t, model.PublicAgentType, newAlert.Type)
	assert.Equal(t, "test", newAlert.Description)
	assert.Nil(t, newAlert.Customer)
	assert.Nil(t, newAlert.PublicAgent)
	assert.Nil(t, newAlert.User)
	assert.Nil(t, newAlert.UsersReceived)
	assert.NotNil(t, newAlert.CreatedAt)
	assert.NotNil(t, newAlert.UpdatedAt)
}

func TestAlertRepository_GetAlert_NotFound(t *testing.T) {
	uRepo := initDBAlert(true)
	defer uRepo.DB.Close()

	newUser, err := uRepo.GetAlert("1111")
	assert.Nil(t, err)
	assert.Nil(t, newUser)
}

func mockAlerts(DB *xorm.Engine) {
	u := model.Alert{
		ID:          "1111",
		Type:        model.PublicAgentType,
		Description: "test",
	}
	DB.InsertOne(u)
}
