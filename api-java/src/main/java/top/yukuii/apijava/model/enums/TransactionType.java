package top.yukuii.apijava.model.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

/**
 * 交易类型枚举
 */
@Getter
@AllArgsConstructor
public enum TransactionType {

    /**
     * 收入
     */
    INCOME("income", "收入"),

    /**
     * 支出
     */
    EXPENSE("expense", "支出");

    /**
     * 枚举编码
     */
    private final String code;

    /**
     * 枚举描述
     */
    private final String description;

    /**
     * 根据编码获取枚举
     */
    public static TransactionType getByCode(String code) {
        for (TransactionType type : values()) {
            if (type.getCode().equals(code)) {
                return type;
            }
        }
        return null;
    }
}