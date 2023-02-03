print("hello world")


def demo():
    print("start")
    for i in range(10):
        if i % 2 == 0:
            continue
        print(i)

def demo2():
    print(type(1),id(1))


def demo3():
    l={"name":"zz","age":20}
    for i in l:
        print(i,l[i])

def demo4():
    try:
        print("try")
    except:
        print("except")

def demo5():
    return 2

def demo06(a,b):
    return a+b

def demo07(a,b):
    return a+b

# demo()
# demo2()
# demo3()
# demo4()
# print(demo5())
# print(type(demo06(1,2)))
# print(type(demo06("1","2")))
print(demo07(2,1))