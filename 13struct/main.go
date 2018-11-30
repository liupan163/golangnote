package _3struct

//结构体类型的某个字段，声明中只有一个类型名，该字段代表了什么？

//字段声明代表了该类型的一个嵌入字段。
//go语言规定：如果一个字段只有字段名而没有字段名称，那么就是嵌入字段，也叫匿名字段。也就是说既是名称，也是类型。

//一、go语言是用嵌入字段实现了继承吗？
//A: go语言中，根本没有继承的概念，所做的就是利用嵌入字段的方式实现了类型组合---利用代码简洁性替换了扩展性

//二、值方法和指针方法都是什么意思？有什么区别？
/*
指针方法
func (cat *Cat) SetName(name string)  {
	cat.name = name
}
值方法
func (cat Cat) SetName(name string)  {
	cat.name = name
}*/


//Q:
//1、我们可以在结构体类型中嵌入某个类型的指针类型吗？需要注意哪些？
//2、字面量struct{}代表了什么？有什么用处？

//A:
//1、可以在结构体内嵌入某个类型的指针类型，他和普通指针类似，默认初始化为nil。用之前需要初始化
//2、空结构体不占用空间，但是具有结构体的一切属性。如可以有方法，写入channel。
