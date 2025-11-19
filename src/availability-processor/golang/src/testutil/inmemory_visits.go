package testutil

import (
	"time"

	"github.com/birdiecare/availability-processor-exercise/src/repositories"
)

type InMemoryVisitRepository struct {
	visits []repositories.Visit
}

func NewInMemoryVisitRepository(initialVisits []repositories.Visit) repositories.VisitRepository {
	visits := make([]repositories.Visit, len(initialVisits))
	copy(visits, initialVisits)
	return &InMemoryVisitRepository{
		visits: visits,
	}
}

func (r *InMemoryVisitRepository) GetCalendar(caregiverId *string, fromTime, toTime time.Time) ([]repositories.Visit, error) {
	var matchingVisits []repositories.Visit

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

func (r *InMemoryVisitRepository) Unassign(visitId string, caregiverId string) error {
	for i, visit := range r.visits {
		if visit.Id == visitId && visit.CaregiverId == caregiverId {
			// Set caregiver ID to empty string to indicate unassignment
			r.visits[i].CaregiverId = ""
			break
		}
	}
	return nil
}
