package top.yukuii.apijava.common;

import lombok.Getter;

@Getter
public class BusinessException extends RuntimeException {

    private Integer code;
    
    public BusinessException(String message) {
        super(message);  // 调用父类构造函数
        this.code = 500;
    }
    
    public BusinessException(Integer code, String message) {
        super(message);
        this.code = code;
    }

}
