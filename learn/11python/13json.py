import json


class JsonDemo:

    def __init__(self):
        pass

    def Demo01(self):
        data = {
            "name": "zz",
            "age": 32,
            "salary": 1000
        }
        str = self.Dumps(data)
        self.Loads(str)

    def Dumps(self, object):
        dumps = json.dumps(object)
        print("json dumps:{0}".format(dumps))
        return dumps

    def Loads(self, str):
        loads = json.loads(str)
        print("json loads:{0}".format(loads))


Jd = JsonDemo()
Jd.Demo01()

a = ["a", 1, {"name": "kx"}]
str = Jd.Dumps(a)
Jd.Loads(str)
