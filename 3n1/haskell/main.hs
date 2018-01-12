-- run it with 
--    ghc main.hs -o main.exe
--    main.exe

module Main where

succ_3n1::Int->Int
succ_3n1 1 = 1
succ_3n1 i = if odd i then 3*i+1 else div i 2

main = do 
    let i = 11
    print (succ_3n1 i)
