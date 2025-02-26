## 安装mariadb
```
brew install mariadb
```
这里如果不指定 mariadb 则安装最新版 mariadb，如果本地 mariadb 包依旧很老，可以使用
```
brew update
```
更新软件包列表，
然后可以通过
```
brew search mariadb
```
查看本地可安装版本
```
==> Formulae
mariadb                    mariadb@10.4               mariadb@11.1
mariadb-connector-c        mariadb@10.5               mariadb@11.2
mariadb-connector-odbc     mariadb@10.6               mariadb@11.4
mariadb@10.10              mariadb@10.9               qt-mariadb
mariadb@10.11 ✔            mariadb@11.0
```
这里以 10.11 为例
```
brew install mariadb@10.11 
```
## 修改端口
可在 /usr/local/etc/my.cnf 修改端口
```
# Default Homebrew MySQL server config
[mysqld]
# Only allow connections from localhost
bind-address = 127.0.0.1
port=3307
```
## 启动服务
后台常驻服务
```
brew services start mariadb@10.11
```
## 修改 root 密码
### 方法一【失败】
可能密码为空，直接登录
ps：执行路径以实际安装版本为准
```
/usr/local/Cellar/mariadb@10.11/10.11.11/bin/mysql -P 3307 -u root -p
```
提示需要输入密码，失败
### 方法二【失败】
通过
```
/usr/local/Cellar/mariadb@10.11/10.11.11/bin/mariadb-secure-installation
```
设置密码
结果执行后依旧提示需要输入密码
### 方法三【成功】
通过执行
```
/usr/local/Cellar/mariadb@10.11/10.11.11/bin/mariadb
```
进入db命令行
再执行
```
ALTER USER root@localhost IDENTIFIED VIA mysql_native_password USING PASSWORD("chenjie");
```
修改密码，这里密码就被修改成了 chenjie
## 重启服务
```
brew services restart mariadb@10.11
```
这个时候就可以用 root / chenjie 正常登录 db 了
## 客户端推荐
这里推荐sequel-ace,专门为 maridb/mysql 定制，属于Sequel Pro的社区精简版，使用了一下还是比较方便的
```
brew install sequel-ace
```


