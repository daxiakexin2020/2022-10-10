import os
import mysql.connector


def PrintInfo():
    print(os.getcwd())

def TestMysql():
    tm=mysql.connector.connect()

PrintInfo()
