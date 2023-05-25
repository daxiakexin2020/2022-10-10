class InnerFuncDemo:
    def __init__(self):
        pass

    def ABS(self):
        num = -1
        i = abs(num)
        print("abs:{0}".format(i))

    def SUM(self):
        a = [1, 2, 3, 4]
        sum1 = sum(a)
        print("sum:{0}".format(sum1))


id = InnerFuncDemo()
id.ABS()

id.SUM()
