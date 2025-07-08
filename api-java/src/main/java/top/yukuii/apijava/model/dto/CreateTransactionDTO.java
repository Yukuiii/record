package top.yukuii.apijava.model.dto;

import java.math.BigDecimal;

import lombok.Data;

@Data
public class CreateTransactionDTO {

    private String type;

    private Long categoryId;

    private BigDecimal amount;

    private String description;

    private String paymentMethod;

    private String remark;

    private String tags;
    
}
