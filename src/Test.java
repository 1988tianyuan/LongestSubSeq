import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
public class Test {
    private static List<Integer> stageList = new ArrayList<Integer>();
    public static void main(String[] args){
        Scanner scanner = new Scanner(System.in);
        int len = scanner.nextInt();
        int[] array = new int[len];
        for(int i = 0; i<len; i++){
            array[i] = scanner.nextInt();
        }
        System.out.println(LongestOrdered(array));
    }
    private static int LongestOrdered(int[] array){
        if(array.length == 0){
            return 0;
        }
        int stage[] = new int[array.length];
        stage[array.length-1] = LongestOrdered(array, array.length-1, stage);
        System.out.println();
        return MaxChase(stage);
    }
    private static int LongestOrdered(int[] array, int k, int[] stage){
        if(1 > k){
            return 1;
        }
        stage[k-1] = LongestOrdered(array, k-1, stage);
        for(int i = 0; i<k; i++){
            if(array[i]<array[k]){
                stageList.add(i);
            }
        }
        if(!stageList.isEmpty()){
            int maxFlag = MaxCheck(stageList, stage);
            stageList.clear();
            return stage[k] = stage[maxFlag] + 1;
        }else {
            stageList.clear();
            return stage[k] = 1;
        }
    }

    private static int MaxCheck(List<Integer>stageList, int[] stage){
        int maxFlag = 0;
        for(int i = 0; i<stageList.size(); i++){
            if(stage[maxFlag]<stage[stageList.get(i)]){
                maxFlag = stageList.get(i);
            }
        }
        return maxFlag;
    }

    private static int MaxChase(int[] array){
        int max = 0;
        for(int i = 0; i<array.length; i++){
            if(array[max]<array[i]){
                max = i;
            }
        }
        return array[max];
    }
}
