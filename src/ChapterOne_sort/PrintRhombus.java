package ChapterOne_sort;

/**
 * 打印一个实心的菱形图案，菱形总层数是2n-1
 */
public class PrintRhombus {
    private static char[] charArray;
    private static int index;

    public static void main(String[] args){
//        printRhombus(6);
        PrintrHollowRhombus.printRhombus(6);
    }

    private static void printRhombus(int n){
        charArray = new char[2*n*n-2*n+1];
        for(int i = 0; i<charArray.length; i++){
            int newChar = i +65;
            charArray[i] = (char) newChar;
        }
        for(int i = 1; i<2*n; i++){
            printLine(n, i);
            System.out.println();
        }
    }
    private static void printLine(int n, int i){
        int blankPrint;
        int wellPrint;
        if(i<=n){
            blankPrint = n - i;
            wellPrint = 2*i - 1;
        }else {
            blankPrint = i - n;
            wellPrint = 2*(2*n-i)-1;
        }
        while (blankPrint>0){
            System.out.print(" ");
            blankPrint--;
        }
        while (wellPrint>0){
            System.out.print(charArray[index++]);
            wellPrint--;
        }
    }

    /**
     * 打印一个空心菱形，总层数是2n-1
     */
    private static class PrintrHollowRhombus{
        private static void printRhombus(int n){
            charArray = new char[4*n-4];
            for(int i = 0; i<charArray.length; i++){
                int newChar = i +65;
                charArray[i] = (char) newChar;
            }
            for(int i = 1; i<2*n; i++){
                printLine(n, i);
                System.out.println();
            }
        }
        private static void printLine(int n, int i){
            int blankPrint1;
            int blankPrint2;
            if(i<=n){
                blankPrint1 = n - i;
                blankPrint2 = (2*i - 3)>0 ? 2*i-3 : 0;
            }else {
                blankPrint1 = i - n;
                blankPrint2 = (4*n - 2*i - 3)>0 ? 4*n - 2*i - 3 : 0;
            }
            while (blankPrint1>0){
                System.out.print(" ");
                blankPrint1--;
            }
            System.out.print(charArray[index++]);
            if(blankPrint2 > 0){
                while (blankPrint2>0){
                    System.out.print(" ");
                    blankPrint2--;
                }
                System.out.print(charArray[index++]);
            }
        }
    }

}
