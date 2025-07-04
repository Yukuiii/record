package top.yukuii.apijava.model.entity;

import lombok.Data;

@Data
public class BaseEntity {

    private Long createTime;

    private Long updateTime;

    private Long createBy;

    private Long updateBy;
    
}
