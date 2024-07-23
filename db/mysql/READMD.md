## sqlx
> 和标准库的sql有相似的api

### 类型

#### handle types
- sqlx.DB 与sql.DB相似，代表一个数据库对象
- sqlx.Tx 与sql.Tx相似，代表一个事务对象
- sqlx.Stmt 与sql.Stmt相似，代表一个预处理语句对象
- sqlx.NameStmt 代表一个预处理语句对象（命名参数）

handle types都嵌入了标准库sql的api，这表示当你调用sqlx.DB.Query等价于调用sql.DB.Query

#### cursor types
- sqlx.Rows 与sql.Rows相同，由Queryx返回的游标
- sqlx.Row 与sql.Row相同，由QueryRowx返回的游标

与handle types相同，sqlx.Rows嵌入了sql.Rows。由于底层实现不可访问，所以sqlx.Row是保留sql.Row标准接口基础上的部分重新实现

