Problem Statement:

Given an array of integers arr[] of size n, find the second largest DISTINCT element in the array.

A distinct element means the value must be different from the largest element.

If the second largest distinct element does not exist, return -1.

The array may contain:
- Positive numbers
- Negative numbers
- Duplicate values

Examples:

Input: arr = [12, 35, 1, 10, 34, 1]
Output: 34
Explanation:
The largest element is 35 and the second largest distinct element is 34.

Input: arr = [10, 5, 10]
Output: 5
Explanation:
The largest element is 10 and the second largest distinct element is 5.

Input: arr = [10, 10, 10]
Output: -1
Explanation:
All elements are the same, so there is no second largest distinct element.

Input: arr = [-10, -20, -30, -5]
Output: -10
Explanation:
The largest element is -5 and the second largest distinct element is -10.