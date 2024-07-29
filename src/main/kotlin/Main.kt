fun main(args: Array<String>) {

    println(twoSum(intArrayOf(2,7,11,15),9).joinToString(", "))
    println(twoSum(intArrayOf(3,2,4),6).joinToString(", "))
    println(twoSum(intArrayOf(3,3),6).joinToString(", "))
    println(twoSum(intArrayOf(1,2,5,4,3),4).joinToString(", "))
    println(twoSum(intArrayOf(4,0 ,3,6),6).joinToString(", "))
    println(twoSum(intArrayOf(0,4,3,0),0).joinToString(", "))
    println(twoSum(intArrayOf(-3,4,3,13),0).joinToString(", "))
}

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Input: nums = [3,3], target = 6
// Output: [0,1]


fun twoSum(nums: IntArray, target: Int): IntArray {
    val map = mutableMapOf<Int,Int>()
    for(i: Int in nums.indices) {
        val reversedNum = target - nums[i]
        val reversedIndex = map[reversedNum]
        if (reversedIndex != null) {
            return intArrayOf(reversedIndex, i)
        } else {
            map[nums[i]] = i
        }
    }
    return intArrayOf(0,0)
}