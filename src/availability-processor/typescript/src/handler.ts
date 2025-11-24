import type { CaregiverPermanentUnavailabilityEvent } from "./events";
import type { VisitRepository } from "./repositories/visits";

/**
 * EventProcessor handles caregiver availability events
 */
export class EventProcessor {
  constructor(private readonly visitRepo: VisitRepository) {}

  async handleEvent(
    event: CaregiverPermanentUnavailabilityEvent
  ): Promise<void> {
    const futureDate = new Date(
      event.effectiveFrom.getTime() + 365 * 24 * 60 * 60 * 1000
    ); // 1 year in the future

    const visits = await this.visitRepo.getCalendar(
      event.caregiverId,
      event.effectiveFrom,
      futureDate
    );

    // Unassign all visits that occur after the permanent unavailability starts
    for (const visit of visits) {
      if (visit.startTime > event.effectiveFrom) {
        await this.visitRepo.unassign(visit.id, event.caregiverId);
      }
    }
  }
}
