package slicey

// Split left, (from left to right) example :
// [] 	[] 	[] 	[]
// ^ s will shared to first (from left) array until it's reach the limit
// if limit has been reached, then continue to fill second array
func Splitl(s []interface{}, n int) [][]interface{} {

	// child define limit of storage child
	// tolerant = 1 | beware !!!
	// child = (len(s) / n) + tolerant
	childlen := (len(s) / n) + 1

	// initialize fixed number of slice
	storage := make([][]interface{}, n)

	// iterate s value and share value to become
	for _, item := range s {
		for index, _ := range storage {
			if len(storage[index]) < childlen {
				storage[index] = append(storage[index], item)
				// always break to fill up child storage from left
				break
			}
		}
	}

	// return multidimentional slice of splitted s data
	return storage
}

// Split full will fill child slice from left to right
// no matter how much len already in x child
func Splitf(s []interface{}, n int) [][]interface{} {

	// initialize fixed number of slice
	storage := make([][]interface{}, n)

	// latest index of s
	index := 0

	for {
		if index >= len(s) {
			// break if index already same or more (rare) than
			// len of given s
			break
		}
		for child, _ := range storage {
			storage[child] = append(storage[child], s[index])
			index += 1
		}
	}

	// return multidimentional slice of splitted s data
	return storage
}
