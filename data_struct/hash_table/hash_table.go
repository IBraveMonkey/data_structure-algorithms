package hash_table

/* Найти разницу между двух строк
На вход подается две строки a и b.
Строка b образовано из строки a путем перемешивания и добавления одной буквы. Неоходимо вернуть эту букву
*/
// a = "abc" b = "beca" => e
// timeComplexity - O(n), по spaceComplexity - O(n)
func ExtraLetter(a, b string) string {
	hTable := make(map[string]int, len(b))

	for _, v := range b {
		hTable[string(v)]++
	}

	for _, v := range a {
		if _, ok := hTable[string(v)]; ok {
			hTable[string(v)]--
		}
	}

	for key, val := range hTable {
		if val > 0 {
			return key
		}
	}

	return ""
}

/* Сумма двух элементов массива
Дан неотсортированный массив целых чисел и некоторое число target. Необходимо написать функцию, которая найдет два таких элемента в массиве, сумма которых будет равна target
*/

func TwoSum(data []int, target int) []int {
	cache := make(map[int]int)

	for i := 0; i < len(data); i++ {
		cache[data[i]] = i
	}

	for i := 0; i < len(data); i++ {
		diff := target - data[i]
		val, ok := cache[diff]
		if ok {
			return []int{i, val}
		}
	}

	return []int{}
}

/* Накормить животных
Массив потребностей в еде животных [3,4,7], массив привезенной еды: [8,1,2] может накормить одно животное

На вход подается 2 массива целых чисел. Первый массив - потребности животных, второй - кол-во привезенной еды. Необходимо вернуть целое число - количество накормленных зверей
*/

func FeedAnimals(animals, food []int) int {
	if len(animals) == 0 || len(food) == 0 {
		return 0
	}

	bubbleSort(animals)
	bubbleSort(food)

	var count int

	for _, v := range food {
		if len(animals) > count {
			if v >= animals[count] {
				count++
			}

			if count == len(animals) {
				break
			}
		}
	}

	return count
}

// o(n^2)
func bubbleSort(arr []int) {
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}

/* Массив Анаграмм
Необходимо найти и сгруппировать
слова-анаграммы вместе

Анаграмм - это слово или фраза полученная путем перестановки букв другого слова или фразы, обычно с использованием всех исходных букв ровно один раз

Дан массив строк strs. Сгруппируйте аннаграмы вместе. Ответ можно вернуть в любом порядке

Входные данные - ["eat", "tea", "tan", "ate", "nat", "bat"]
Выходные - [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
*/

func GroupAnagrams(strs []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, v := range strs {
		sV := selectionSort(v)

		_, ok := anagrams[sV]
		if ok {
			anagrams[sV] = append(anagrams[sV], v)
		} else {
			anagrams[sV] = []string{v}
		}
	}

	return anagrams
}

func selectionSort(arr string) string {
	runes := []rune(arr)

	n := len(runes)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if runes[j] < runes[minIndex] {
				minIndex = j
			}
		}

		runes[i], runes[minIndex] = runes[minIndex], runes[i]
	}

	return string(runes)
}
