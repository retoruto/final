课程地址：</br>
https://jersey.github.io/documentation/latest/rx-client.html#d0e5556
</br>练习要求：</br>

**1.依据文档图6-1，用中文描述 Reactive 动机**</br>
图6-1：
![这里写图片描述](http://img.blog.csdn.net/20180103134259242?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
简单描述为：一个用户同时申请多个服务，直观上可以为每一个服务都发送一次申请，但是这样不利于效率工作。使用无扩展的客户服务之后，用户只需要为所有的服务申请一次即可得到所有的响应，提高了效率。

---------------------------------------------------
**2.使用 go HTTPClient 实现图 6-2 的 Naive Approach**</br>
图6-2：
![这里写图片描述](http://img.blog.csdn.net/20180103134734002?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

------------------------------
**3.为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3**</br>
图6-3：
![这里写图片描述](http://img.blog.csdn.net/20180103134839227?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

**4.对比两种实现，用数据说明 go 异步 REST 服务协作的优势**</br>
通过2和3的实现，运行代码的结果如下：</br>
使用第一种方法：</br>

![这里写图片描述](http://img.blog.csdn.net/20180103135819738?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
</br>使用第二种方法：</br>
![这里写图片描述](http://img.blog.csdn.net/20180103135844851?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
</br>对比：</br>
第一种方法比第二种方法明显处理时间要长许多，在运行过程可以观察到第一种方法得出结果之前需要几秒的延迟，而第二种方法几乎没有延迟。
