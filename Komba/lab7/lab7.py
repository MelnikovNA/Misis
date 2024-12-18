from collections import defaultdict  # Импортируем defaultdict из модуля collections для удобного создания словаря с автоматическими значениями по умолчанию

# Определение класса графа
class Graph:
    def __init__(self, vertices):
        # Конструктор для инициализации графа
        self.V = vertices  # количество вершин в графе
        self.graph = defaultdict(list)  # представление графа в виде списка смежности
        self.capacity = {}  # словарь для хранения пропускных способностей ребер

    def add_edge(self, u, v, capacity):
        # Метод для добавления ребра с заданной пропускной способностью
        self.graph[u].append(v)  # добавляем вершину v в список смежности вершины u
        self.graph[v].append(u)  # добавляем вершину u в список смежности вершины v (граф двунаправленный)
        self.capacity[(u, v)] = capacity  # устанавливаем пропускную способность для ребра (u, v)
        self.capacity[(v, u)] = capacity  # для обратного ребра (v, u) пропускная способность такая же

    def _dfs(self, source, sink, visited, flow):
        # Приватный метод для поиска пути из источника в сток с использованием DFS
        visited.add(source)  # отмечаем текущую вершину как посещенную
        if source == sink:  # если достигли стока
            return flow  # возвращаем текущий поток как возможный

        # Итерация по всем соседям текущей вершины
        for neighbor in self.graph[source]:
            # Проверяем, что сосед не посещен и пропускная способность ребра больше 0
            if neighbor not in visited and self.capacity[(source, neighbor)] > 0:
                # Находим минимальный поток на пути
                min_flow = min(flow, self.capacity[(source, neighbor)])
                # Рекурсивно запускаем DFS для поиска пути с минимальным потоком
                result = self._dfs(neighbor, sink, visited, min_flow)

                # Если нашли путь до стока, обновляем остаточную сеть
                if result > 0:
                    self.capacity[(source, neighbor)] -= result  # уменьшаем пропускную способность ребра (source, neighbor)
                    self.capacity[(neighbor, source)] += result  # увеличиваем пропускную способность обратного ребра (neighbor, source)
                    return result  # возвращаем найденный поток

        return 0  # если путь не найден, возвращаем 0

    def ford_fulkerson(self, source, sink):
        # Основной метод для вычисления максимального потока с использованием алгоритма Форда-Фалкерсона
        max_flow = 0  # инициализируем максимальный поток как 0

        while True:  # бесконечный цикл для поиска путей
            visited = set()  # множество для хранения посещенных вершин
            flow = self._dfs(source, sink, visited, float('Inf'))  # запускаем DFS для поиска пути с максимальным потоком

            if flow == 0:  # если больше нет путей, которые можно использовать
                break  # выходим из цикла

            max_flow += flow  # добавляем найденный поток к общему максимальному потоку

        return max_flow  # возвращаем максимальный поток

# Пример использования
if __name__ == "__main__":
    import networkx as nx  # Импортируем библиотеку для работы с графами
    import matplotlib.pyplot as plt  # Импортируем библиотеку для построения графиков

    # Создаем граф с 8 вершинами
    graph = Graph(8)  # 8 вершин
    edges = [
        (1, 2, 8), (1, 4, 3),
        (2, 1, 8), (2, 3, 2), (2, 4, 4), (2, 5, 5),
        (3, 2, 2), (3, 5, 4),
        (4, 1, 3), (4, 2, 4), (4, 7, 5), (4, 6, 5),
        (5, 3, 4), (5, 2, 5), (5, 7, 2), (5, 8, 3),
        (6, 4, 5), (6, 7, 4),
        (7, 6, 4), (7, 4, 5), (7, 5, 2), (7, 8, 8),
        (8, 7, 8), (8, 5, 3),
    ]


    # Добавляем рёбра в граф
    for u, v, capacity in edges:
        graph.add_edge(u, v, capacity)

    source, sink = 6, 3  # Задаем источник и сток
    max_flow = graph.ford_fulkerson(source, sink)  # Вычисляем максимальный поток
    print(f"Максимальный поток: {max_flow}")  # Выводим результат


    # Визуализация графа
    G = nx.DiGraph()  # Создаем направленный граф с помощью библиотеки networkx
    for u, v, capacity in edges:
        G.add_edge(u, v, capacity=capacity)  # добавляем ребро (u, v) с пропускной способностью
        G.add_edge(v, u, capacity=capacity)  # добавляем обратное ребро (v, u)

    pos = nx.spring_layout(G)  # Генерируем расположение вершин графа
    edge_labels = nx.get_edge_attributes(G, 'capacity')  # Получаем пропускные способности рёбер
    # Рисуем граф
    nx.draw(G, pos, with_labels=True, node_size=700, node_color='skyblue', font_size=15, font_weight='bold', arrows=False)
    # Добавляем подписи к рёбрам
    nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels, font_size=12)
    plt.title("Исходный граф")  # Устанавливаем заголовок графа
    plt.show()  # Отображаем граф