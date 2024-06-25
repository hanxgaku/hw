from ctypes import c_int
from typing import Any

#a=[1,2,3]
#a.append(42)

#[1,2,3,42]


#      1 2 3 . . . . . 
#_ _ _ _ _ _ _ _ _ _ _ _ ...
#                    ^

# __iter__ + yield
# fix __repr__
class IntegerArray: 
    def __init__(self):
        self.array = ((c_int)*10)()
        self.length = 0
        self.capacity = 10
    
    def __getitem__(self, n : int) -> int:
        if n >= self.length:
            raise IndexError("list index out of bounds")
        else:
            return self.array[n]
            
    def __delitem__(self, n : int) -> None:
        del self.array[n]
        self.length = self.length - 1
        return
    
    def __setitem__(self, id : int, n : int) -> None:
        if id >= self.length:
            raise IndexError("list index out of bounds")
        else:
            self.array[id] = n
        return
    
    def __str__(self) -> str:
        s = "["
        for i in range(self.length):
            if i != self.length - 1:
                s = s + str(self.array[i]) + " "
            else:
                s = s + str(self.array[i])
        return s + "]"
    
    def __repr__(self) -> str:
        s = "slice(["
        for i in range(self.length):
            if i != self.length - 1:
                s = s + str(self.array[i]) + " "
            else:
                s = s + str(self.array[i])
        return s + "], length=" + str(self.length) + ", capacity=" + str(self.capacity) + ")"
    
    def __format__(self, spec: str) -> str:
        s = "["
        for i in range(self.length):
            if i != self.length - 1:
                s = s + str(self.array[i]) + " "
            else:
                s = s + str(self.array[i])
        return s + "]"
    
    def __doc__(self) -> str:
        return "array of integers object"
    
    def __eq__(self, other: object) -> bool:
        if isinstance(other, IntegerArray):
            if self.length != other.length and self.capacity != other.capacity:
                return False
            else:
                for i in range(self.length):
                    if self.array[i] != other.array[i]:
                        return False
        return True

    def __add__(self, other: object) -> object:
        if isinstance(other, IntegerArray):
            new_arr = IntegerArray()
            for i in range(self.length):
                new_arr.append(self[i])
            for i in range(other.length):
                new_arr.append(other[i])
        return new_arr
    
    def __mul__(self, other: int) -> object:
        temp_arr = IntegerArray()
        while other > 0:
            for i in range(self.length):
                temp_arr.append(self.array[i])
            other = other - 1
        return temp_arr
    
    
    def __iadd__(self, other: object) -> object:
        if isinstance(other, IntegerArray):
            for i in range(other.length):
                self.append(other[i])
        return self
    
    def __imul__(self, other: int) -> object:
        temp_arr = IntegerArray()
        while other > 0:
            for i in range(self.length):
                temp_arr.append(self.array[i])
            other = other - 1
        self = temp_arr
        return self
    
    def __len__(self) -> int:
        return self.length
    
    def __reversed__(self) -> object:
        temp_arr = IntegerArray()
        for i in range(self.length):
            temp_arr.append(self.array[self.length - i - 1])
        return temp_arr
    
    def append(self, n : int) -> None:
        if self.length < self.capacity:
            self.array[self.length] = n
        else:
            self.capacity = self.capacity * 10
            temp_arr = self.array
            self.array = ((c_int)*self.capacity)()
            for i in range(self.length):
                self.array[i] = temp_arr[i]
            self.array[self.length] = n
        self.length = self.length + 1
        return
    
        
a = IntegerArray()
a.append(0)
a.append(0)
a.append(0)
for i in range(10):
    a.append(i)

b = IntegerArray()
b.append(100)
b.append(101)
    
print(a[0])
print(a[1])
print(a[2])
print(a)

print(a + b)

print(b.__reversed__())