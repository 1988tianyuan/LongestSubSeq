package ChapterTwo_string;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

/**
 * 给定两个字符串A和B及他们的长度，请返回一个bool值，代表他们是否互为变形词
 */
public class TransformWord {

    /**
     * 用Hashmap遍历两个字符串并保存各字符出现次数，然后判断两字符串的字符和出现次数是否都相同
     * @param A
     * @param lena
     * @param B
     * @param lenb
     * @return
     */
    public static boolean chkTransform(String A, int lena, String B, int lenb) {
        if(A == null || B == null || lena != A.length() || lenb != B.length() || lena != lenb){
            return false;
        }
        Map<Character, Integer> map1 = new HashMap<>((int)Character.MAX_VALUE);
        Map<Character, Integer> map2 = new HashMap<>((int)Character.MAX_VALUE);
        for(Character c:A.toCharArray()){
            Integer oldValue = map1.put(c, 1);
                if(oldValue != null){
                map1.put(c, ++oldValue);
            }
        }
        for(Character c:B.toCharArray()){
            Integer oldValue = map2.put(c, 1);
            if(oldValue != null){
                map2.put(c, ++oldValue);
            }
        }
        Set<Map.Entry<Character, Integer>> set = map1.entrySet();
        for(Map.Entry entry:set){
            Character key = (Character) entry.getKey();
            Integer value2 = map2.get(key);
            if(value2 == null || !value2.equals(entry.getValue())){
                return false;
            }
        }
        return true;
    }

    /**
     * 用一个新数组arrayACount来保存String A出现过的字符的出现次数，数组下标即为出现字符的ASCII码
     * 然后遍历String B，以B的字符为下标在arrayACount中查找，若查找结果为0说明A中没有该字符
     * 若查找结果不为0则说明A中有该字符，则将arrayACount中该字符出现次数减1，保证正确校验B中字符的出现次数
     *
     * @param A
     * @param lena
     * @param B
     * @param lenb
     * @return
     */
    public static boolean chkTransformByArray(String A, int lena, String B, int lenb){
        if(A == null || B == null || lena != A.length() || lenb != B.length() || lena != lenb){
            return false;
        }
        int[] arrayACount = new int[Character.MAX_VALUE];
        char[] charArrayA = A.toCharArray();
        for(int i = 0; i<lena; i++){
            int charA = charArrayA[i];
            arrayACount[charA]++;
        }
        char[] charArrayB = B.toCharArray();
        for(int j = lenb-1; j>=0; j--){
            int charB = charArrayB[j];
            if(arrayACount[charB] == 0){
                return false;
            }
            arrayACount[charB]--;
        }
        return true;
    }

}
