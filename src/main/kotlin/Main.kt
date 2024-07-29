fun main(args: Array<String>) {
    println(easyWay(arrayOf("flower","flow","flight")))
    println(easyWay(arrayOf("dog","racecar","car")))
    println(easyWay(arrayOf("ab", "a")))
}

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.

fun longestCommonPrefix(strs: Array<String>): String {
    if (strs.size == 1) return strs[0]
    // find min size element in first step, because prefix cannot be bigger than element size
    var baseElementIndex = 0
    var baseElement = strs[baseElementIndex] // no exception by task rules
    for (i: Int in 1 until strs.size) {
        if (strs[i].length < baseElement.length) {
            baseElement = strs[i]
            baseElementIndex = i
        }
    }
    strs.toMutableList().apply { removeAt(baseElementIndex) }.forEach { stringi ->
        for (i: Int in baseElement.indices) {
            if (baseElement[i] != stringi[i]) {
                baseElement = baseElement.slice(IntRange(0,i-1))
                break
            }
        }
    }

    return baseElement
}

fun easyWay(strs: Array<String>): String {
    if (strs.size == 1) return strs[0]
    val firstStringa = strs[0]
    var out = ""
    for (c: Int in firstStringa.indices) {
        for (i: Int in 1 until strs.size) {
                if (strs[i].length <= c) return out
                if (firstStringa[c] != strs[i][c]) return out
        }
        out += firstStringa[c]
    }
    return out
}
