package solution

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type Node struct {
	label       string
	neighbors   []string
	isSmallCave bool
}

func (s *Solution) Day12Part1(fn string) (ret int) {
	lines := toLines(fn)

	// ptrs to all nodes
	nodes := make(map[string]Node)
	for _, l := range lines {
		split := strings.Split(l, "-")
		l := split[0]
		r := split[1]

		small := func(label string) bool {
			if label == "start" || label == "end" {
				return false
			}
			r, _ := utf8.DecodeRuneInString(label)
			return unicode.IsLower(r)
		}
		addNode := func(label string, neighbor string, nodes map[string]Node) {
			n := nodes[label]
			n.label = label
			n.neighbors = append(n.neighbors, neighbor)
			n.isSmallCave = small(label)
			nodes[label] = n
		}
		addNode(l, r, nodes)
		addNode(r, l, nodes)
	}

	// find all paths
	var dfs func(n Node, path map[string]bool, nodes map[string]Node) int
	dfs = func(n Node, path map[string]bool, nodes map[string]Node) int {
		if n.label == "end" {
			return 1
		}
		if (n.isSmallCave || n.label == "start") && path[n.label] {
			return 0
		}
		cnt := 0
		for _, s := range n.neighbors {
			path[n.label] = true
			cnt += dfs(nodes[s], path, nodes)
			path[n.label] = false
		}
		return cnt
	}

	return dfs(nodes["start"], make(map[string]bool), nodes)
}

func (s *Solution) Day12Part2(fn string) (ret int) {
	lines := toLines(fn)

	// ptrs to all nodes
	nodes := make(map[string]Node)
	for _, l := range lines {
		split := strings.Split(l, "-")
		l := split[0]
		r := split[1]

		small := func(label string) bool {
			if label == "start" || label == "end" {
				return false
			}
			r, _ := utf8.DecodeRuneInString(label)
			return unicode.IsLower(r)
		}
		addNode := func(label string, neighbor string, nodes map[string]Node) {
			n := nodes[label]
			n.label = label
			n.neighbors = append(n.neighbors, neighbor)
			n.isSmallCave = small(label)
			nodes[label] = n
		}
		addNode(l, r, nodes)
		addNode(r, l, nodes)
	}

	// find all paths
	var dfs func(n Node, path map[string]int, nodes map[string]Node) int
	dfs = func(n Node, path map[string]int, nodes map[string]Node) int {
		if n.label == "start" && path[n.label] == 1 {
			return 0
		}
		if n.isSmallCave && path[n.label] >= 2 {
			return 0
		}
		twice := false
		for k, v := range path {
			if v >= 2 && twice && nodes[k].isSmallCave {
				return 0
			}
			if v >= 2 && nodes[k].isSmallCave {
				twice = true
			}
		}
		if n.label == "end" {
			return 1
		}

		cnt := 0
		for _, s := range n.neighbors {
			if s == "start" {
				continue
			}
			path[n.label] = path[n.label] + 1
			cnt += dfs(nodes[s], path, nodes)
			path[n.label] = path[n.label] - 1
		}
		return cnt
	}

	return dfs(nodes["start"], make(map[string]int), nodes)
}
