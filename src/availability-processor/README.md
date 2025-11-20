# Caregiver Availability Processor

Birdie's platform includes a rostering solution where caregivers are assigned to patient visits. The system processes availability events to automatically manage visit assignments when caregivers become unavailable.

## Getting Started

```bash

cd golang

# Run tests to see current functionality
go test ./...

# Run with verbose output to see test details
go test -v ./src
```

## Exercise Tasks

The system currently handles **permanent unavailability events**. When a caregiver becomes permanently unavailable (e.g., leaves the company, goes on extended leave), the system automatically unassigns them from all upcoming visits starting from the effective date.

Starting from the top, we'll work through the tasks below and expand the processor to handle more requirements. We don't expect to complete all of the tasks - we're much more interested in understanding how you approach a problem.

Whilst working, please keep in mind the following non-functional requirements:

* We value test-driven development and would like to see this used throughout the exercise today.
* Reliability is vitally important to us. A mistake could lead to a missed visit.
* We love pairing and will appreciate if you communicate your thinking and progress as you go, as well as adopting a gradual approach which is easy for your pair to understand.

### Task 1: Add Multi-Tenancy Support

The current implementation doesn't handle multiple tenants (care agencies). When a caregiver is made inactive in one agency, this should not affect their visits with another agency.

Let's add multi-tenancy support so only visits for a given agency are unassigned. You'll see that `Visit` and `CaregiverPermanentUnavailabilityEvent` already have a `TenantId` property.

### Task 2: Process Absence Events

Currently, only permanent unavailability is supported. We need to add support for **temporary absences** using the `CaregiverAbsenceBookedEvent`. These absences have a start and an end date. If a caregiver is absent, they should be unassigned from any visits during that period.

### Task 3: Invalid Events

We cannot rely on all the events that come to us being valid. If an invalid event arrives, we should exit with a non-zero status.
