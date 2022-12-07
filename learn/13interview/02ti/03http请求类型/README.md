body
1. application/x-www-form-urlencoded
请求参数格式key1=val1&key2=val2的方式进行拼接，并放到请求实体里面，如果是中文或特殊字符等会自动进行URL转码。一般用于表单提交

2. multipart/form-data
与application/x-www-form-urlencoded不同，它会将表单的数据处理为一条消息，
以标签为单元，用分隔符 boundary分开。
既可以上传键值对，也可以上传文件。
当上传的字段是文件时，会有Content-Type来表名文件类型,
content-disposition用来说明字段的一些信息,最后以隔符 boundary–为结束标识。
multipart/form-data支持文件上传的格式，一般需要上传文件的表单则用该类型

3.application/json
application/json 作为响应头比较常见。
实际上，现在越来越多的人把它作为请求头，用来告诉服务端消息主体是序列化后的 JSON 字符串，
其中一个好处就是JSON 格式支持比键值对复杂得多的结构化数据。
由于 JSON 规范的流行，除了低版本 IE 之外的各大浏览器都原生支持JSON.stringify，服务端语言也都有处理 JSON 的函数，使用起来没有困难。

4.application/xml 和 text/xml  text/html, text/plain, text/css, text/javascript, image/jpeg, image/png, image/gif