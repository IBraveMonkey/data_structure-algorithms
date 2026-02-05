package main

// Reads from nil-map -> ok
func readFromNilMap() {
	var data map[int]int
	_ = data[100]
}

// Deletes from nil-map -> nothing (safe)
func deleteFromNilMap() {
	var data map[int]int
	delete(data, 100)
}

// Writes to nil-map -> panic
func writeToNilMap() {
	var data map[int]int
	data[100] = 100
}

// Iterates over nil-map -> nothing (safe)
func rangeByNilMap() {
	var data map[int]int
	for range data {
	}
}

// Overwrites existing key -> OK (value will be updated)
func rewriteExistingKey() {
	data := make(map[int]int)
	data[100] = 500
	data[100] = 1000
}

// Deletes non-existing key -> nothing (safe)
func deleteNonExistingKey() {
	data := make(map[int]int)
	delete(data, 100)
}

func basic_map() {

}
