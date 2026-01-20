import heapq

arr = [1, 3, 2, 8, 5, 11, 10, 19, 25, 15]

# convert an array into a heap in place
heapq.heapify(arr)

# push 0 to the heap
heapq.heappush(arr, 0)

min_element = heapq.heappop(arr)

print(min_element)
