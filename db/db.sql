# 删除表
DROP TABLE user
# 创建表
CREATE TABLE user (id integer PRIMARY KEY,email TEXT, password TEXT,name TEXT,updated_at integer,created_at integer);
# 添加记录
Insert into user (email ,password,name ,updated_at,created_at) VALUES ('michael.chu@test.com','123456','chu',12345,123456);
# 查询记录
SELECT count(id) FROM user WHERE email=LOWER("michael.chu@test.com") LIMIT 1
# 查询记录
SELECT id, email, password, name, updated_at, created_at FROM user WHERE email=LOWER("michael.chu@test.com") LIMIT 1

