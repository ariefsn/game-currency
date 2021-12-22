package repositories_test

import (
	"ariefsn/game-currency/global/helper"
	"ariefsn/game-currency/services/currency-service/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseRepository(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		helper.MockDb()

		db := repositories.DB()

		assert.NotNil(t, db)
	})
}
