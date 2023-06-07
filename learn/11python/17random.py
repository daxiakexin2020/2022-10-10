import random


class RandomDemo:
    def __init__(self):
        pass

    def Demo01(self):
        random.seed()
        r = random.random()
        print("random:{0}".format(r))

    def Demo02(self):
        randrange = random.randrange(1, 10)
        print("range:{0}".format(randrange))


demo = RandomDemo()
demo.Demo01()
demo.Demo02()
