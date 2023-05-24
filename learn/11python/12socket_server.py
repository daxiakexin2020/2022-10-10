import socket


class SocketDemo:
    def __init__(self):
        pass

    def Demo01(self):

        serversocket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

        # hostname = socket.gethostname()  # 本地主机名
        hostname = "0.0.0.0"
        port = 9999

        serversocket.bind((hostname, port))

        serversocket.listen(5)

        while True:
            conn, addr = serversocket.accept()
            print("accept:{0},addr:{1}".format(conn, addr))
            self.handleConn(conn)

    def handleConn(self, conn):
        msg = "welcome cai " + "\r\n"
        conn.send(msg.encode('utf-8'))
        conn.close()


sd = SocketDemo()
sd.Demo01()
