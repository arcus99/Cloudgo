#Cloudgo


##框架选择
这里所选择使用的框架是课程群里推荐的Martini。Martini 是一个非常新的 Go 语言的 Web 框架，使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，你可使用自己的 DB 层、会话管理和模板。

特性：

**·**   使用非常简单
**·**   无侵入设计
**·**   可与其他 Go 的包配合工作
**·**   超棒的路径匹配和路由
**·**   模块化设计，可轻松添加工具
**·**   大量很好的处理器和中间件
**·**   很棒的开箱即用特性
**·**   完全兼容 http.HandlerFunc 接口
##框架安装
直接使用
`go get github.com/go-martini/martini`

即可。
##测试
####**·**   启动服务器
`go run main.go -p9090`
然后在浏览器中（不关闭终端）打开
`http://localhost:9090/hello/主机名`
可以看到页面上出现service中的字符：
![浏览器](http://img.blog.csdn.net/20171114223654673?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
此时终端中的监听信息如下
![这里写图片描述](http://img.blog.csdn.net/20171114223722445?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
####**·**   进行curl测试
在不关闭当前终端的情况下开启另一个终端，并在该终端上使用curl命令
`curl -v http://localhost:9090/hello/主机名`
![这里写图片描述](http://img.blog.csdn.net/20171114223949301?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
####**·**   进行ApacheBench测试
使用命令
`ab -n 1000 -c 100 http://localhost:9090/hello/主机名`
测试结果如下
![这里写图片描述](http://img.blog.csdn.net/20171114224242549?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
![这里写图片描述](http://img.blog.csdn.net/20171114224253357?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
从结果中我们可以得到，总共发送了1000个请求，失败0个，耗时0.207秒，其中50%的请求可以在19ms内被响应，90%的请求可以在25ms内被响应，最长的响应时间为35ms。