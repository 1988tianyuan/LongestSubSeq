import ChapterTwo_string.TransformWord;
import leetcode_China.*;

import java.text.SimpleDateFormat;
import java.util.*;

import static ChapterTwo_string.TransformWord.chkTransform;

public class Test {
    private static Random random = new Random();
    private static SimpleDateFormat simpleDateFormat = new SimpleDateFormat("hh:mm:ss");
    public static void main(String[] args){
//        int[][] nums = new int[][]{new int[]{1,3,1}, new int[]{1,5,1},new int[]{4,2,1}};
        int[][] nums = new int[][]{new int[]{1,3,1}, new int[]{1,5,1}};
        MinPathSum minPathSum = new MinPathSum();
        System.out.println(minPathSum.minPathSum(nums));
    }




}

