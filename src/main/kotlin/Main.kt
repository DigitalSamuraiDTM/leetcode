
fun main(args: Array<String>) {

    assert(maxArea(intArrayOf(1,8,6,2,5,4,8,3,7)) == 49)
    assert(maxArea(intArrayOf(1,1)) == 1)
    assert(maxArea(intArrayOf(1,9,1,1,1,1,1,1,1)) == 8)
    assert(maxArea(intArrayOf(1,9,1,1,1,1,1,1,100,100)) == 100)
    assert(maxArea(intArrayOf(1,2,3,1,1,100,100,1,1,3,2,1)) == 100 )
    assert(maxArea(intArrayOf(1,2,100,2,1,1,1,1,2,2,2,1)) == 18)
}

fun maxArea(height: IntArray): Int {
    var leftPointer = 0
    var rightPointer = height.size - 1
    var maxSize = 0
    while (rightPointer - leftPointer != 0) {
        val leftElement = height[leftPointer]
        val rightElement = height[rightPointer]
        val containerHeight = if (leftElement>rightElement) rightElement else leftElement
        val containerSize = containerHeight * (rightPointer - leftPointer)
        if (containerSize > maxSize) {
            maxSize = containerSize
        }
//        println(containerSize)

        if (leftElement > rightElement) {
            rightPointer --
        } else if (leftElement < rightElement) {
            leftPointer ++
        } else {
            leftPointer ++
        }
    }
//    println("RETURN: ${maxSize}")
    return maxSize
}