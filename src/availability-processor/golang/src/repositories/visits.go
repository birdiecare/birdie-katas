package repositories

import "time"

type Visit struct {
	Id          string    // Unique identifier for the visit
	TenantId    string    // Unique identifier for the tenant (care agency)
	PatientId   string    // Unique identifier for the patient
	CaregiverId string    // Unique identifier for the caregiver
	StartTime   time.Time // Start time of the visit
	EndTime     time.Time // End time of the visit
}

type VisitRepository interface {
	GetCalendar(caregiverId *string, fromTime time.Time, toTime time.Time) ([]Visit, error)
	Unassign(visitId string, caregiverId string) error
}
