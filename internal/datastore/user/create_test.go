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
	"be/internal/model"
)

func TestDB_CreateUser(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	r, err := testutil.SetupTestDataBaseWithFixtures(t)
	require.NoError(t, err)

	store := userStorage.New(r.DB)

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
			name: "creates user",
			args: args{
				user: model.User{
					ProfileID: "profile_id",
					Email:     "some@gmail.com",
					Phone:     "79990001010",
					Role:      model.Regular,
					Employee: model.Employee{
						FirstName:  lib.String("first_name"),
						MiddleName: lib.String("middle_name"),
						LastName:   lib.String("last_name"),
						Avatar:     lib.String("avatar"),
						Email:      lib.String("email"),
						Phone:      lib.String("phone"),
						PortalCode: 111111,
					},
				},
			},
			want: &model.User{
				ProfileID: "profile_id",
				Email:     "some@gmail.com",
				Phone:     "79990001010",
				Role:      model.Regular,
				Employee: model.Employee{
					FirstName:  lib.String("first_name"),
					MiddleName: lib.String("middle_name"),
					LastName:   lib.String("last_name"),
					Avatar:     lib.String("avatar"),
					Email:      lib.String("email"),
					Phone:      lib.String("phone"),
					PortalCode: 111111,
				},
			},
			wantErr: nil,
		},
		{
			name: "creates user if profile id is empty",
			args: args{
				user: model.User{
					ProfileID: "",
					Email:     "some@gmail.com",
					Phone:     "79990001010",
					Role:      model.Regular,
					Employee: model.Employee{
						FirstName:  lib.String("first_name"),
						MiddleName: lib.String("middle_name"),
						LastName:   lib.String("last_name"),
						Avatar:     lib.String("avatar"),
						Email:      lib.String("email"),
						Phone:      lib.String("phone"),
						PortalCode: 111112,
					},
				},
			},
			want: &model.User{
				ProfileID: "",
				Email:     "some@gmail.com",
				Phone:     "79990001010",
				Role:      model.Regular,
				Employee: model.Employee{
					FirstName:  lib.String("first_name"),
					MiddleName: lib.String("middle_name"),
					LastName:   lib.String("last_name"),
					Avatar:     lib.String("avatar"),
					Email:      lib.String("email"),
					Phone:      lib.String("phone"),
					PortalCode: 111112,
				},
			},
			wantErr: nil,
		},
		{
			name: "return portal code constraint error",
			args: args{
				user: model.User{
					ProfileID: "901234_profile_id",
					Email:     "some@gmail.com",
					Phone:     "79990001010",
					Role:      model.Regular,
					Employee:  model.Employee{PortalCode: 901234},
				},
			},
			wantErr: base.ErrUserPortalCodeAlreadyExists,
		},
	}
	for _, tst := range tests {
		tt := tst

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, gotErr := store.CreateUser(ctx, tt.args.user)
			if tt.wantErr != nil {
				require.Error(t, gotErr)
				assert.EqualError(t, tt.wantErr, gotErr.Error())
				return
			}

			require.NoError(t, gotErr)
			assertUser(t, *tt.want, *got, "ID", "CreatedAt", "UpdatedAt")
			assertDBUser(t, store, *got)
		})
	}
}
