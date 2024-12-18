import heapq
import matplotlib.pyplot as plt
import networkx as nx

# Алгоритм Дейкстры
def dijkstra(graph, start, target):
    distances = {node: float("inf") for node in graph} # хранит минимальные расстояния от стартового узла до всех остальных узлов.
    distances[start] = 0
    priority_queue = [(0, start)] # используется для хранения узлов с их текущими минимальными расстояниями.
    predecessors = {node: None for node in graph} # хранит предшествующие узлы для восстановления пути.

    while priority_queue:
        current_distance, current_node = heapq.heappop(priority_queue)

        if current_node == target:   # Алгоритм пришёл в финальную точку
            break

# После извлечения узла из очереди алгоритм больше не пытается обновлять его расстояние. Это обеспечивается строкой:
        if current_distance > distances[current_node]:
            continue

        for neighbor, weight in graph[current_node]:
            distance = current_distance + weight
            if distance < distances[neighbor]:  #Узел добавляется в очередь только с меньшим расстоянием,
                distances[neighbor] = distance  #чем записанное ранее. Это гарантирует, что более длинные пути до этого узла не будут обрабатываться.
                predecessors[neighbor] = current_node
                heapq.heappush(priority_queue, (distance, neighbor)) # Узел с минимальным текущим расстоянием всегда извлекается первым.

# После завершения основного цикла путь восстанавливается с помощью предшественников (predecessors):
    path = []
    current = target
    while current is not None:
        path.append(current)
        current = predecessors[current]

    path.reverse()
    return distances[target], path

# Граф с весами рёбер
graph = {
    1: [(2, 1), (3, 6), (6, 6)],
    2: [(1, 1), (3, 0), (4, 2), (5, 5), (7, 7)],
    3: [(1, 6), (2, 0), (4, 4)],
    4: [(2, 2), (3, 4), (7, 7)],
    5: [(2, 5), (6, 0), (7, 2)],
    6: [(1, 6), (5, 0)],
    7: [(2, 7), (4, 7), (5, 2)],
    8: []
}

start_node = 1
target_node = 7

# Найти кратчайший путь
shortest_distance, shortest_path = dijkstra(graph, start_node, target_node)
print(f"Кратчайшее расстояние от {start_node} до {target_node}: {shortest_distance}")
print(f"Кратчайший путь: {shortest_path}")

# Визуализация графа
def visualize_graph(graph, shortest_path):
    G = nx.DiGraph()

    # Добавляем вершины и рёбра с весами
    for node, edges in graph.items():
        for neighbor, weight in edges:
            G.add_edge(node, neighbor, weight=weight)

    pos = nx.spring_layout(G)  # Компоновка графа

    plt.figure(figsize=(12, 8))

    # Рисуем вершины и рёбра
    nx.draw_networkx_nodes(G, pos, node_color='lightblue', node_size=2000)
    nx.draw_networkx_labels(G, pos, font_size=14, font_color='black')
    nx.draw_networkx_edges(G, pos, edge_color='gray', arrows=True)

    # Выделяем кратчайший путь
    path_edges = list(zip(shortest_path, shortest_path[1:]))
    nx.draw_networkx_edges(G, pos, edgelist=path_edges, edge_color='red', width=2.5, arrows=True)

    # Добавляем веса рёбер
    edge_labels = nx.get_edge_attributes(G, 'weight')
    nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels, font_size=12)

    plt.title("Граф с кратчайшим путём (Дейкстра)", fontsize=16)
    plt.show()

# Визуализация графа с кратчайшим путём
visualize_graph(graph, shortest_path)
print(f"Кратчайшее расстояние от {start_node} до {target_node}: {shortest_distance}")
print(f"Кратчайший путь: {shortest_path}")