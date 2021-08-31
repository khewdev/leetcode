class Solution {
    public int[][] transpose(int[][] matrix) {
        int row = matrix.length;
        int column = matrix[0].length;
        
        int[][] newMatrix = new int[column][row];
        
        for(int rowIndex = 0; rowIndex < row; rowIndex++) {
            for(int columnIndex = 0; columnIndex < column; columnIndex++) {
                newMatrix[columnIndex][rowIndex] = matrix[rowIndex][columnIndex];
            }
        }
        
        return newMatrix;
    }
}