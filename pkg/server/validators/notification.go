package validators

import (
	"github.com/pkg/errors"
	bookingpb "be/proto"

	"be/internal/model"
)

func Notification(notification *bookingpb.Notification) (model.Notification, error) {
	if notification == nil {
		return model.Notification{}, errors.New("slot must be not empty")
	}

	fragments := notification.GetFragments()

	if fragments > model.MaxFragment {
		return model.Notification{}, errors.New("fragments must be from 0 to 99")
	}

	return model.NewNotificationUnitsAndFragments(notification.GetUnits(), fragments), nil
}
