import ChapterTwo_string.TransformWord;
import leetcode_China.*;

import java.text.SimpleDateFormat;
import java.util.*;

import static ChapterTwo_string.TransformWord.chkTransform;

public class Test {
    private static Random random = new Random();
    private static SimpleDateFormat simpleDateFormat = new SimpleDateFormat("hh:mm:ss");
    public static void main(String[] args){
        BiggestSquare biggestSquare = new BiggestSquare();
        char[][]matrix = {{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','1','1','1'}};
        System.out.println(biggestSquare.maximalSquare(matrix));


    }




}
