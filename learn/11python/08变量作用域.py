num = 1  # global


class NamespaceDemo:
    def __init__(self):
        pass

    def Demo01(self):
        global num  #如果不声明，仅仅是内部修改，不会影像外部的使用
        num = 3
        print(num)


nd = NamespaceDemo()
nd.Demo01()

print(num)
