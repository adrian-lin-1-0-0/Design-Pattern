module MemCache where
    import Model
    import Control.Concurrent
    import Data.Map qualified as Map

    type Key = String
    type Store = MVar (Map.Map Key Model)

    memStore :: IO Store
    memStore = newMVar Map.empty

    storeValue :: Store -> Key -> Model -> IO ()
    storeValue store key value = modifyMVar_ store $ \storeMap ->
        return $ Map.insert key value storeMap

    retrieveValue :: Store -> Key -> IO (Maybe Model)
    retrieveValue store key = withMVar store $ \storeMap ->
        return $ Map.lookup key storeMap