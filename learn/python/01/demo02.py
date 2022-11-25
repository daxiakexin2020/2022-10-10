class People:
    color = "yellow"

    def __int__(self):
        self.color = "red"

    def eat(self):
        print("eating--")

    def sleep(self):
        print("sleeping--")


p = People()
p.sleep()
p.eat()
p.color = "黑色"
print(p.color)

p = People()
p.sleep()
p.eat()
print(p.color)


class Animal:
    name = ""

    '''todo __私有属性'''
    __age = 0
    color = "黑色"

    def ShowColor(self):
        print(self.color)

    def ShowName(self):
        print(self.name)

    '''todo __: 私有方法'''
    def __ShowAge(self):
        print(self.__age)

    def ShowAge(self):
        print(self.__age)

a=Animal()
a.ShowAge()

class Dog(Animal):
    def Jiao(self):
        print("dog jiao")

    def SetColor(self,color):
        self.color=color


d=Dog()
d.SetColor("hei")
d.ShowColor()

