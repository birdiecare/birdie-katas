# Mars Rover Kata

(Taken with gratitude from: https://kata-log.rocks/mars-rover-kata)

## Your Task

You’re part of the team that explores Mars by sending remotely controlled vehicles to the surface of the planet. Develop an API that translates the commands sent from earth to instructions that are understood by the rover.

The rover operates like a very basic tank. It can move forwards or backwards in a straight line, and it can turn on the spot in 90-degree (quarter circle) turns.

The rover boots up once it lands on a planet, and it is given its landing coordinates (x,y) and the direction it is facing.

## Requirements

- You are given the initial starting point (x,y) of a rover and the direction (N,S,E,W) it is facing.
- The rover receives a character array of commands.
- Implement commands that move the rover forward/backward (f,b).
- Implement commands that turn the rover left/right (l,r).
- Implement wrapping at edges. But be careful, planets are spheres. Connect the x edge to the other x edge, so (1,1) for x-1 to (5,1), but connect vertical edges towards themselves in inverted coordinates, so (1,1) for y-1 connects to (5,1).
- Implement obstacle detection before each move to a new square. If a given sequence of commands encounters an obstacle, the rover moves up to the last possible point, aborts the sequence and reports the obstacle.

## Rules

- TDD please.
- Consider edge cases and exceptions.
