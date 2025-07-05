package top.yukuii.apijava.model.vo;

import lombok.Data;

import java.math.BigDecimal;

@Data
public class GetTransactionVO {

    private Long id;

    private String type;

    private Long categoryId;

    private String categoryName;

    private BigDecimal amount;

    private String description;

    private Long transactionDate;

    private String paymentMethod;

    private String status;

    private String location;

    private String currency;

    private String tags;

    private String remark;

    private Long createTime;

    private Long updateTime;

}