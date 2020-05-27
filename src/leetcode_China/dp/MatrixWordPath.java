package leetcode_China.dp;

/**
 * 请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
 * 路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子
 * 在下面的3×4的矩阵中包含一条字符串“bfce”的路径:
 * [['a','b','c','e'],
 * ['s','f','c','s'],
 * ['a','d','e','e']]
 * 但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
 * 
 * 提示：暴力遍历 + 深度优先搜索 + 剪枝 
 */
public class MatrixWordPath {
	
	public static void main(String[] args) {
		test1();
		test2();
		test3();
	}
	
	private static void test1() {
		char[] chars1 = {'a','b','c','e'};
		char[] chars2 = {'s','f','e','s'};
		char[] chars3 = {'a','d','e','e'};
		char[][] board = new char[3][4];
		board[0] = chars1;
		board[1] = chars2;
		board[2] = chars3;
		System.out.println(exist(board, "abcefsadeese"));
	}
	
	private static void test2() {
		char[] chars = {'a','a'};
		char[][] board = new char[1][2];
		board[0] = chars;
		System.out.println(exist(board, "aaa"));
	}
	
	private static void test3() {
		char[] chars1 = {'a','a','a','a'};
		char[] chars2 = {'a','a','a','a'};
		char[] chars3 = {'a','a','a','a'};
		char[][] board = new char[3][4];
		board[0] = chars1;
		board[1] = chars2;
		board[2] = chars3;
		System.out.println(exist(board, "aaaaaaaaaaaaa"));
	}
	
	private static boolean search(int col, int row, char[] chars, char[][] board, int curCharIndex, boolean[][] flagMatrix) {
		if (col < 0 || col >= board[0].length || row < 0 || row >= board.length || flagMatrix[row][col]) {
			return false;
		}
		boolean searched = false;
		if (chars[curCharIndex] == board[row][col]) {
			int nextCharIndex = curCharIndex + 1;
			flagMatrix[row][col] = true;
			if (nextCharIndex >= chars.length) {
				searched = true;
			} else {
				if (search(col + 1, row, chars, board, nextCharIndex, flagMatrix)) {
					searched = true;
				} else if (search(col, row + 1, chars, board, nextCharIndex, flagMatrix)) {
					searched = true;
				} else if (search(col - 1, row, chars, board, nextCharIndex, flagMatrix)) {
					searched = true;
				} else if (search(col, row - 1, chars, board, nextCharIndex, flagMatrix)) {
					searched = true;
				}
			}
		}
		flagMatrix[row][col] = searched;
		return searched;
	}
	
	public static boolean exist(char[][] board, String word) {
		char[] chars = new char[word.length()];
		word.getChars(0, word.length(), chars, 0);
		for (int rowIndex = 0; rowIndex < board.length; rowIndex++) {
			for (int colIndex = 0; colIndex < board[0].length; colIndex++) {
				boolean[][] flagMatrix = new boolean[board.length][board[0].length];
				if (search(colIndex, rowIndex, chars, board, 0, flagMatrix)) {
					return true;
				}
			}
		}
		return false;
	}
}
