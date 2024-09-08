# Service

## 创建实体
```bash
# 会生成 ent/schema/user.go 文件
ent new User
```

## 修改生成 user.go 文件，定义表字段
```go
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("username"),
		field.String("email"),
		field.String("password"),
		field.String("bio").Default(""),
		field.String("image").Default(""),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
		field.Time("deleted_at").Optional(),
	}
}
```

## 生成数据库文件
```bash
# 生成数据库文件
go generate ./ent

# 重新生成项目下所有自动生成的文件
make generate
```