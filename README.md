Go微服务+领域驱动+K8s

### 目录结构
```
.
└── src
    ├── domain                  领域层
    │   ├── aggregation         聚合
    │   └── model               模型
    │       ├── course
    │       ├── price
    │       └── repository      仓储接口
    ├── infrastructure          基础设施层
    │   ├── GormDao             dao层
    │   └── errors              自定义错误
    └── test                    单元测试
        └── internal            内部
            ├── aggregation     聚合测试
            ├── api             api测试
            │   └── ctrl        控制器
            ├── lib
            └── model           模型测试

```