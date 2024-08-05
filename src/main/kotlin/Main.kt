import java.nio.charset.Charset

fun main(args: Array<String>) {
    println(lengthOfLongestSubstring("abcaeba"))
    println(lengthOfLongestSubstring("abcabcbb"))
    println(lengthOfLongestSubstring("bbbbb"))
    println(lengthOfLongestSubstring("pwwkew"))
    println(lengthOfLongestSubstring(extraString))

}
// Example 1:
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.

//Example 2:
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.

//Example 3:
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

fun lengthOfLongestSubstring(s: String): Int {
    if (s.isEmpty()) return 0
     var stringSize = s.toSet().size
    while (stringSize != 1) {
        val iterations = s.length - stringSize
        //launch search strings
        for (i: Int in 0 .. iterations) {
            val subString = s.slice(IntRange(i,i+stringSize - 1))
            val outSet = mutableSetOf<Char>()
            subString.forEach { char ->
                if (char in outSet) {
                    return@forEach
                } else {
                    outSet.add(char)
                }
            }
            if (outSet.size == subString.length) {
                return subString.length
            }
        }
        stringSize -= 1
    }
    return 1
}