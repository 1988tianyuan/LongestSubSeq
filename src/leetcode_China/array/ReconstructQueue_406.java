package leetcode_China.array;

import java.util.Arrays;

public class ReconstructQueue_406 {

    //输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
    //输出：people = [[7,0],[7,1],[6,1],[5,0],[5,2],[1,4]]
    //输出：result = [[7,0],[],[],[],[],[]]
    //输出：result = [[7,0],[7,1],[],[],[],[]]
    //输出：result = [[7,0],[6,1],[7,1],[],[],[]]
    //输出：result = [[5,0],[7,0],[6,1],[7,1],[],[]]
    //输出：result = [[5,0],[7,0],[5,2],[6,1],[7,1],[]]
    //输出：result = [[5,0],[7,0],[5,2],[6,1],[1,4],[7,1]]
    public int[][] reconstructQueue(int[][] people) {
        Tuple[] tuples = new Tuple[people.length];
        for (int i = 0; i < people.length; i++) {
            tuples[i] = new Tuple(people[i][0], people[i][1]);
        }
        Arrays.sort(tuples);
        int[][] result = new int[people.length][2];
        for (Tuple tuple : tuples) {
            insertIntoResult(result, tuple.pair[1], tuple.pair);
        }
        return result;
    }

    private void insertIntoResult(int[][] result, int index, int[] pair) {
        int[] tmp = pair;
        while (index < result.length) {
            if (result[index][0] == 0 && result[index][1] == 0) {
                result[index] = tmp;
                return;
            } else {
                int[] tmp2 = result[index];
                result[index] = tmp;
                tmp = tmp2;
                index++;
            }
        }
    }

    private static class Tuple implements Comparable<Tuple> {
        public Tuple(int h, int k) {
            this.pair = new int[2];
            this.pair[0] = h;
            this.pair[1] = k;
        }
        private int[] pair;
        @Override
        public int compareTo(Tuple t2) {
            if (pair[0] > t2.pair[0]) {
                return -1;
            } else if (pair[0] == t2.pair[0]) {
                if (pair[1] >= t2.pair[1]) {
                    return 1;
                } else {
                    return -1;
                }
            }
            return 1;
        }
    }

    public static void main(String[] args) {
        ReconstructQueue_406 reconstructQueue_406 = new ReconstructQueue_406();
        int[][] people = new int[][]{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}};
        int[][] result = reconstructQueue_406.reconstructQueue(people);
        System.out.println("");
        int[][] people2 = new int[][]{{9,0},{7,0},{1,9},{3,0},{2,7},{5,3},{6,0},{3,4},{6,2},{5,2}};
        // [[3,0],[6,0],[7,0],[5,2],[3,4],[6,2],[5,3],[2,7],[9,0],[1,9]]
        // [[3,0],[6,0],[7,0],[5,2],[3,4],[5,3],[6,2],[2,7],[9,0],[1,9]]
        int[][] result2 = reconstructQueue_406.reconstructQueue(people2);
        System.out.println("");
    }
}
