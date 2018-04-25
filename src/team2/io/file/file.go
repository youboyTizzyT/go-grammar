package file

import (
	"os"
	"fmt"
)

func File() {
	//writeFile()
	//renameFile()
	//sameFile()
	readFile()
}
/**
写文件
 */
func writeFile()  {
	// 文件路径
	userFile:="./src/team2/io/res/file.txt"
	fout,err:=os.Create(userFile)
	defer fout.Close()
	if err!=nil {
		fmt.Print(userFile,err)
		return 
	}
	fout.WriteString("helloworld!\n")
	fout.Write([]byte("abcd!\n"))
	fmt.Println(fout.Fd())
}
/**
读文件
 */
func readFile() {
	// 文件路径
	userFile:="./src/team2/io/res/file.txt"
	// 打开文件,返回File的内存地址
	fin,err:=os.Open(userFile)
	// 延迟关闭资源
	defer fin.Close()
	if err != nil {
		fmt.Println(userFile,err)
	}
	// 创建一个切片容量为1024 作为缓冲容器
	buf:=make([]byte,2)
	for{
		// 循环读文件数据到缓冲容器中,返回读取到的个数
		n,_:=fin.Read(buf)
		if 0==n {
			// 如果读到个数为0,则读取完毕,跳出循环
			break
		}
		// 从缓冲切片中
		os.Stdout.Write(buf[:n])
		os.Stdout.WriteString("\n123123\n")
	}
}
/**
更改文件夹名字
 */
func renameFile() {
	err:=os.Rename("./src/team2/io/res/file1.txt","./src/team2/io/res/file2.txt")
	fmt.Println(err)
	if err!=nil {
		if os.IsExist(err) {
			fmt.Println("文件已经存在")
			os.Rename("./src/team2/io/res/file.txt","./src/team2/io/res/file1.txt")
		}
	}
}
func sameFile(){
	f1,_:=os.Stat("./src/team2/io/res/file2.txt")
	f2,_:=os.Stat("./src/team2/io/res/file.txt")
	f3,err:=os.Readlink("./src/team2/io/res/file.txt")
	fmt.Println(f3,err)
	if os.SameFile(f1, f2) {
		fmt.Println("两个文件夹相同")
	}else {
		fmt.Println("两个文件夹不一样")
	}
}