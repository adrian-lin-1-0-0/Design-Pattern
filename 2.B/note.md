## Force
- player 出牌時檢查是否符合規則,不符合則回傳錯誤訊息,並重新出牌
  > class Player  內需要有 class CardPattern ?
- 如果 class Player 內沒有 class CardPattern , 那就需要在 player 出牌後,handcard 產生變動後,再檢查是否符合規則
  > 如果不符合規則, handCard 如何回到原本的狀態?