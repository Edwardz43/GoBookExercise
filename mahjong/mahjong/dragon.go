package mahjong

type Dragon struct {
	Name string
	Rank int
}

func RedDragon() Dragon {
	return Dragon{"red_dragon", 0}
}

func GreenDragon() Dragon {
	return Dragon{"green_dragon", 0}
}

func WhiteDragon() Dragon {
	return Dragon{"white_dragon", 0}
}
