from abc import abstractmethod, ABC
from typing import Iterator


class User:
    def __init__(self, name, age) -> None:
        self.name = name
        self.age = age


class DBAccess(ABC):
    @abstractmethod
    def getUserInfo(self) -> dict:
        pass


class SQLiteDB(DBAccess):
    def getUserInfo(self) -> dict:
        print("Get User Info From SQLite.....")
        return [{"name": "zz", "age": 33}, {"name": "kx", "age": 32}]


class MongoDB(DBAccess):
    def getUserInfo(self) -> dict:
        print("Get User Info From MongoDB.....")
        return [{"name": "xiaojiahuo", "age": 60}, {"name": "ljh", "age": 65}]


class UserRepository(ABC):
    @abstractmethod
    def getUserInfo(self) -> Iterator[User]:
        pass


class DBUserRepository(UserRepository):
    def __init__(self, db: DBAccess) -> None:
        self.db = db

    def getUserInfo(self) -> Iterator[User]:
        for entry in self.db.getUserInfo():
            user = User(**entry)
            yield user


class UserService:
    def __init__(self, user_repositiry: UserRepository) -> None:
        self.user_repository = user_repositiry

    def getUserInfo(self):
        return [user for user in self.user_repository.getUserInfo()]


def main():
    db: DBAccess = SQLiteDB()
    db: DBAccess = MongoDB()
    user_repositorry = DBUserRepository(db)
    user_service = UserService(user_repositorry)
    users = user_service.getUserInfo()
    for user in users:
        print(user.name, user.age)


if __name__ == "__main__":
    main()
