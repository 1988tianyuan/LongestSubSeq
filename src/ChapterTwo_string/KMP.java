package ChapterTwo_string;

/**
 * 经典的KMP算法
 */
public class KMP {
    public static boolean kmp(String str1, String str2){
        int i = 0;
        int j = 0;
        int[] next = getNext(str2);
        while (i<str1.length()){
            if(str1.charAt(i) == str2.charAt(j)){
                i++;
                j++;
                if(j == str2.length()){
                    break;
                }
            }else {
                j = next[j];
                if(j < 0){
                    j = 0;
                    i++;
                }
            }
        }
        return j == str2.length();
    }

    private static int[] getNext(String str2){
        int[] next = new int[str2.length()];
        int j = 0;
        int k = -1;
        next[0] = -1;
        while (j<str2.length()-1){
            if(k<0 || str2.charAt(j) == str2.charAt(k)){
                next[++j] = ++k;
            }else {
                k = next[k];
            }
        }
        return next;
    }
    public static void main(String[] args){
        String str1 = "ABCDAABCABDSDASDQ";
        String str2 = "ABCABD";
        System.out.println(kmp(str1, str2));
    }
}
