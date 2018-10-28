# 删除表
DROP TABLE user
# 创建表
CREATE TABLE user (id integer PRIMARY KEY,email character varying,password character varying,name character varying,updated_at integer,created_at integer);
# 添加记录
Insert user (id ,email ,password,name ,updated_at,created_at) VALUES (1,'michael.chu@test.com','123456');
# 查询记录
SELECT count(id) FROM user WHERE email=LOWER("michael.chu@test.com") LIMIT 1
# 查询记录
SELECT id, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER("michael.chu@test.com") LIMIT 1"

