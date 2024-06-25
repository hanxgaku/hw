import copy

# bind to interface with moving pieces
# read about en passant
class Board:
    def __init__(self):
        self.board = [["brook", "bknight", "bbishop", "bqueen", "bking", "bbishop", "bknight", "brook"],
        ["bpawn","bpawn","bpawn","bpawn","bpawn","bpawn","bpawn","bpawn"],
        ["_","_","_","_","_","_","_","_"],
        ["_","_","_","_","_","_","_","_"],
        ["_","_","_","_","_","_","_","_"],
        ["_","_","_","_","_","_","_","_"],
        ["wpawn","wpawn","wpawn","wpawn","wpawn","wpawn","wpawn","wpawn"],
        ["wrook", "wknight", "wbishop", "wqueen", "wking", "wbishop", "wknight", "wrook"]]
         
        # initialize threat map
        self.is_wl_castling_available = True
        self.is_wr_castling_available = True
        self.is_bl_castling_available = True
        self.is_br_castling_available = True
        
        # store last move
        self.last_move = (0, 0, 0, 0)

    def __len__(self) -> int:
        return len(self.board)
    
    def __getitem__(self, c : (int, int)) -> str:
        return self.board[c[1]][c[0]]
        
    def __setitem__(self, c : (int, int), s : str) -> None:
        self.board[c[1]][c[0]] = s
    
    def make_move(self, c : (int, int, int, int)) -> "Board":
        new_Board = Board()
        temp = copy.deepcopy(self.board)
        is_wl_castling_available = self.is_wl_castling_available
        is_wr_castling_available = self.is_wr_castling_available
        is_bl_castling_available = self.is_bl_castling_available
        is_br_castling_available = self.is_br_castling_available
        if self.board[c[1]][c[0]] in {"bking", "wking"} and abs(c[2] - c[0]) > 1 and abs(c[3] - c[1]) > 1:
            king = temp[c[1]][c[0]]
            rook = temp[c[3]][c[2]]
            temp[c[3]][c[2]] = king
            temp[c[1]][c[0]] = rook
            if c[3] == 0 and c[2] == 0:
                is_bl_castling_available = False
            elif c[3] == 0 and c[2] == 7:
                is_br_castling_available = False
            elif c[3] == 7 and c[2] == 0:
                is_wl_castling_available = False
            elif c[3] == 7 and c[2] == 7:
                is_wr_castling_available = False
        elif self.board[c[1]][c[0]] == "brook":
            if c[1] == 0 and c[0] == 0:
                is_bl_castling_available = False
            elif c[1] == 0 and c[0] == 7:
                is_br_castling_available = False
            piece = temp[c[1]][c[0]]
            temp[c[3]][c[2]] = piece
            temp[c[1]][c[0]] = "_"
        elif self.board[c[1]][c[0]] == "bking" and c[1] == 0 and c[0] == 4:
            is_bl_castling_available = False
            is_br_castling_available = False
            
            piece = temp[c[1]][c[0]]
            temp[c[3]][c[2]] = piece
            temp[c[1]][c[0]] = "_"
        elif self.board[c[1]][c[0]] in {"bpawn", "wpawn"} and self.board[c[2]][c[3]] == "_" and abs(c[2] - c[0]) == 1 and abs(c[3] - c[1]) == 1:
            piece = temp[c[1]][c[0]]
            temp[c[3]][c[2]] = piece
            temp[c[1]][c[0]] = "_"
            temp[c[1]][c[2]] = "_"
        else:
            piece = temp[c[1]][c[0]]
            temp[c[3]][c[2]] = piece
            temp[c[1]][c[0]] = "_"
        
        new_Board.board = temp
        new_Board.is_wl_castling_available = is_wl_castling_available
        new_Board.is_wr_castling_available = is_wr_castling_available
        new_Board.is_bl_castling_available = is_bl_castling_available
        new_Board.is_br_castling_available = is_br_castling_available
        new_Board.last_move = (c[0], c[1], c[2], c[3])
        return new_Board
    

threat_map = [[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],
[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0]]

# Common Lisp
# CLOS
# Clojure
#x.f() ad hoc полиморфизм
#f(x,y,z)

pieces = {
    "brook" : "♜",
    "bknight" : "♞",
    "bbishop" : "♝",
    "bqueen" : "♛",
    "bking" : "♚",
    "bpawn" : "♟",
    "wrook" : "♖",
    "wknight" : "♘",
    "wbishop" : "♗",
    "wqueen" : "♕",
    "wking" : "♔",
    "wpawn" : "♙",
    "_" : "_"
}


# print chess board ascii
def print_board(board : Board) -> None:
    for i in range(len(board)):
        for j in range(len(board)):
            if j == len(board) - 1:
                print(pieces[board[j, i]])
            else:
                print(pieces[board[j, i]], end=' ')

# check that coordinates in array are not out of bounds
def check_coordinates(c_arr : [[int, int]]) -> [[int, int]]:
    new_arr = []
    for i in range(len(c_arr)):
        if c_arr[i][0] >= 0 and c_arr[i][1] >= 0 and c_arr[i][0] < 8 and c_arr[i][1] < 8:
            new_arr = new_arr + [[c_arr[i][0], c_arr[i][1]]]
    return new_arr

def check_coordinates2(c_arr : [(int, int, int, int)]) -> [(int, int, int, int)]:
    new_arr = []
    for i in range(len(c_arr)):
        if c_arr[i][3] >= 0 and c_arr[i][2] >= 0 and c_arr[i][3] < 8 and c_arr[i][2] < 8:
            new_arr = new_arr + [(c_arr[i][0], c_arr[i][1], c_arr[i][2], c_arr[i][3])]
    return new_arr

# functions to fill threat map with 1s as potential threats
# combine into one function
def fill_threat_map_line(board : Board, tm : [[int]], x : int, y : int, dx : int, dy : int) -> [[int]]:
    x_init = x
    y_init = y
    while x >= 0 and x < 8 and y >= 0 and y < 8:
        if board[x, y] != "_" and x != x_init and y != y_init and get_cell_color(board, x, y) == get_cell_color(board, x_init, y_init):
            return tm
        elif board[x, y] != "_" and x != x_init and y != y_init and get_cell_color(board, x, y) != get_cell_color(board, x_init, y_init):
            tm[y][x] = 1
            return tm
        elif x != x_init or y != y_init:
            tm[y][x] = 1
        x = x + dx
        y = y + dy
    return tm

# functions to fill array with coordinates to use for potential piece moves
def fill_coordinates_line(board : Board, x : int, y : int, dx : int, dy : int) -> [(int, int, int, int)]:
    c_arr = []
    x_init = x
    y_init = y
    x = x + dx
    y = y + dy
    while x >= 0 and x < 8 and y >= 0 and y < 8:
        if board[x, y] != "_" and get_cell_color(board, x, y) == get_cell_color(board, x_init, y_init):
            break
        elif board[x, y] != "_" and get_cell_color(board, x, y) != get_cell_color(board, x_init, y_init):
            c_arr += [(x_init, y_init, x, y)]
            break
        else:
            c_arr += [(x_init, y_init, x, y)]
        x = x + dx
        y = y + dy
    return c_arr

# threats and moves for pieces
def king_threats(tm : [[int]], x : int, y : int) -> [[int]]:
    c_arr = check_coordinates([[y + 1, x], [y + 1, x + 1], [y + 1, x - 1], [y, x + 1], [y, x - 1], [y - 1, x], [y - 1, x + 1], [y - 1, x - 1]])
    for i in range(len(c_arr)):
        tm[c_arr[i][0]][c_arr[i][1]] = 1
    return tm
    
def king_moves(board : Board, x : int, y : int) -> [[int, int, int, int]]:
    c_arr = []
    t_arr = check_coordinates2([[x, y, x, y + 1], [x, y, x + 1, y + 1], [x, y, x - 1, y + 1], [x, y, x + 1, y], [x, y, x - 1, y], [x, y, x, y - 1], [x, y, x + 1, y - 1], [x, y, x - 1, y - 1]])
    for i in range(len(t_arr)):
        if get_cell_color(board, t_arr[i][2], t_arr[i][3]) != get_cell_color(board, x, y):
            c_arr += [t_arr[i]]
    return c_arr
    
def queen_threats(board : Board, tm : [[int]], x : int, y : int) -> [[int]]:
    tm = fill_threat_map_line(board, tm, x, y, 1, 0)
    tm = fill_threat_map_line(board, tm, x, y, 1, 1)
    tm = fill_threat_map_line(board, tm, x, y, 1, -1)
    tm = fill_threat_map_line(board, tm, x, y, 0, 1)
    tm = fill_threat_map_line(board, tm, x, y, 0, -1)
    tm = fill_threat_map_line(board, tm, x, y, -1, 0)
    tm = fill_threat_map_line(board, tm, x, y, -1, 1)
    tm = fill_threat_map_line(board, tm, x, y, -1, -1)
    return tm

def queen_moves(board : Board, x : int, y : int) -> [(int, int, int, int)]:
    c_arr = []
    c_arr += fill_coordinates_line(board, x, y, 1, 0)
    c_arr += fill_coordinates_line(board, x, y, 1, 1)
    c_arr += fill_coordinates_line(board, x, y, 1, -1)
    c_arr += fill_coordinates_line(board, x, y, 0, 1)
    c_arr += fill_coordinates_line(board, x, y, 0, -1)
    c_arr += fill_coordinates_line(board, x, y, -1, 0)
    c_arr += fill_coordinates_line(board, x, y, -1, 1)
    c_arr += fill_coordinates_line(board, x, y, -1, -1)
    return c_arr
    
def wpawn_threats(tm : [[int]], x : int, y : int) -> [[int]]:
    c_arr = check_coordinates([[y - 1, x - 1], [y - 1, x + 1]])
    for i in range(len(c_arr)):
        tm[c_arr[i][0]][c_arr[i][1]] = 1
    return tm

def wpawn_moves(board : Board, x : int, y : int) -> [(int, int, int, int)]:
    c_arr = []
    t_arr = check_coordinates2([[x, y, x - 1, y - 1], [x, y, x + 1, y - 1]])
    for i in range(len(t_arr)):
        if get_cell_color(board, t_arr[i][2], t_arr[i][3]) == flip_color(get_cell_color(board, x, y)):
            c_arr += [t_arr[i]]
    if y == 6 and get_cell_color(board, x, 5) == "n" and get_cell_color(board, x, 4) == "n":
        c_arr += [(x, y, x, 4)]
    if get_cell_color(board, x, y - 1) == "n":
        c_arr += [(x, y, x, y - 1)]
    if x > 0 and x < 7 and y == 3:
        if board[x + 1, y] == "bpawn" and board.last_move == (x + 1, 1, x + 1, 3):
            c_arr += [(x, 3, x + 1, 2)]
        if board[x - 1, y] == "bpawn" and board.last_move == (x - 1, 1, x - 1, 3):
            c_arr += [(x, 3, x - 1, 2)]
    return c_arr

def bpawn_threats(tm : [[int]], x : int, y : int) -> [[int]]:
    c_arr = check_coordinates([[y + 1, x - 1], [y + 1, x + 1]])
    for i in range(len(c_arr)):
        tm[c_arr[i][0]][c_arr[i][1]] = 1
    return tm

def bpawn_moves(board : Board, x : int, y : int) -> [(int, int, int, int)]:
    c_arr = []
    t_arr = check_coordinates2([[x, y, x - 1, y + 1], [x, y, x + 1, y + 1]])
    for i in range(len(t_arr)):
        if get_cell_color(board, t_arr[i][2], t_arr[i][3]) == flip_color(get_cell_color(board, x, y)):
            c_arr += [t_arr[i]]
    if y == 1 and get_cell_color(board, x, 2) == "n" and get_cell_color(board, x, 3) == "n":
        c_arr += [(x, y, x, 3)]
    if y < 7 and get_cell_color(board, x, y + 1) == "n":
            c_arr += [(x, y, x, y + 1)]
    if x > 0 and x < 7 and y == 4:
        if board[x + 1, y] == "wpawn" and board.last_move == (x + 1, 6, x + 1, 4):
            c_arr += [(x, 4, x + 1, 5)]
        if board[x - 1, y] == "wpawn" and board.last_move == (x - 1, 6, x - 1, 4):
            c_arr += [(x, 4, x - 1, 5)]
    return c_arr
    
def knight_threats(tm : [[int]], x : int, y : int) -> [[int]]:
    c_arr = check_coordinates([[y - 2, x - 1], [y - 2, x + 1],[y - 1, x - 2],[y - 1, x + 2],[y + 1, x - 2],[y + 1, x + 2],[y + 2, x - 1],[y + 2, x + 1]])
    for i in range(len(c_arr)):
        tm[c_arr[i][0]][c_arr[i][1]] = 1
    return tm

def knight_moves(board : Board, x : int, y : int) -> [[int, int, int, int]]:
    new_arr = []
    c_arr = check_coordinates2([[x, y, x - 1, y - 2], [x, y, x + 1, y - 2],[x, y, x - 2, y - 1],[x, y, x + 2, y - 1],[x, y, x - 2, y + 1],[x, y, x + 2, y + 1],[x, y, x - 1, y + 2],[x, y, x + 1, y + 2]])
    for i in range(len(c_arr)):
        if get_cell_color(board, c_arr[i][2], c_arr[i][3]) != get_cell_color(board, x, y):
            new_arr += [c_arr[i]]
    return new_arr

def bishop_threats(board : Board, tm : [[int]], x : int, y : int) -> [[int]]:
    tm = fill_threat_map_line(board, tm, x, y, 1, -1)
    tm = fill_threat_map_line(board, tm, x, y, 1, 1)
    tm = fill_threat_map_line(board, tm, x, y, -1, -1)
    tm = fill_threat_map_line(board, tm, x, y, -1, 1)
    return tm

def bishop_moves(board : Board, x : int, y : int) -> [(int, int, int, int)]:
    c_arr = []
    c_arr += fill_coordinates_line(board, x, y, 1, -1)
    c_arr += fill_coordinates_line(board, x, y, 1, 1)
    c_arr += fill_coordinates_line(board, x, y, -1, -1)
    c_arr += fill_coordinates_line(board, x, y, -1, 1)
    return c_arr
    
def rook_threats(board : Board, tm : [[int]], x : int, y : int) -> [[int]]:
    tm = fill_threat_map_line(board, tm, x, y, 1, 0)
    tm = fill_threat_map_line(board, tm, x, y, 0, 1)
    tm = fill_threat_map_line(board, tm, x, y, -1, 0)
    tm = fill_threat_map_line(board, tm, x, y, 0, -1)
    return tm

def rook_moves(board : Board, x : int, y : int) -> [(int, int, int, int)]:
    c_arr = []
    c_arr += fill_coordinates_line(board, x, y, 1, 0)
    c_arr += fill_coordinates_line(board, x, y, 0, 1)
    c_arr += fill_coordinates_line(board, x, y, -1, 0)
    c_arr += fill_coordinates_line(board, x, y, 0, -1)
    return c_arr
    
def get_cell_color(board : Board, x : int, y : int) -> str:
    if board[x, y] in {"brook", "bknight", "bbishop", "bking", "bqueen", "bpawn"}:
        return "b"
    elif board[x, y] in {"wrook", "wknight", "wbishop", "wking", "wqueen", "wpawn"}:
        return "w"
    else:
        return "n"
        
def flip_color(color : str) -> str:
    colors = {
        "w": "b",
        "b": "w"
    }
    return colors.get(color, "n")
        
def get_king_coordinates(board : Board, color : str) -> (int, int): 
    for i in range(len(board)):
        for j in range(len(board)):
            if get_cell_color(board, j, i) == color and board[j, i] in {"wking", "bking"}:
                return (j, i)
                
def is_check(board : Board, color : str) -> bool:
    x, y = get_king_coordinates(board, color)
    threat_map = [[0]*8 for i in range(8)]
    threat_map = calculate_threats(board, threat_map, flip_color(color))
    return threat_map[y][x] == 1
    
def calculate_threats(board : Board, tm : [[int]], color : str) -> [[int]]:
    for i in range(len(board)):
        for j in range(len(board)):
            if get_cell_color(board, j, i) == color:
                if board[j, i] == "brook" or board[j, i] == "wrook":
                    tm = rook_threats(board, tm, j, i)
                elif board[j, i] == "bknight" or board[j, i] == "wknight":
                    tm = knight_threats(tm, j, i)
                elif board[j, i] == "bbishop" or board[j, i] == "wbishop":
                    tm = bishop_threats(board, tm, j, i)
                elif board[j, i] == "bking" or board[j, i] == "wking":
                    tm = king_threats(tm, j, i)
                elif board[j, i] == "bqueen" or board[j, i] == "wqueen":
                    tm = queen_threats(board, tm, j, i)
                elif board[j, i] == "bpawn":
                    tm = bpawn_threats(tm, j, i) 
                elif board[j, i] == "wpawn":
                    tm = wpawn_threats(tm, j, i)
    return tm
                
def print_threat_map(tm : [[int]]) -> None:
    for i in range(len(tm)):
        for j in range(len(tm)):
            if j == len(tm) - 1:
                print(tm[i][j])
            else:
                print(tm[i][j], end=' ')
                
def filter_moves(board : Board, color : str, coordinates : [(int, int, int, int)]) -> [(int, int, int, int)]:
    c_arr = []
    for i in range(len(coordinates)):
        new_board = board.make_move(coordinates[i])
        if not is_check(new_board, color):
            c_arr += [coordinates[i]]
    return c_arr
    
def get_possible_castlings(board : Board, color : str) -> [(int, int, int, int)]:
    c_arr = []
    if color == "w":
        if board.is_wl_castling_available and board[7, 1] == board[7, 2] == board[7, 3] == "_":
            c_arr += [(7, 4, 7, 0)]
        if board.is_wr_castling_available and board[7, 5] == board[7, 6] == "_":
            c_arr += [(7, 4, 7, 7)]
    elif color == "b":
        if board.is_bl_castling_available and board[0, 1] == board[0, 2] == board[0, 3] == "_":
            c_arr += [(0, 4, 0, 0)]
        if board.is_br_castling_available and board[0, 5] == board[0, 6] == "_":
            c_arr += [(0, 4, 0, 7)]
    return c_arr

def get_possible_moves(board : Board, color : str) -> [(int, int, int, int)]:
    c_arr = []
    for i in range(len(board)):
        for j in range(len(board)):
            if get_cell_color(board, j, i) == color:
                if board[j, i] == "brook" or board[j, i] == "wrook":
                    c_arr += rook_moves(board, j, i)
                elif board[j, i] == "bknight" or board[j, i] == "wknight":
                    c_arr += knight_moves(board, j, i)
                elif board[j, i] == "bbishop" or board[j, i] == "wbishop":
                    c_arr += bishop_moves(board, j, i)
                elif board[j, i] == "bking" or board[j, i] == "wking":
                    c_arr += king_moves(board, j, i)
                elif board[j, i] == "bqueen" or board[j, i] == "wqueen":
                    c_arr += queen_moves(board, j, i)
                elif board[j, i] == "bpawn":
                    c_arr += bpawn_moves(board, j, i)
                elif board[j, i] == "wpawn":
                    c_arr += wpawn_moves(board, j, i)
    c_arr += get_possible_castlings(board, color)
    moves = filter_moves(board, color, c_arr)
    #for i in range(len(moves)):
        #print_board(make_move(board, moves[i]))
        #print(" ")
    return moves
    
def count_boards(board : Board, color : str, r : int) -> int:
    if r == 0:
        return 1
    c = 0
    moves = get_possible_moves(board, color)
    for i in range(len(moves)):
        new_board = board.make_move(moves[i])
        c = c + count_boards(new_board, flip_color(color), r - 1)
    return c
    
def print_rec_table(board : Board) -> None:
    for i in range(7):
        print(str(i) + " - " + str(count_boards(board, "w", i)))
        
#board = Board()
#print(count_boards(board, "w", 1))
#print_rec_table(board)



#print_threat_map(fill_diagonal_dl(board, threat_map, 4,4))
#print_threat_map(calculate_threats(board, threat_map, "b"))
#print(get_king_coordinates(board, "w"))
#print(is_check(board, "w"))
#print_board(board)
        
#print(get_cell_color(board, 4, 7))
#get_possible_moves(board, "w")