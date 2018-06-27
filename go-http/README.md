# go-http

> 练习net/http模块搭建一个简单web服务

## 知识点

* 依赖管理glide
* http的handler使用
* 静态资源访问
* http拦截
* 重定向
* cookie使用
* template使用
* 数据库gorm

## 依赖管理

go的依赖管理很多，试下下感觉glide顺手

安装

```bash
go get github.com/Masterminds/glide
```

使用说明

```bash
glide -h
```

## http服务

过程大概是：开启端口监听http请求，根据url调用定义的回调函数处理

```go
//开启端口监听
http.ListenAndServe(":9090", nil)

//注册回调的两种方式http.Handle()和http.HandleFunc()
//回调函数
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
http.HandleFunc("/test", SayHello)

//func Handle(pattern string, handler Handler)
http.Handle("/test", http.HandlerFunc(SayHello))
```

## 静态资源访问

静态资源都放在assets下

```go
//url路径不带前辍
http.Handle("/", http.Dir("assets"))
//带前辍得这样处理，不然会404
http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
```

## http拦截

做权限的时候会对收到的http请求过滤，类似filter或interceptor，go里函数可以传递，这样可能传http.HandlerFunc做一些处理再返回一个http.HandlerFunc

```go
func (h *Handler) Authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie(AuthKey)
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	}
}
```

## 重定向

提供了http.Redirect用于重定向，坑有：
* http.Redirect必须要return,不然后面的代码会执行
* 状态码要为http.StatusFound，302。

```go
//func Redirect(w ResponseWriter, r *Request, url string, code int)
if !isLogin {
	http.Redirect(w, r, "/login", http.StatusFound)
	return
}

```

## cookie使用

操作cookie无非就是set，get和remove
```go
//set
//func SetCookie(w ResponseWriter, cookie *Cookie)
expires := time.Now().AddDate(1, 0, 0)
cookie := &http.Cookie{Name: "username", Value: username, Expires: expires}
http.SetCookie(w, cookie)
//get
//func (r *Request) Cookie(name string) (*Cookie, error)
cookie, err := r.Cookie("username")
//remove
//这个没有提供，让Expires就相当于删除了
cookie, _ := r.Cookie("username")
if cookie != nil {
	cookie.Expires = time.Now().AddDate(-1, 0, 0)
}
http.SetCookie(w, cookie)
```

## template使用

文档http://docscn.studygolang.com/pkg/html/template/

## 数据库

* 使用mysql安装`github.com/go-sql-driver/mysql`
* 使用sqlite安装`github.com/mattn/go-sqlite3`

gorm使用先安装上面的，再安装`github.com/jinzhu/gorm`

```go
//先导入驱动程序
//import _ "github.com/jinzhu/gorm/dialects/mysql"
// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"

//连接
//mysql 
db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
//sqlite
db, err := gorm.Open("sqlite3", "sqlite.db")
```
