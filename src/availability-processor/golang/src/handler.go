package processor

import (
	"github.com/birdiecare/availability-processor-exercise/src/repositories"
)

// EventProcessor handles caregiver availability events
type EventProcessor struct {
	visitRepo repositories.VisitRepository
}

// NewEventProcessor creates a new event processor
func NewEventProcessor(visitRepo repositories.VisitRepository) *EventProcessor {
	return &EventProcessor{
		visitRepo: visitRepo,
	}
}

func (p *EventProcessor) HandleEvent(event CaregiverPermanentUnavailabilityEvent) error {
	futureDate := event.EffectiveFrom.AddDate(1, 0, 0) // 1 year in the future

	visits, err := p.visitRepo.GetCalendar(
		&event.CaregiverId,
		event.EffectiveFrom,
		futureDate,
	)
	if err != nil {
		return err
	}

	// Unassign all visits that occur after the permanent unavailability starts
	for _, visit := range visits {
		if visit.StartTime.After(event.EffectiveFrom) || visit.StartTime.Equal(event.EffectiveFrom) {
			err := p.visitRepo.Unassign(visit.Id, event.CaregiverId)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
