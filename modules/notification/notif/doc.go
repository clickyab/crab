package notif

import (
	"time"
)

// NotificationType is the type of notification
type (
	// NotificationType is the user active status
	// @Enum{
	// }
	NotificationType string
)

const (
	// NotificationTypeEmail is email notif type
	NotificationTypeEmail NotificationType = "email"
	// NotificationTypeSMS is sms notif type
	NotificationTypeSMS NotificationType = "sms"
)

// Notification user model in database
// @Model {
//		table = notification
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Notification struct {
	ID        int64            `json:"id" db:"id"`
	UserID    int64            `json:"user_id" db:"user_id"`
	Title     string           `json:"title" db:"title"`
	Message   string           `json:"message" db:"message"`
	Type      NotificationType `json:"type" db:"type"`
	CreatedAt time.Time        `json:"created_at" db:"created_at"`
}
