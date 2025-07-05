package top.yukuii.apijava.model.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

/**
 * 货币类型枚举
 */
@Getter
@AllArgsConstructor
public enum CurrencyType {

    /**
     * 人民币
     */
    CNY("CNY", "人民币", "¥"),

    /**
     * 美元
     */
    USD("USD", "美元", "$"),

    /**
     * 欧元
     */
    EUR("EUR", "欧元", "€"),

    /**
     * 英镑
     */
    GBP("GBP", "英镑", "£"),

    /**
     * 日元
     */
    JPY("JPY", "日元", "¥"),

    /**
     * 港币
     */
    HKD("HKD", "港币", "HK$"),

    /**
     * 新台币
     */
    TWD("TWD", "新台币", "NT$"),

    /**
     * 韩元
     */
    KRW("KRW", "韩元", "₩");

    /**
     * 货币代码
     */
    private final String code;

    /**
     * 货币名称
     */
    private final String name;

    /**
     * 货币符号
     */
    private final String symbol;

    /**
     * 根据代码获取枚举
     */
    public static CurrencyType getByCode(String code) {
        for (CurrencyType currency : values()) {
            if (currency.getCode().equals(code)) {
                return currency;
            }
        }
        return null;
    }
}