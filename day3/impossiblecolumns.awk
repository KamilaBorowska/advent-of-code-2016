function inssort(arr, len,   i, j, value) {
	for (i = 2; i <= len; i++) {
		value = arr[i]
		j = i - 1
		while (j > 0 && arr[j] > value) {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = value
	}
}

BEGIN {
	count = 0
	row = 0
}

{
	row++
	a[row] = $1
	b[row] = $2
	c[row] = $3
}

row == 3 {
	inssort(a, 3)
	if (a[1] + a[2] > a[3]) count++

	inssort(b, 3)
	if (b[1] + b[2] > b[3]) count++

	inssort(c, 3)
	if (c[1] + c[2] > c[3]) count++

	row = 0
}

END {
	if (row != 0) {
		print "Not enough rows"
	}
	print count
}
