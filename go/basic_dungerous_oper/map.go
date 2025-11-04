package main

// Читает из nil-мапы -> нормально
func readFromNilMap() {
	var data map[int]int
	_ = data[100]
}

// Удаляет из nil-мапы -> ничего (безопасно)
func deleteFromNilMap() {
	var data map[int]int
	delete(data, 100)
}

// Пишет в nil-мапу -> panic
func writeToNilMap() {
	var data map[int]int
	data[100] = 100
}

// Итерирует по nil-мапе -> ничего (безопасно)
func rangeByNilMap() {
	var data map[int]int
	for range data {
	}
}

// Перезаписывает существующий ключ -> OK (значение обновится)
func rewriteExistingKey() {
	data := make(map[int]int)
	data[100] = 500
	data[100] = 1000
}

// Удаляет несуществующий ключ -> ничего (безопасно)
func deleteNonExistingKey() {
	data := make(map[int]int)
	delete(data, 100)
}

func main() {

}
