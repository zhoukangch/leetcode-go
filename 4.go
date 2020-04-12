package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	var grayIndex, curIndex, curPreIndex, stepIndex int
	var grayNums, curNums, curPreNums []int
	grayIndex = -1
	if len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] <= nums2[0] {
			curNums = nums1
			grayNums = nums2
		} else {
			curNums = nums2
			grayNums = nums1
		}
	} else if len(nums1) > 0 {
		curNums = nums1
		grayNums = nums2
	} else {
		curNums = nums2
		grayNums = nums1
	}
	for stepIndex = 1; stepIndex <= totalLen/2; stepIndex++ {
		curPreIndex = curIndex
		curPreNums = curNums
		if curIndex < len(curNums)-1 && grayIndex < len(grayNums)-1 {
			if curNums[curIndex+1] <= grayNums[grayIndex+1] {
				curIndex++
			} else {
				curIndex, grayIndex = swapIndex(curIndex, grayIndex+1)
				curNums, grayNums = swapNums(curNums, grayNums)
			}
		} else if curIndex < len(curNums)-1 {
			curIndex++
		} else {
			curIndex, grayIndex = swapIndex(curIndex, grayIndex+1)
			curNums, grayNums = swapNums(curNums, grayNums)
		}
	}
	if totalLen%2 == 0 {
		return float64(curPreNums[curPreIndex]+curNums[curIndex]) / 2.0
	} else {
		return float64(curNums[curIndex])
	}
}

func swapIndex(index1 int, index2 int) (int, int) {
	return index2, index1
}

func swapNums(nums1 []int, nums2 []int) ([]int, []int) {
	return nums2, nums1
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
