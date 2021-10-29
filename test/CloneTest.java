// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/CloneTest
public class CloneTest implements Cloneable {

    private double pi = 3.14;

    @Override
    public CloneTest clone() {
        try {
            return (CloneTest) super.clone();
        } catch (CloneNotSupportedException e) {
            throw new RuntimeException(e);
        }
    }

    public static void main(String[] args) {
        CloneTest obj1 = new CloneTest();
        CloneTest obj2 = obj1.clone();
        obj1.pi = 3.1415926;
        System.out.println(obj1.pi);
        System.out.println(obj2.pi);
    }

}
/**
 public class CloneTest implements java.lang.Cloneable {
     public CloneTest();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: aload_0
             5: ldc2_w        #2                  // double 3.14d
             8: putfield      #4                  // Field pi:D
             11: return

     public CloneTest clone();
         Code:
             0: aload_0
             1: invokespecial #5                  // Method java/lang/Object.clone:()Ljava/lang/Object;
             4: checkcast     #6                  // class CloneTest
             7: areturn
             8: astore_1
             9: new           #8                  // class java/lang/RuntimeException
             12: dup
             13: aload_1
             14: invokespecial #9                  // Method java/lang/RuntimeException."<init>":(Ljava/lang/Throwable;)V
             17: athrow
         Exception table:
            from    to  target type
                0     7     8   Class java/lang/CloneNotSupportedException

     public static void main(java.lang.String[]);
         Code:
             0: new           #6                  // class CloneTest
             3: dup
             4: invokespecial #10                 // Method "<init>":()V
             7: astore_1
             8: aload_1
             9: invokevirtual #11                 // Method clone:()LCloneTest;
             12: astore_2
             13: aload_1
             14: ldc2_w        #12                 // double 3.1415926d
             17: putfield      #4                  // Field pi:D
             20: getstatic     #14                 // Field java/lang/System.out:Ljava/io/PrintStream;
             23: aload_1
             24: getfield      #4                  // Field pi:D
             27: invokevirtual #15                 // Method java/io/PrintStream.println:(D)V
             30: getstatic     #14                 // Field java/lang/System.out:Ljava/io/PrintStream;
             33: aload_2
             34: getfield      #4                  // Field pi:D
             37: invokevirtual #15                 // Method java/io/PrintStream.println:(D)V
             40: return

     public java.lang.Object clone() throws java.lang.CloneNotSupportedException;
         Code:
             0: aload_0
             1: invokevirtual #11                 // Method clone:()LCloneTest;
             4: areturn
 }
 */
