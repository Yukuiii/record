package top.yukuii.apijava.common;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class Result<T> {

    private Integer code;
    private String message;
    private T data;

    public static <T> Result<T> success(T data) {
        return new Result<T>(200, "success", data);
    }

    public static <T> Result<T> success() {
        return new Result<T>(200, "success", null);
    }

    public static <T> Result<T> error(String message) {
        return new Result<T>(500, message, null);
    }

    public static <T> Result<T> error(Integer code, String message) {
        return new Result<T>(code, message, null);
    }

    public static <T> Result<T> error(Integer code, String message, T data) {
        return new Result<T>(code, message, data);
    }

}
