import copy


board = [["_","_","_"],["_","_","_"],["_","_","_"]]
#board = [["X","O","X"],["O","O","X"],["_","_","_"]]

# [["X","X","X"],["_","_","_"],["_","_","_"]]
# [["X","_","_"],["_","_","_"],["_","_","_"]]
# [["X","X","X"],["_","_","_"],["_","_","_"]]

def printboard(m : [[str]]) -> None:
    for i in range(len(m)):
        for j in range(len(m)):
            if j == len(m) - 1:
                print(m[i][j])
            else:
                print(m[i][j], end=' ')

def printboard_num(m : [[str]]) -> None:
    k = 1
    for i in range(len(m)):
        for j in range(len(m)):
            if m[i][j] == "_":
                slot = k
                k = k + 1
            else:
                slot = m[i][j]
            if j == len(m) - 1:
                print(slot)
            else:
                print(slot, end=' ')
                
def same_elements_array(a : [str]) -> bool:
    for i in range(len(a) - 1):
        if (a[i] != a[i + 1]):
            return False
    return True

def get_column(m : [[str]], n : int) -> [str]:
    c = []
    for i in range(len(m)):
        c = c + [m[i][n]]
    return c
    
def get_diagonal_l(m : [[str]]) -> [str]:
    d = []
    for i in range(len(m)):
        d = d + [m[i][i]]
    return d
        
def get_diagonal_r(m : [[str]]) -> [str]:
    d = []
    for i in range(len(m)):
        d = d + [m[i][len(m) - i - 1]]
    return d
    
def is_filled(m : [[str]]) -> bool:
    for i in range(len(m)):
        for j in range(len(m)):
            if m[i][j] == "_":
                return False
    return True
    
def get_possible_turns(m : [[str]], s : str) -> [[[str]]]:
    p = []
    for i in range(len(m)):
        for j in range(len(m)):
            if m[i][j] == "_":
                n = copy.deepcopy(m)
                n[i][j] = s
                p = p + [n]
    return p
    
def print_vertical(ms : [[[str]]]) -> None:
    for i in range(len(ms)):
        printboard(ms[i])
        print(" ")
        
def print_horizontal(ms : [[[str]]]) -> None:
    for k in range(len(ms[0])):
        for i in range(len(ms)):
            for j in range(len(ms[i])):
                if i == len(ms) - 1 and j == len(ms[i]) - 1:
                    print(ms[i][k][j])
                elif j == len(ms[i]) - 1:
                    print(ms[i][k][j], end='  ')
                else:
                    print(ms[i][k][j], end=' ')
                    
def insert(m : [[str]], s : str, n : int) -> [[str]]:
    k = 1
    for i in range(len(m)):
       for j in range(len(m)):
           if m[i][j] == "_":
                if k == n:
                    m[i][j] = s
                    return m
                else:
                    k = k + 1
    return m

# Печать в виде досок массива досок возможных ходов по горизонтали и по вертикали
# Печать доски с нумерованием ячеек
# Интерфейс для игры
# minimax алгоритм

def flip_turn(s : str) -> str:
    if s == "X":
        return "O"
    else:
        return "X"

def count_boards(m : [[str]], s : str, r : int) -> int:
    print("recursion level " + str(r))
    if win(m) == "X" or win(m) == "O" or is_filled(m):
        #print("leaf")
        #printboard(m)
        return 1
    c = 1
    s_next = flip_turn(s)
    boards = get_possible_turns(m, s)
    #print_horizontal(boards)
    for i in range(len(boards)):
        c = c + count_boards(boards[i], s_next, r + 1)
    return c

def win(m : [[str]]) -> str:
    for i in range(len(m)):
        if same_elements_array(m[i]):
            return m[i][0]
        elif same_elements_array(get_column(m, i)):
            return m[0][i]
    if same_elements_array(get_diagonal_l(m)):
        return m[0][0]
    elif same_elements_array(get_diagonal_r(m)):
        return m[0][len(m) - 1]
    else:
        return "_"
    
def min_element(a : [int]) -> int:
    k = a[0]
    for i in range(len(a)):
        if a[i] < k:
           k = a[i]
    return k
    

def max_element(a : [int]) -> int:
    k = a[0]
    for i in range(len(a)):
        if a[i] > k:
           k = a[i]
    return k
    
def get_turn_num(a : [int]) -> int:
    if len(a) == 0:
        return 0
    n = a[0]
    k = 0
    for i in range(len(a)):
        if a[i] > n:
           n = a[i]
           k = i
    return k

def get_optimal_turn(m : [[str]]) -> [[str]]:
    boards = get_possible_turns(m, "O")
    res = []
    for i in range(len(boards)):
        res = res + [minimax(boards[i], "X")]
    max_k = get_turn_num(res)
    return boards[max_k]

# computer - O
def minimax(m : [[str]], s : str) -> int:
    if win(m) == "X":
        return -1
    elif win(m) == "O":
        return 1
    elif is_filled(m):
        return 0
    res = []
    boards = get_possible_turns(m, s)
    for i in range(len(boards)):
        res = res + [minimax(boards[i], flip_turn(s))]
    if s == "X":
        return min_element(res)
    else:
        return max_element(res)
                
def is_game_over(m : [[str]]) -> bool:
    if is_filled(m) and win(m) != "X" and win(m) != "O":
        print("DRAW")
        return True
    elif win(m) == "X":
        print("PLAYER WIN")
        return True
    elif win(m) == "O":
        print("COMPUTER WIN")
        return True
    else:
        return False
        
def game(m : [[str]]) -> None:
    while True: 
        printboard_num(m)
        if is_game_over(m):
            break
        val = int(input("Enter value "))
        m = insert(m, "X", val)
        printboard(m)
        if is_game_over(m):
            break
        m = get_optimal_turn(m)
        
        
#printboard(board)

#print(same_elements_array(["X","X","X"]))

#print(get_column(board, 0))

#print(print_horizontal(get_possible_turns(board, "X")))

#print(insert(board, "O", 3))
#print(count_boards(board, "X", 0))
#printboard(get_optimal_turn(board))
#printboard(insert(board, "X", 1))
game(board)