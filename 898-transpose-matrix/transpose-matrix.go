func transpose(matrix [][]int) [][]int {
    mat := make([][]int, len(matrix[0]))

    for i :=0; i<len(mat);i++ {
        mat[i] = make([]int, len(matrix))
    }

    for i:=0; i<len(matrix); i++{
        for j:=0; j<len(matrix[i]); j++{
            mat[j][i]=matrix[i][j]
        }
    }
    return mat
    
}