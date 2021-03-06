package main

// 常见的内置函数
/**
  常 见的 内置函数
名称	说明
close	用于管道通信
len、cap
  len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；
  cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
new、make	new 和 make 均是用于分配内存：
  new 用于值类型和用户定义的类型，如自定义结构，
  make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作）new() 是一个函数，不要忘记它的括号
copy、append	用于复制和连接切片
panic、recover	两者均用于错误处理机制
      recover 必须与defer使用；defer必须在panic代码之前定义，否则不会被执行（一般都在函数的首行定义defer）。
print、println	底层打印函数，在部署环境中建议使用 fmt 包
complex、real imag	用于
*/
