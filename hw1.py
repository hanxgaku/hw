def reverse(a):
    for i in range(len(a) // 2):
        j = len(a) - i - 1
        n = a[j]
        a[j] = a[i]
        a[i] = n
    return a

def reverse_rec(a):
    if len(a) == 0:
        return []
    else:
        return reverse_rec(a[1:]) + [a[0]]
    
def join_arr(a, b):
    x = []
    i = 0
    j = 0
    while (i <= len(a)) and (j <= len(b)):
        if i == len(a):
            x = x + b[j:]
            break
        elif j == len(b):
            x = x + a[i:]
            break
        elif (a[i] < b[j]):
            x = x + [a[i]]
            i = i + 1
            #print(x)
        else:
            x = x + [b[j]]
            j = j + 1
            #print(x)
    return x

#4 1 8 3 2 9 0
#[4 1 8]  [3 2 9 0]
#[4] [1 8] [3 2] [9 0]
#[4] [1] [8] [3] [2] [9] [0]

#3 2 9 0

#[42]  -- [] [42] -- [] [42] --...

def msort(a : [int]) -> [int]:
    if not len(a):
        return []
    if len(a) == 1:
        return [a[0]]
    k = len(a) // 2
    x = a[:k]
    y = a[k:]
    return join_arr(msort(x), msort(y))

#print(msort([4,1,8,3,2,9,0]))


#1 2 3 5 7
#3 5 2 7 1
def bsort(a : [int]) -> [int]:
    k = 0
    for i in range(len(a) - 1):
        if a[i] > a[i + 1]:
            n = a[i]
            a[i] = a[i + 1]
            a[i + 1] = n
            k = k + 1
    if k == 0:
        return a
    else:
        return bsort(a)

def bsort_while(a: [int]) -> [int]:
    while True:
        k = 0
        for i in range(len(a) - 1):
            if a[i] > a[i + 1]:
                n = a[i]
                a[i] = a[i + 1]
                a[i + 1] = n
                k = k + 1
        if k == 0:
           return a
         

def bsort_while_1(a: [int]) -> [int]:
    while True:
        k = 0
        i = 0
        while i < len(a) - 1:
            if a[i] > a[i + 1]:
                n = a[i]
                a[i] = a[i + 1]
                a[i + 1] = n
                k = k + 1
                if i > 0:
                    i = i - 2
            i = i + 1
            print(a)
        if k == 0:
           return a
         
               
#3 6 1 8 4
#3 6 1 8 4
#3 1 6 8 4
#1 3 6 8 4
#1 3 6 8 4
#1 3 6 8 4
#1 3 6 4 8
#1 3 4 6 8
#1 3 4 6 8

#1) сортровка с отступанием
#2) вспомнить git, рег на гитхаб
#3) перевод на Java



#print(bsort_while_1([3,5,2,7,1]))

def join_arr_rec(a, b):
    if len(a) == 0:
        return b
    elif len(b) == 0:
        return a
    elif a[0] < b[0]:
        return [a[0]] + join_arr_rec(a[1:], b)
    else:
        return [b[0]] + join_arr_rec(a, b[1:])
            
def pow(n, m):
    x = 1
    for i in range(m):
        x = x * n
    return x
    
def pow_rec(n, m):
    if m == 0:
        return 1
    else:
        return n * pow(n, m - 1)

#print(join_arr([1,2,4],[3,4,5]))
        
        
        
#5 3 7 2 9 6 1
#3 2 1 - 1 2 3 
#7 9 6 - 6 7 9
#1 2 3 5 6 7 9

#[4 1 2 3]
#[1 2 3]
#[]

def qsort(a : [int]) -> [int]:
    if not len(a):
        return []
    if len(a) == 1:
        return [a[0]]
    k = a[0]
    x = []
    y = []
    for i in range(1, len(a)):
        if a[i] < k:
            x += [a[i]]
        else:
            y += [a[i]]
    return qsort(x) + [k] + qsort(y)





#print(qsort([5,3,7,2,9,6,1]))



#"abc" -> "abc","acb", "cab", "cba", "bac", "bca"
#"bc" -> ["bc", "cb"] -> ["abc", "acb"]
#"ac" -> ["ac", "ca"] -> ["bac", "bca"]
#"ab" -> ["ab", "ba"] -> ["cab", "cba"]

def permutations(s : str, d: int) -> [str]:
    print(s, d)
    if not s:
        return [""]
    x = []
    for i in range(len(s)):
         p = permutations(s[:i] + s[i + 1:], d + 1)
         x = x + [s[i] + j for j in p]
    print(x, d)
    return x
    
#print(permutations("abc", 0))

def permutations2(s: str) -> [str]:
    if not s:
        return [""]
    elif len(s) == 1:
        return [s]
    x = []
    p = []
    p = permutations2(s[1:])
    x = x + [s[0] + "|" + j for j in p]
    x = x + [s[0] + j for j in p]
    return x
    
#abcd -> abc|d, ab|cd, a|bcd, a|bc|d, ab|c|d, a|b|cd, a|b|c|d
#bcd -> b|cd, bc|d, b|c|d

print(permutations2("abcd"))