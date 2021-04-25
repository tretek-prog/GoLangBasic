package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)


// Структура для обработки json.
type reguestJson struct {
	Array string `json:"arr"`
	Action string `json:"action"`
}

// Обработка запроса и ответа
func answer(w http.ResponseWriter, r *http.Request) {

	// Проверка метода запроса
	if r.Method == http.MethodPost {

		// Запись json в структуру.
		var requestData reguestJson
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestData)
		closeError := r.Body.Close()
		println(closeError)

		// Проверка записи на ошибки.
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		// Передача данных в функцию калькулятора.
		q, calcError := calculator(requestData.Array, requestData.Action)

		// Обработка ошибок калькулятора.
		if calcError != "" {
			resp, _ := json.Marshal(calcError)
			_, _ = w.Write(resp)
			w.WriteHeader(http.StatusOK)

		} else {
			resp, _ := json.Marshal(q)
			_, _ = w.Write(resp)
			w.WriteHeader(http.StatusOK)
		}

	} else {
		_, _ = w.Write([]byte(r.Method + " Вы ввели что-то не то. Попробуйте POST запрос."))
	}
}

// Функция умножения.
func mult (numbers []string) (ans int64) {
	println(numbers)
	ans = 1
	for _, number := range numbers {
		i, _ := strconv.ParseInt(number, 16, 16)
		ans *= i
	}
	return ans
}

// Функция сложения.
func sum (numbers []string) (ans int64) {
	println(numbers)
	ans = 0
	for _, number := range numbers {
		i, _ := strconv.ParseInt(number, 16, 16)
		ans += i
	}
	return ans
}

// Функция калькулятора.
func calculator (stringArray string, action string) (number int64, calcError string) {

	// Перевод строки списка в список строк.
	stringArray = stringArray[1 : len(stringArray) - 1]
	numbers := strings.Split(stringArray, ", ")

	// Определение действий со списком и вывод результата.
	if action == "mult" {
		return mult(numbers), ""
	} else if action == "sum" {
		return sum(numbers), ""
	} else {
		return -1, action + " Вы ввели что-то не то. Используйте sum или mult."
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", answer)
	err := http.ListenAndServe(":3000", mux)
	println(err)
}
