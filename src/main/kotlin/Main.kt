fun main(args: Array<String>) {
    println(lengthOfLastWord2("Hello World"))
    println(lengthOfLastWord2("   fly me   to   the moon  "))
    println(lengthOfLastWord2("luffy is still joyboy"))
}

// Example 1:
//Input: s = "Hello World"
//Output: 5
//Explanation: The last word is "World" with length 5.

//Example 2:
//Input: s = "   fly me   to   the moon  "
//Output: 4
//Explanation: The last word is "moon" with length 4.

//Example 3:
//Input: s = "luffy is still joyboy"
//Output: 6
//Explanation: The last word is "joyboy" with length 6.

fun lengthOfLastWord(s: String): Int {
    return s.split(" ").last { it.isNotBlank() }.length
}

fun lengthOfLastWord2(s: String): Int {
    var indexator = 0
    var isLastWord = false
    s.forEach { c ->
        if (c == ' ') {
            isLastWord = true
        } else {
            if (isLastWord) {
                isLastWord = false
                indexator = 0
            }
            indexator++
        }
    }
    return indexator
}
