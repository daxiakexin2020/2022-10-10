import time


class TimeDemo:
    def __init__(self):
        pass

    def Demo01(self):
        time_time = time.time()
        print("time.time():{0}".format(time_time))

    def Demo02LocalTime(self):
        print("time localtime str:{0}".format(time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())))


td = TimeDemo()
td.Demo01()
td.Demo02LocalTime()
