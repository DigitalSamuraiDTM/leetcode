package main

func main() {
	println(climbStairs(6))
	println(climbStairs(7))
	println(climbStairs(8))
	println(climbStairs(9))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	var ne = 2
	if n == ne {
		return ne
	}
	var i = 3
	var current = 1 + ne
	for n != i {
		current, ne = ne+current, current
		i++
	}
	return current
}
