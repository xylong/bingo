```
.
├── api
│   └── v1
├── application         应用层
│   ├── assembler       装配层(领域对象和响应dto之间的类型转换、数据填充)
│   └── dto             dto层
├── cmd
├── domain              领域层
│   ├── aggregation     聚合层
│   ├── repository      仓储层
│   └── model           模型层
├── infrastructure      基础实施层
│   ├── GormDao         dao层
│   └── error           错误封装
├── lib
│   └── db
└── middleware
```

###