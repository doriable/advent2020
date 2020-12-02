# Day 1

## Problem 1

### Notes

- A faster solution can be found by creating a set using `map[int]struct{}` of the values in the input, then indexing on the difference of each input
  - However, this solution does not account for 2 duplicate entries adding up to 2020 (e.g. `1010 + 1010 = 2020; 1010^2`)
