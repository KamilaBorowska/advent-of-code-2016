function inssort(arr, len,   i, j, value) {
	for (i = 2; i <= len; i++) {
		value = arr[i]
		j = i - 1
		while (j > 0 && array[j] > value) {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = value
	}
}

BEGIN {
	count = 0
}

{
	array[1] = $1
	array[2] = $2
	array[3] = $3
	inssort(array, 3)
}

array[1] + array[2] > array[3] {
	count++
}

END {
	print count
}
