import time


class Demo01:
    def ListDemo(self):
        l = [1, 2, "3"]
        for item in l:
            print(item)

        print(len(l))

    def SetDemo(self):
        s = {1, 2, 1, "a"}  # todo 天然去重
        for item in s:
            print(item)

    def DictDemo(self):
        d = {"a": "1", "b": "2", "a": 11}  # key必须唯一,如有重复，后key会覆盖前key
        for k, v in d.items():
            print(k, v)

    def ByteDemo(self):
        b = bytes("a")
        print(b)

    def ListDemo2(self):
        l = [["a", 1], ["b", 2]]
        for item in l:
            for data in item:
                print(data)

    def ListDemo3(self):

        l2 = []
        d1 = {"start_info": {"value": "2023-05-23", "address": "beijing"}}
        d2 = {"end_info": {"value": "2023-05-25", "address": "tianjin"}}
        l2.append(d1)
        l2.append(d2)
        print("l2:", l2)

        for item in l2:
            print("item:", item)
            for k1, v1 in item.items():
                print("k1:", k1, v1)
                for k2, v2 in v1.items():
                    print("k2:", k2, v2)

    def DictDemo2(self):
        d = {"a": ["a1", "2"], "b": "b11"}
        for key, item in d.items():
            for data in item:
                print(data)


def NoSelf():
    print("no self")


class Demo02:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def Demo(self):
        print(self.name, self.age)

d = Demo01()
# d.ListDemo()
# d.SetDemo()
# d.DictDemo()
# d.ByteDemo()
# d.ListDemo2()
d.ListDemo3()
# d.DictDemo2()
#
# d2 = Demo02("zhang-san", 33)
# d2.Demo()
#
# NoSelf()
