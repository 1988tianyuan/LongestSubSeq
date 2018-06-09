import algorithmPracticeChapterOne.ShellSort;

import java.text.SimpleDateFormat;
import java.util.*;

public class Test {
    private static Random random = new Random();
    private static SimpleDateFormat simpleDateFormat = new SimpleDateFormat("hh:mm:ss");
    public static void main(String[] args){
        int[]array = new int[]{54,35,48,36,27,12,44,44,8,14,26,17,28};
        for(Integer i: ShellSort.shellSort(array, array.length)){
            System.out.println(i+" ");
        }
    }




}

