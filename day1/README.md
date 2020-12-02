# Day 1

## Problem 1

### Notes

- A faster solution can be found by creating a set using `map[int]struct{}` of the values in the input, then indexing on the difference of each input
  - However, this solution does not account for 2 duplicate entries adding up to 2020 (e.g. `1010 + 1010 = 2020; 1010^2`)
- Is there a cleaner way using recursion to combine part 1 + 2? Maybe.

### Performance

```
$ time go run solution.go
the entries for part 1 are: 492, 1528. the product is 751776
the entries for part 2 are: 1258, 715, 47. the product is 42275090

real    0m0.255s
user    0m0.094s
sys     0m0.156s
```
