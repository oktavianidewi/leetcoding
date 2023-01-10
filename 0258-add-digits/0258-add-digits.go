func addDigits(num int) int {

	for {

		add := 0
		for num > 0 {
			add += num % 10
			num = num / 10
		}

		if add < 10 {
			return add
		}
		num = add
	}

}
