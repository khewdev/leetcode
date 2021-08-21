/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function(n) {
    let output_arr = [];
    backtrack(output_arr, "", 0, 0, n);
    
    return output_arr;
};

let backtrack = (output_arr, str, open, close, max) => {
    if(str.length === max * 2) {
        output_arr.push(str);
        return;
    }
    
    if(open < max) backtrack(output_arr, str+"(", open+1, close, max);
    if(close < open) backtrack(output_arr, str+")", open, close+1, max);
}



