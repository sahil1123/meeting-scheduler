package main

import (
	"fmt"
	"meeting-scheduler/models"
	"meeting-scheduler/service"
)

func main() {
	meetings := []service.Meeting{
		{
			Users: []models.User{
				{
					Id:    1,
					Name:  "first",
					Email: "first@mail",
				},
				{
					Id:    2,
					Name:  "second",
					Email: "second@mail",
				},
			},
			StartTime: 2,
			EndTime:   4,
		},
		{
			Users: []models.User{
				{
					Id:    0,
					Name:  "first",
					Email: "first@mail",
				},
				{
					Id:    2,
					Name:  "second",
					Email: "second@mail",
				},
			},
			StartTime: 2,
			EndTime:   4,
		},
	}

	meetingRooms := []service.MeetingRoom{
		{
			Id:       1,
			Capacity: 1,
		},
		{
			Id:       2,
			Capacity: 2,
		},
		{
			Id:       3,
			Capacity: 0,
		},
	}

	req := service.Request{
		Meetings:    meetings,
		MeetingRoom: meetingRooms,
	}

	emailSvc := service.NewEmailNotificationService()
	resp := service.NewMeetingScheduler(emailSvc).Schedule(req)
	for _, r := range resp.MeetingResponse {
		if r.Error {
			fmt.Printf("error: %s, req: %v \n", r.ErrorMessage, r.Meeting)
		} else {
			fmt.Printf("MeetingId: %d, MeetingRoomId: %d \n", r.MeetingId, r.MeetingRoomId)
		}
	}
}

//Questions an Interviewer Might Ask:
//Scalability:
//How would your design handle 10,000 meeting rooms and millions of meeting requests daily? Can you optimize the search for an available room?

//Concurrency:
//How would you handle concurrent booking requests? What changes would you make to ensure thread safety?

//Multi-Day Scheduling:
//How would you extend the design to handle meetings spanning multiple days?
