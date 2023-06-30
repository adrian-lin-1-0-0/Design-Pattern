module Model where
    type Model = [[Float]]

    multiplyMatrix :: Model -> [Float] -> [Float]
    multiplyMatrix matrix vector = map (sum . zipWith (*) vector) matrix

