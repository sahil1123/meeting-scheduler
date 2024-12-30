package service

import (
	"meeting-scheduler/models"
	"sync"
)

type Meeting struct {
	Users              []models.User
	StartTime, EndTime int
}
type MeetingRoom struct {
	Id       int
	Capacity int
}

type Request struct {
	Meetings    []Meeting
	MeetingRoom []MeetingRoom
}

type MeetingResponse struct {
	MeetingId     int
	MeetingRoomId int
	Error         bool
	ErrorMessage  string
	Meeting       Meeting
}

type Response struct {
	MeetingResponse []MeetingResponse
}

type MeetingScheduler struct {
	wg sync.WaitGroup
	mu sync.Mutex // Protects the response
	//Without a lock, simultaneous writes to resp.MeetingResponse could lead to:
	//Race conditions: Goroutines might interfere with each other's writes, leading to corrupted or incorrect data.
	//Panic: The Go runtime could panic due to concurrent writes to a slice
	emailSvc BulkEmailService
}

func NewMeetingScheduler(emailSvc BulkEmailService) MeetingScheduler {
	return MeetingScheduler{
		wg:       sync.WaitGroup{},
		mu:       sync.Mutex{},
		emailSvc: emailSvc,
	}
}

func (ms MeetingScheduler) Schedule(req Request) Response {
	var (
		meetingRooms = make([]models.MeetingRoom, 0)
		//users = make([]models.User, 0)
		resp = Response{}
	)
	for _, m := range req.MeetingRoom {
		meetingRooms = append(meetingRooms, models.GetNewMeetingRoom(m.Id, m.Capacity))
	}

	meetingId := 1
	for _, m := range req.Meetings {
		ms.wg.Add(1)

		go func(meeting Meeting) {
			defer ms.wg.Done()

			var found bool
			for _, r := range meetingRooms {
				if r.GetCapacity() >= len(m.Users) &&
					r.IsAvailable(m.StartTime, m.EndTime) {

					//book meeting room
					r.Book(m.StartTime, m.EndTime)

					ms.mu.Lock()
					resp.MeetingResponse = append(resp.MeetingResponse, MeetingResponse{
						MeetingId:     meetingId,
						MeetingRoomId: r.Id,
						Error:         false,
						ErrorMessage:  "",
						Meeting:       m,
					})

					meetingId++
					ms.mu.Unlock()
					//The mu.Lock() ensures that the update to resp.MeetingResponse and the increment of meetingId are atomic operations.
					//meetingId is shared across all goroutines. Without synchronization, concurrent reads/writes to meetingId could result in inconsistent values.

					//ms.emailSvc.SendBulkEmails(users, meeting)  SEND EMAIL HERE

					found = true
					break
				}
			}
			if !found {
				ms.mu.Lock()
				resp.MeetingResponse = append(resp.MeetingResponse, MeetingResponse{
					MeetingId:     -1,
					MeetingRoomId: -1,
					Error:         true,
					ErrorMessage:  "NO MEETING ROOM AVAILABLE",
					Meeting:       m,
				})
				ms.mu.Unlock()
			}
		}(m)

	}
	ms.wg.Wait()

	return resp
}
