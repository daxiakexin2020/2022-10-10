import os
import sys


class FileDemo:
    def __init__(self):
        pass

    def Demo01(self):
        try:
            path = os.getcwd() + "/test01.txt"
            # os.remove(path)
            fd = open(path, 'a+')
            fd.write("123\n")
            fd.close()
        except OSError as err:
            print("err:{0}".format(err))
        except:
            print("Unexpected error:", sys.exc_info()[0])
        finally:
            print("finally")

fd = FileDemo()
fd.Demo01()
