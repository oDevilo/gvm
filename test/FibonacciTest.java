// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/FibonacciTest
public class FibonacciTest {

    public static void main(String[] args) {
        long x = fibonacci(30);
        System.out.println(x);
    }

    private static long fibonacci(long n) {
        if (n <= 1) {
            return n;
        } else {
            return fibonacci(n - 1) + fibonacci(n - 2);
        }
    }

}
/**
 Compiled from "FibonacciTest.java"
 public class FibonacciTest {
     public FibonacciTest();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: return

     public static void main(java.lang.String[]);
         Code:
             0: ldc2_w        #2                  // long 30l
             3: invokestatic  #4                  // Method fibonacci:(J)J
             6: lstore_1
             7: getstatic     #5                  // Field java/lang/System.out:Ljava/io/PrintStream;
             10: lload_1
             11: invokevirtual #6                  // Method java/io/PrintStream.println:(J)V
             14: return
 }

 */
