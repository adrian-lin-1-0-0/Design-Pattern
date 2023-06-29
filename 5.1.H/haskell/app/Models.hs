module Models where
    import Model
    import MemCache
    import Control.Concurrent
    import Control.Monad
    import Data.Map qualified as Map
    import System.IO
    import Data.Array

    readModel :: Store -> String -> IO Model
    readModel store modelName = do
        maybeModel <- retrieveValue store modelName
        case maybeModel of
            Just m -> return m
            Nothing -> do
                print "Reading file"
                let fileName = modelName ++ ".mat"
                contents <- readFile fileName
                let rows = lines contents
                    newModel = map (map read . words) rows
                storeValue store modelName newModel
                return newModel

    getModel :: Store -> String -> IO Model
    getModel store modelName = do readModel store modelName

    multiplyMatrix :: [[Float]] -> [Float] -> [Float]
    multiplyMatrix matrix vector = map (sum . zipWith (*) vector) matrix
