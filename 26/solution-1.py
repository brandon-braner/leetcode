from typing import List


from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        """
        the constraint of in place makes this harder then it needs to be
        basically it switches the time and space complexity
        the below is O(n) time because of the loop and space of O(1) since we don't make a copy
        if we could use a set we would have O(1) time and O(n) space but less code
        """
        i = 1
        size = len(nums)

        for j in range(1, size):
            if nums[j - 1] != nums[j]:
                nums[i] = nums[j]
                i = i + 1
        return i


arr = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]


solution = Solution()
solution.removeDuplicates(arr)
