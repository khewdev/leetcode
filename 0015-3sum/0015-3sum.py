class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        nums.sort()
        
        for i, a in enumerate(nums):
            # if current number index > 0 and equals to previous number
            if i > 0 and a == nums[i - 1]:
                continue
                
            l, r = i + 1, len(nums) - 1
            # l and r can't be equal
            while l < r:
                threeSum = a + nums[l] + nums[r]
                
                if threeSum > 0:
                    # decrementing r since threeSum is bigger than 0
                    r -= 1 
                elif threeSum < 0:
                    # incrementing l since threeSum is smaller than 0
                    l += 1
                else:
                    res.append([a, nums[l], nums[r]])
                    
                    # only have to update one pointer, and the two conditions
                    # up there will take care of the rest
                    l += 1
                    
                    while nums[l] == nums[l - 1] and l < r:
                        l += 1
                    
        return res