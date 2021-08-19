/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var threeSumClosest = function(nums, target) {
    let result = nums[0] + nums[1] + nums[nums.length-1];
    
    nums.sort((a,b)=>a-b);
    
    for(let i = 0; i < nums.length; i++) {
        let ptrA = i + 1;
        let ptrB = nums.length - 1;
        
        while (ptrA < ptrB) {
            let currentSum = nums[i] + nums[ptrA] + nums[ptrB];
            
            if (currentSum > target) {
                ptrB -= 1;
            } else {
                ptrA += 1;
            }
            
            if(Math.abs(currentSum - target) < Math.abs(result - target)) {
                result = currentSum;
            }
        }
    }
    
    return result;
    
    
};