package letcode

func IsValidSudoku(board [][]byte) bool {
	hashMap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row := i
			column := j
			current_val := string(board[i][j])
			if current_val == "." {
				continue
			}
			_, ok1 := hashMap[current_val+"found in row"+string(rune(row))]
			_, ok2 := hashMap[current_val+"found in column"+string(rune(column))]
			_, ok3 := hashMap[current_val+"found in grid"+string(rune(i/3))+"-"+string(rune(j/3))]
			if ok1 || ok2 || ok3 {

				return false
			} else {
				hashMap[current_val+"found in row"+string(rune(row))] = true
				hashMap[current_val+"found in column"+string(rune(column))] = true
				hashMap[current_val+"found in grid"+string(rune(i/3))+"-"+string(rune(j/3))] = true
			}
		}
	}
	return true
}
