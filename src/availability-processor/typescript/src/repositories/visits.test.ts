import { VisitRepository, type Visit } from "./visits";

// Import date extensions to allow for use of methods like addHours, subtractMinutes, etc.
import "../date-extensions";

// Helper function to create a test visit
function createTestVisit(
  id: string,
  tenantId: string,
  patientId: string,
  caregiverId: string,
  startTime: Date,
  endTime: Date
): Visit {
  return {
    id,
    tenantId,
    patientId,
    caregiverId,
    startTime,
    endTime,
  };
}

describe("VisitRepository", () => {
  describe("constructor", () => {
    it("creates repository with initial visits", () => {
      const visits = [
        createTestVisit(
          "1",
          "tenant1",
          "patient1",
          "caregiver1",
          new Date(),
          new Date().addHours(1)
        ),
      ];
      const repo = new VisitRepository(visits);
      expect(repo).toBeDefined();
    });
  });

  describe("getCalendar", () => {
    const now = new Date("2023-11-20T09:00:00.000Z");

    const visits = [
      createTestVisit(
        "1",
        "tenant1",
        "patient1",
        "caregiver1",
        now,
        now.addHours(1)
      ),
      createTestVisit(
        "2",
        "tenant1",
        "patient2",
        "caregiver1",
        now.addHours(2),
        now.addHours(3)
      ),
      createTestVisit(
        "3",
        "tenant1",
        "patient3",
        "caregiver2",
        now.addHours(1),
        now.addHours(2)
      ),
      createTestVisit(
        "4",
        "tenant1",
        "patient4",
        "caregiver1",
        now.subtractHours(1),
        now.subtractMinutes(30)
      ),
      createTestVisit(
        "5",
        "tenant1",
        "patient5",
        "caregiver2",
        now.addHours(4),
        now.addHours(5)
      ),
    ];

    it("returns all visits when no caregiver filter applied", async () => {
      const repo = new VisitRepository(visits);
      const fromTime = now.subtractHours(2);
      const toTime = now.addHours(6);

      const result = await repo.getCalendar(null, fromTime, toTime);

      expect(result).toHaveLength(5);
    });

    it("filters visits by caregiver ID", async () => {
      const repo = new VisitRepository(visits);
      const fromTime = now.subtractHours(2);
      const toTime = now.addHours(6);
      const caregiverId = "caregiver1";

      const result = await repo.getCalendar(caregiverId, fromTime, toTime);

      expect(result).toHaveLength(3);
      for (const visit of result) {
        expect(visit.caregiverId).toBe("caregiver1");
      }
    });

    it("filters visits by time range", async () => {
      const repo = new VisitRepository(visits);
      // Time range that should only include visits 1, 2, and 3
      const fromTime = now.subtractMinutes(30);
      const toTime = now.addHours(3);

      const result = await repo.getCalendar(null, fromTime, toTime);

      expect(result).toHaveLength(3);

      // Verify the correct visits are returned
      const visitIds = new Set(result.map((visit) => visit.id));
      expect(visitIds.has("1")).toBe(true);
      expect(visitIds.has("2")).toBe(true);
      expect(visitIds.has("3")).toBe(true);
    });

    it("filters visits by both caregiver and time range", async () => {
      const repo = new VisitRepository(visits);
      const fromTime = now;
      const toTime = now.addHours(3);
      const caregiverId = "caregiver1";

      const result = await repo.getCalendar(caregiverId, fromTime, toTime);

      expect(result).toHaveLength(2);
      for (const visit of result) {
        expect(visit.caregiverId).toBe("caregiver1");
      }
    });

    it("returns empty list when no visits match criteria", async () => {
      const repo = new VisitRepository(visits);
      const fromTime = now.addHours(10);
      const toTime = now.addHours(12);
      const caregiverId = "nonexistent_caregiver";

      const result = await repo.getCalendar(caregiverId, fromTime, toTime);

      expect(result).toHaveLength(0);
    });

    it("includes visits that partially overlap with time range", async () => {
      // Test visit that starts before range but ends within range
      const partialVisits = [
        createTestVisit(
          "partial1",
          "tenant1",
          "patient1",
          "caregiver1",
          now.subtractMinutes(30),
          now.addMinutes(30)
        ),
        createTestVisit(
          "partial2",
          "tenant1",
          "patient2",
          "caregiver1",
          now.addMinutes(30),
          now.addMinutes(90)
        ),
      ];

      const repo = new VisitRepository(partialVisits);
      const fromTime = now;
      const toTime = now.addHours(1);

      const result = await repo.getCalendar(null, fromTime, toTime);

      expect(result).toHaveLength(2);
    });

    it("includes visits with times exactly matching range boundaries", async () => {
      const exactVisits = [
        createTestVisit(
          "exact1",
          "tenant1",
          "patient1",
          "caregiver1",
          now,
          now.addHours(1)
        ),
      ];

      const repo = new VisitRepository(exactVisits);
      const fromTime = now;
      const toTime = now.addHours(1);

      const result = await repo.getCalendar(null, fromTime, toTime);

      expect(result).toHaveLength(1); // Visit should be included as it overlaps with the range
    });

    it("excludes visits that don't overlap with time range", async () => {
      const nonOverlappingVisits = [
        // Visit that ends exactly when range starts
        createTestVisit(
          "before",
          "tenant1",
          "patient1",
          "caregiver1",
          now.subtractHours(1),
          now
        ),
        // Visit that starts exactly when range ends
        createTestVisit(
          "after",
          "tenant1",
          "patient2",
          "caregiver1",
          now.addHours(1),
          now.addHours(2)
        ),
      ];

      const repo = new VisitRepository(nonOverlappingVisits);
      const fromTime = now;
      const toTime = now.addHours(1);

      const result = await repo.getCalendar(null, fromTime, toTime);

      expect(result).toHaveLength(0); // No visits should be included
    });
  });

  describe("unassign", () => {
    it("unassigns caregiver from visit", async () => {
      const now = new Date();
      const visits = [
        createTestVisit(
          "1",
          "tenant1",
          "patient1",
          "caregiver1",
          now,
          now.addHours(1)
        ),
        createTestVisit(
          "2",
          "tenant1",
          "patient2",
          "caregiver2",
          now.addHours(1),
          now.addHours(2)
        ),
      ];

      const repo = new VisitRepository(visits);

      await repo.unassign("1", "caregiver1");

      // Verify the visit was unassigned by checking calendar
      const result = await repo.getCalendar(
        "caregiver1",
        now.subtractHours(1),
        now.addHours(3)
      );
      expect(result).toHaveLength(0); // Should find no visits for caregiver1 now

      // Verify other visits are unaffected
      const result2 = await repo.getCalendar(
        "caregiver2",
        now.subtractHours(1),
        now.addHours(3)
      );
      expect(result2).toHaveLength(1);
    });

    it("does nothing when visit ID doesn't exist", async () => {
      const now = new Date();
      const visits = [
        createTestVisit(
          "1",
          "tenant1",
          "patient1",
          "caregiver1",
          now,
          now.addHours(1)
        ),
      ];

      const repo = new VisitRepository(visits);

      await repo.unassign("nonexistent", "caregiver1");

      // Verify original visit is still assigned
      const result = await repo.getCalendar(
        "caregiver1",
        now.subtractHours(1),
        now.addHours(3)
      );
      expect(result).toHaveLength(1);
    });

    it("does nothing when caregiver ID doesn't match", async () => {
      const now = new Date();
      const visits = [
        createTestVisit(
          "1",
          "tenant1",
          "patient1",
          "caregiver1",
          now,
          now.addHours(1)
        ),
      ];

      const repo = new VisitRepository(visits);

      await repo.unassign("1", "wrong_caregiver");

      // Verify visit is still assigned to original caregiver
      const result = await repo.getCalendar(
        "caregiver1",
        now.subtractHours(1),
        now.addHours(3)
      );
      expect(result).toHaveLength(1);
      expect(result[0].caregiverId).toBe("caregiver1");
    });

    it("unassigns only the matching visit when multiple visits exist", async () => {
      const now = new Date();
      const visits = [
        createTestVisit(
          "1",
          "tenant1",
          "patient1",
          "caregiver1",
          now,
          now.addHours(1)
        ),
        createTestVisit(
          "2",
          "tenant1",
          "patient1",
          "caregiver1",
          now.addHours(1),
          now.addHours(2)
        ),
      ];

      const repo = new VisitRepository(visits);

      await repo.unassign("1", "caregiver1");

      // Verify only one visit remains assigned
      const result = await repo.getCalendar(
        "caregiver1",
        now.subtractHours(1),
        now.addHours(3)
      );
      expect(result).toHaveLength(1);
      expect(result[0].id).toBe("2");
    });
  });
});
