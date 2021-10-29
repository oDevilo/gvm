// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/MyObject
public class MyObject {
    public static int staticVar;
    public int instanceVar;
    public static void main(String[] args) {
        int x = 32768; // ldc
        MyObject myObj = new MyObject(); // new
        MyObject.staticVar = x; // putstatic
        x = MyObject.staticVar; // getstatic
        myObj.instanceVar = x; // putfield
        x = myObj.instanceVar; // getfield
        Object obj = myObj;
        if (obj instanceof MyObject) { // instanceof
            myObj = (MyObject) obj; // checkcast
            System.out.println(myObj.instanceVar);
        }
    }
}

/**
 Compiled from "MyObject.java"
 public class MyObject {
 public static int staticVar;

 public int instanceVar;

 public MyObject();
     Code:
         0: aload_0
         1: invokespecial #1                  // Method java/lang/Object."<init>":()V
         4: return

 public static void main(java.lang.String[]);
     Code:
         0: ldc           #2                  // int 32768
         2: istore_1
         3: new           #3                  // class MyObject
         6: dup
         7: invokespecial #4                  // Method "<init>":()V
         10: astore_2
         11: iload_1
         12: putstatic     #5                  // Field staticVar:I
         15: getstatic     #5                  // Field staticVar:I
         18: istore_1
         19: aload_2
         20: iload_1
         21: putfield      #6                  // Field instanceVar:I
         24: aload_2
         25: getfield      #6                  // Field instanceVar:I
         28: istore_1
         29: aload_2
         30: astore_3
         31: aload_3
         32: instanceof    #3                  // class MyObject
         35: ifeq          53
         38: aload_3
         39: checkcast     #3                  // class MyObject
         42: astore_2
         43: getstatic     #7                  // Field java/lang/System.out:Ljava/io/PrintStream;
         46: aload_2
         47: getfield      #6                  // Field instanceVar:I
         50: invokevirtual #8                  // Method java/io/PrintStream.println:(I)V
         53: return
 }
 */
