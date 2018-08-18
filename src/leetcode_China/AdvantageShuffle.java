package leetcode_China;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

/**
 * 给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。
 返回 A 的任意排列，使其相对于 B 的优势最大化。
 */
public class AdvantageShuffle {

    public int[] advantageCount(int[] A, int[] B) {
        Arrays.sort(A);
        int[] flagA = new int[A.length];
        int[] newA = new int[A.length];
        int minIndex = 0;
        for(int i = 0; i < B.length; i++){
            boolean hasFound = false;
            for(int j = 0; j < A.length; j++){
                if(B[i] < A[j] && flagA[j] == 0){
                    newA[i] = A[j];
                    flagA[j] = 1;
                    hasFound = true;
                    break;
                }
            }
            if(!hasFound){
                while (flagA[minIndex] != 0){
                    minIndex++;
                }
                newA[i] = A[minIndex];
                flagA[minIndex] = 1;
            }
        }
        return newA;

    }


}
