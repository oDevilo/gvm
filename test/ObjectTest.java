// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/ObjectTest
public class ObjectTest {

    public static void main(String[] args) {
        Object obj1 = new ObjectTest();
        Object obj2 = new ObjectTest();
        System.out.println(obj1.hashCode());
        System.out.println(obj1.toString());
        System.out.println(obj1.equals(obj2));
        System.out.println(obj1.equals(obj1));
    }

}
/**
 public class ObjectTest {
     public ObjectTest();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: return

     public static void main(java.lang.String[]);
         Code:
             0: new           #2                  // class ObjectTest
             3: dup
             4: invokespecial #3                  // Method "<init>":()V
             7: astore_1
             8: new           #2                  // class ObjectTest
             11: dup
             12: invokespecial #3                  // Method "<init>":()V
             15: astore_2
             16: getstatic     #4                  // Field java/lang/System.out:Ljava/io/PrintStream;
             19: aload_1
             20: invokevirtual #5                  // Method java/lang/Object.hashCode:()I
             23: invokevirtual #6                  // Method java/io/PrintStream.println:(I)V
             26: getstatic     #4                  // Field java/lang/System.out:Ljava/io/PrintStream;
             29: aload_1
             30: invokevirtual #7                  // Method java/lang/Object.toString:()Ljava/lang/String;
             33: invokevirtual #8                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             36: getstatic     #4                  // Field java/lang/System.out:Ljava/io/PrintStream;
             39: aload_1
             40: aload_2
             41: invokevirtual #9                  // Method java/lang/Object.equals:(Ljava/lang/Object;)Z
             44: invokevirtual #10                 // Method java/io/PrintStream.println:(Z)V
             47: getstatic     #4                  // Field java/lang/System.out:Ljava/io/PrintStream;
             50: aload_1
             51: aload_1
             52: invokevirtual #9                  // Method java/lang/Object.equals:(Ljava/lang/Object;)Z
             55: invokevirtual #10                 // Method java/io/PrintStream.println:(Z)V
             58: return
 }
 */
