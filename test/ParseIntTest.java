// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/ParseIntTest 123
// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/ParseIntTest abc
// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/ParseIntTest
public class ParseIntTest {

    public static void main(String[] args) {
        foo(args);
    }

    private static void foo(String[] args) {
        try {
            bar(args);
        } catch (NumberFormatException e) {
            System.out.println(e.getMessage());
        }
    }

    private static void bar(String[] args) {
        if (args.length == 0) {
            throw new IndexOutOfBoundsException("no args!");
        }
        int x = Integer.parseInt(args[0]);
        System.out.println(x);
    }

}
/**
 public class ParseIntTest {
     public ParseIntTest();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: return

     public static void main(java.lang.String[]);
         Code:
             0: aload_0
             1: invokestatic  #2                  // Method foo:([Ljava/lang/String;)V
             4: return
 }
 */
