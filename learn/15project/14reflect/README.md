https://www.jianshu.com/p/c76367756d78

    type Type interface {

        // 变量的内存对齐，返回 rtype.align
        Align() int
    
        // struct 字段的内存对齐，返回 rtype.fieldAlign
        FieldAlign() int
    
        // 根据传入的 i，返回方法实例，表示类型的第 i 个方法
        Method(int) Method
    
        // 根据名字返回方法实例，这个比较常用
        MethodByName(string) (Method, bool)
    
        // 返回类型方法集中可导出的方法的数量
        NumMethod() int
    
        // 只返回类型名，不含包名
        Name() string
    
        // 返回导入路径，即 import 路径
        PkgPath() string
    
        // 返回 rtype.size 即类型大小，单位是字节数
        Size() uintptr
    
        // 返回类型名字，实际就是 PkgPath() + Name()
        String() string
    
        // 返回 rtype.kind，描述一种基础类型
        Kind() Kind
    
        // 检查当前类型有没有实现接口 u
        Implements(u Type) bool
    
        // 检查当前类型能不能赋值给接口 u
        AssignableTo(u Type) bool
    
        // 检查当前类型能不能转换成接口 u 类型
        ConvertibleTo(u Type) bool
    
        // 检查当前类型能不能做比较运算，其实就是看这个类型底层有没有绑定 typeAlg 的 equal 方法。
        // 打住！不要去搜 typeAlg 是什么，不然你会陷进去的！先把本文看完。
        Comparable() bool
    
        // 返回类型的位大小，但不是所有类型都能调这个方法，不能调的会 panic
        Bits() int
    
        // 返回 channel 类型的方向，如果不是 channel，会 panic
        ChanDir() ChanDir
    
        // 返回函数类型的最后一个参数是不是可变数量的，"..." 就这样的，同样，如果不是函数类型，会 panic
        IsVariadic() bool
    
        // 返回所包含元素的类型，只有 Array, Chan, Map, Ptr, Slice 这些才能调，其他类型会 panic。
        // 这不是废话吗。。其他类型也没有包含元素一说。
        Elem() Type
    
        // 返回 struct 类型的第 i 个字段，不是 struct 会 panic，i 越界也会 panic
        Field(i int) StructField
    
        // 跟上边一样，不过是嵌套调用的，比如 [1, 2] 就是说返回当前 struct 的第1个struct 的第2个字段，适用于 struct 本身嵌套的类型
        FieldByIndex(index []int) StructField
    
        // 按名字找 struct 字段，第二个返回值 ok 表示有没有
        FieldByName(name string) (StructField, bool)
    
        // 按函数名找 struct 字段，因为 struct 里也可能有类型是 func 的嘛
        FieldByNameFunc(match func(string) bool) (StructField, bool)
        
        // 返回函数第 i 个参数的类型，不是 func 会 panic
        In(i int) Type
    
        // 返回 map 的 key 的类型，不是 map 会 panic
        Key() Type
    
        // 返回 array 的长度，不是 array 会 panic
        Len() int
    
        // 返回 struct 字段数量，不是 struct 会 panic
        NumField() int
    
        // 返回函数的参数数量，不是 func 会 panic
        NumIn() int
    
        // 返回函数的返回值数量，不是 func 会 panic
        NumOut() int
    
        // 返回函数第 i 个返回值的类型，不是 func 会 panic
        Out(i int) Type
}