package info.macias;

// Sample program to explore the limits of Java Threads creation
// CAUTION!! May hang your laptop
public class SpawnThreadsUntilOOM {
    public static void main(String[] args) {
        Runtime runtime = Runtime.getRuntime();
        while (true) {
            new Thread(() -> {
                System.out.println("active threads = " + Thread.activeCount()
                        + " (total memory: " + (runtime.totalMemory() / 1_000_000) + " MB)");
                try {
                    Thread.sleep(1_000_000_000);
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }
            }).start();
        }
    }
}
