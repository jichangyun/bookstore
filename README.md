# bookstore

# xx

# test
## 验证 flush 功能
浏览器访问 http://localhost:8888/check?book=go-zero
请求结果
Start return value.
Return value 0 after sleep 0 second.
Return value 1 after sleep 1 second.
Return value 2 after sleep 2 second.
Return value 3 after sleep 3 second.
Return value 4 after sleep 4 second.
{"found":false,"price":0}
现象：
0-4依次显示，并不是一次显示全部。
