package leetcode_China.array;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

public class SubSets_78 {

    public static List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        result.add(new ArrayList<>());
        for (int num : nums) {
            int resultSize = result.size();
            for (int i = 0; i < resultSize; i++) {
                List<Integer> resultE = result.get(i);
                List<Integer> newE = new ArrayList<>(resultE);
                newE.add(num);
                result.add(newE);
            }
        }
        return result;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1,2,3};
        System.out.println(subsets(nums));
    }

    private void sub(List<Integer> tmp, int start, int[] nums, List<List<Integer>> result) {
        if (start >= nums.length) {
            return;
        }
        for (int i = start; i < nums.length; i++) {
            List<Integer> newTmp = new LinkedList<>();
            newTmp.addAll(tmp);
            newTmp.add(nums[i]);
            result.add(newTmp);
            sub(newTmp, i + 1, nums, result);
        }
    }
}
