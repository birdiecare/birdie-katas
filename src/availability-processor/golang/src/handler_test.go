package processor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/birdiecare/availability-processor-exercise/src/repositories"
)

const (
	testTenantID    = "tenant-1"
	testCaregiverID = "caregiver-1"
)

func TestProcessPermanentUnavailabilityEvent(t *testing.T) {
	visits := []repositories.Visit{
		{
			// This visit is before the permanent unavailability and should remain assigned
			Id:          "visit-1",
			TenantId:    testTenantID,
			PatientId:   "patient-1",
			CaregiverId: testCaregiverID,
			StartTime:   time.Date(2025, 11, 6, 10, 0, 0, 0, time.UTC), // Day before
			EndTime:     time.Date(2025, 11, 6, 11, 0, 0, 0, time.UTC),
		},
		{
			// This visit is after the permanent unavailability and should be unassigned
			Id:          "visit-2",
			TenantId:    testTenantID,
			PatientId:   "patient-2",
			CaregiverId: testCaregiverID,
			StartTime:   time.Date(2025, 11, 8, 10, 0, 0, 0, time.UTC), // Day after
			EndTime:     time.Date(2025, 11, 8, 11, 0, 0, 0, time.UTC),
		},
	}

	// Create repository and processor
	repo := repositories.NewVisitRepository(visits)
	eventProcessor := NewEventProcessor(repo)

	// Create a permanent unavailability event starting on Nov 7th
	unavailabilityEvent := CaregiverPermanentUnavailabilityEvent{
		Id:            "unavailability-1",
		TenantId:      testTenantID,
		CaregiverId:   testCaregiverID,
		EffectiveFrom: time.Date(2025, 11, 7, 0, 0, 0, 0, time.UTC), // Start of Nov 7th
	}

	// Process the permanent unavailability event
	err := eventProcessor.HandleEvent(unavailabilityEvent)
	assert.NoError(t, err)

	// Check the results - get all visits by not specifying caregiver ID
	allVisits, err := repo.GetCalendar(nil, time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC), time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC))
	assert.NoError(t, err)

	// Visits should be in the same order as they were created
	assert.Len(t, allVisits, 2, "Should have exactly 2 visits")

	// Check visit-1 (position 0, before unavailability) remains assigned
	assert.Equal(t, "visit-1", allVisits[0].Id)
	assert.Equal(t, testCaregiverID, allVisits[0].CaregiverId)

	// Check visit-2 (position 1, after unavailability) was unassigned
	assert.Equal(t, "visit-2", allVisits[1].Id)
	assert.Empty(t, allVisits[1].CaregiverId)
}

func TestProcessPermanentUnavailabilityEventStartingAtSameTime(t *testing.T) {
	t.Skip("TODO: Implement test for visits starting exactly at unavailability time")

	// When a visit starts exactly at the same time as the permanent unavailability,
	// that visit should also be unassigned.

	startTime := time.Date(2025, 11, 7, 0, 0, 0, 0, time.UTC)

	visits := []repositories.Visit{
		{
			// This visit starts exactly at the unavailability time and should be unassigned
			Id:          "visit-1",
			TenantId:    testTenantID,
			PatientId:   "patient-1",
			CaregiverId: testCaregiverID,
			StartTime:   startTime,
			EndTime:     startTime.Add(1 * time.Hour),
		},
	}

	// Create repository and processor
	repo := repositories.NewVisitRepository(visits)
	eventProcessor := NewEventProcessor(repo)

	// Create a permanent unavailability event starting on Nov 7th
	unavailabilityEvent := CaregiverPermanentUnavailabilityEvent{
		Id:            "unavailability-1",
		TenantId:      testTenantID,
		CaregiverId:   testCaregiverID,
		EffectiveFrom: startTime,
	}

	// Process the permanent unavailability event
	err := eventProcessor.HandleEvent(unavailabilityEvent)
	assert.NoError(t, err)

	// Check the results - get all visits by not specifying caregiver ID
	allVisits, err := repo.GetCalendar(nil, time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC), time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC))
	assert.NoError(t, err)

	// Check visit-1 was unassigned
	assert.Equal(t, "visit-1", allVisits[0].Id)
	assert.Empty(t, allVisits[0].CaregiverId)
}
