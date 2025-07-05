package top.yukuii.apijava.model.vo;

import lombok.Data;

import java.util.List;

/**
 * 分类树结构VO
 */
@Data
public class CategoryTreeVO {

    private Long id;

    private String name;

    private String code;

    private Long parentId;

    private Integer level;

    private String type;

    private String icon;

    private String color;

    private Integer sortOrder;

    private Integer isSystem;

    private Integer status;

    private String description;

    /**
     * 子分类列表
     */
    private List<CategoryTreeVO> children;

}