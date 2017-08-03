package main

import (
	"fmt"
)

const (
	// 未訪問
	novisite = 0
	// 訪問済みかつ未訪問の隣接点あり
	visited = 1
)

var graph map[string][]*node
var start *node
var goal *node
var queue []*node

type node struct {
	weight int
	name   string
	color  int
}

func (n *node) String() string {
	return n.name
}

func (n *node) mark() {
	n.color = visited
}

func (n *node) isVisited() bool {
	if n.color == novisite {
		return false
	}
	return true
}

func init() {
	// 探索用キュー
	queue = make([]*node, 0)
	// n0が開始点, n99が到達目標
	start = &node{0, "n0", novisite}
	goal = &node{0, "n99", novisite}

	// 重みはとりあえずすべて0
	graph = make(map[string][]*node, 0)
	graph["n0"] = []*node{&node{0, "n1", novisite}, &node{0, "n6", novisite}, &node{0, "n8", novisite}}
	graph["n1"] = []*node{start, &node{0, "n2", novisite}, &node{0, "n3", novisite}}
	graph["n2"] = []*node{&node{0, "n1", novisite}, &node{0, "n11", novisite}, &node{0, "n10", novisite}}
	graph["n3"] = []*node{&node{0, "n1", novisite}, &node{0, "n12", novisite}, &node{0, "n4", novisite}}
	graph["n4"] = []*node{&node{0, "n3", novisite}, &node{0, "n13", novisite}, &node{0, "n5", novisite}}
	graph["n5"] = []*node{&node{0, "n4", novisite}, &node{0, "n6", novisite}, &node{0, "n9", novisite}}
	graph["n6"] = []*node{&node{0, "n5", novisite}, start, &node{0, "n7", novisite}}
	graph["n7"] = []*node{&node{0, "n6", novisite}, &node{0, "n9", novisite}, &node{0, "n8", novisite}}
	graph["n8"] = []*node{&node{0, "n7", novisite}, start, &node{0, "n14", novisite}}
	graph["n9"] = []*node{&node{0, "n5", novisite}, &node{0, "n7", novisite}, goal}
	graph["n10"] = []*node{&node{0, "n2", novisite}}
	graph["n11"] = []*node{&node{0, "n2", novisite}}
	graph["n12"] = []*node{&node{0, "n3", novisite}}
	graph["n13"] = []*node{&node{0, "n4", novisite}}
	graph["n14"] = []*node{&node{0, "n8", novisite}}
	graph["n99"] = []*node{&node{0, "n9", novisite}}
}

func main() {
	start.mark()
	queue = append(queue, start)
	breadthFirstSearch(graph, start, goal)
}

func breadthFirstSearch(graph map[string][]*node, start *node, goal *node) {

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		fmt.Println(v)

		if v.name == goal.name {
			fmt.Println("Goal!")
			return
		}
		for _, v2 := range graph[v.name] {
			if v2.isVisited() == false {
				v2.mark()
				queue = append(queue, v2)
			}
		}
	}
}
