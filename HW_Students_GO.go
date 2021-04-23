package main

/* Подключаем библиотеки */
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/* Функция для перемешивания массива */
func swap(a, b *int) {
	var tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	var tic int

	fmt.Println("Введите колл-во билетов")
	fmt.Scanf("%d", &tic)

	var key int

	fmt.Println("Введите ключ")
	fmt.Scanf("%d", &key)

	/* Открываем файл со списком студентов */
	file, err := os.Open("Students")
	if err != nil {
		log.Fatal(err)
	}

	var n int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n++
	}

	/* Создаем динамический массив */
	arr := make([]int, n)
	defer file.Close()

	/* Заполняем массив значениями от 1 до ticket */
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	/* Перемешиваем массив */
	for i := 0; i < tic-1; i++ {
		swap(&arr[(i^2+key+i)%n], &arr[((i+1)^2+key)%(n-1)])
	}

	/* Получаем данные из файла и выводим на консоль с указанием номера билета*/
	var i int = 0
	file.Close()
	file_sec, err := os.Open("Students")
	scan := bufio.NewScanner(file_sec)
	for scan.Scan() {
		fmt.Println(scan.Text(), arr[i]%tic+1)
		i++
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}
