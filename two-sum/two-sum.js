/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let map = new Map();
    
    for(let i = 0;i < nums.length; i++) {
        let complement = target - nums[i];
        if(!map.has(complement))
            map.set(nums[i], i);
        else
            return [map.get(complement), i];
    }
    
    return false;
};