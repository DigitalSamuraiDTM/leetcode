fun main(args: Array<String>) {
    println(strStr("mississippi","pi")) // 9
    println(strStr("qweeerty","eee")) // 2
    println(strStr("mississippi","issip")) // 4
    println(strStr("sadbutsad","sad")) // 0
    println(strStr("leetcode","leeto"))// -1
    println(strStr("qwerty","r")) // 3
    println(strStr("qwerty","rerty")) // -1
    println(strStr("qwerty","y")) // 5
    println(strStr("qqqqqqqqqqqqqqqqq","qqqqq")) // 0
}

fun strStr(haystack: String, needle: String): Int {
    var index = 0
    var out = 0
    var i = 0
    while(i != haystack.length) {
        val char = haystack[i]
        if (char == needle[index]) {
            if (index == 0) {
                out = i
            }
            index++
        } else {
            if (index != 0) {
                i = out + 1
                index = 0
                out = 0
                continue
            }
        }
        if (index == needle.length) {
            return out
        }
        i++
    }
    return -1
}