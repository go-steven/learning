// run it with: (add . to CLASSPATH)
//   download log4j-api-2.8.2.jar, log4j-api-2.8.2.jar to current directory
//   javac -cp ./log4j-api-2.8.2.jar;./log4j-core-2.8.2.jar JavaLearning.java
//   java -cp ./log4j-api-2.8.2.jar;./log4j-core-2.8.2.jar;./ JavaLearning
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class JavaLearning {
    private static final Logger logger = LogManager.getLogger("JavaLearning");
    
    public static void main(String[] args) throws Exception {
        logger.info("Hello, World!");
    }
}