from threading import Thread

i = 0

def thread1():
    for j in range(1000000):
        global i
        i += 1

def thread2():
    for k in range(1000000):
        global i
        i -= 1

def main():
    t1 = Thread(target = thread1, args = (),)
    t2 = Thread(target = thread2, args = (),)
    t1.start()
    t2.start()
    t1.join()
    t2.join()
    print(i)


main()
