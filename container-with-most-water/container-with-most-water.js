/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
    let ptrA = 0;
    let ptrB = height.length - 1;
    let maxArea = 0;
    
    while(ptrA < ptrB) {
        if (height[ptrA] < height[ptrB]) {
            maxArea = Math.max(maxArea, height[ptrA] * (ptrB - ptrA));
            ptrA++;
        } else {
            maxArea = Math.max(maxArea, height[ptrB] * (ptrB - ptrA));
            ptrB--;
        }
    }
    
    return maxArea;
};