def demo01():
    i1 = 1
    i2 = 2
    i3 = 1.1
    print(i1 + i2 + i3)
    print("*********************************************")


def demo02():
    s1 = "s1"
    s2 = "s2"
    s3 = s1+s2

    for i in s3:
        print(i)
    print("*********************************************")

def demo03():
    l=[1,2,3,"zz","kx"]
    l.append("append01")
    for i in l:
        print(i)
    print(len(l))
    print("*********************************************")

demo01()
demo02()
demo03()
