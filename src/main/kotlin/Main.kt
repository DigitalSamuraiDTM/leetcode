fun main(args: Array<String>) {
    println(searchInsert(intArrayOf(1,3,5,6), 7))
    println("MY CASES")
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 1))
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 3))
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 5))
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 6))
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 9))
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 12))
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 14))
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 166))
    println(searchInsert(intArrayOf(1, 3, 5, 6, 9, 12, 14, 166, 180), 180))

    println("OUTSIDE")

    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 0))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 2))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 4))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 6))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 8))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 10))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 11))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 13))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 15))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 165))
    println(searchInsert(intArrayOf(1, 3, 5, 7, 9, 12, 14, 166, 180), 181))

}

// Example 1:
//Input: nums = [1,3,5,6], target = 5
//Output: 2

//Example 2:
//Input: nums = [1,3,5,6], target = 2
//Output: 1

//Example 3:
//Input: nums = [1,3,5,6], target = 7
//Output: 4

fun searchInsert(nums: IntArray, target: Int): Int {
    var left = 0
    var right = nums.size
    if (nums[0] > target) return 0
    if (nums[nums.size - 1] < target) return nums.size
    while (true) {
        val currentIndex = left + (right - left) / 2
        val current = nums[currentIndex]
        if (current == target) {
            return currentIndex
        } else {
            if (target > current) {
                if (currentIndex == nums.size -1) return currentIndex + 1
                if (target < nums[currentIndex + 1]) return currentIndex + 1
                left = currentIndex
            } else if (target < current) {
                if (currentIndex == 0) return 0
                if (target > nums[currentIndex - 1]) return currentIndex
                right = currentIndex
            }
        }
    }
}

// не совсем доделанное рекурсивное решение. Проблемой в этом решении является запоминание того, в какой части массива находится текущий срез массива
//fun recursive(nums: IntArray, target: Int, lastIndex: Int): Int {
//    val currentIndex = nums.size / 2
//    val currentNum = nums[currentIndex]
//    if (target > nums[nums.size - 1]) return lastIndex + 1
//    if (target < nums[0]) return lastIndex - 1
//    return if (currentNum == target) {
//        currentIndex
//    } else if (target > currentNum) {
//        recursive(nums.sliceArray(IntRange(currentIndex, nums.size-1)), target, 0)
//    } else {
//        recursive(nums.sliceArray(IntRange(0, currentIndex)), target, )
//    }
//}