import ChapterTwo_string.TransformWord;
import leetcode_China.*;

import java.text.SimpleDateFormat;
import java.util.*;

import static ChapterTwo_string.TransformWord.chkTransform;

public class Test {
    private static Random random = new Random();
    private static SimpleDateFormat simpleDateFormat = new SimpleDateFormat("hh:mm:ss");
    public static void main(String[] args){
        int[] nums = new int[]{2,7,9,3,1};
        HouseRobber robber = new HouseRobber();
        System.out.println(robber.rob(nums));
    }




}

