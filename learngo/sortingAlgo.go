package learngo

/*
{1,5,7,9,3,2} --> {1,5,7,9,3,2} --> {1,3,7,9,5,2}
									{1,2,7,9,5,3} --> {1,2,5,9,7,3}
													  {1,2,3,9,7,5} --> .. --> {1,2,3,5,7,9}
Bubble Sort:
Characteristics: Simple and straightforward. It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.
Pros: Easy to understand and implement.
Cons: Inefficient for large datasets. Time complexity of O(n^2).
Use Case: Not typically used in practical scenarios due to its inefficiency, more suitable for educational purposes.


Selection Sort:
Characteristics: Divides the input list into two parts: the sublist of items already sorted and the sublist of items remaining to be sorted. Finds the smallest element from the unsorted sublist and swaps it with the leftmost unsorted element.
Pros: Simple to implement, performs well for small datasets.
Cons: Inefficient for large datasets. Time complexity of O(n^2) regardless of input.
Use Case: Suitable for small datasets or situations where memory is limited.

Insertion Sort:
Characteristics: Builds the final sorted array one item at a time by repeatedly taking the next item and inserting it into the correct position in the already sorted part of the array.
Pros: Efficient for small datasets, works well with nearly sorted arrays.
Cons: Inefficient for large datasets. Time complexity of O(n^2).
Use Case: Suitable for small datasets or partially sorted arrays.

Merge Sort:
Characteristics: Uses the divide-and-conquer strategy by dividing the unsorted list into two halves, sorting each half, and merging them back together.
Pros: Efficient and stable. Guaranteed time complexity of O(n log n). Suitable for large datasets.
Cons: Requires additional memory space for the merging process.
Use Case: General-purpose sorting algorithm for large datasets. Good choice for sorting linked lists.

Quick Sort:
Characteristics: Uses the divide-and-conquer strategy by selecting a 'pivot' element and partitioning the array around the pivot, such that elements less than the pivot are placed before it, and elements greater than the pivot are placed after it.
Pros: Efficient with an average-case time complexity of O(n log n). In-place sorting (space-efficient).
Cons: Worst-case time complexity of O(n^2) can occur, depending on the choice of the pivot.
Use Case: General-purpose sorting algorithm for large datasets. Often preferred over other algorithms due to its average-case efficiency.

Heap Sort:
Characteristics: Builds a heap data structure and repeatedly extracts the maximum (for a max-heap) element from it, resulting in a sorted array.
Pros: Efficient with a guaranteed time complexity of O(n log n). In-place sorting.
Cons: Not stable. Slower than quicksort for most inputs.
Use Case: Suitable for scenarios requiring a stable sort with a guaranteed worst-case time complexity. Often used in priority queue implementations.
*/
