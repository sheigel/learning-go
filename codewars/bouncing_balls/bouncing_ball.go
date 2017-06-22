package kata

func BouncingBall(height, bounce, window float64) int {
	switch {
	case height <= window:
		return -1
	case height <= 0:
		return -1
	case bounce <= 0 || bounce >= 1:
		return -1
	}

	return BallCount(height, bounce, window)
}


func BallCount(height, bounce, window float64) int {
	switch{
	case height == window:
		return 1
	case height < window:
		return -1
	case height <= 0:
		return 0
	}

	return 2 + BallCount(height * bounce, bounce, window)
}