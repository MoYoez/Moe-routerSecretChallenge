### Moe's Secret Challenge

- 大概算是第一次自己写谜题，不过使用了自己熟悉的后端语言（


#### Router 1

使用 [Text-waterMark](https://github.com/MoYoez/Text-WaterMark)


基础原理如下:

[Be careful what you copy: Invisibly inserting usernames into text with Zero-Width Characters](https://medium.com/@umpox/be-careful-what-you-copy-invisibly-inserting-usernames-into-text-with-zero-width-characters-18b4e6f17b66)

![image.png](https://s2.loli.net/2023/08/19/StqCixnp738zM1D.png)

任何类型的Unicode字符本质上仍然是二进制, 所以说

我们可以使用对于在页面上无法直接显示的字符,将其代表为0 或者是 1 

同时unicode字符中 零宽字符可以满足我们的需求

![image.png](https://s2.loli.net/2023/08/19/BcQuVf5pDTsHeZE.png)

- Example Program Lang: Go
```Go
	for _, char := range encodingMsg {
		switch char {
		case '\u200B': // zero-width space
			binaryBuilder += "1"
		case '\u200C': // zero-width non-joiner
			binaryBuilder += "0"
		default:
			binaryBuilder += " " // add single space
		}
	}
```

将我们需要的任意 Unicode 字符优先转成 binary Code

此处的解密需要使用审查元素:

![image.png](https://s2.loli.net/2023/08/19/P9RFHSyifKANm72.png)

其中可以发现 &zwnj || &ZeroWidthSpace

可以通过拿到原始数据，而不是将他们直接转换成Unicode字符

这一部分建议自行编写工具~因为现成实例工具事实上不能为这个使用（

此处的结果是:

> Achieved The First Goal. | Router To next: time | Requests Args: ?timeunix=?

#### Router 2 && Router 3

放水的，通过提示一个是根据 GMT +8 时间戳 + POST 请求得到 Router 3

Router 3 Char 反转即可.

### License

MIT