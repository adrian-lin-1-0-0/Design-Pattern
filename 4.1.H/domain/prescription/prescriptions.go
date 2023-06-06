package prescription

var (
	Covid19 = NewPrescription(
		"清冠一號",
		"COVID-19",
		[]Medicine{
			NewMedicine("清冠一號"),
		},
		"將相關藥材裝入茶包裡，使用500 mL 溫、熱水沖泡悶煮1~3 分鐘後即可飲用。",
	)
	Attractive = NewPrescription(
		"青春抑制劑",
		"Attractive",
		[]Medicine{
			NewMedicine("假鬢角"),
			NewMedicine("臭味"),
		},
		"把假鬢角黏在臉的兩側，讓自己異性緣差一點，自然就不會有人想妳了。",
	)
	SleepApneaSyndrome = NewPrescription(
		"打呼抑制劑",
		"SleepApneaSyndrome",
		[]Medicine{
			NewMedicine("一捲膠帶"),
		},
		"睡覺時，撕下兩塊膠帶，將兩塊膠帶交錯黏在關閉的嘴巴上，就不會打呼了。",
	)
)
