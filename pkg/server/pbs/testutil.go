package pbs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	bookingpb "be/proto"
)

func ProtoEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	opts := cmpopts.IgnoreUnexported(
		bookingpb.GetListResponse{},
		bookingpb.GetBookingResponse{},
		bookingpb.Booking{},
		bookingpb.User{},
		bookingpb.Timeline{},
		bookingpb.Department{},
		bookingpb.Employee{},
		bookingpb.Notification{},
		bookingpb.BookingReport{},
		bookingpb.ReportPeriod{},
		bookingpb.FinalReport{},
		bookingpb.Outmember{},
		bookingpb.Comment{},
		bookingpb.FilterSlot{},
		bookingpb.Issue{},
	)

	assert.True(t, cmp.Equal(expected, actual, opts), cmp.Diff(expected, actual, opts))
}
