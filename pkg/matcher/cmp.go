package matcher

import (
	"fmt"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"be/internal/model"
)

type CmpMatcher struct {
	i interface{}
}

func NewCmpMatcher(tx interface{}) gomock.Matcher {
	return CmpMatcher{tx}
}

func (d CmpMatcher) Matches(got interface{}) bool {
	opts := cmp.Options{
		cmpopts.IgnoreUnexported(time.Time{}, model.Attachment{}),
		cmpopts.EquateEmpty(),
	}

	isEq := cmp.Equal(d.i, got, opts...)
	if !isEq {
		fmt.Println(isEq, cmp.Diff(d.i, got, opts...)) //nolint:forbidigo
	}

	return isEq
}

func (d CmpMatcher) String() string {
	return fmt.Sprintf("%v", d.i)
}
