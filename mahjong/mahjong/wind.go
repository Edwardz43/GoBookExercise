package mahjong

type Wind struct {
	Name string
	Rank int
}

func EastWind() Wind {
	return Wind{
		"east_wind",
		0,
	}
}

func WestWind() Wind {
	return Wind{
		"west_wind",
		0,
	}
}

func SouthWind() Wind {
	return Wind{
		"south_wind",
		0,
	}
}

func NorthWind() Wind {
	return Wind{
		"north_wind",
		0,
	}
}
