package api

import (
	db "github.com/BiPwL/quiz-base-api/db/sqlc"
	"github.com/BiPwL/quiz-base-api/util"
)

func randomUser() db.User {
	return db.User{
		ID:             util.RandomInt(1, 1000),
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomPasswordStr(10),
	}
}
