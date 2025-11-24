export interface CaregiverPermanentUnavailabilityEvent {
  /** Unique identifier for the permanent unavailability event */
  id: string;
  /** Unique identifier for the tenant (care agency) */
  tenantId: string;
  /** Unique identifier for the caregiver */
  caregiverId: string;
  /** Time when the permanent unavailability starts */
  effectiveFrom: Date;
}

export interface CaregiverAbsenceBookedEvent {
  /** Unique identifier for the absence event */
  id: string;
  /** Unique identifier for the tenant (care agency) */
  tenantId: string;
  /** Unique identifier for the caregiver */
  caregiverId: string;
  /** Start time of the absence */
  startTime: Date;
  /** End time of the absence */
  endTime: Date;
}
