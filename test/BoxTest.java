import java.util.ArrayList;
import java.util.List;
// 自动装箱/拆箱
// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/BoxTest
public class BoxTest {

    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        list.add(3);
        System.out.println(list.toString());
        for (int x : list) {
            System.out.println(x);
        }
    }

}
/**
 java中自动拆箱装箱不是虚拟机的事情，而是编译器在编译的时候已经处理，如下面
 public class BoxTest {
     public BoxTest();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: return

     public static void main(java.lang.String[]);
         Code:
             0: new           #2                  // class java/util/ArrayList
             3: dup
             4: invokespecial #3                  // Method java/util/ArrayList."<init>":()V
             7: astore_1
             8: aload_1
             9: iconst_1
             // 利用Integer.valueOf 装箱
             10: invokestatic  #4                  // Method java/lang/Integer.valueOf:(I)Ljava/lang/Integer;
             13: invokeinterface #5,  2            // InterfaceMethod java/util/List.add:(Ljava/lang/Object;)Z
             18: pop
             19: aload_1
             20: iconst_2
             21: invokestatic  #4                  // Method java/lang/Integer.valueOf:(I)Ljava/lang/Integer;
             24: invokeinterface #5,  2            // InterfaceMethod java/util/List.add:(Ljava/lang/Object;)Z
             29: pop
             30: aload_1
             31: iconst_3
             32: invokestatic  #4                  // Method java/lang/Integer.valueOf:(I)Ljava/lang/Integer;
             35: invokeinterface #5,  2            // InterfaceMethod java/util/List.add:(Ljava/lang/Object;)Z
             40: pop
             41: getstatic     #6                  // Field java/lang/System.out:Ljava/io/PrintStream;
             44: aload_1
             45: invokevirtual #7                  // Method java/lang/Object.toString:()Ljava/lang/String;
             48: invokevirtual #8                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             51: aload_1
             52: invokeinterface #9,  1            // InterfaceMethod java/util/List.iterator:()Ljava/util/Iterator;
             57: astore_2
             58: aload_2
             59: invokeinterface #10,  1           // InterfaceMethod java/util/Iterator.hasNext:()Z
             64: ifeq          90
             67: aload_2
             68: invokeinterface #11,  1           // InterfaceMethod java/util/Iterator.next:()Ljava/lang/Object;
             73: checkcast     #12                 // class java/lang/Integer
             // 利用Integer.intValue 拆箱
             76: invokevirtual #13                 // Method java/lang/Integer.intValue:()I
             79: istore_3
             80: getstatic     #6                  // Field java/lang/System.out:Ljava/io/PrintStream;
             83: iload_3
             84: invokevirtual #14                 // Method java/io/PrintStream.println:(I)V
             87: goto          58
             90: return
 }
 */
