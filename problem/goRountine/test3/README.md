# 测试goroutine 在cpu中的迁移
将goroutine绑定到cpu上
结论
* goroutine在每次退让cpu时会换cpu
* 每次sleep都会退让cpu 每次退让后下一次线程执行时将会重新选择cpu

