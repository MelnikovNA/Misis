package main

import "fmt"

//Andrey 0 (1, 0, 3, 2)
//Boris 1	(1, 2, 3, 0)
//Volodya 2 (3, 2, 0, 1)
//Gocha 3	(1, 0, 2, 3)

//Elena 0 (3, 2, 1, 0)
//Zhana 1 (3, 2, 1, 0)
//Zina 2  (0, 1, 2, 3)
//Irina 3 (2, 3, 1, 0)

// Структура для описания мужчины
type Man struct {
	name        string
	preferences []int
	isEngaged   bool
}

// Структура для описания женщины
type Woman struct {
	name           string
	preferences    []int
	currentPartner int
}

// проверка предпочтений женщины
func prefers(preferences []int, newMan, currentMan int) bool {
	for _, man := range preferences {
		if man == newMan {
			return true
		}
		if man == currentMan {
			return false
		}
	}
	return false
}

func main() {
	//Мужчины и женщины с их предпочтениями
	men := []Man{
		{name: "Andrey", preferences: []int{1, 0, 3, 2}, isEngaged: false},
		{name: "Boris", preferences: []int{1, 2, 3, 0}, isEngaged: false},
		{name: "Volodya", preferences: []int{3, 2, 0, 1}, isEngaged: false},
		{name: "Gocha", preferences: []int{1, 0, 2, 3}, isEngaged: false},
	}

	women := []Woman{
		{name: "Elena", preferences: []int{3, 2, 1, 0}, currentPartner: -1},
		{name: "Zhana", preferences: []int{3, 2, 1, 0}, currentPartner: -1},
		{name: "Zina", preferences: []int{0, 1, 2, 3}, currentPartner: -1},
		{name: "Irina", preferences: []int{2, 3, 1, 0}, currentPartner: -1},
	}
	married := stableMarriage(men, women)
	fmt.Println("Стабильные пары")
	for _, woman := range married {
		fmt.Printf("Женщина %s в браке с мужчиной %s\n", woman.name, men[woman.currentPartner].name)
	}
}

// Алгоритм предложений
func stableMarriage(men []Man, women []Woman) []Woman {
	unmatchedMen := len(men)

	for unmatchedMen > 0 {
		for i := range men {
			if men[i].isEngaged {
				continue
			}

			// Шаг 1: Мужчина делает предложение лучшей женщине
			if len(men[i].preferences) == 0 {
				continue // Если список пуст, пропускаем
			}
			womanIndex := men[i].preferences[0]         // Берём первую женщину из списка
			men[i].preferences = men[i].preferences[1:] // Вычёркиваем её

			// Женщина, к которой сделано предложение
			woman := &women[womanIndex]

			if woman.currentPartner == -1 {
				// Если женщина незамужем, она принимает предложение
				woman.currentPartner = i
				men[i].isEngaged = true
				unmatchedMen--
			} else {
				// Если женщина замужем, она сравнивает нового кандидата с текущим мужем
				currentMan := woman.currentPartner
				newMan := i

				if prefers(woman.preferences, newMan, currentMan) {
					// Если новый кандидат предпочтительнее, развод и брак с ним
					men[currentMan].isEngaged = false
					woman.currentPartner = newMan
					men[newMan].isEngaged = true
				}
			}
		}
	}
	return women
}
