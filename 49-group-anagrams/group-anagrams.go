func groupAnagrams(strs []string) [][]string {
   m:= make(map[[26]int][]string)

   for _,s := range strs{
    count:=[26]int{}

    for i:=0; i<len(s); i++{
        count[s[i]-'a']++
    }
    m[count]=append(m[count], s)
   }

   res:= make([][]string, 0, len(m))
   for _, gr:=range m{
    res=append(res, gr)
   }

   return res
    
}