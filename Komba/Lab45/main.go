package main

import (
	"fmt"
)

// Edge представление ребра графа
type Edge struct {
	NodeTo int
	Weight int
}

type Graph map[int][]Edge

// Функция для добавления вершины в черпак
func addToBucket(buckets map[int][]int, vertex, distance int) {
	buckets[distance] = append(buckets[distance], vertex)
}

// Функция для удаления вершины из черпака
func removeFromBucket(buckets map[int][]int, vertex, distance int) {
	for i, v := range buckets[distance] {
		if v == vertex {
			buckets[distance] = append(buckets[distance][:i], buckets[distance][i+1:]...)
			break
		}
	}
}

// Функция для получения вершины с минимальным расстоянием
func getMinDistanceVertex(buckets map[int][]int) (int, bool) {
	for distance := range buckets {
		if len(buckets[distance]) > 0 {
			vertex := buckets[distance][0]
			buckets[distance] = buckets[distance][1:]
			return vertex, true
		}
	}
	return -1, false // Очередь пуста
}

func Dijkstra(graph Graph, start int) map[int]int {
	dist := make(map[int]int) //инициализация расстояний
	prev := make(map[int]int) // инициализация предыдущих вершин

	for v := range graph {
		dist[v] = 1000000 // Бесконечность
	}
	dist[start] = 0

	//Черпаки для быстрого поиска вершин с минимальным расстоянием
	buckets := make(map[int][]int)
	buckets[0] = []int{start}
	for len(buckets) > 0 {
		// Получаем вершину с минимальным расстоянием
		u, ok := getMinDistanceVertex(buckets)
		if !ok {
			break // Очередь пуста
		}

		// Рассматриваем соседей текущей вершины
		for _, neighbor := range graph[u] {
			v, weight := neighbor.NodeTo, neighbor.Weight
			alt := dist[u] + weight
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u

				// Обновляем черпаки
				removeFromBucket(buckets, v, dist[v])
				addToBucket(buckets, v, alt)
			}
		}
	}

	return dist
}

func main() {
	graph := Graph{
		1: []Edge{{2, 1}, {3, 6}, {6, 6}},
		2: []Edge{{1, 1}, {3, 0}, {4, 2}, {5, 5}, {7, 7}},
		3: []Edge{{1, 6}, {2, 0}, {4, 4}},
		4: []Edge{{2, 2}, {3, 4}, {7, 7}},
		5: []Edge{{2, 5}, {6, 0}, {7, 2}},
		6: []Edge{{1, 6}, {5, 0}},
		7: []Edge{{2, 7}, {4, 7}, {5, 2}},
	}

	distances := Dijkstra(graph, 1)
	fmt.Println(distances)
}
