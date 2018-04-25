import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

public class Test {
    public static void main(String[] args){


        Scanner scanner = new Scanner(System.in);
        int len = scanner.nextInt();
        int[] array = new int[len];
        for(int i = 0; i<len; i++){
            array[i] = scanner.nextInt();
        }

        System.out.println(LongestOrdered(array, array.length-1));



    }
    public static int LongestOrdered(int[] array, int k){
        if(2 > k){
            return 1;
        }

        int longestNum = LongestOrdered(array, k-1);

        if(MaxCheck(array, k)){
            return longestNum + 1;
        }

        return longestNum;
    }

    public static boolean MaxCheck(int[] array,int k){
        boolean flag = true;
        for(int i = 0; i<k; i++){
            if(array[k] <= array[i]){
                flag = false;
            }
        }
        return flag;
    }

}
