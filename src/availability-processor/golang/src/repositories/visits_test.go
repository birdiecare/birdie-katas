package repositories_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/birdiecare/availability-processor-exercise/src/repositories"
)

// Helper function to create a test visit
func createTestVisit(id, tenantId, patientId, caregiverId string, startTime, endTime time.Time) repositories.Visit {
	return repositories.Visit{
		Id:          id,
		TenantId:    tenantId,
		PatientId:   patientId,
		CaregiverId: caregiverId,
		StartTime:   startTime,
		EndTime:     endTime,
	}
}

// Helper function to get pointer to string
func stringPtr(s string) *string {
	return &s
}

func TestNewVisitRepository(t *testing.T) {
	t.Run("should create repository with empty visits list", func(t *testing.T) {
		repo := repositories.NewVisitRepository([]repositories.Visit{})
		assert.NotNil(t, repo)
	})

	t.Run("should create repository with initial visits", func(t *testing.T) {
		visits := []repositories.Visit{
			createTestVisit("1", "tenant1", "patient1", "caregiver1",
				time.Now(), time.Now().Add(time.Hour)),
		}
		repo := repositories.NewVisitRepository(visits)
		assert.NotNil(t, repo)
	})
}

func TestVisitRepositoryGetCalendar(t *testing.T) {
	// Setup test data
	baseTime := time.Date(2023, 11, 20, 9, 0, 0, 0, time.UTC)

	visits := []repositories.Visit{
		createTestVisit("1", "tenant1", "patient1", "caregiver1",
			baseTime, baseTime.Add(time.Hour)),
		createTestVisit("2", "tenant1", "patient2", "caregiver1",
			baseTime.Add(2*time.Hour), baseTime.Add(3*time.Hour)),
		createTestVisit("3", "tenant1", "patient3", "caregiver2",
			baseTime.Add(time.Hour), baseTime.Add(2*time.Hour)),
		createTestVisit("4", "tenant1", "patient4", "caregiver1",
			baseTime.Add(-time.Hour), baseTime.Add(-30*time.Minute)),
		createTestVisit("5", "tenant1", "patient5", "caregiver2",
			baseTime.Add(4*time.Hour), baseTime.Add(5*time.Hour)),
	}

	repo := repositories.NewVisitRepository(visits)

	t.Run("should return all visits when no caregiver filter applied", func(t *testing.T) {
		fromTime := baseTime.Add(-2 * time.Hour)
		toTime := baseTime.Add(6 * time.Hour)

		result, err := repo.GetCalendar(nil, fromTime, toTime)

		require.NoError(t, err)
		assert.Len(t, result, 5)
	})

	t.Run("should filter visits by caregiver ID", func(t *testing.T) {
		fromTime := baseTime.Add(-2 * time.Hour)
		toTime := baseTime.Add(6 * time.Hour)
		caregiverId := "caregiver1"

		result, err := repo.GetCalendar(&caregiverId, fromTime, toTime)

		require.NoError(t, err)
		assert.Len(t, result, 3)
		for _, visit := range result {
			assert.Equal(t, "caregiver1", visit.CaregiverId)
		}
	})

	t.Run("should filter visits by time range", func(t *testing.T) {
		// Time range that should only include visits 1, 2, and 3
		fromTime := baseTime.Add(-30 * time.Minute)
		toTime := baseTime.Add(3 * time.Hour)

		result, err := repo.GetCalendar(nil, fromTime, toTime)

		require.NoError(t, err)
		assert.Len(t, result, 3)

		// Verify the correct visits are returned
		visitIds := make(map[string]bool)
		for _, visit := range result {
			visitIds[visit.Id] = true
		}
		assert.True(t, visitIds["1"])
		assert.True(t, visitIds["2"])
		assert.True(t, visitIds["3"])
	})

	t.Run("should filter visits by both caregiver and time range", func(t *testing.T) {
		fromTime := baseTime
		toTime := baseTime.Add(3 * time.Hour)
		caregiverId := "caregiver1"

		result, err := repo.GetCalendar(&caregiverId, fromTime, toTime)

		require.NoError(t, err)
		assert.Len(t, result, 2)
		for _, visit := range result {
			assert.Equal(t, "caregiver1", visit.CaregiverId)
		}
	})

	t.Run("should return empty list when no visits match criteria", func(t *testing.T) {
		fromTime := baseTime.Add(10 * time.Hour)
		toTime := baseTime.Add(12 * time.Hour)
		caregiverId := "nonexistent_caregiver"

		result, err := repo.GetCalendar(&caregiverId, fromTime, toTime)

		require.NoError(t, err)
		assert.Len(t, result, 0)
	})

	t.Run("should handle visits that partially overlap with time range", func(t *testing.T) {
		// Test visit that starts before range but ends within range
		partialVisits := []repositories.Visit{
			createTestVisit("partial1", "tenant1", "patient1", "caregiver1",
				baseTime.Add(-30*time.Minute), baseTime.Add(30*time.Minute)),
			createTestVisit("partial2", "tenant1", "patient2", "caregiver1",
				baseTime.Add(30*time.Minute), baseTime.Add(90*time.Minute)),
		}

		partialRepo := repositories.NewVisitRepository(partialVisits)
		fromTime := baseTime
		toTime := baseTime.Add(time.Hour)

		result, err := partialRepo.GetCalendar(nil, fromTime, toTime)

		require.NoError(t, err)
		assert.Len(t, result, 2)
	})

	t.Run("should handle edge case where visit times exactly match range boundaries", func(t *testing.T) {
		exactVisits := []repositories.Visit{
			createTestVisit("exact1", "tenant1", "patient1", "caregiver1",
				baseTime, baseTime.Add(time.Hour)),
		}

		exactRepo := repositories.NewVisitRepository(exactVisits)
		fromTime := baseTime
		toTime := baseTime.Add(time.Hour)

		result, err := exactRepo.GetCalendar(nil, fromTime, toTime)

		require.NoError(t, err)
		assert.Len(t, result, 1) // Visit should be included as it overlaps with the range
	})

	t.Run("should exclude visits that don't overlap with time range", func(t *testing.T) {
		nonOverlappingVisits := []repositories.Visit{
			// Visit that ends exactly when range starts
			createTestVisit("before", "tenant1", "patient1", "caregiver1",
				baseTime.Add(-time.Hour), baseTime),
			// Visit that starts exactly when range ends
			createTestVisit("after", "tenant1", "patient2", "caregiver1",
				baseTime.Add(time.Hour), baseTime.Add(2*time.Hour)),
		}

		noOverlapRepo := repositories.NewVisitRepository(nonOverlappingVisits)
		fromTime := baseTime
		toTime := baseTime.Add(time.Hour)

		result, err := noOverlapRepo.GetCalendar(nil, fromTime, toTime)

		require.NoError(t, err)
		assert.Len(t, result, 0) // No visits should be included
	})
}

func TestVisitRepositoryUnassign(t *testing.T) {
	t.Run("should unassign caregiver from visit", func(t *testing.T) {
		visits := []repositories.Visit{
			createTestVisit("1", "tenant1", "patient1", "caregiver1",
				time.Now(), time.Now().Add(time.Hour)),
			createTestVisit("2", "tenant1", "patient2", "caregiver2",
				time.Now().Add(time.Hour), time.Now().Add(2*time.Hour)),
		}

		repo := repositories.NewVisitRepository(visits)

		err := repo.Unassign("1", "caregiver1")

		require.NoError(t, err)

		// Verify the visit was unassigned by checking calendar
		result, err := repo.GetCalendar(stringPtr("caregiver1"),
			time.Now().Add(-time.Hour), time.Now().Add(3*time.Hour))
		require.NoError(t, err)
		assert.Len(t, result, 0) // Should find no visits for caregiver1 now

		// Verify other visits are unaffected
		result2, err := repo.GetCalendar(stringPtr("caregiver2"),
			time.Now().Add(-time.Hour), time.Now().Add(3*time.Hour))
		require.NoError(t, err)
		assert.Len(t, result2, 1)
	})

	t.Run("should do nothing when visit ID doesn't exist", func(t *testing.T) {
		visits := []repositories.Visit{
			createTestVisit("1", "tenant1", "patient1", "caregiver1",
				time.Now(), time.Now().Add(time.Hour)),
		}

		repo := repositories.NewVisitRepository(visits)

		err := repo.Unassign("nonexistent", "caregiver1")

		require.NoError(t, err)

		// Verify original visit is still assigned
		result, err := repo.GetCalendar(stringPtr("caregiver1"),
			time.Now().Add(-time.Hour), time.Now().Add(3*time.Hour))
		require.NoError(t, err)
		assert.Len(t, result, 1)
	})

	t.Run("should do nothing when caregiver ID doesn't match", func(t *testing.T) {
		visits := []repositories.Visit{
			createTestVisit("1", "tenant1", "patient1", "caregiver1",
				time.Now(), time.Now().Add(time.Hour)),
		}

		repo := repositories.NewVisitRepository(visits)

		err := repo.Unassign("1", "wrong_caregiver")

		require.NoError(t, err)

		// Verify visit is still assigned to original caregiver
		result, err := repo.GetCalendar(stringPtr("caregiver1"),
			time.Now().Add(-time.Hour), time.Now().Add(3*time.Hour))
		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "caregiver1", result[0].CaregiverId)
	})

	t.Run("should unassign only the matching visit when multiple visits exist", func(t *testing.T) {
		visits := []repositories.Visit{
			createTestVisit("1", "tenant1", "patient1", "caregiver1",
				time.Now(), time.Now().Add(time.Hour)),
			createTestVisit("2", "tenant1", "patient1", "caregiver1",
				time.Now().Add(time.Hour), time.Now().Add(2*time.Hour)),
		}

		repo := repositories.NewVisitRepository(visits)

		err := repo.Unassign("1", "caregiver1")

		require.NoError(t, err)

		// Verify only one visit remains assigned
		result, err := repo.GetCalendar(stringPtr("caregiver1"),
			time.Now().Add(-time.Hour), time.Now().Add(3*time.Hour))
		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "2", result[0].Id)
	})
}

func TestVisitStruct(t *testing.T) {
	t.Run("should create visit with all fields", func(t *testing.T) {
		startTime := time.Now()
		endTime := startTime.Add(time.Hour)

		visit := repositories.Visit{
			Id:          "test-id",
			TenantId:    "tenant-123",
			PatientId:   "patient-456",
			CaregiverId: "caregiver-789",
			StartTime:   startTime,
			EndTime:     endTime,
		}

		assert.Equal(t, "test-id", visit.Id)
		assert.Equal(t, "tenant-123", visit.TenantId)
		assert.Equal(t, "patient-456", visit.PatientId)
		assert.Equal(t, "caregiver-789", visit.CaregiverId)
		assert.Equal(t, startTime, visit.StartTime)
		assert.Equal(t, endTime, visit.EndTime)
	})
}
