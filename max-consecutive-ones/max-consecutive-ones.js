/**
 * @param {number[]} nums
 * @return {number}
 */

var findMaxConsecutiveOnes = function(nums) {
    let max = 0;
    let current_number_of_ones = 0;
    for(let i=0; i<nums.length;i++) {
        if(nums[i]==1) {
            current_number_of_ones+=1;
            max = Math.max(max, current_number_of_ones);
        }
        else {
            current_number_of_ones=0;
        }
    }
    return max;
};