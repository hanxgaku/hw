def factorial(x):
    n = 1
    while x > 0:
        n = n * x
        x = x - 1
    return n

def factorial_rec(x):
    if x <= 0:
        return 1
    else:
        return factorial_rec(x - 1) * x   # non-tail recursion

def factorial_rec2(x, n):   # tail recursion,  TCO - Scheme, Haskell
    if x > 0:
        return factorial_rec2(x - 1, n * x)
    else: 
        return n
        
def is_sorted_array(a):
    for i in range(len(a) - 1):
        if a[i] > a[i + 1]:
            return False
    return True
    
def is_sorted_array_rec(a):
    if len(a) == 1 or len(a) == 0:
        return True
    elif a[0] > a[1]:
        return False
    else:
        return is_sorted_array_rec(a[1:])

print(is_sorted_array_rec([]))


#1!=1
#2!=2
#3!=1*2*3=6
#4!=1*2*3*4=24=3!*4
#n!=n*(n-1)!