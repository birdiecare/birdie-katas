# Lanternfish kata

> Note: this is the day 6 of the Advent of Code 2021. The instructions are simplified to make them easier to understand. You can find the original instructions [here](https://adventofcode.com/2021/day/6).

## The brief

You're a marine biologist, studying lanternfish. These are famous for being found in massive numbers: they must spawn quickly to reach such large numbers, maybe exponentially quickly?

After some research, you find out that they have a very stable reproduction cycle: a lanternfish spawns a new lanternfish every 7 days. There's one exception: a _newborn_ lanternfish takes an additional two days before its first reproduction cycle, therefore taking 9 days for its first cycle.

You can model each fish as a single number that represents the number of days until it creates a new lanternfish. Suppose you have a lanternfish with an internal timer value of 3:

- After one day, its internal timer would become 2.
- After another day, its internal timer would become 1.
- After another day, its internal timer would become 0.
- After another day, its internal timer would reset to 6, and it would create a new lanternfish with an internal timer of 8.
- After another day, the first lanternfish would have an internal timer of 5, and the second lanternfish would have an internal timer of 7.

## Your task

You observed a sample of 5 lanternfish, you know how many days are left until they each spawn a new lanternfish: respectively `3,4,3,1,2`.

By hand, you simulate the growth of this school for a few days:


```
Initial state: 3,4,3,1,2
After  1 day:  2,3,2,0,1
After  2 days: 1,2,1,6,0,8
After  3 days: 0,1,0,5,6,7,8
After  4 days: 6,0,6,4,5,6,7,8,8
After  5 days: 5,6,5,3,4,5,6,7,7,8
After  6 days: 4,5,4,2,3,4,5,6,6,7
After  7 days: 3,4,3,1,2,3,4,5,5,6
After  8 days: 2,3,2,0,1,2,3,4,4,5
After  9 days: 1,2,1,6,0,1,2,3,3,4,8
After 10 days: 0,1,0,5,6,0,1,2,2,3,7,8
After 11 days: 6,0,6,4,5,6,0,1,1,2,6,7,8,8,8
After 12 days: 5,6,5,3,4,5,6,0,0,1,5,6,7,7,7,8,8
After 13 days: 4,5,4,2,3,4,5,6,6,0,4,5,6,6,6,7,7,8,8
After 14 days: 3,4,3,1,2,3,4,5,5,6,3,4,5,5,5,6,6,7,7,8
After 15 days: 2,3,2,0,1,2,3,4,4,5,2,3,4,4,4,5,5,6,6,7
After 16 days: 1,2,1,6,0,1,2,3,3,4,1,2,3,3,3,4,4,5,5,6,8
After 17 days: 0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8
After 18 days: 6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8
```

It's getting unwieldy to do by hand but you're curious: how many lanternfish would there be after 80 days?
