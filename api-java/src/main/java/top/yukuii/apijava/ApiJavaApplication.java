package top.yukuii.apijava;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("top.yukuii.apijava.mapper")
public class ApiJavaApplication {

    public static void main(String[] args) {
        SpringApplication.run(ApiJavaApplication.class, args);
    }

}
