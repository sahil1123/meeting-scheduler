package service

import (
	"fmt"
	"meeting-scheduler/models"
)

type NotificationService interface {
	SendNotification(user models.User, meeting models.Meeting)
}

type msgNotificationService struct{}

func NewMessageNotificationService() NotificationService {
	return msgNotificationService{}
}

func (n msgNotificationService) SendNotification(user models.User, meeting models.Meeting) {
	//TODO implement me
}

type emailNotificationService struct{}

type BulkEmailService interface {
	NotificationService
	SendBulkEmails(users []models.User, meeting models.Meeting)
}

func NewEmailNotificationService() BulkEmailService {
	return emailNotificationService{}
}

func (e emailNotificationService) SendNotification(user models.User, meeting models.Meeting) {
	//TODO implement me
}

func (e emailNotificationService) SendBulkEmails(users []models.User, meeting models.Meeting) {
	// Implement bulk email sending logic here
	for _, user := range users {
		// Logic for sending email to each user
		fmt.Println(user)
	}
}
