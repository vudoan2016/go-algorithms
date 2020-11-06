package main

import "fmt"

type Id int
type Weight int

// Directed weighted graph
type Node struct {
	id      Id
	visited bool
	edges   map[Id]Weight
}

type Graph struct {
	nodes map[Id]*Node
}

func New() *Graph {
	return &Graph{
		nodes: map[Id]*Node{},
	}
}

func (g *Graph) addNode(id Id) {
	g.nodes[id] = &Node{id: id, edges: make(map[Id]Weight)}
}

// Edge points from 'fromId' to 'toId'
func (g *Graph) addEdge(fromId Id, toId Id, weight Weight) {
	g.nodes[fromId].edges[toId] = weight
}

func (g Graph) print() {
	for _, n := range g.nodes {
		fmt.Printf("Node %d: ", n.id)
		for toId, w := range n.edges {
			fmt.Printf("[%d, %d] ", toId, w)
		}
		fmt.Println()
	}
}

func (g Graph) traverseHelper(source Id, stack *[]Id) {
	if _, found := g.nodes[source]; !found {
		return
	}
	g.nodes[source].visited = true
	*stack = append(*stack, source)

	for neighbor := range g.nodes[source].edges {
		if !g.nodes[neighbor].visited {
			g.traverseHelper(neighbor, stack)
			if (*stack)[0] == source {
				fmt.Println(*stack)
				*stack = (*stack)[:1]
			}
		}
	}
}

func (g Graph) dfTraverse(root Id) {
	for _, n := range g.nodes {
		n.visited = false
	}
	var stack []Id
	g.traverseHelper(root, &stack)
}

func (g Graph) dfsHelper(root, key Id) bool {
	if root == key {
		return true
	}
	g.nodes[root].visited = true
	for neighbor := range g.nodes[root].edges {
		if !g.nodes[neighbor].visited {
			if g.dfsHelper(neighbor, key) {
				return true
			}
		}
	}
	return false
}

func (g Graph) dfs(root, key Id) bool {
	for _, n := range g.nodes {
		n.visited = false
	}
	return g.dfsHelper(root, key)
}

// BFS is used to find SP between 2 vertices of an unweighted graph
func (g Graph) bfsTraverse(root Id) {
	for _, n := range g.nodes {
		n.visited = false
	}
	q := []Id{}
	q = append(q, root)
	for len(q) > 0 {
		fmt.Printf("%d ", q[0])
		for neighbor := range g.nodes[q[0]].edges {
			if !g.nodes[neighbor].visited {
				g.nodes[q[0]].visited = true
				q = append(q, neighbor)
			}
		}
		q = q[1:]
	}
	fmt.Println()
}

func graphTest() {
	g := New()
	g.addNode(7)
	g.addNode(11)
	g.addNode(9)
	g.addNode(28)
	g.addNode(5)
	g.addEdge(7, 11, 4)
	g.addEdge(11, 9, 6)
	g.addEdge(9, 28, 8)
	g.addEdge(7, 5, 6)
	g.addEdge(5, 28, 2)
	g.print()
	g.dfTraverse(7)
	root := Id(11)
	key := Id(5)
	fmt.Printf("Key %d found: %t\n", key, g.dfs(root, key))

	g.bfsTraverse(7)
}

// with memoization
func fibo(n int, cnt *int, cache *[]int) int {
	*cnt += 1
	if n == 0 || n == 1 {
		return n
	} else if *cache != nil {
		if (*cache)[n-1] == 0 {
			(*cache)[n-1] = fibo(n-1, cnt, cache)
		}
		if (*cache)[n-2] == 0 {
			(*cache)[n-2] = fibo(n-2, cnt, cache)
		}

		return (*cache)[n-1] + (*cache)[n-2]
	} else {
		return fibo(n-1, cnt, cache) + fibo(n-2, cnt, cache)
	}
}

func fiboTest() {
	n, cnt := 10, 0
	cache := make([]int, n) // array with non-constant length
	f := fibo(n, &cnt, &cache)
	fmt.Printf("fibo(%d) = %d, recursed %d times\n", n, f, cnt)
	cache = nil
	f = fibo(n, &cnt, &cache)
	fmt.Printf("fibo(%d) = %d, recursed %d times\n", n, f, cnt)
}

// longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	max := 0
	for i := 0; i < len(s); i++ {
		count := 0
		for j := i; j < len(s); j++ {
			if m[s[j]] == 0 {
				m[s[j]] = 1
				count += 1
			} else {
				break
			}
			if count > max {
				max = count
			}
		}
		m = make(map[byte]int)
	}
	return max
}

func lengthOfLongestSubstringTest() {
	s := ""
	fmt.Printf("Longest substring of '%s' = %d\n", s, lengthOfLongestSubstring(s))
}

// find all anagrams of 'a' in 's'
func findAnagrams(s string, a string) []int {
	var result []int
	m := make(map[rune]int)
	for i := 0; (i + len(a)) <= len(s); i++ {
		isAnagram := true
		for _, c := range a {
			m[c] += 1
		}
		for _, c := range s[i : i+len(a)] {
			m[c] -= 1
		}
		for _, c := range a {
			if m[c] != 0 {
				isAnagram = false
			}
			m[c] = 0
		}
		if isAnagram {
			result = append(result, i)
		}
	}
	return result
}

func anagramTest() {
	strs := []string{"cbaebabacd", "cbaebaaacd", "abaebabaab", "abab", ""}
	a := []string{"abc", "aaa", "aab", "ab", ""}

	for i, s := range strs {
		fmt.Println("Anagrams @", findAnagrams(s, a[i]))
	}
}

// print out all permutations of rune 'r'
func permutation(r []rune, i int, cnt *int, result map[string]bool) {
	if i >= len(r) {
		if _, found := result[string(r)]; found {
			panic(fmt.Sprintf("Duplicate permutation %s", string(r)))
		}
		*cnt++
		return
	} else {
		for j := 0; j <= i; j++ {
			r[j], r[i] = r[i], r[j]
			permutation(r, i+1, cnt, result)
			r[j], r[i] = r[i], r[j]
		}
	}
}

func fact(x int) int {
	if x <= 1 {
		return 1
	}
	return x * fact(x-1)
}

func permutationTest() {
	strs := []string{"ABCDEFGHIJKLMNOPQRSTUV", "A", ""}
	for _, s := range strs {
		cnt := 0
		result := make(map[string]bool)
		permutation([]rune(s), 1, &cnt, result)
		expect := fact(len(s))
		if cnt != expect {
			panic(fmt.Sprintf("cnt = %d, expected = %d", cnt, expect))
		}
	}
}

func main() {
	//lengthOfLongestSubstringTest()
	//anagramTest()
	//permutationTest()
	fiboTest()
	//graphTest()
}
