package learngo

// W.A.P to find if 5 exists in the arry
// a := []int{1, 2, 3, 4, 5, 6, 7}
// fmt.Println(learngo.FindValueInArray(5, a))
//findValueInArray(int,[]int) bool

// for loop --> n --> linear search,
func FindValueInArray(findVal int, arr []int) (bool, int) {
	for i, value := range arr {
		if value == findVal {
			return true, i
		}
	}
	return false, -1
}

// binary search , sorted array
// 1. left pointer that holds index of left most value  (lp),
//2.  Right Pointer--> index of right most value. (Rp)
// 3. find middle and check below things:
// check if lp == rp --> exit.
// 3.1 is target == value at middle, if no then
// 3.2 check if it is less then middle value -- if true --> then check in left side., Else..
// 3.3 check in right side.

// n,n/2,n/4,..,n/2^n :-> log(n)with base 2
