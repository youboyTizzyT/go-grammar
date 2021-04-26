# 测试goroutine 在cpu中的迁移


输出结果
    
    worker: 7, CPU: 4
    worker: 2, CPU: 4
    worker: 1, CPU: 4
    worker: 0, CPU: 7
    worker: 6, CPU: 4
    worker: 1, CPU: 4
    worker: 0, CPU: 3
    worker: 0, CPU: 3
    worker: 3, CPU: 2
    worker: 7, CPU: 3
    worker: 5, CPU: 3
    worker: 6, CPU: 3
    worker: 6, CPU: 3
    worker: 5, CPU: 3
    worker: 2, CPU: 0
    worker: 0, CPU: 6
    worker: 7, CPU: 0
    worker: 6, CPU: 0
    worker: 3, CPU: 2
    worker: 1, CPU: 4
    worker: 4, CPU: 4
    worker: 5, CPU: 2
    worker: 3, CPU: 0
    worker: 6, CPU: 2
    worker: 4, CPU: 4
    worker: 0, CPU: 2
    worker: 1, CPU: 7
    worker: 1, CPU: 7
    worker: 2, CPU: 4
    worker: 4, CPU: 4
    worker: 7, CPU: 4
    worker: 5, CPU: 1
    worker: 0, CPU: 1
    worker: 0, CPU: 1
    worker: 6, CPU: 1
    worker: 2, CPU: 6
    worker: 6, CPU: 6
    worker: 7, CPU: 0
    worker: 5, CPU: 6
    worker: 3, CPU: 1
    worker: 7, CPU: 0
    worker: 7, CPU: 0
    worker: 1, CPU: 6
    worker: 3, CPU: 0
    worker: 6, CPU: 0
    worker: 4, CPU: 6
    worker: 0, CPU: 3
    worker: 5, CPU: 6
    worker: 7, CPU: 3
    worker: 1, CPU: 6
    worker: 6, CPU: 2
    worker: 5, CPU: 6
    worker: 2, CPU: 2
    worker: 3, CPU: 4
    worker: 7, CPU: 2
    worker: 0, CPU: 4
    worker: 6, CPU: 2
    worker: 5, CPU: 4
    worker: 4, CPU: 2
    worker: 1, CPU: 2
    worker: 0, CPU: 4
    worker: 2, CPU: 0
    worker: 7, CPU: 6
    worker: 3, CPU: 0
    worker: 6, CPU: 6
    worker: 6, CPU: 0
    worker: 1, CPU: 1
    worker: 6, CPU: 0
    worker: 0, CPU: 0
    worker: 7, CPU: 3
    worker: 7, CPU: 2
    worker: 5, CPU: 3
    worker: 4, CPU: 2
    worker: 2, CPU: 0
    worker: 3, CPU: 2
    worker: 1, CPU: 0
    worker: 6, CPU: 4
    worker: 5, CPU: 4
    worker: 6, CPU: 2
    worker: 3, CPU: 3
    worker: 7, CPU: 2
    worker: 3, CPU: 7
    worker: 0, CPU: 4

    `

结论
* goroutine在每次退让cpu时会换cpu
* 每次sleep都会退让cpu 每次退让后下一次线程执行时将会重新选择cpu

