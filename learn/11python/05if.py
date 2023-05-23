class IFDemo:
    def __init__(self):
        pass

    def Demo01(self):
        num = 1
        if num % 2 == 0:
            print("偶数")
        else:
            print("奇数")

    def Demo02(self):
        num = 1
        if num == 1:
            print("1")
        elif num == 2:
            print("2")
        elif num == 3:
            print("3")
        else:
            print("超出限制")


d = IFDemo()
d.Demo01()
