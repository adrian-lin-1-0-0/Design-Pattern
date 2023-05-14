```mermaid
---
title: Big 2 
---
class Big-2{

}

class Player{
    - name : String
    - cards : Card[]
    - play(cards : Card[])
    - pass()
    - draw(cards : Card[])
}

class Card{
    - suit : String
    - rank : String
}

class Deck{
    - cards : Card[]
    - shuffle()
    - deal()
}

class Game{
    - players : Player[]
    - deck : Deck
    - play()
    - pass()
    - draw()
}

class Hand{
    - cards : Card[]
    - isStraight()
    - isFlush()
    - isFullHouse()
    - isFourOfAKind()
    - isStraightFlush()
    - isRoyalFlush()
}

class HandType{
    - name : String
    - isMatch(hand : Hand)
}

class Straight extends HandType{
    - name : String
    - isMatch(hand : Hand)
}

class Flush extends HandType{
    - name : String
    - isMatch(hand : Hand)
}

class FullHouse extends HandType{
    - name : String
    - isMatch(hand : Hand)
}

class FourOfAKind extends HandType{
    - name : String
    - isMatch(hand : Hand)
}

class StraightFlush extends HandType{
    - name : String
    - isMatch(hand : Hand)
}

class RoyalFlush extends HandType{
    - name : String
    - isMatch(hand : Hand)
}

class HandTypeFactory{
    - handTypes : HandType[]
    - getHandType(hand : Hand)
}

class HandTypeComparator{
    - handTypes : HandType[]
    - compare(hand1 : Hand, hand2 : Hand)
}

class HandComparator{
    - handTypeComparator : HandTypeComparator
    - compare(hand1 : Hand, hand2 : Hand)
}

class HandRankComparator{
    - handTypeComparator : HandTypeComparator
    - compare(hand1 : Hand, hand2 : Hand)
}

class HandSuitComparator{
    - handTypeComparator : HandTypeComparator
    - compare(hand1 : Hand, hand2 : Hand)
}

class HandRankSuitComparator{
    - handTypeComparator : HandTypeComparator
    - compare(hand1 : Hand, hand2 : Hand)
}



```