package task1

import (
	"bufio"
	"fmt"
	"os"
)

type light struct {
	dir string
	n   int
	m   int
}

// Task## - решение задачи # (вариант #). Запускать через main.go в папке cmd
func Task11() {
	// Раскомментить при запуске на своей машине, закомментить при отправке на платформу
	// В папку tests скопировать тесты с платформы
	// Использовать для тестирования на своей машине, используя данные из указанного файла
	file, err := os.Open("../internal/task1/tests/1")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %s\n", err.Error())
	}
	in := bufio.NewReader(file)

	// Раскомментить при отправке на платформу, закомментить при запуске
	// in := bufio.NewReader(os.Stdin)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Основное решение задачи # (вариант #)
	Run11(in, out)
}

// Run## - основное решение задачи # (вариант #)
// Менять какие-либо аргументы или возвращаемые значения не требуется
func Run11(in *bufio.Reader, out *bufio.Writer) {
	// t - количество входных данных (подзадач)
	var t int
	fmt.Fscanln(in, &t)

	// Решение t подзадач
	for i := 0; i < t; i++ {
		// Примеры чтения параметров подзадачи i

		// Чтение из одной строки нескольких значений, разделённых пробелом
		var n, m int
		fmt.Fscanln(in, &n, &m)

		// Чтение целой строки
		// str1, _ := in.ReadString('\n')
		// Удаление символа \n (на некоторых машинах надо удалить два символа \r\n)
		// str := strings.Trim(str1, "\n")
		// str := strings.Trim(str1, "\r\n")
		// Для избежания ошибки выведем str (закомментить при использовании шаблона)
		// fmt.Println(str)

		// Запуск и вывод в out решения подзадачи t
		// В зависимости от условия задачи алгоритм вывода может потребовать доработки
		result := tTaskSolving11(n, m)
		fmt.Fprintln(out, len(result))
		for _, j := range result {
			fmt.Fprintln(out, j.m, j.n, j.dir)
		}
	}
}

// tTaskSolving## - функция для решения подзадачи t задачи # (вариант #)
// В зависимости от условия задачи, необходимо указать требуемые аргументы и возвращаемое значение функции
func tTaskSolving11(n, m int) []light {
	result := make([]light, 0)
	switch {
	case n == 1 && m != 1:
		result = append(result, light{"R", 1, 1})
	case m == 1:
		result = append(result, light{"D", 1, 1})
	default:
		result = append(result, light{"R", 1, 1})
		result = append(result, light{"D", 2, 1})
	}
	return result
}
