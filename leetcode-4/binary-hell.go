package main

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		idxHead := binarySearch(nums1, nums2[len(nums2)-1])
		idxTail := binarySearch(nums1, nums2[0])
		mid := len(nums1) + (idxHead-idxTail+1)*len(nums2)
		if (len(nums1)+len(nums2))%2 == 0 {
			if mid > len(nums1)-1 {
				mid = mid - (len(nums1) - 1)
				return (float64(nums2[mid]) + float64(nums2[mid-1])) / 2
			}
			return (float64(nums1[mid]) + float64(nums1[mid-1])) / 2
		} else {
			if mid > len(nums1)-1 {
				mid = mid - (len(nums1) - 1)
				return float64(nums2[mid])
			}
			return float64(nums1[mid])
		}
	}

	idxHead := binarySearch(nums2, nums1[len(nums1)-1])
	idxTail := binarySearch(nums2, nums1[0])

	mid := len(nums2) + (idxHead-idxTail+1)*len(nums1)
	if (len(nums1)+len(nums2))%2 == 0 {
		if mid > len(nums2)-1 {
			mid = mid - (len(nums2) - 1)
			return (float64(nums1[mid]) + float64(nums1[mid-1])) / 2
		}
		return (float64(nums2[mid]) + float64(nums2[mid-1])) / 2
	} else {
		if mid > len(nums2) {
			mid = mid - len(nums2)
			return (float64(nums1[mid]))
		}
		return float64(nums2[mid])
	}
}
