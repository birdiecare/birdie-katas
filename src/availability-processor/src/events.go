package processor

import "time"

type CaregiverPermanentUnavailabilityEvent struct {
	Id            string    // Unique identifier for the permanent unavailability event
	TenantId      string    // Unique identifier for the tenant (care agency)
	CaregiverId   string    // Unique identifier for the caregiver
	EffectiveFrom time.Time // Time when the permanent unavailability starts
}

type CaregiverAbsenceBookedEvent struct {
	Id          string    // Unique identifier for the absence event
	TenantId    string    // Unique identifier for the tenant (care agency)
	CaregiverId string    // Unique identifier for the caregiver
	StartTime   time.Time // Start time of the absence
	EndTime     time.Time // End time of the absence
}
