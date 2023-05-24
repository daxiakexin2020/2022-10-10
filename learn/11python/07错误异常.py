import sys

class ErrDemo:

    def __init__(self):
        pass

    def Demo01(self):
        try:
            num = 1
            num2 = 0
            print(num / num2)
        except:
            print("Unexpected error:", sys.exc_info()[0])
        finally:
            print("finally")


rd = ErrDemo()
rd.Demo01()
