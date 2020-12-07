package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type inner struct {
	colour string
	num    int
}

func main() {
	raw, err := ioutil.ReadFile("example2")
	if err != nil {
		os.Exit(1)
	}

	data := strings.Split(strings.TrimSpace(string(raw)), "\n")
	g := map[string][]inner{}

	for _, line := range data {
		tokens := strings.Split(line, " ")
		colour := fmt.Sprintf("%v %v", tokens[0], tokens[1])
		addVertices(g, colour, tokens[4:])
	}

	allPaths := map[string][][]inner{}
	for colour := range g {
		paths := dfs(g, []inner{inner{colour: colour}}, [][]inner{})
		allPaths[colour] = paths
	}

	shinyGold := []inner{}
	for _, pathGroup := range allPaths {
		for _, path := range pathGroup {
			for i, bag := range path {
				if i > 0 && bag.colour == "shiny gold" {
					shinyGold = append(shinyGold, path[:i]...)
				}
			}
		}
	}

	set := map[string]struct{}{}
	for _, s := range shinyGold {
		set[s.colour] = struct{}{}
	}

	// for k, v := range getPathMap(allPaths["shiny gold"]) {
	// 	fmt.Printf("k: %v, v: %v\n", k, v)
	// }

	fmt.Printf("number of paths that contain shiny gold: %v\n", len(set))
	fmt.Printf("each shiny gold bag contains: %v bags\n", getCount(
		getPathMap(allPaths["shiny gold"]),
		[]inner{inner{
			colour: "shiny gold",
			num:    0,
		}},
		0,
	))

	os.Exit(0)
}

func getCount(data map[inner][]inner, running []inner, total int) int {
	fmt.Println("Start")
	fmt.Println(running)
	d := running[len(running)-1]
	fmt.Println(d)
	fmt.Println(total)

	if v, ok := data[d]; ok {
		for _, bag := range v {
			if d.num == 0 {
				d.num = 1
			}

			running = append(running, bag)
			for i := 0; i < d.num; i++ {
				total += getCount(data, running, total)
			}
		}
	} else {
		return d.num
	}

	return total
}

func getPathMap(paths [][]inner) map[inner][]inner {
	pathMap := map[inner][]inner{}
	for _, path := range paths {
		directPaths(path, pathMap)
	}
	return pathMap
}

func directPaths(paths []inner, pathMap map[inner][]inner) map[inner][]inner {
	if len(paths) == 1 {
		return pathMap
	}

	v, ok := pathMap[paths[0]]
	if ok {
		if !in(paths[1], v) {
			pathMap[paths[0]] = append(v, paths[1])
		}
	} else {
		pathMap[paths[0]] = []inner{paths[1]}
	}

	return directPaths(paths[1:], pathMap)
}

func in(n inner, v []inner) bool {
	for _, i := range v {
		if n.colour == i.colour {
			return true
		}
	}
	return false
}

func addVertices(g map[string][]inner, colour string, tokens []string) {
	if len(tokens) == 0 || tokens[0] == "no" {
		return
	}

	if v, ok := g[colour]; ok {
		num, err := strconv.ParseInt(tokens[0], 10, 64)
		if err != nil {
			fmt.Println("as.dlkfjlksajdf")
		}
		g[colour] = append(v, inner{
			colour: fmt.Sprintf("%v %v", tokens[1], tokens[2]),
			num:    int(num),
		})
	} else {
		num, err := strconv.ParseInt(tokens[0], 10, 64)
		if err != nil {
			fmt.Println("sldkjflksdf")
		}
		g[colour] = []inner{inner{
			colour: fmt.Sprintf("%v %v", tokens[1], tokens[2]),
			num:    int(num),
		}}
	}

	addVertices(g, colour, tokens[4:])
}

func dfs(data map[string][]inner, path []inner, paths [][]inner) [][]inner {
	d := path[len(path)-1]
	if v, ok := data[d.colour]; ok {
		for _, i := range v {
			path := append(path, i)
			paths = dfs(data, path, paths)
		}
	} else {
		paths = append(paths, path)
	}
	return paths
}
