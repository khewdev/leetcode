/**
 * @param {string} s
 * @return {boolean}
 */
var checkValidString = function(s) {
    let leftWild = 0;
    let left = 0;

    for(let i = 0; i < s.length; i++) {
        if(s[i] === ')') {
            if(leftWild === 0) return false;
            leftWild--;
        } else {
            leftWild++;
        }

        if(s[i] === '(') {
            left++;
        } else {
            left = Math.max(0, left - 1);
        }
    }
    
    return left === 0;
};