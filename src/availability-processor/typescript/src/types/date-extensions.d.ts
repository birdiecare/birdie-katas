// Extend Date prototype with helper methods
declare global {
  interface Date {
    addMinutes(minutes: number): Date;
    addHours(hours: number): Date;
    subtractMinutes(minutes: number): Date;
    subtractHours(hours: number): Date;
  }
}

export {};
