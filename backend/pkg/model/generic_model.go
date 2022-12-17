package model

type GenericModel interface {
	User | Client | Employee | Vehicle | Salary | Rental | Repair | Notification
	GetId() int
}
