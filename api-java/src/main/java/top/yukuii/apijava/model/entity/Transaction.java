package top.yukuii.apijava.model.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;

import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;

import java.math.BigDecimal;

@Data
@EqualsAndHashCode(callSuper=false)
@Builder
@TableName("transaction")
public class Transaction extends BaseEntity {

    /**
     * 交易记录主键ID
     */
    @TableId(type = IdType.AUTO)
    private Long id;

    /**
     * 用户ID，关联用户表
     */
    private Long userId;

    /**
     * 交易类型：income-收入，expense-支出
     */
    private String type;

    /**
     * 交易分类ID，关联分类表
     */
    private Long categoryId;

    /**
     * 交易金额
     */
    private BigDecimal amount;

    /**
     * 交易描述信息
     */
    private String description;

    /**
     * 交易日期时间戳
     */
    private Long transactionDate;

    /**
     * 支付方式：现金、银行卡、支付宝、微信等
     */
    private String paymentMethod;

    /**
     * 交易状态：pending-待处理，completed-已完成，cancelled-已取消
     */
    private String status;

    /**
     * 交易地点
     */
    private String location;

    /**
     * 货币类型：CNY、USD、EUR等
     */
    private String currency;

    /**
     * 交易标签，多个标签用逗号分隔
     */
    private String tags;

    /**
     * 交易备注信息
     */
    private String remark;

}