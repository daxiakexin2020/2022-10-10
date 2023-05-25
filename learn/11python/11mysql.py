import mysql.connector


class MysqlDemo:

    def __init__(self):
        mydb = mysql.connector.connect(
            host="127.0.0.1",
            user="root",
            passwd="123",
            database="test",
            auth_plugin="mysql_native_password"  # mysql8.0有坑，验证方式改变了
        )
        self.db = mydb

    def Demo1(self):
        mycursor = self.db.cursor()
        execute = mycursor.execute("SHOW DATABASES")
        for x in mycursor:
            print(x)

    def Demo02(self):
        cursor = self.db.cursor()

        createSql = "CREATE TABLE sites (name VARCHAR(255), url VARCHAR(255))"
        cursor.execute(createSql)

        setPriKeySql = "ALTER TABLE sites ADD COLUMN id INT AUTO_INCREMENT PRIMARY KEY"
        cursor.execute(setPriKeySql)

    def Demo03(self):
        cursor = self.db.cursor()
        insertSql = "insert into sites (name,url) VALUES (%s,%s)"
        val = ('zhang-san', 'http://baidu.com')
        execute = cursor.execute(insertSql, val)
        self.db.commit()  # 有数据更新，必须使用该方法
        print("insert result:{0}".format(execute))


md = MysqlDemo()
# md.Demo1()
# md.Demo02()
md.Demo03()
