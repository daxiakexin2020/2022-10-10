import socket


class SocketClient:
    def __init__(self):
        pass

    def Demo01(self):
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        # hostname = socket.gethostname()
        hostname = "0.0.0.0"
        port = 9999
        s.connect((hostname, port))
        msg = s.recv(1024)
        s.close()
        print(msg.decode('utf-8'))


sc = SocketClient()
sc.Demo01()
