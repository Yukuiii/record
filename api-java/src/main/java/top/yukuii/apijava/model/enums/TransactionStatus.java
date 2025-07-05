package top.yukuii.apijava.model.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

/**
 * 交易状态枚举
 */
@Getter
@AllArgsConstructor
public enum TransactionStatus {

    /**
     * 待处理
     */
    PENDING("pending", "待处理"),

    /**
     * 已完成
     */
    COMPLETED("completed", "已完成"),

    /**
     * 已取消
     */
    CANCELLED("cancelled", "已取消"),

    /**
     * 处理中
     */
    PROCESSING("processing", "处理中"),

    /**
     * 失败
     */
    FAILED("failed", "失败");

    /**
     * 状态编码
     */
    private final String code;

    /**
     * 状态描述
     */
    private final String description;

    /**
     * 根据编码获取枚举
     */
    public static TransactionStatus getByCode(String code) {
        for (TransactionStatus status : values()) {
            if (status.getCode().equals(code)) {
                return status;
            }
        }
        return null;
    }
}