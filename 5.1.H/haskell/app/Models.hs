module Models where
    import Model
    import MemCache
    import System.IO
    import Data.Array

    getModel :: String -> IO Model
    getModel modelName = do
        store <- memStore
        maybeModel <- retrieveValue store modelName
        case maybeModel of
            Just m -> return m
            Nothing -> do readModel modelName

    readModel :: String -> IO Model
    readModel modelName = do
        print $ "Reading model " ++ modelName ++ " from disk"
        let fileName = modelName ++ ".mat"
        contents <- readFile fileName
        let rows = lines contents
            newModel = map (map read . words) rows
        return newModel