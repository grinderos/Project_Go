package main

import (
	"fmt"
)

const (
	StatusOK = 200
)

func main() {
	// var discountPercent float64 = 10.0
	// var productPrice float64 = 59.99
	// discountAmount := discountPercent * productPrice / 100
	// finalPrice := math.Floor((productPrice-discountAmount)*100) / 100
	// fmt.Println(finalPrice)

	// // str2 := "Hello, World!"
	// var str3 = "Hello, World!"
	// // var str4 string = "Hello, World!"
	// // str5 = "Hello, World!"
	// fmt.Println(str3)
	// str := "𝓗𝓮𝓵𝓵𝓸, мой друг."
	// fmt.Println(str, len(str), utf8.RuneCountInString(str))

	// var message string
	// message = " Go - это не просто язык, это СТИЛЬ ЖИЗНИ! "
	// message = strings.TrimSpace(message)
	// fmt.Println(message)
	// message = strings.ToLower(message)
	// fmt.Println(message)
	// fmt.Println(strings.HasPrefix(message, "go"))

	// i := 555
	// s := strconv.Itoa(i)
	// fmt.Println(s)

	// // var price float64 = 62.231413
	// var price float64 = 23.43753424
	// fmt.Println(strconv.FormatFloat(price, 'f', 3, 64))

	// priceStr := "19.22"
	// quantityStr := "19"
	// price, err1 := strconv.ParseFloat(priceStr, 64)
	// quantity, err2 := strconv.Atoi(quantityStr)
	// if err1 != nil || err2 != nil {
	// 	fmt.Println("Ошибка конвертации")
	// 	return
	// }
	// fmt.Println(strconv.FormatFloat(price*float64(quantity), 'f', 2, 64))

	// var city string = "Москва"
	// var temp int = 25
	// var weather string = "солнечно"
	// fmt.Printf("В городе %s температура %d°C, и %s.\n", city, temp, weather)

	// var num int = 875394
	// fmt.Printf("Запись числа %d в разных системах счисления:\nДесятичная: %d\nДвоичная: %b\nВосьмеричная: %o\nШестнадцатеричная: %X\n", num, num, num, num, num)
	//// ----------------------------------------------------
	////fmt.Scanln() - для чтения строки:

	// fmt.Println("Введите текст:")
	// // подготавливаем переменные для ввода
	// var w1, w2, w3 string
	// // ожидаем ввода
	// n1, err := fmt.Scanln(&w1, &w2, &w3)
	// // проверяем, не было ли ошибки
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // выводим данные
	// fmt.Printf("Количество считанных элементов: %d\n", n1)
	// fmt.Printf("Считанная строка: %s %s %s\n", w1, w2, w3)

	//// ----------------------------------------------------
	//// fmt.Scan() - для чтения нескольких строк:

	// fmt.Println("Введите текст:")
	// // подготавливаем переменные для ввода
	// var w4, w5, w6 string
	// // ожидаем ввода
	// n2, err1 := fmt.Scan(&w4, &w5, &w6)
	// // проверяем, не было ли ошибки
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }
	// // выводим данные
	// fmt.Printf("Количество считанных элементов: %d\n", n2)
	// fmt.Printf("Считанный текст: %s %s %s\n", w4, w5, w6)

	//// ----------------------------------------------------
	//// bufio.Reader или bufio.Scanner - для чтения одной строки текста:
	//// bufio.Reader:

	// fmt.Println("Введите текст:")
	// // создаем читателя, который будет читать из консоли
	// reader := bufio.NewReader(os.Stdin)
	// // читаем до определенного символа включительно (\n - это символ переноса строки)
	// line, err2 := reader.ReadString('\n')
	// // проверяем, не было ли ошибки
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// // выводим данные
	// fmt.Printf("Считанная строка: %s\n", line)
	// // --------
	// // bufio.Scanner:
	// fmt.Println("Введите текст:")
	// // создаем сканер, который будет читать из консоли
	// scanner1 := bufio.NewScanner(os.Stdin)
	// // читаем ввод данных
	// scanner1.Scan()
	// // проверяем, не было ли ошибки
	// err3 := scanner1.Err()
	// if err3 != nil {
	// 	log.Fatal(err3)
	// }
	// // выводим данные
	// fmt.Printf("Считанная строка: %s\n", scanner1.Text())

	//// ----------------------------------------------------

	//// Для чтений нескольких строк также можно использовать bufio.Reader или bufio.Scanner.
	//// bufio.Reader:

	// fmt.Println("Введите текст:")
	// // создаем читателя, который будет читать из консоли
	// reader1 := bufio.NewReader(os.Stdin)

	// // в этот слайс записываем все введенные строки
	// var lines []string
	// for {
	// 	// читаем строку, вплоть до переноса строки
	// 	line, err := reader1.ReadString('\n')
	// 	// проверяем, что не было ошибки
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// если не было введено ничего (пробелы не считаем), то выходим из цикла
	// 	if len(strings.TrimSpace(line)) == 0 {
	// 		break
	// 	}
	// 	// добавляем введенную строку в слайс
	// 	lines = append(lines, line)
	// }

	// // выводим каждую введенную строку
	// fmt.Println("Вывод:")
	// for _, l := range lines {
	// 	fmt.Printf(l)
	// }

	//// bufio.Scanner:

	// fmt.Println("Введите текст:")
	// // создаем сканнер, который будет читать из консоли
	// scanner := bufio.NewScanner(os.Stdin)

	// // слайс, в котором будут храниться все введенные строки
	// var lines []string
	// for {
	// 	// ожидаем ввода
	// 	scanner.Scan()
	// 	// получаем введенный текст
	// 	line := scanner.Text()
	// 	// если строка пустая - выходим из цикла
	// 	if len(line) == 0 {
	// 		break
	// 	}
	// 	// добавляем строку в слайс линий
	// 	lines = append(lines, line)
	// }

	// // проверяем, не было ли ошибки во время ввода
	// err := scanner.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // выводим каждую введенную строку
	// fmt.Println("Вывод:")
	// for _, l := range lines {
	// 	fmt.Println(l)
	// }

	//// ----------------------------------------------------

	//// Чтение одного символа
	//// Для чтений просто одного символа можно использовать bufio.Reader или fmt.Scanf().
	//// bufio.Reader:

	// fmt.Println("Введите текст:")
	// // создаем читателя, который будет читать из консоли
	// reader := bufio.NewReader(os.Stdin)
	// // читаем один символ из того, что было введено
	// char, _, err := reader.ReadRune()
	// // проверяем, не было ли ошибки
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // выводим результат
	// fmt.Printf("Считанный символ: %c\n", char)

	//// Функция fmt.Scanf() интересней, она может читать один символ, либо вычитывать определенные символы из строки,
	////  причем ввод должен быть непременно таким, какой ожидается, иначе получим ошибку.

	// fmt.Println("Введите символ:")
	// // создаем переменную для введенного символа
	// var char rune
	// // сканируем ввод данных
	// // %c - означает, что мы ожидаем один символ
	// // &char - место, куда записываем ввденный символ
	// _, err := fmt.Scanf("%c", &char)
	// // проверяем, не было ли ошибки
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // выводим данные
	// fmt.Printf("Считанный символ: %c\n", char)

	//// ----------------------------------------------------
	//// Чтение числа

	// fmt.Println("Введите число:")
	// // создаем переменную, в которую положим введенное число
	// var number int64

	// // читаем ввод пользователя
	// // %d - означает, что ждем целое число
	// // &number - указатель, в какую переменную записывать введенное значение
	// _, err := fmt.Scanf("%d", &number)
	// // либо можно использовать fmt.Scan():
	// // _, err := fmt.Scan(&number)

	// // проверяем, не было ли ошибки
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // выводим данные
	// fmt.Printf("Считанное число: %d\n", number)

	////----------------------------------------------------

	// fmt.Println("Введите Ваши 1) Имя и 2) Фамилию:")
	// // подготавливаем переменные для ввода
	// var w1, w2 string
	// var d1 int
	// // ожидаем ввода
	// fmt.Scanln(&w1, &w2)
	// fmt.Println("Введите Ваш возраст числом:")
	// fmt.Scanln(&d1)
	// fmt.Printf("Приятно познакомиться, %s. Я 5 лет назад познакомился с человеком, у которого тоже фамилия %s, "+
	// 	"вам тогда было %d. Как молоды мы были!", w1, w2, d1-5)
	////----------------------------------------------------

	// random := math.Floor((rand.Float64()*100)*10) / 10

	// fmt.Printf("Исходное число: %.1f\n", random)
	// fmt.Printf("Исходное число, увеличенное на 10%%: %.5f\n", random+(random*0.1))
	// fmt.Printf("Исходное число является четным: %t\n", math.Mod(random, 2) == 0)
	// fmt.Printf("Предпоследняя цифра целой части исходного числа: %d\n", int(math.Mod(math.Floor(random/10), 10)))

	// num := 10
	// num++
	// fmt.Println(num)
	// string1 := "Hello"
	// string2 := ", World!"
	// string1 += string2
	// fmt.Println(string1)

	// fmt.Println(num < 10+5)
	////----------------------------------------------------

	// age := 42
	// status := "paused"
	// role := "officer"
	// if (role == "admin" || role == "moderator") || (status == "active" && age >= 18 && role == "user") {
	// 	fmt.Println(true)
	// } else {
	// 	fmt.Println(false)
	// }

	// num := 5
	// fmt.Println(num << 2)
	// fmt.Println(num >> 1)
	// fmt.Println(num & 3)
	// fmt.Println(num | 2)
	// fmt.Println(num ^ 2)
	// fmt.Println(^num)
	////----------------------------------------------------
	// temp := 55
	// if temp < 0 {
	// 	fmt.Println("Город замерзает! Верните лето.")
	// } else if temp >= 0 && temp <= 35 {
	// 	fmt.Println("Температура в норме. Продолжаем писать код.")
	// } else {
	// 	fmt.Println("Город в огне! Яичницу можно жарить на асфальте.")
	// }
	////----------------------------------------------------
	////задание 4.5

	// fmt.Println("Введите время в часах (например: 14):")
	// var hour int
	// // ожидаем ввода
	// _, err := fmt.Scan(&hour)
	// // проверяем, не было ли ошибки
	// if err != nil {
	// 	log.Fatal("Неверный формат времени.")
	// }
	// if hour >= 0 && hour < 6 || hour == 23 {
	// 	fmt.Printf("Сейчас %dч. - ночь", hour)
	// } else if hour >= 6 && hour < 12 {
	// 	fmt.Printf("Сейчас %dч. - утро", hour)
	// } else if hour >= 12 && hour < 18 {
	// 	fmt.Printf("Сейчас %dч. - день", hour)
	// } else if hour >= 18 && hour < 23 {
	// 	fmt.Printf("Сейчас %dч. - вечер", hour)
	// } else {
	// 	fmt.Println("В сутках только 24 часа! Введите значение от 0 до 23.")
	// }
	////----------------------------------------------------
	////задание 4.6

	// fmt.Printf("Введите ваш вес (кг): ")
	// var weight float64
	// // ожидаем ввода
	// _, err1 := fmt.Scan(&weight)
	// fmt.Printf("Введите ваш рост (см): ")
	// var height float64
	// // ожидаем ввода
	// _, err2 := fmt.Scan(&height)
	// if err1 != nil || err2 != nil {
	// 	log.Fatal("Неверный формат введенных данных")
	// }
	// index := weight / math.Pow(height/100, 2)
	// fmt.Printf("Ваш ИМТ: %.2f\n", index)

	// var category string
	// switch {
	// case index < 18.5:
	// 	category = "Недостаточный вес"
	// case index >= 18.5 && index < 25:
	// 	category = "Нормальный вес"
	// case index >= 25 && index < 30:
	// 	category = "Избыточный вес"
	// default:
	// 	category = "Ожирение"
	// }
	// fmt.Printf("Категория: %s", category)

	////----------------------------------------------------
	////задание 4.7

	// products := map[string]float64{"Клавиатура JZ9": 19200, "Наушники N40": 9600, "Смартфон S10": 55000}
	// fmt.Println("Введите название товара:")
	// // создаем читателя, который будет читать из консоли
	// reader := bufio.NewReader(os.Stdin)
	// // читаем до определенного символа включительно (\n - это символ переноса строки)
	// line, err := reader.ReadString('\n')
	// // проверяем, не было ли ошибки
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Вы ввели:", line)
	// line = strings.ToLower(strings.TrimSpace(line))
	// for key, value := range products {
	// 	if strings.EqualFold(key, line) {
	// 		fmt.Println(key, ": ", value)
	// 		return
	// 	}
	// 	getLowerKey := strings.ToLower(key)
	// 	if len(line) > 2 && strings.Contains(getLowerKey, line) {
	// 		fmt.Println(key, ": ", value)
	// 		return
	// 	}
	// }
	// fmt.Println("Товар не найден")

	////----------------------------------------------------
	// line, err := UserProfileToString(" Asd", 32)
	// fmt.Println(line, err)
	// result, err := calculate(10, 3, "divide")
	// fmt.Println(result, err)

	// //----------------------------------------------------
	// fmt.Println(userProfile(""))
	// //----------------------------------------------------
	// add := adder(10)
	// fmt.Println(add(5))
	// fmt.Println(add(10))
	////----------------------------------------------------
	// fmt.Println(sumOfDigits(-456))
	////----------------------------------------------------
	// leter, err := letterBall(105)
	// leter, err := letterBall(95)
	// leter, err := letterBall(75)
	// leter, err := letterBall(65)
	// leter, err := letterBall(40)
	// leter, err := letterBall(-15)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Буквенный балл: %s\n", leter)
	// }
	////----------------------------------------------------
	PrintReplaced("Кукушка")
}

// func UserProfileToString(name string, age int) (string, error) {
// 	if name == "" {
// 		return "", errors.New("empty name")
// 	}
// 	trimmedName := strings.TrimSpace(name)
// 	if trimmedName == "" {
// 		return "", errors.New("name cannot contain only spaces")
// 	}
// 	if age < 0 {
// 		return "", errors.New("negative age")
// 	}
// 	return fmt.Sprintf("Имя человека: %s, возраст: %d.", trimmedName, age), nil
// }

// func calculate(first float64, second float64, operator string) (float64, error) {
// 	switch operator {
// 	case "add":
// 		return first + second, nil
// 	case "subtract":
// 		return first - second, nil
// 	case "multiply":
// 		return first * second, nil
// 	case "divide":
// 		if second == 0 {
// 			return 0, errors.New("division by zero")
// 		}
// 		return first / second, nil
// 	default:
// 		return 0, errors.New("unknown operation")
// 	}
// }

// func userProfile(id string) (string, error) {
// 	result, err := fetchUserInfo(id)
// 	if err != nil {
// 		return "", fmt.Errorf("fetch error: %w", err)
// 	}
// 	rub := float64(result) / 100
// 	return fmt.Sprintf("Пользователь с id %s имеет на счету %0.2f руб.", id, rub), nil
// }

// func fetchUserInfo(id string) (int, error) {
// 	if id == "" {
// 		return 0, errors.New("empty id")
// 	}
// 	return 155, nil
// }

// func adder(num int) func(i int) int {
// 	return func(i int) int {
// 		num += i
// 		return num
// 	}
// }

// type Day int

// const (
// 	_ Day = iota
// 	Monday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// 	Sunday
// )

// func isWeekend(day Day) bool {
// 	return day == Saturday || day == Sunday
// }

// func sumOfDigits(num int) int {
// 	if num < 0 {
// 		num = -num
// 	}
// 	if num < 10 {
// 		return num
// 	}
// 	return num%10 + sumOfDigits(num/10)
// }

func letterBall(num int) (leter string, err error) {

	switch {
	case num >= 90 && num <= 100:
		return "A", nil
	case num >= 80 && num < 90:
		return "B", nil
	case num >= 70 && num < 80:
		return "C", nil
	case num >= 60 && num < 70:
		return "D", nil
	case num >= 0 && num < 60:
		return "F", nil
	default:
		return "", fmt.Errorf("Числовой балл %d вне диапазона 0-100.", num)
	}
}

func PrintReplaced(str string) {
	runes := []rune(str)
	for i, char := range runes {
		if char == 'у' {
			runes[i] = 'а'
		}
	}
	fmt.Println(string(runes))
}
