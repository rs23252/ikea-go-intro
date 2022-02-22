## Looping / Iteration
Go has a single for loop construct that combines
- `while` condition { … }
- `do { … } while` condition
- `do { … } until` condition

into one syntax.

- `for (init statement); condition; (post statement) { … }`

The parts of a for statement are:
- `init statement`: used to initalise the loop variable; i = 0.
- `condition`: user to test if the loop is done; i < 10, true means keep looping.
- `post statement`: user to increment the loop variable; i++, i = i - 1.