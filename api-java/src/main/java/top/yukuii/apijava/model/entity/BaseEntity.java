package top.yukuii.apijava.model.entity;

import lombok.Data;

@Data
public class BaseEntity {

    /**
     * 创建时间戳
     */
    private Long createTime;

    /**
     * 更新时间戳
     */
    private Long updateTime;

    /**
     * 创建者用户ID
     */
    private Long createBy;

    /**
     * 更新者用户ID
     */
    private Long updateBy;
    
}
