package pbs

import (
	bookingpb "be/proto"

	"be/internal/model"
)

func BookingReport(report model.ReportBooking) *bookingpb.BookingReport {
	return &bookingpb.BookingReport{
		Id: report.ID.String(),
		ReportPeriod: &bookingpb.ReportPeriod{
			Month: report.Period.Month,
			Year:  report.Period.Year,
		},
		Slot:  toNotification(report.GetSlot()),
		EndDate: ToUTCString(report.GetEndAt()),
		Events:  report.GetEvents(),
		Reasons: report.GetReasons(),
		Comment: report.GetComment(),
		Status:  bookingpb.ReportStatus(report.Status),
	}
}

func BookingReportsList(reports []model.ReportBooking) []*bookingpb.BookingReport {
	res := make([]*bookingpb.BookingReport, 0, len(reports))
	for _, report := range reports {
		res = append(res, BookingReport(report))
	}

	return res
}

func RoomReport(report model.ReportRoom) *bookingpb.RoomReport {
	return &bookingpb.RoomReport{
		Uuid:     report.ID.String(),
		Comment:  *report.Comment,
		Timeline: timelinePtr(&report.Timeline),
		Slots:  PbSlotList(report.Slots),
		Equipments:  PbEquipmentList(report.Equipments),
		Releases: PbRoomReleaseList(report.Releases),
	}
}

func RoomReportsList(reports []model.ReportRoom) []*bookingpb.RoomReport {
	res := make([]*bookingpb.RoomReport, 0, len(reports))
	for _, report := range reports {
		res = append(res, RoomReport(report))
	}

	return res
}
