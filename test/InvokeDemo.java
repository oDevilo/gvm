// go run *.go -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_202.jdk/Contents/Home/jre" test/InvokeDemo
public class InvokeDemo implements Runnable {
    public static void main(String[] args) {
        new InvokeDemo().test();
    }

    public void test() {
        InvokeDemo.staticMethod(); // invokestatic
        InvokeDemo demo = new InvokeDemo(); // invokespecial
        demo.instanceMethod(); // invokespecial
        super.equals(null); // invokespecial
        this.run(); // invokevirtual
        ((Runnable) demo).run(); // invokeinterface
    }
    public static void staticMethod() {}
    private void instanceMethod() {}
    @Override public void run() {}
}
/**
Compiled from "InvokeDemo.java"
public class InvokeDemo implements java.lang.Runnable {
  public InvokeDemo();
    Code:
       0: aload_0
       1: invokespecial #1                  // Method java/lang/Object."<init>":()V
       4: return

  public static void main(java.lang.String[]);
    Code:
       0: new           #2                  // class InvokeDemo
       3: dup
       4: invokespecial #3                  // Method "<init>":()V
       7: invokevirtual #4                  // Method test:()V
      10: return

  public void test();
    Code:
       0: invokestatic  #5                  // Method staticMethod:()V
       3: new           #2                  // class InvokeDemo
       6: dup
       7: invokespecial #3                  // Method "<init>":()V
      10: astore_1
      11: aload_1
      12: invokespecial #6                  // Method instanceMethod:()V
      15: aload_0
      16: aconst_null
      17: invokespecial #7                  // Method java/lang/Object.equals:(Ljava/lang/Object;)Z
      20: pop
      21: aload_0
      22: invokevirtual #8                  // Method run:()V
      25: aload_1
      26: invokeinterface #9,  1            // InterfaceMethod java/lang/Runnable.run:()V
      31: return

  public static void staticMethod();
    Code:
       0: return

  public void run();
    Code:
       0: return
}

*/
