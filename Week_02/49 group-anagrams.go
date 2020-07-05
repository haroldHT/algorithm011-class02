func groupAnagrams(strs []string) [][]string {
primes := [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
    
    m:=make(map[int][]string)
    var i int
    for i=0;i<len(strs);i++{
        sum:=1
        for _,t:=range strs[i]{
            sum*=primes[int(t-'a')]
        }
            m[sum]=append(m[sum],strs[i])
        
    }
    
    var arr [][]string
    
    for _,v:=range m{
        arr=append(arr,v)
    }
    
    return arr
}