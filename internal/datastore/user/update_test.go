package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"be/internal/datastore/base"
	"be/internal/datastore/testutil"
	userStorage "be/internal/datastore/user"
	"be/internal/lib"
	"be/internal/model"
)

func TestDB_UpdateUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r, err := testutil.SetupTestDataBaseWithFixtures(t)
	require.NoError(t, err)

	store := userStorage.New(r.DB)

	existedUser, err := store.FindUserByPK(ctx, testutil.ExistedUserID)
	require.NoError(t, err)

	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr error
	}{
		{
			name: "updates user",
			args: args{
				user: model.User{
					ID:        existedUser.ID,
					CreatedAt: existedUser.CreatedAt,
					UpdatedAt: time.Time{},
					ProfileID: "u_" + existedUser.ProfileID,
					Email:     "u_" + existedUser.Email,
					Phone:     "90" + existedUser.Phone,
					Role:      existedUser.Role,
					Employee: model.Employee{
						FirstName:  lib.String("new_first_name"),
						MiddleName: lib.String("new_middle_name"),
						LastName:   lib.String("new_last_name"),
						Avatar:     lib.String("new_avatar"),
						Email:      lib.String("new_email"),
						Phone:      lib.String("new_phone"),
						PortalCode: 2 + existedUser.Employee.PortalCode,
					},
				},
			},
			want: &model.User{
				ID:        existedUser.ID,
				CreatedAt: existedUser.CreatedAt,
				UpdatedAt: time.Time{},
				ProfileID: "u_" + existedUser.ProfileID,
				Email:     "u_" + existedUser.Email,
				Phone:     "90" + existedUser.Phone,
				Role:      existedUser.Role,
				Employee: model.Employee{
					FirstName:  lib.String("new_first_name"),
					MiddleName: lib.String("new_middle_name"),
					LastName:   lib.String("new_last_name"),
					Avatar:     lib.String("new_avatar"),
					Email:      lib.String("new_email"),
					Phone:      lib.String("new_phone"),
					PortalCode: 2 + existedUser.Employee.PortalCode,
				},
			},
			wantErr: nil,
		},
		{
			name: "return error if user does not exists",
			args: args{
				user: model.User{
					ID: lib.UUID(),
				},
			},
			wantErr: base.ErrNotFound,
		},
	}
	for _, tst := range tests {
		tt := tst

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, gotErr := store.UpdateUser(ctx, tt.args.user)
			if tt.wantErr != nil {
				require.Error(t, gotErr)
				assert.EqualError(t, tt.wantErr, gotErr.Error())
				return
			}

			require.NoError(t, gotErr)
			assertUser(t, *tt.want, *got, "UpdatedAt")
			assertDBUser(t, store, *tt.want, "UpdatedAt")
		})
	}
}
