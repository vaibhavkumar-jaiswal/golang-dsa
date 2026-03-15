Problem Statement:

Given an array of integers arr[] of size n, find the third largest DISTINCT element in the array.

A distinct element means the value must be different from the first and second largest elements.

If the third largest distinct element does not exist, return -1.

The array may contain:
- Positive numbers
- Negative numbers
- Duplicate values

Examples:

Input: arr = [1, 14, 2, 16, 10, 20]
Output: 14
Explanation:
The largest element is 20, the second largest element is 16, and the third largest distinct element is 14.

Input: arr = [19, -10, 20, 14, 2, 16, 10]
Output: 16
Explanation:
The largest element is 20, the second largest element is 19, and the third largest distinct element is 16.

Input: arr = [10, 10, 10]
Output: -1
Explanation:
All elements are the same, so there is no third largest distinct element.

Input: arr = [-10, -20, -30, -5]
Output: -20
Explanation:
The largest element is -5, the second largest element is -10, and the third largest distinct element is -20.