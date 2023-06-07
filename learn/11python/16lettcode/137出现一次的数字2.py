class Solution(object):
    def singleNumber(self, nums):
        flag = {}
        for num in nums:
            if num in flag:
                flag[num] += 1
            else:
                flag[num] = 1
        for k, v in flag.items():
            if v == 1:
                return k
        return 0

    def Start(self, data):
        number = self.singleNumber(data)
        print("number:{0}".format(number))

s = Solution()
nums = [1, 1, 1, 2]
s.Start(nums)
