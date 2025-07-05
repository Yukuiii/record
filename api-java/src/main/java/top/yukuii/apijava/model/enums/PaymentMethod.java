package top.yukuii.apijava.model.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

/**
 * 支付方式枚举
 */
@Getter
@AllArgsConstructor
public enum PaymentMethod {

    /**
     * 现金
     */
    CASH("cash", "现金"),

    /**
     * 银行卡
     */
    BANK_CARD("bank_card", "银行卡"),

    /**
     * 信用卡
     */
    CREDIT_CARD("credit_card", "信用卡"),

    /**
     * 支付宝
     */
    ALIPAY("alipay", "支付宝"),

    /**
     * 微信支付
     */
    WECHAT_PAY("wechat_pay", "微信支付"),

    /**
     * 云闪付
     */
    UNION_PAY("union_pay", "云闪付"),

    /**
     * PayPal
     */
    PAYPAL("paypal", "PayPal"),

    /**
     * Apple Pay
     */
    APPLE_PAY("apple_pay", "Apple Pay"),

    /**
     * Google Pay
     */
    GOOGLE_PAY("google_pay", "Google Pay"),

    /**
     * 其他
     */
    OTHER("other", "其他");

    /**
     * 支付方式编码
     */
    private final String code;

    /**
     * 支付方式名称
     */
    private final String name;

    /**
     * 根据编码获取枚举
     */
    public static PaymentMethod getByCode(String code) {
        for (PaymentMethod method : values()) {
            if (method.getCode().equals(code)) {
                return method;
            }
        }
        return null;
    }
}