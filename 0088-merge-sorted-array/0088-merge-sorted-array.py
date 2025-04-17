class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        # iteration decrease in python
        i = m - 1
        j = n - 1
        k = m + n - 1
        
        while j > -1:
            if (nums2[j] > nums1[i]):
                nums1[k] = nums2[j]
                j = j - 1
                # print(f"nums1 -> {nums1}")
            else:
                if i >= 0 :
                    nums1[k] = nums1[i]
                    i = i - 1
                else: 
                    nums1[k] = nums2[j]
                    j = j - 1
            
            k = k - 1
        
        return nums1