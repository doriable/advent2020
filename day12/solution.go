package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var compass = []rune{'E', 'S', 'W', 'N'}

type ship struct {
	eastWest   int
	northSouth int
	w          *waypoint
}

type waypoint struct {
	left       int
	leftUnits  int
	right      int
	rightUnits int
}

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	data := strings.Split(strings.TrimSpace(string(raw)), "\n")

	w := &waypoint{
		left:       0,
		leftUnits:  10,
		right:      3,
		rightUnits: 1,
	}

	s := &ship{
		eastWest:   0,
		northSouth: 0,
		w:          w,
	}

	for _, line := range data {
		fmt.Println(line)
		if err := s.move(line); err != nil {
			os.Exit(1)
		}
	}

	fmt.Printf(
		"the manhanttan distance: %v\n",
		math.Abs(float64(s.eastWest))+math.Abs(float64(s.northSouth)),
	)

	os.Exit(0)
}

func (s *ship) move(cmd string) error {
	n, err := strconv.ParseInt(cmd[1:], 10, 64)
	if err != nil {
		return err
	}

	d := rune(cmd[0])

	if d == 'F' {
		s.moveShip(compass[s.w.left], s.w.leftUnits*int(n))
		s.moveShip(compass[s.w.right], s.w.rightUnits*int(n))
		return nil
	}

	if d == 'L' {
		s.w.left = turnLeft(s.w.left, int(n))
		s.w.right = turnLeft(s.w.right, int(n))
		return nil
	}

	if d == 'R' {
		s.w.left = turnRight(s.w.left, int(n))
		s.w.right = turnRight(s.w.right, int(n))
		return nil
	}

	return s.moveWaypoint(d, int(n))
}

func turnLeft(l, n int) int {
	i := (l - n/90) % 4
	if i < 0 {
		i = len(compass) + i
	}

	return i
}

func turnRight(l, n int) int {
	return (l + n/90) % 4
}

func (s *ship) moveShip(d rune, n int) {
	switch d {
	case 'N':
		s.northSouth += int(n)

	case 'S':
		s.northSouth -= int(n)

	case 'E':
		s.eastWest += int(n)

	case 'W':
		s.eastWest -= int(n)
	}

}

func (s *ship) moveWaypoint(d rune, n int) error {
	left := compass[s.w.left]
	right := compass[s.w.right]

	switch d {
	case 'E':
		if left == 'E' {
			s.w.leftUnits += n
			return nil
		}

		if left == 'W' {
			s.w.leftUnits -= n
			return nil
		}

		if right == 'E' {
			s.w.rightUnits += n
			return nil
		}

		if right == 'W' {
			s.w.rightUnits -= n
			return nil
		}

	case 'W':
		if left == 'W' {
			s.w.leftUnits += n
			return nil
		}

		if left == 'E' {
			s.w.leftUnits -= n
			return nil
		}

		if right == 'W' {
			s.w.rightUnits += n
			return nil
		}

		if right == 'E' {
			s.w.rightUnits -= n
			return nil
		}

	case 'N':
		if left == 'N' {
			s.w.leftUnits += n
			return nil
		}

		if left == 'S' {
			s.w.leftUnits -= n
			return nil
		}

		if right == 'N' {
			s.w.rightUnits += n
			return nil
		}

		if right == 'S' {
			s.w.rightUnits -= n
			return nil
		}

	case 'S':
		if left == 'S' {
			s.w.leftUnits += n
			return nil
		}

		if left == 'N' {
			s.w.leftUnits -= n
			return nil
		}

		if right == 'S' {
			s.w.rightUnits += n
			return nil
		}

		if right == 'N' {
			s.w.rightUnits -= n
			return nil
		}
	}

	return fmt.Errorf("invalid waypoint: %v\n", d)
}
