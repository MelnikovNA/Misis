package main

// Структура для представления ребра
type Edge struct {
	u, v, weight int
}

// Структура для ребра
type Graph struct {
	vertices int
	edges    []Edge
}

// Структура для представления множества с объединением и путём сжатия
type DisjoinSet struct {
	parent []int
	rank   []int
}

// Функция для инициализации графа
func NewGraph(vertices int) *Graph {
	return &Graph{vertices: vertices, edges: []Edge{}}
}

func (g *Graph) addEdge(u, v, weight int) {
	g.edges = append(g.edges, Edge{u: u, v: v, weight: weight})
}

// Функция для инициализации множества (MakeSet)
func NewDisjoinSet(n int) *DisjoinSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i //каждая вершина является своим родителем
	}
	return &DisjoinSet{parent: parent, rank: rank}
}

func (ds *DisjoinSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) //Сжатие пути
	}
	return ds.parent[x]
}

//fucn :=

func main() {
	g := NewGraph(4) // Создаём граф с 4 вершинами

	// Добавляем рёбра в граф
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 6)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 4)
}
