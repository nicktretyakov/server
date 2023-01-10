package user_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	userStorage "be/internal/datastore/user"
	"be/internal/model"
)

func assertDBUser(t *testing.T, store *userStorage.Storage, want model.User, ignore ...string) {
	t.Helper()

	got, err := store.FindUserByPK(context.Background(), want.ID)
	require.NoError(t, err)

	assertUser(t, want, *got, ignore...)
}

func assertUser(t *testing.T, want, got model.User, ignore ...string) {
	t.Helper()
	opts := cmpopts.IgnoreFields(model.User{}, ignore...)

	assert.True(t, cmp.Equal(want, got, opts), cmp.Diff(want, got, opts))
}
