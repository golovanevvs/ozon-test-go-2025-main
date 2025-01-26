package task2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task## - решение задачи # (вариант #). Запускать через main.go в папке cmd
func Task21() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	// file, err := os.Open("../internal/task2/tests/8")
	// if err != nil {
	// 	fmt.Printf("Ошибка открытия файла: %s\n", err.Error())
	// }
	// in := bufio.NewReader(file)

	// Раскомментить при отправке на платформу, закомментить при запуске
	in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Основное решение задачи # (вариант #)
	Run21(in, out)
}

// Run## - основное решение задачи # (вариант #)
// Менять какие-либо аргументы или возвращаемые значения не требуется
func Run21(in *bufio.Reader, out *bufio.Writer) {
	// t - количество входных данных (подзадач)
	var t int
	fmt.Fscanln(in, &t)

	// Решение t подзадач
here:
	for i := 0; i < t; i++ {
		// Примеры чтения параметров подзадачи i

		// Чтение из одной строки нескольких значений, разделённых пробелом
		var n int
		fmt.Fscanln(in, &n)

		store := make(map[string]string)

		for j := 0; j < n; j++ {
			// Чтение целой строки
			// str1, _ := in.ReadString('\n')
			// Удаление символа \n (на некоторых машинах надо удалить два символа \r\n)
			// str := strings.Trim(str1, "\n")
			// str := strings.Trim(str1, "\r\n")
			// Для избежания ошибки выведем str (закомментить при использовании шаблона)
			// fmt.Println(str)
			var name, price string
			fmt.Fscanln(in, &name, &price)
			store[name] = price
		}

		expectedStore := make(map[string]string)
		var expectedStr string
		fmt.Fscanln(in, &expectedStr)
		expected1 := strings.Split(expectedStr, ",")
		for _, k := range expected1 {
			expected2 := strings.Split(k, ":")
			if _, ok := expectedStore[expected2[0]]; ok {
				fmt.Fprintln(out, "NO")
				continue here
			}
			if len(expected2) > 1 && len(expected2) < 3 {
				expectedStore[expected2[0]] = expected2[1]
			} else {
				fmt.Fprintln(out, "NO")
				continue here
			}
		}

		// Запуск и вывод в out решения подзадачи t
		// В зависимости от условия задачи алгоритм вывода может потребовать доработки
		fmt.Fprintln(out, tTaskSolving21(store, expectedStore))
	}
}

// tTaskSolving## - функция для решения подзадачи t задачи # (вариант #)
// В зависимости от условия задачи, необходимо указать требуемые аргументы и возвращаемое значение функции
func tTaskSolving21(store, expectedStore map[string]string) string {
	uniquePrice := make(map[string]string)

	for exName, exPrice := range expectedStore {
		if price, ok := store[exName]; !ok || exPrice != price {
			return "NO"
		}

		if _, ok := uniquePrice[exPrice]; ok {
			return "NO"
		}

		uniquePrice[exPrice] = exName
	}

	for _, price := range store {
		if _, ok := uniquePrice[price]; !ok {
			return "NO"
		}
	}

	return "YES"
}
