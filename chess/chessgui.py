import chess
from tkinter import *
from functools import partial

t = Tk()
for i in range(64):
    txt = str(i)
    b = Button(text=txt, command=partial(lambda n: print(n), i))   
    b.grid(row=i//8, column=i%8)

board = chess.Board()


mainloop()