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

type VisitRepository struct {
	visits []Visit
	// GetCalendar(caregiverId *string, fromTime time.Time, toTime time.Time) ([]Visit, error)
	// Unassign(visitId string, caregiverId string) error
}

func NewVisitRepository(visits []Visit) VisitRepository {
	return VisitRepository{
		visits: visits,
	}
}

func (r VisitRepository) GetCalendar(caregiverId *string, fromTime, toTime time.Time) ([]Visit, error) {
	var matchingVisits []Visit

	for _, visit := range r.visits {
		// Filter by caregiver only if caregiverId is provided
		if caregiverId != nil && visit.CaregiverId != *caregiverId {
			continue
		}

		// Filter by time range - visit overlaps with requested period
		if visit.StartTime.Before(toTime) && visit.EndTime.After(fromTime) {
			matchingVisits = append(matchingVisits, visit)
		}
	}

	return matchingVisits, nil
}

func (r VisitRepository) Unassign(visitId string, caregiverId string) error {
	for i, visit := range r.visits {
		if visit.Id == visitId && visit.CaregiverId == caregiverId {
			// Set caregiver ID to empty string to indicate unassignment
			r.visits[i].CaregiverId = ""
			break
		}
	}
	return nil
}
