package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"
)

func AssertProtoErrorEqual(t *testing.T, want status.Status, got error) {
	t.Helper()

	if got == nil {
		assert.Fail(t, "got is nil")
		return
	}

	gotStatus, ok := status.FromError(got)
	require.Truef(t, ok, "%s is not grpc error", got.Error())

	assert.EqualValuesf(t,
		want.Code(), gotStatus.Code(),
		"want=%s, got=%s", want.Code().String(), gotStatus.Code().String())
	assert.EqualValues(t, want.Message(), gotStatus.Message())
}
