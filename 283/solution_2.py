# This should run in O(n) time being a better solution
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """

        n = len(nums)
    
    # Track the boundary of the zero region
        zero_bound = 0
        for i in range(n):
            if nums[i] != 0:
                nums[zero_bound], nums[i] = nums[i], nums[zero_bound]
                zero_bound += 1