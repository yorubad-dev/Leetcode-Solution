package main

func increasingTriplet(nums []int) bool {
    point1, point2 := int(^uint(0)>>1), int(^uint(0)>>1)

    for _, num := range nums {
        if num <= point1 {
            point1 = num
        }else if num <= point2 {
            point2 = num
        }else{
            return true
        }
    }

    return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	println(increasingTriplet(nums))
}