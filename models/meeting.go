package models

type Meeting struct {
	Users         []User
	StartTime     int
	EndTime       int
	MeetingRoomId int
}

func GetNewMeeting(users []User, sT, eT, mRId int) Meeting {
	return Meeting{
		Users:         users,
		StartTime:     sT,
		EndTime:       eT,
		MeetingRoomId: mRId,
	}
}
