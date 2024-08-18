import kotlin.math.max
import kotlin.math.min

fun main(args: Array<String>) {
    println("RESULT11: " + longestPalindrome("abaab")) // baab
    println("RESULT2: " + longestPalindrome("bab")) // bab
    println("RESULT0: " + longestPalindrome("abqwertyud")) // a
    println("RESULT5: " + longestPalindrome("abaaaaaaad")) // 7x a
    println("RESULT1: " + longestPalindrome("babad")) // aba
    println("RESULT3: " + longestPalindrome("zz")) // zz
    println("RESULT4: " + longestPalindrome("v")) // v
    println("RESULT6: " + longestPalindrome("abcbbcbba")) // bcbbcb
    println("RESULT7: " + longestPalindrome("88888888")) // 88888888
    println("RESULT8: " + longestPalindrome("7777777")) // 7777777
    println("RESULT9: " + longestPalindrome("abba")) // abba
    println("RESULT10: " + longestPalindrome("baaba")) // baab
}

fun longestPalindrome(s: String): String {
    if (s.length == 1) return s
    if (s.length == 2) return if (s[0] == s[1]) s else s[0].toString()

    val isLengthEven = s.length % 2 == 0
    // odd searching
    var halfSize = s.length / 2
    val oddString = s + (if (isLengthEven) "_" else "")

    var maxPalindrome = s[0].toString()
    for (i: Int in 0..halfSize - 1) {
//        println("I IS: ${i}")
        var left = oddString[halfSize - i].toString()
        var right = oddString[halfSize + i].toString()
        var isRightPalindromeEnded = false
        var isLeftPalindromeEnded = false
        for (j: Int in 1..halfSize - i) {
            // check left
            if (oddString[halfSize - i - j] == oddString[halfSize - i + j] && !isLeftPalindromeEnded) {
                left = (oddString[halfSize - i - j] + left + oddString[halfSize - i + j])
            } else {
                isLeftPalindromeEnded = true
            }
            // check right
            if (oddString[halfSize + i - j] == oddString[halfSize + i + j] && !isRightPalindromeEnded) {
                right = (oddString[halfSize + i - j] + right + oddString[halfSize + i + j])
            } else {
                isRightPalindromeEnded = true
            }
//            println(left + "|" + right)
            if (isLeftPalindromeEnded && isRightPalindromeEnded) break
        }
        if (left.length > maxPalindrome.length) maxPalindrome = left
        if (right.length > maxPalindrome.length) maxPalindrome = right

    }

    // even searching
    halfSize += (if (!isLengthEven) 1 else 0)
    val evenString = s + (if (!isLengthEven) "_" else "")
    for (i: Int in 0..halfSize - 1) {
        var left = ""
        var right = ""
        var isRightPalindromeEnded = false
        var isLeftPalindromeEnded = false
//        println("I IS: ${i}")
        for (j: Int in 0 until halfSize - i) {
            // check left
//            println("J: ${j}")
            if (evenString[halfSize - 1 - i - j] == evenString[halfSize - i + j] && !isLeftPalindromeEnded) {
                left = (evenString[halfSize - 1 - i - j] + left + evenString[halfSize - i + j])
            } else {
                isLeftPalindromeEnded = true
            }
            if (evenString[halfSize - 1 + i - j] == evenString[halfSize + i + j] && !isRightPalindromeEnded) {
                right = (evenString[halfSize - 1 + i - j] + right + evenString[halfSize + i + j])
            } else {
                isRightPalindromeEnded = true
            }
            if (isLeftPalindromeEnded && isRightPalindromeEnded) break
//            println(left + "|" + right)
        }
        if (left.length > maxPalindrome.length) maxPalindrome = left
        if (right.length > maxPalindrome.length) maxPalindrome = right
    }
    return maxPalindrome
}

// TODO: Решение данной задачи оказывается решается через "Алгоритм Манакера", о котором я до момента своего решения ничего не знал
// Это решение кринж, довольно медленное, смотри истину тут https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_%D0%9C%D0%B0%D0%BD%D0%B0%D0%BA%D0%B5%D1%80%D0%B0