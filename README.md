# seckill

互联网实践-秒杀系统实践
### 部署环境
** MacOS Big Sur 11.1 **

### 操作步骤
**配置apache路径**
```
vi /etc/apache2/httpd.conf
```
配置vhost
```
Include /private/etc/apache2/extra/httpd-vhosts.conf
```
```
vi /private/etc/apache2/extra/httpd-vhosts.conf
```
自定后端义项目路径
```
<VirtualHost *:80>
    ServerAdmin webmaster@localhost
    DocumentRoot "/Users/xuhaotian/gitrepo/seckill/backend/src/public"
    <Directory "/Users/xuhaotian/gitrepo/seckill/backend/src/public">
        AllowOverride All
        Allow from all
    </Directory>
    ServerName localhost
    ServerAlias localhost
    ErrorLog "/private/var/log/apache2/myhost-error_log"
    CustomLog "/private/var/log/apache2/myhost-access_log" common
</VirtualHost>
```

加载rewrite模块
```
LoadModule rewrite_module libexec/apache2/mod_rewrite.so
```
加载php模块
```
LoadModule php7_module libexec/apache2/libphp7.so
```
重启apache服务
```
sudo apachectl restart
```
确保项目文件夹拥有足够权限
```
chmod -R ugo+rwx * 
```

配置mysql
```
sudo mysql.server start
```
添加mysql账户和数据库，设置用户名和密码与database.php中相同，如下，并给予账户权限。
```
'database'        => 'phpdemo',
// 用户名
'username'        => 'phpdemo',
// 密码
'password'        => '4wj7fGAKJR2ddXDx',
```
初始化数据库中数据，（运行GO前安装import所需库，且go env环境正常,安装gorm，go-redis等）
安装redis方式见下文
在golang目下
```
go run src/*.go
```

登陆后台管理员账号
```
http://localhost/index.php/admin/login/index.html
```
账号可用用户名密码如下
```
admin2
admin2
```
可以进行添加删除商品、更改密码等操作


**配置前端**
进入目录front
安装node环境，使用npm构建vue环境
安装依赖模块vuex、axios等
```
npm i
npm install --save vuex
npm install axios --save
npm install qs --save-dev
```
运行vue前端, 运行在localhost:8080
```
npm run dev
```

**配置RabbitMQ**
下载安装rabbitmq并启动
```
brew install rabbitmq
```
查看rabbitmq是否成功启动
```
http://localhost:15672/
```

**配置Redis**
在golang目录下
```
go mod init github.com/seckill
go get github.com/go-redis/redis/v8
go build src/*.go
```

**运行Go服务后登陆前端页面**
```
http://localhost:8080/#/login
```
登陆账户密码可用
```
user1
upw1
```


