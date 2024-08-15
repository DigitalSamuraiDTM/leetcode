fun main(args: Array<String>) {
    val test1 = intArrayOf(1,1,2)
    val removed1 = removeDuplicates(test1)
    println("REMOVED: ${removed1} :: " + test1.joinToString(", "))

    val test2 = intArrayOf(0,0,1,1,1,2,2,3,3,4)
    val removed2 = removeDuplicates(test2)
    println("REMOVED: ${removed2} :: " + test2.joinToString(", "))

    val test3 = intArrayOf(0,2,5,7,9,11,11,11,11,11,12)
    val removed3 = removeDuplicates(test3)
    println("REMOVED: ${removed3} :: " + test3.joinToString(", "))

    val test4 = intArrayOf(1)
    val removed4 = removeDuplicates(test4)
    println("REMOVED: ${removed4} :: " + test4.joinToString(", "))
}

fun removeDuplicates(nums: IntArray): Int {
    if (nums.size == 1) return 1
    var replacerIndex = 0
    var forwarderIndex = 1
    var index = 1
    while (replacerIndex <= nums.size - 1) {
        if (nums[replacerIndex] != nums[forwarderIndex]) {
            index++
            replacerIndex++
            nums[replacerIndex] = nums[forwarderIndex]
        }
        if (forwarderIndex == nums.size - 1) {
            return index
        }
        forwarderIndex++
    }
    return index
}