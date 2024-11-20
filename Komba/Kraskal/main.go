package main

import (
	"fmt"
	"sort"
)

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
	parent []int // Родительский элемент для каждой вершины
	rank   []int // Для оптимизации по рангу
}

// Функция для инициализации графа
func NewGraph(vertices int) *Graph {
	return &Graph{vertices: vertices, edges: []Edge{}}
}

func (g *Graph) AddEdge(u, v, weight int) {
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
		ds.parent[x] = ds.Find(ds.parent[x]) // сжатие пути
	}
	return ds.parent[x]
}

func (ds *DisjoinSet) Union(x, y int) bool {
	rootX := ds.Find(x)
	rootY := ds.Find(y)
	if rootX != rootY {
		// Объединение по рангу
		if ds.rank[rootX] > ds.rank[rootY] {
			ds.parent[rootY] = rootX
		} else if ds.rank[rootX] < ds.rank[rootY] {
			ds.parent[rootX] = rootY
		} else {
			ds.parent[rootY] = rootX
			ds.rank[rootX]++
		}
		return true
	}
	return false
}

func Kraskal(g *Graph) []Edge {

	// Сортировка рёбер по весу
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	// Инициализация множества (MakeSet)
	ds := NewDisjoinSet(g.vertices)

	// Массив для хранения рёбер минимального остова
	var mst []Edge

	// Процесс добавления рёбер в минимальное остовное дерево
	for _, edge := range g.edges {
		//проверяем на цикл через Union
		if ds.Union(edge.u, edge.v) {
			mst = append(mst, edge)
		}
	}
	return mst
}

func main() {
	g := NewGraph(4) // Создаём граф с 4 вершинами

	// Добавляем рёбра в граф
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 6)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 4)

	//применение алгоритма Краскала
	mst := Kraskal(g)
	fmt.Println("Рёбра минимального остова:")
	for _, edge := range mst {
		fmt.Printf("%d - %d : %d\n", edge.u, edge.v, edge.weight)
	}
}
