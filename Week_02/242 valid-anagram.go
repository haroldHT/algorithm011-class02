func isAnagram(s string, t string) bool {
    var i int
    
    var arr [26]int
    
    if len(s)!=len(t){
        return false
    }
    
    for i=0;i<len(s);i++{
        arr[s[i]-'a']++
        arr[t[i]-'a']--
    }
    
    for i=0;i<len(arr);i++{
        if arr[i]!=0{
            return false
        }
    }
    
    return true
}