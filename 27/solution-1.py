from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        try:
            while val in nums:
                nums.remove(val)
        except ValueError:
            pass
        return len(nums)


# nums = [3, 2, 2, 3]
# val = 3

nums = [0, 1, 2, 2, 3, 0, 4, 2]
val = 2

solution = Solution()
print(solution.removeElement(nums, val))
