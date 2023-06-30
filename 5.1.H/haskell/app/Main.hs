module Main where
import Models
import MemCache
import Model

main :: IO ()
main = do
  reflection <- getModel "Reflection"
  reflection2 <- getModel "Reflection"
  -- Reading file from disk only once
  let vector = replicate 1000 1.0 :: [Float]
  let result = multiplyMatrix reflection vector
  print result
  -- replicate 1000 -1.0
  scaling <- getModel "Scaling"
  -- Reading file from disk
  let result2 = multiplyMatrix scaling vector
  print result2
  -- replicate 1000 2.0
  shrinking <- getModel "Shrinking"
  -- Reading file from disk
  let result3 = multiplyMatrix shrinking vector
  print result3
  -- replicate 1000 0.5
