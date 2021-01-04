package leetcode_China.array;

import java.util.HashMap;
import java.util.Map;

public class NumIslands_200 {

    private static class Solution {
        public int numIslands(char[][] grid) {
            IslandNode[][] islandNodeMap = new IslandNode[grid.length][grid[0].length];
            for (int i = 0; i < grid.length; i++) {
                for (int j = 0; j < grid[0].length; j++) {
                    if (grid[i][j] == '0') {
                        continue;
                    }
                    islandNodeMap[i][j] = new IslandNode();
                }
            }
            Map<IslandNode, Byte> rootMap = new HashMap<>();
            for (int i = 0; i < grid.length; i++) {
                for (int j = 0; j < grid[0].length; j++) {
                    if (grid[i][j] == '0') {
                        continue;
                    }
                    IslandNode curNode = islandNodeMap[i][j];
                    IslandNode upNode = null;
                    if (i != 0 && grid[i-1][j] == '1') {
                        upNode = islandNodeMap[i-1][j];
                    }
                    IslandNode leftNode = null;
                    if (j != 0 && grid[i][j-1] == '1') {
                        leftNode = islandNodeMap[i][j-1];
                    }
                    IslandNode newRoot = unionNode(upNode, leftNode, rootMap);
                    if (newRoot != null) {
                        curNode.father = newRoot;
                    } else {
                        rootMap.put(curNode, (byte)1);
                    }
                }
            }
            return rootMap.size();
        }

        private IslandNode unionNode(IslandNode node1, IslandNode node2, Map<IslandNode, Byte> rootMap) {
            IslandNode newRoot;
            IslandNode node1Root = null;
            int node1Deep = 0;
            if (node1 != null) {
                node1Root = node1.father;
                if (node1Root == null) {
                    node1Root = node1;
                } else {
                    node1Deep = 1;
                    while (node1Root.father != null) {
                        node1Root = node1Root.father;
                        node1Deep++;
                    }
                }
            }
            IslandNode node2Root = null;
            int node2Deep = 0;
            if (node2 != null) {
                node2Root = node2.father;
                if (node2Root == null) {
                    node2Root = node2;
                } else {
                    node2Deep = 1;
                    while (node2Root.father != null) {
                        node2Root = node2Root.father;
                        node2Deep++;
                    }
                }
            }
            if (node1Root != null && node2Root != null) {
                if (node1Root == node2Root) {
                    return node1Root;
                }
                if (node1Deep > node2Deep) {
                    node2Root.father = node1Root;
                    rootMap.remove(node2Root);
                    newRoot = node1Root;
                } else {
                    node1Root.father = node2Root;
                    rootMap.remove(node1Root);
                    newRoot = node2Root;
                }
            } else {
                newRoot = node1Root != null ? node1Root : node2Root;
            }
            if (newRoot != null) {
                if (node1 != null) {
                    IslandNode fatherNode;
                    while (node1.father != null) {
                        fatherNode = node1.father;
                        node1.father = newRoot;
                        node1 = fatherNode;
                    }
                }
                if (node2 != null) {
                    IslandNode fatherNode;
                    while (node2.father != null) {
                        fatherNode = node2.father;
                        node2.father = newRoot;
                        node2 = fatherNode;
                    }
                }
            }
            return newRoot;
        }

        private static class IslandNode {
            private IslandNode father;
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        char[][] grid = {
                {'1','1','0','1','1'},
                {'1','1','1','1','1'},
                {'0','0','0','0','0'},
                {'1','1','1','1','1'},
                {'0','1','0','1','1'},
        };
        System.out.println(solution.numIslands(grid));
        char[][] grid2 = {
                {'1','1','0','1','1'},
                {'1','1','1','1','1'},
                {'0','1','1','1','0'},
                {'1','1','1','1','1'},
                {'0','1','0','1','1'},
        };
        System.out.println(solution.numIslands(grid2));
        char[][] grid3 = {
                {'1','1','1','1','0'},
                {'1','1','0','1','0'},
                {'1','1','0','0','0'},
                {'0','0','0','0','0'},
        };
        System.out.println(solution.numIslands(grid3));
        char[][] grid4 = {
                {'1'},
                {'1'},
        };
        System.out.println(solution.numIslands(grid4));
        char[][] grid5 = {
                {'1','1','0','0','0'},
                {'1','1','0','0','0'},
                {'0','0','1','0','0'},
                {'0','0','0','1','1'},
        };
        System.out.println(solution.numIslands(grid5));
    }
}
