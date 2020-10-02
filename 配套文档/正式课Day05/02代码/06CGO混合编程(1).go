package main


//所有的C语言代码 需要使用注释写入

/*
#include <stdio.h>
void SayHello()
{
	printf("大家好，我是法师");
}
*/
import "C"

func main() {
	//调用C语言代码
	C.SayHello()
}
