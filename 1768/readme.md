# Problem:

You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

[Link to the problem] (https://leetcode.com/problems/merge-strings-alternately/)


# Solution:

This is an O(n^2)
- check to make sure both strings are not empty
- find the length of the longest string
- create a new string to store the merged string
- loop over both strings and add the characters to the new string 
- once the shortest string is exhausted, add the remaining characters from the longer string