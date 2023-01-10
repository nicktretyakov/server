package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"be/internal/datastore/base"
	"be/internal/datastore/testutil"
	userStorage "be/internal/datastore/user"
	"be/internal/lib"
)

func TestDB_FindUserByProfileID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r, err := testutil.SetupTestDataBaseWithFixtures(t)
	require.NoError(t, err)

	store := userStorage.New(r.DB)
	fixtures := testutil.NewFixtureStore(t)

	want, err := fixtures.UserByID(testutil.ExistedUserID)
	require.NoError(t, err)

	t.Run("return user", func(t *testing.T) {
		t.Parallel()

		got, err := store.FindUserByPK(ctx, testutil.ExistedUserID)
		require.NoError(t, err)

		assertUser(t, want, *got, "CreatedAt", "UpdatedAt")
	})

	t.Run("return user", func(t *testing.T) {
		t.Parallel()

		got, err := store.FindUserByPortalCode(ctx, want.Employee.PortalCode)
		require.NoError(t, err)
		assertUser(t, want, *got, "CreatedAt", "UpdatedAt")
	})

	t.Run("return pgx.ErrNoRows", func(t *testing.T) {
		t.Parallel()

		got, err := store.FindUserByPortalCode(ctx, 912222)
		require.Nil(t, got)
		require.Error(t, err)

		assert.ErrorIs(t, err, base.ErrNotFound)
	})

	t.Run("return pgx.ErrNoRows", func(t *testing.T) {
		t.Parallel()

		got, err := store.FindUserByPK(ctx, lib.UUID())
		require.Nil(t, got)
		require.Error(t, err)

		assert.ErrorIs(t, err, base.ErrNotFound)
	})
}
