import contextlib
from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        with contextlib.suppress(ValueError):
            while val in nums:
                nums.remove(val)
        return len(nums)


# nums = [3, 2, 2, 3]
# val = 3

nums = [0, 1, 2, 2, 3, 0, 4, 2]
val = 2

solution = Solution()
print(solution.removeElement(nums, val))
