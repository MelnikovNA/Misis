package main

import (
	"testing"
)

func TestStableMarriage5to4(t *testing.T) {
	// Входные данные
	men := []Man{
		{name: "Andrey", preferences: []int{1, 0, 3, 2}, isEngaged: false},
		{name: "Boris", preferences: []int{1, 2, 3, 0}, isEngaged: false},
		{name: "Volodya", preferences: []int{3, 2, 0, 1}, isEngaged: false},
		{name: "Gocha", preferences: []int{1, 0, 2, 3}, isEngaged: false},
		{name: "Kolya", preferences: []int{1, 0, 2, 3}, isEngaged: false},
	}

	women := []Woman{
		{name: "Elena", preferences: []int{3, 2, 1, 0, 4}, currentPartner: -1},
		{name: "Zhana", preferences: []int{3, 2, 1, 0, 4}, currentPartner: -1},
		{name: "Zina", preferences: []int{0, 1, 2, 3, 4}, currentPartner: -1},
		{name: "Irina", preferences: []int{2, 3, 1, 0, 4}, currentPartner: -1},
	}

	// Запуск алгоритма
	finalMen, finalWomen := stableMarriage(men, women)

	// Ожидаемые результаты
	expectedPairs := map[string]string{
		"Andrey":  "Elena",
		"Boris":   "Zina",
		"Volodya": "Irina",
		"Gocha":   "Zhana",
	}
	expectedUnmatched := "Kolya"

	// Проверка пар
	for _, woman := range finalWomen {
		if woman.currentPartner != -1 {
			manName := finalMen[woman.currentPartner].name
			if expectedPairs[manName] != woman.name {
				t.Errorf("Ожидалось: %s - %s, получено: %s - %s", manName, expectedPairs[manName], manName, woman.name)
			}
		}
	}

	// Проверка мужчин без пары
	for _, man := range finalMen {
		if !man.isEngaged && man.name != expectedUnmatched {
			t.Errorf("Ожидалось, что мужчина %s останется без пары, но это не так", man.name)
		}
	}
}

func TestStableMarriage4to4(t *testing.T) {
	// Входные данные
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

	// Запуск алгоритма
	finalMen, finalWomen := stableMarriage(men, women)

	// Ожидаемые результаты
	expectedPairs := map[string]string{
		"Andrey":  "Elena",
		"Boris":   "Zina",
		"Volodya": "Irina",
		"Gocha":   "Zhana",
		"Kolya":   "Alisa",
	}

	// Проверка пар
	for _, woman := range finalWomen {
		if woman.currentPartner != -1 {
			manName := finalMen[woman.currentPartner].name
			if expectedPairs[manName] != woman.name {
				t.Errorf("Ожидалось: %s - %s, получено: %s - %s", manName, expectedPairs[manName], manName, woman.name)
			}
		}
	}
}

func TestStableMarriage5to5(t *testing.T) {
	// Входные данные
	men := []Man{
		{name: "Andrey", preferences: []int{1, 0, 4, 3, 2}, isEngaged: false},
		{name: "Boris", preferences: []int{1, 4, 2, 3, 0}, isEngaged: false},
		{name: "Volodya", preferences: []int{3, 2, 0, 1}, isEngaged: false},
		{name: "Gocha", preferences: []int{1, 4, 0, 2, 3}, isEngaged: false},
		{name: "Kolya", preferences: []int{1, 0, 2, 3, 4}, isEngaged: false},
	}

	women := []Woman{
		{name: "Elena", preferences: []int{3, 2, 1, 0, 4}, currentPartner: -1},
		{name: "Zhana", preferences: []int{3, 2, 1, 0, 4}, currentPartner: -1},
		{name: "Zina", preferences: []int{0, 1, 2, 3, 4}, currentPartner: -1},
		{name: "Irina", preferences: []int{2, 3, 1, 0, 4}, currentPartner: -1},
		{name: "Alisa", preferences: []int{2, 4, 1, 0, 3}, currentPartner: -1},
	}

	// Запуск алгоритма
	finalMen, finalWomen := stableMarriage(men, women)

	// Ожидаемые результаты
	expectedPairs := map[string]string{
		"Andrey":  "Elena",
		"Boris":   "Alisa",
		"Volodya": "Irina",
		"Gocha":   "Zhana",
		"Kolya":   "Zina",
	}

	// Проверка пар
	for _, woman := range finalWomen {
		if woman.currentPartner != -1 {
			manName := finalMen[woman.currentPartner].name
			if expectedPairs[manName] != woman.name {
				t.Errorf("Ожидалось: %s - %s, получено: %s - %s", manName, expectedPairs[manName], manName, woman.name)
			}
		}
	}
}

func TestEqualPreferences(t *testing.T) {
	// Входные данные с одинаковыми предпочтениями
	men := []Man{
		{name: "Andrey", preferences: []int{0, 1, 2, 3}, isEngaged: false},
		{name: "Boris", preferences: []int{0, 1, 2, 3}, isEngaged: false},
		{name: "Volodya", preferences: []int{0, 1, 2, 3}, isEngaged: false},
		{name: "Gocha", preferences: []int{0, 1, 2, 3}, isEngaged: false},
		{name: "Kolya", preferences: []int{0, 1, 2, 3}, isEngaged: false},
	}

	women := []Woman{
		{name: "Elena", preferences: []int{0, 1, 2, 3, 4}, currentPartner: -1},
		{name: "Zhana", preferences: []int{0, 1, 2, 3, 4}, currentPartner: -1},
		{name: "Zina", preferences: []int{0, 1, 2, 3, 4}, currentPartner: -1},
		{name: "Irina", preferences: []int{0, 1, 2, 3, 4}, currentPartner: -1},
	}

	// Запуск алгоритма
	finalMen, _ := stableMarriage(men, women)

	// Проверка количества мужчин без пары
	unmatchedCount := 0
	for _, man := range finalMen {
		if !man.isEngaged {
			unmatchedCount++
		}
	}

	if unmatchedCount != 1 {
		t.Errorf("Ожидалось, что один мужчина останется без пары, получено: %d", unmatchedCount)
	}
}
