class ADD1:
    def __init__(self):
        pass

    def plusOne(self, digits):
        n = len(digits)
        for frontIndex in range(n):
            behindIndex = n - frontIndex - 1
            print(frontIndex, behindIndex)
            if digits[behindIndex] != 9:
                digits[behindIndex] += 1
                j = behindIndex + 1
                while j < n:
                    digits[j] = 0
                    j += 1
                return digits

        ret = []
        for i in range(n + 1):
            if i == 0:
                ret.append(1)
            else:
                ret.append(0)
        return ret


ad = ADD1()
digits = [9, 9, 9, 9, 9]
one = ad.plusOne(digits)
print("one:{0}".format(one))
