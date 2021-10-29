// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/GetClassTest
public class GetClassTest {

    public static void main(String[] args) {
        System.out.println(void.class.getName()); // void
        System.out.println(boolean.class.getName()); // boolean
        System.out.println(byte.class.getName()); // byte
        System.out.println(char.class.getName()); // char
        System.out.println(short.class.getName()); // short
        System.out.println(int.class.getName()); // int
        System.out.println(long.class.getName()); // long
        System.out.println(float.class.getName()); // float
        System.out.println(double.class.getName()); // double
        System.out.println(Object.class.getName()); // java.lang.Object
        System.out.println(GetClassTest.class.getName()); // jvmgo.book.ch09.GetClassTest
        System.out.println(int[].class.getName()); // [I
        System.out.println(int[][].class.getName()); // [[I
        System.out.println(Object[].class.getName()); // [Ljava.lang.Object;
        System.out.println(Object[][].class.getName()); // [[Ljava.lang.Object;
        System.out.println(Runnable.class.getName()); // java.lang.Runnable
        System.out.println("abc".getClass().getName()); // java.lang.String
        System.out.println(new double[0].getClass().getName()); // [D
        System.out.println(new String[0].getClass().getName()); // [Ljava.lang.String;
    }

}
/**
 public class GetClassTest {
     public GetClassTest();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: return

     public static void main(java.lang.String[]);
         Code:
             0: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             3: getstatic     #3                  // Field java/lang/Void.TYPE:Ljava/lang/Class;
             6: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             9: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             12: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             15: getstatic     #6                  // Field java/lang/Boolean.TYPE:Ljava/lang/Class;
             18: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             21: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             24: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             27: getstatic     #7                  // Field java/lang/Byte.TYPE:Ljava/lang/Class;
             30: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             33: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             36: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             39: getstatic     #8                  // Field java/lang/Character.TYPE:Ljava/lang/Class;
             42: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             45: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             48: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             51: getstatic     #9                  // Field java/lang/Short.TYPE:Ljava/lang/Class;
             54: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             57: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             60: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             63: getstatic     #10                 // Field java/lang/Integer.TYPE:Ljava/lang/Class;
             66: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             69: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             72: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             75: getstatic     #11                 // Field java/lang/Long.TYPE:Ljava/lang/Class;
             78: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             81: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             84: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             87: getstatic     #12                 // Field java/lang/Float.TYPE:Ljava/lang/Class;
             90: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             93: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             96: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             99: getstatic     #13                 // Field java/lang/Double.TYPE:Ljava/lang/Class;
             102: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             105: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             108: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             111: ldc           #14                 // class java/lang/Object
             113: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             116: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             119: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             122: ldc           #15                 // class GetClassTest
             124: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             127: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             130: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             133: ldc           #16                 // class "[I"
             135: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             138: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             141: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             144: ldc           #17                 // class "[[I"
             146: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             149: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             152: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             155: ldc           #18                 // class "[Ljava/lang/Object;"
             157: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             160: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             163: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             166: ldc           #19                 // class "[[Ljava/lang/Object;"
             168: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             171: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             174: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             177: ldc           #20                 // class java/lang/Runnable
             179: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             182: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             185: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             188: ldc           #21                 // String abc
             190: invokevirtual #22                 // Method java/lang/Object.getClass:()Ljava/lang/Class;
             193: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             196: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             199: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             202: iconst_0
             203: newarray       double
             205: invokevirtual #22                 // Method java/lang/Object.getClass:()Ljava/lang/Class;
             208: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             211: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             214: getstatic     #2                  // Field java/lang/System.out:Ljava/io/PrintStream;
             217: iconst_0
             218: anewarray     #23                 // class java/lang/String
             221: invokevirtual #22                 // Method java/lang/Object.getClass:()Ljava/lang/Class;
             224: invokevirtual #4                  // Method java/lang/Class.getName:()Ljava/lang/String;
             227: invokevirtual #5                  // Method java/io/PrintStream.println:(Ljava/lang/String;)V
             230: return
 }
 */
