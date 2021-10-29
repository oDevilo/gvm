// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/PrintArgs foo bar 你好，世界！
public class PrintArgs {

    public static void main(String[] args) {
        for (String arg : args) {
            System.out.println(arg);
        }
    }

}
/**
 public class PrintArgs {
     public PrintArgs();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: return

     public static void main(java.lang.String[]);
         Code:
             0: aload_0
             1: astore_1
             2: aload_1
             3: arraylength
             4: istore_2
             5: iconst_0
             6: istore_3
             7: iload_3
             8: iload_2
             9: if_icmpge     31
             12: aload_1
             13: iload_3
             14: aaload
             15: astore        4
             17: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             20: aload         4
             22: invokevirtual #3                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             25: iinc          3, 1
             28: goto          7
             31: return
 }
*/
