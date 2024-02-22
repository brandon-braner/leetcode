# this is O(n^2) we can do better
def bubble_sort_with_zeros(arr):
	n = len(arr)
	for i in range(n - 1):
		swapped = False
		for j in range(n - i - 1):
			if arr[j] == 0 and arr[j + 1] != 0:
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				swapped = True
		if not swapped:
			break
	return arr