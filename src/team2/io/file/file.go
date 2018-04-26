package file

import (
	os "os"
	"fmt"
	"io"
	"bufio"
	"io/ioutil"
)

func File() {
	//writeFile()
	//renameFile()
	//sameFile()
	//readFile()
	//copyFile()
	//bufIoCopyFile()
	ioutilCopyFile()

}
/**
写文件
 */
func writeFile()  {
	// 文件路径
	userFile:="./src/team2/io/res/file.txt"
	fout,err:=os.Create(userFile)
	defer fout.Close()
	// 判断返回值部位nil
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
/**
判断两个文件夹是否相同
 */
func sameFile(){
	f1,_:=os.Stat("./src/team2/io/res/file2.txt")
	f2,_:=os.Stat("./src/team2/io/res/file.txt")
	if os.SameFile(f1, f2) {
		fmt.Println("两个文件夹相同")
	}else {
		fmt.Println("两个文件夹不一样")
	}
}
func copyFile() {
	// 打开要复制的文件
	fi,err:=os.Open("./src/team2/io/res/file2.txt")
	if err != nil {
		panic(err)
	}
	// 使用defer关闭
	defer fi.Close()
	// 创建一个文件,进行复制
	fo,err:=os.Create("./src/team2/io/res/file3.txt")
	if err!=nil {
		panic(err)
	}
	// 使用defer关闭
	defer fo.Close(	)
	// 创建一个缓冲
	buf:=make([]byte,1024)
	for {
		n,err:=fi.Read(buf)
		if err!=nil&&err!=io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if n2,err:=fo.Write(buf[:n]);err!=nil {
			panic(err)
		}else if n2!=n {
			panic("error in writing")
		}
	}
}
/**
通过文件的缓冲流进行复制文件
 */
func bufIoCopyFile() {
	// 打开要复制的文件
	fi,err:=os.Open("./src/team2/io/res/file2.txt")
	if err != nil {
		panic(err)
	}
	// 使用defer关闭
	defer fi.Close()
	//创建一个读取缓冲流
	r:=bufio.NewReader(fi)
	// 创建一个文件,进行复制
	fo,err:=os.Create("./src/team2/io/res/file4.txt")
	if err!=nil {
		panic(err)
	}
	// 使用defer关闭
	defer fo.Close()
	//创建一个写取缓冲流
	w:=bufio.NewWriter(fo)
	// 创建一个缓冲取
	buf:=make([]byte,1024)
	for {
		n,err:=r.Read(buf)
		if err != nil &&err!=io.EOF{
			panic(err)
		}
		if n==0 {
			break
		}
		if n2,err:=w.Write(buf[:n]);err!=nil{
			panic(err)
		}else if n2!=n {
			panic("error in writing")
		}
	}
	if err=w.Flush();err!=nil {
		panic(err)
	}
}
/**
通过ioutil进行文件复制
 */

func ioutilCopyFile() {
	b, err := ioutil.ReadFile("./src/team2/io/res/file3.txt")//读文件
	fmt.Println(string(b))
	if err != nil { panic(err) }
	// 后面的644 为权限
	err = ioutil.WriteFile("./src/team2/io/res/file5.txt", b, 0644)//写文件
	if err != nil { panic(err) }

}

