#服务器tcp抓包


前提：
* 需要知道服务器上对应的客户端ip地址。
* 客户端连接到的服务器网卡名
* 安装tcpdump



安装tcpdump
``yum install tcpdump``

查看服务器本地网卡名称
``tcpdump -D``

抓包
``tcpdump -X -i 网卡名 host 客户端ip``-X是以16进制ASCII打印包内容。