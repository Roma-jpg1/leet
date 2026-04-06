func diagonalSum(mat [][]int) int {
    s:=0
    if len(mat)==1{
        return mat[0][0]
    }

    for i:=0;i<len(mat); i++ {
        s+=mat[i][i]
        s+=mat[i][len(mat[i])-1-i] 

        
    }
    if len(mat)%2!=0{
        s-=mat[len(mat)/2][len(mat)/2]

    }
    return s
    
}