package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

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

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}
