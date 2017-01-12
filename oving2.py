import threading

i = 0
Lock = threading.Lock()


def thread1():
    for j in range(1000000):
        Lock.acquire()
        global i
        i += 1
        Lock.release()

def thread2():
    for k in range(1000001):
        Lock.acquire()
        global i
        i -= 1
        Lock.release()

def main():
    t1 = threading.Thread(target = thread1, args = (),)
    t2 = threading.Thread(target = thread2, args = (),)
    t1.start()
    t2.start()
    t1.join()
    t2.join()
    print(i)


main()
