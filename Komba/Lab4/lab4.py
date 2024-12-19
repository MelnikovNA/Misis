
from collections import deque # Очередь через либу

def bfs(graph, start):
    # queue — очередь для хранения вершин, которые нужно посетить (сначала добавляется стартовая вершина).
    # visited — множество для отслеживания уже посещенных вершин, чтобы избежать бесконечного цикла.
    # order — список, в который записывается порядок обхода вершин.
   
    # Очередь для обхода и множество посещённых вершин
    queue = deque([start])
    visited = set() # множество для отслеживания уже посещенных вершин, чтобы избежать бесконечного цикла.
    order = []  # Порядок обхода вершин

    visited.add(start)# Начинаем с начальной вершины

# Пока есть вершины в очереди: Извлекается текущая вершина (current) из очереди. Она добавляется в список order.
    while queue:
        # Извлекаем вершину из очереди
        current = queue.popleft()
        order.append(current)

        # Обходим всех соседей текущей вершины
        for neighbor in graph[current]:
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append(neighbor)

    return order

graph = {
    1: [2, 4, 6],
    2: [1, 3, 4, 5, 7],
    3: [2, 4],
    4: [1, 2, 3],
    5: [2, 6, 7],
    6: [1, 5],
    7: [2, 5],
    8: []
}

start_vertex = 1
print("Порядок обхода:", bfs(graph, start_vertex))