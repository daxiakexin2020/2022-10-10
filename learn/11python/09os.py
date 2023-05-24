import os


class OSDemo:
    def __init__(self):
        pass

    def Demo01(self):
        currentPath = os.getcwd()  # 当前的工作目录
        print("currentPath:{0}".format(currentPath))

        systemRes = os.system('ls -al')  # 执行系统命令
        print("systemRes:{0},type:{1}".format(systemRes, type(systemRes)))

        path = currentPath + "/test.txt"
        fd = os.open(path, os.O_RDWR | os.O_CREAT | os.O_APPEND)  # 打开文件  写入文件
        os.write(fd, str.encode("abc\n"))
        os.close(fd)

        stat = os.stat(path)  # 文件信息，大小，时间等信息
        print("stat:{0}".format(stat))
        print("stat info,size:{0},mode:{1}".format(stat.st_size, stat.st_mode))


osd = OSDemo()
osd.Demo01()
