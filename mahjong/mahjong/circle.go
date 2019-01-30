package mahjong

type Circle struct {
	Name  string
	Rank  int
	Point int
}

func Circle1() Circle {
	return Circle{
		"circle_1",
		0,
		1,
	}
}

func Circle2() Circle {
	return Circle{
		"circle_2",
		0,
		2,
	}
}

func Circle3() Circle {
	return Circle{
		"circle_3",
		0,
		3,
	}
}

func Circle4() Circle {
	return Circle{
		"circle_4",
		0,
		4,
	}
}

func Circle5() Circle {
	return Circle{
		"circle_5",
		0,
		5,
	}
}

func Circle6() Circle {
	return Circle{
		"circle_6",
		0,
		6,
	}
}

func Circle7() Circle {
	return Circle{
		"circle_7",
		0,
		7,
	}
}

func Circle8() Circle {
	return Circle{
		"circle_8",
		0,
		8,
	}
}

func Circle9() Circle {
	return Circle{
		"circle_9",
		0,
		9,
	}
}

func CircleUndifined() Circle {
	return Circle{
		"circle_undefined",
		0,
		-1,
	}
}
