# -*- coding: utf-8 -*-

if __name__ == "__main__":
    import ctypes

    lib = ctypes.CDLL('./libadd.dll')
    print lib.Add(7, 11)
