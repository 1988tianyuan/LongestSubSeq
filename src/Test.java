import ChapterTwo_string.TransformWord;
import leetcode_China.*;

import java.text.SimpleDateFormat;
import java.util.*;

import static ChapterTwo_string.TransformWord.chkTransform;

public class Test {
    private static Random random = new Random();
    private static SimpleDateFormat simpleDateFormat = new SimpleDateFormat("hh:mm:ss");
    public static void main(String[] args){
        LongestIncreasSubSeq lis = new LongestIncreasSubSeq();
        System.out.println(lis.lengthOfLIS(new int[]{4,10,4,3,8,9}));
    }




}

