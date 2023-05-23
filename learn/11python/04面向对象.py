class People:
    def __init__(self, name, age):
        self.name = name
        self.age = age
        self.___priID = "test01"
        self.publicCommon = "common"
        print("People called")

    def priid(self):
        return self.___priID

    def showName(self):
        print("People Print=>>>>>>name:%s,age:%d" % (self.name, self.age))


class SchoolStudent(People):
    def __init__(self, name, age, level):
        People.__init__(self, name, age)
        self.level = level

    def showName(self):
        print("School Student Print=>>>>>>level:%d" % (self.level))
        print("School Student Print=>>>>>>name:%s,age:%d,level:%d" % (self.name, self.age, self.level))
        super().showName()
        # print("School Student Print=>>>>>>priID" % (self.___priID)) 错误，父类私有属性
        print("School Student Print=>>>>>>publicCommon:%s" % (self.publicCommon))
        print("no call parent class")

    def PriID(self):
        return super().priid()


ss = SchoolStudent("zhang-san", 33, 2)
# ss.showName()

id = ss.PriID()
print("ID:%s" % (id))
