// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/StringTest
public class StringTest {

    public static void main(String[] args) {
        String s1 = "abc1";
        String s2 = "abc1";
        System.out.println(s1 == s2);

        int x = 1;
        String s3 = "abc" + x;
        System.out.println(s1 == s3);

        s3 = s3.intern();
        System.out.println(s1 == s3);
    }

}
/**
 public class StringTest {
     public StringTest();
         Code:
             0: aload_0
             1: invokespecial #1                  // Method java/lang/Object."<init>":()V
             4: return

     public static void main(java.lang.String[]);
         Code:
             0: ldc           #2                  // String abc1
             2: astore_1
             3: ldc           #2                  // String abc1
             5: astore_2
             6: getstatic     #3                  // Field java/lang/System.out:Ljava/io/PrintStream;
             9: aload_1
             10: aload_2
             11: if_acmpne     18
             14: iconst_1
             15: goto          19
             18: iconst_0
             19: invokevirtual #4                  // Method java/io/PrintStream.println:(Z)V
             22: iconst_1
             23: istore_3
             24: new           #5                  // class java/lang/StringBuilder
             27: dup
             28: invokespecial #6                  // Method java/lang/StringBuilder."<init>":()V
             31: ldc           #7                  // String abc
             33: invokevirtual #8                  // Method java/lang/StringBuilder.append:(Ljava/lang/String;)Ljava/lang/StringBuilder;
             36: iload_3
             37: invokevirtual #9                  // Method java/lang/StringBuilder.append:(I)Ljava/lang/StringBuilder;
             40: invokevirtual #10                 // Method java/lang/StringBuilder.toString:()Ljava/lang/String;
             43: astore        4
             45: getstatic     #3                  // Field java/lang/System.out:Ljava/io/PrintStream;
             48: aload_1
             49: aload         4
             51: if_acmpne     58
             54: iconst_1
             55: goto          59
             58: iconst_0
             59: invokevirtual #4                  // Method java/io/PrintStream.println:(Z)V
             62: aload         4
             64: invokevirtual #11                 // Method java/lang/String.intern:()Ljava/lang/String;
             67: astore        4
             69: getstatic     #3                  // Field java/lang/System.out:Ljava/io/PrintStream;
             72: aload_1
             73: aload         4
             75: if_acmpne     82
             78: iconst_1
             79: goto          83
             82: iconst_0
             83: invokevirtual #4                  // Method java/io/PrintStream.println:(Z)V
             86: return
 }
 */
