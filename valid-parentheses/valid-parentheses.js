/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    if (s.length === 1 || s.length % 2 !== 0) return false;
    
    let stack = [];
    
    for(let i = 0; i < s.length; i++) {
        if(s[i] === '(')
            stack.push(')');
        else if(s[i] === '{')
            stack.push('}');
        else if(s[i] === '[')
            stack.push(']');
        else if(stack.length === 0 || stack.pop() != s[i])
            return false;
    }
    
    if(stack.length === 0) return true;
    else return false;
};