import { EventProcessor } from "./handler";
import type { CaregiverPermanentUnavailabilityEvent } from "./events";
import { VisitRepository, type Visit } from "./repositories/visits";

const testTenantId = "tenant-1";
const testCaregiverId = "caregiver-1";

describe("EventProcessor", () => {
  describe("handleEvent", () => {
    it("processes permanent unavailability event correctly", async () => {
      const visits: Visit[] = [
        {
          // This visit is before the permanent unavailability and should remain assigned
          id: "visit-1",
          tenantId: testTenantId,
          patientId: "patient-1",
          caregiverId: testCaregiverId,
          startTime: new Date("2025-11-06T10:00:00.000Z"), // Day before
          endTime: new Date("2025-11-06T11:00:00.000Z"),
        },
        {
          // This visit is after the permanent unavailability and should be unassigned
          id: "visit-2",
          tenantId: testTenantId,
          patientId: "patient-2",
          caregiverId: testCaregiverId,
          startTime: new Date("2025-11-08T10:00:00.000Z"), // Day after
          endTime: new Date("2025-11-08T11:00:00.000Z"),
        },
      ];

      // Create repository and processor
      const repo = new VisitRepository(visits);
      const eventProcessor = new EventProcessor(repo);

      // Create a permanent unavailability event starting on Nov 7th
      const unavailabilityEvent: CaregiverPermanentUnavailabilityEvent = {
        id: "unavailability-1",
        tenantId: testTenantId,
        caregiverId: testCaregiverId,
        effectiveFrom: new Date("2025-11-07T00:00:00.000Z"), // Start of Nov 7th
      };

      // Process the permanent unavailability event
      await eventProcessor.handleEvent(unavailabilityEvent);

      // Check the results - get all visits by not specifying caregiver ID
      const allVisits = await repo.getCalendar(
        null,
        new Date("2025-11-01T00:00:00.000Z"),
        new Date("2026-11-01T00:00:00.000Z")
      );

      // Visits should be in the same order as they were created
      expect(allVisits).toHaveLength(2);

      // Check visit-1 (position 0, before unavailability) remains assigned
      expect(allVisits[0].id).toBe("visit-1");
      expect(allVisits[0].caregiverId).toBe(testCaregiverId);

      // Check visit-2 (position 1, after unavailability) was unassigned
      expect(allVisits[1].id).toBe("visit-2");
      expect(allVisits[1].caregiverId).toBe("");
    });

    it.todo(
      "handles visits starting exactly at unavailability time",
      async () => {
        // When a visit starts exactly at the same time as the permanent unavailability,
        // that visit should also be unassigned.

        const startTime = new Date("2025-11-07T00:00:00.000Z");

        const visits: Visit[] = [
          {
            // This visit starts exactly at the unavailability time and should be unassigned
            id: "visit-1",
            tenantId: testTenantId,
            patientId: "patient-1",
            caregiverId: testCaregiverId,
            startTime,
            endTime: new Date(startTime.getTime() + 60 * 60 * 1000), // 1 hour later
          },
        ];

        // Create repository and processor
        const repo = new VisitRepository(visits);
        const eventProcessor = new EventProcessor(repo);

        // Create a permanent unavailability event starting on Nov 7th
        const unavailabilityEvent: CaregiverPermanentUnavailabilityEvent = {
          id: "unavailability-1",
          tenantId: testTenantId,
          caregiverId: testCaregiverId,
          effectiveFrom: startTime,
        };

        // Process the permanent unavailability event
        await eventProcessor.handleEvent(unavailabilityEvent);

        // Check the results - get all visits by not specifying caregiver ID
        const allVisits = await repo.getCalendar(
          null,
          new Date("2025-11-01T00:00:00.000Z"),
          new Date("2026-11-01T00:00:00.000Z")
        );

        // Check visit-1 was unassigned
        expect(allVisits[0].id).toBe("visit-1");
        expect(allVisits[0].caregiverId).toBe("");
      }
    );
  });
});
