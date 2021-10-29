// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/BubbleSortTest
public class BubbleSortTest {

    public static void main(String[] args) {
        int[] arr = {
            22, 84, 77, 11, 95,  9, 78, 56,
            36, 97, 65, 36, 10, 24 ,92, 48
        };

        //printArray(arr);
        bubbleSort(arr);
        //System.out.println(123456789);
        printArray(arr);
    }

    private static void bubbleSort(int[] arr) {
        boolean swapped = true;
        int j = 0;
        int tmp;
        while (swapped) {
            swapped = false;
            j++;
            for (int i = 0; i < arr.length - j; i++) {
                if (arr[i] > arr[i + 1]) {
                    tmp = arr[i];
                    arr[i] = arr[i + 1];
                    arr[i + 1] = tmp;
                    swapped = true;
                }
            }
        }
    }

    private static void printArray(int[] arr) {
        for (int i : arr) {
            System.out.println(i);
        }
    }

}
/**
 public class BubbleSortTest {
     public BubbleSortTest();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: return

     public static void main(java.lang.String[]);
         Code:
             0: bipush        16
             2: newarray       int
             4: dup
             5: iconst_0
             6: bipush        22
             8: iastore
             9: dup
             10: iconst_1
             11: bipush        84
             13: iastore
             14: dup
             15: iconst_2
             16: bipush        77
             18: iastore
             19: dup
             20: iconst_3
             21: bipush        11
             23: iastore
             24: dup
             25: iconst_4
             26: bipush        95
             28: iastore
             29: dup
             30: iconst_5
             31: bipush        9
             33: iastore
             34: dup
             35: bipush        6
             37: bipush        78
             39: iastore
             40: dup
             41: bipush        7
             43: bipush        56
             45: iastore
             46: dup
             47: bipush        8
             49: bipush        36
             51: iastore
             52: dup
             53: bipush        9
             55: bipush        97
             57: iastore
             58: dup
             59: bipush        10
             61: bipush        65
             63: iastore
             64: dup
             65: bipush        11
             67: bipush        36
             69: iastore
             70: dup
             71: bipush        12
             73: bipush        10
             75: iastore
             76: dup
             77: bipush        13
             79: bipush        24
             81: iastore
             82: dup
             83: bipush        14
             85: bipush        92
             87: iastore
             88: dup
             89: bipush        15
             91: bipush        48
             93: iastore
             94: astore_1
             95: aload_1
             96: invokestatic  #2                  // Method bubbleSort:([I)V
             99: aload_1
             100: invokestatic  #3                  // Method printArray:([I)V
             103: return
 }
*/
