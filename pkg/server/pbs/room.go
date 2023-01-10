package pbs

import (
	bookingpb "be/proto"

	"be/internal/model"
)

func PbRoom(room model.Room) *bookingpb.Room {
	return &bookingpb.Room{
		Uuid:        room.ID.String(),
		Title:       room.Title,
		Description: room.Description,
		Number:      room.Number,
		Assignee:    PbUser(room.Employee),
		Author:      PbUser(room.Author),
		Owner:       PbUser(room.Owner),
		Slots:     PbRoomSlotList(room.Slots),
		Equipments:     PbRoomEquipmentList(room.Equipments),

		TargetAudience: room.TargetAudience,
		Status:         bookingpb.AddressStatus(room.Status),
		CreationDate:   ToUTCString(room.CreationDate),
		CreatedAt:      ToUTCString(room.CreatedAt),
		UpdatedAt:      ToUTCString(room.UpdatedAt),
		Outmembers:     OutmembersList(room.Outmembers),
		Links:          PbLinkList(room.Links),
		Releases:       PbRoomReleaseList(room.Releases),
		Attachments:    AttachmentList(room.Attachments),
		Bookings:       PbBookings(room.Bookings),
		Reports:        RoomReportsList(room.Reports),
		Participants:   PbParticipantsList(room.Participants),
		State:          bookingpb.AddressState(room.State),
	}
}

func PbRooms(rooms []model.Room) []*bookingpb.Room {
	pbRooms := make([]*bookingpb.Room, 0, len(rooms))
	for _, room := range rooms {
		pbRooms = append(pbRooms, PbRoom(room))
	}

	return pbRooms
}

func PbRoomSlotList(slots []model.Slot) []*bookingpb.RoomSlot {
	pbSlots := make([]*bookingpb.RoomSlot, 0, len(slots))
	for _, slot := range slots {
		pbSlots = append(pbSlots, PbRoomSlot(slot))
	}

	return pbSlots
}

func PbRoomSlot(slot model.Slot) *bookingpb.RoomSlot {
	return &bookingpb.RoomSlot{
		Uuid: slot.ID.String(),
		Timeline: &bookingpb.Timeline{
			Start: ToUTCString(slot.Timeline.StartAt),
			End:   ToUTCString(slot.Timeline.EndAt),
		},
		PlanSlot: toNotification(slot.PlanSlot),
		FactSlot: toNotification(slot.FactSlot),
		CreatedAt:  ToUTCString(slot.CreatedAt),
	}
}

func PbSlotList(slots []model.Slot) []*bookingpb.Slot {
	pbSlots := make([]*bookingpb.Slot, 0, len(slots))

	for _, slot := range slots {
		pbSlots = append(pbSlots, &bookingpb.Slot{
			Timeline: &bookingpb.Timeline{
				Start: ToUTCString(slot.Timeline.StartAt),
				End:   ToUTCString(slot.Timeline.EndAt),
			},
			PlanSlot: toNotification(slot.PlanSlot),
			FactSlot: toNotification(slot.FactSlot),
		})
	}

	return pbSlots
}

func PbRoomEquipmentList(equipments []model.Equipment) []*bookingpb.RoomEquipment {
	pbEquipments := make([]*bookingpb.RoomEquipment, 0, len(equipments))
	for _, equipment := range equipments {
		pbEquipments = append(pbEquipments, PbRoomEquipment(equipment))
	}

	return pbEquipments
}

func PbRoomEquipment(equipment model.Equipment) *bookingpb.RoomEquipment {
	return &bookingpb.RoomEquipment{
		Uuid:  equipment.ID.String(),
		Title: equipment.Title,
		Timeline: &bookingpb.Timeline{
			Start: ToUTCString(equipment.Timeline.StartAt),
			End:   ToUTCString(equipment.Timeline.EndAt),
		},
		Description: equipment.Description,
		PlanValue:   equipment.PlanValue,
		FactValue:   equipment.FactValue,
		CreatedAt:   ToUTCString(equipment.CreatedAt),
	}
}

func PbEquipmentList(equipments []model.Equipment) []*bookingpb.Equipment {
	pbEquipments := make([]*bookingpb.Equipment, 0, len(equipments))

	for _, equipment := range equipments {
		pbEquipments = append(pbEquipments, &bookingpb.Equipment{
			Title: equipment.Title,
			Timeline: &bookingpb.Timeline{
				Start: ToUTCString(equipment.Timeline.StartAt),
				End:   ToUTCString(equipment.Timeline.EndAt),
			},
			Description: equipment.Description,
			PlanValue:   equipment.PlanValue,
			FactValue:   equipment.FactValue,
		})
	}

	return pbEquipments
}

func PbRoomReleaseList(releases []model.Release) []*bookingpb.Release {
	pbRelease := make([]*bookingpb.Release, 0, len(releases))
	for _, release := range releases {
		pbRelease = append(pbRelease, PbRoomRelease(release))
	}

	return pbRelease
}

func PbRoomRelease(release model.Release) *bookingpb.Release {
	return &bookingpb.Release{
		Uuid:        release.ID.String(),
		Title:       release.Title,
		Description: release.Description,
		Date:        ToUTCString(release.Date),
		FactSlot:  toNotification(release.FactSlot),
		RoomID:   release.RoomID.String(),
	}
}

func PbParticipant(participant *model.User, role string) *bookingpb.Participant {
	return &bookingpb.Participant{
		Participant: PbUser(participant),
		Role:        role,
	}
}

func PbParticipantsList(participants []model.Participant) []*bookingpb.Participant {
	pbParticipants := make([]*bookingpb.Participant, 0, len(participants))
	for _, participant := range participants {
		pbParticipants = append(pbParticipants, PbParticipant(&participant.User, participant.Role))
	}

	return pbParticipants
}
