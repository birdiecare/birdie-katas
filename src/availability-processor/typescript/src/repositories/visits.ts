export interface Visit {
  /** Unique identifier for the visit */
  id: string;
  /** Unique identifier for the tenant (care agency) */
  tenantId: string;
  /** Unique identifier for the patient */
  patientId: string;
  /** Unique identifier for the caregiver */
  caregiverId: string;
  /** Start time of the visit */
  startTime: Date;
  /** End time of the visit */
  endTime: Date;
}

export class VisitRepository {
  private visits: Visit[];

  constructor(visits: Visit[] = []) {
    this.visits = visits;
  }

  async getCalendar(
    caregiverId: string | null,
    fromTime: Date,
    toTime: Date
  ): Promise<Visit[]> {
    const matchingVisits: Visit[] = [];

    for (const visit of this.visits) {
      // Filter by caregiver only if caregiverId is provided
      if (caregiverId !== null && visit.caregiverId !== caregiverId) {
        continue;
      }

      // Filter by time range - visit overlaps with requested period
      if (visit.startTime < toTime && visit.endTime > fromTime) {
        matchingVisits.push(visit);
      }
    }

    return matchingVisits;
  }

  async unassign(visitId: string, caregiverId: string): Promise<void> {
    for (const visit of this.visits) {
      if (visit.id === visitId && visit.caregiverId === caregiverId) {
        // Set caregiver ID to empty string to indicate unassignment
        visit.caregiverId = "";
        break;
      }
    }
  }
}
