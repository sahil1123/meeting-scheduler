package models

import "sync"

type MeetingRoom struct {
	Id       int
	Capacity int
	Calendar calendar
	mu       sync.RWMutex
}

func GetNewMeetingRoom(id int, cap int) MeetingRoom {
	hrs := make([]bool, 24) //by default false
	return MeetingRoom{
		Id:       id,
		Capacity: cap,
		Calendar: calendar{
			Interval: hrs,
		},
	}
}

func (m MeetingRoom) IsAvailable(s, e int) bool {
	m.mu.RLock() //this will make sure that when Lock() is taken in Book(), this gets locked & only read after Book() lock is released
	defer m.mu.RUnlock()

	for hour := s; hour < e; hour++ {
		if m.Calendar.Interval[hour] {
			return false
		}
	}
	return true
}

func (m MeetingRoom) GetCapacity() int {
	return m.Capacity
}

func (m MeetingRoom) Book(s, e int) {
	//use lock here - for concurrency handling

	m.mu.Lock()
	defer m.mu.Unlock()

	for hour := s; hour < e; hour++ {
		m.Calendar.Interval[hour] = true
	}
}
