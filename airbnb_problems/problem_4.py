"""This problem was asked by Airbnb.
Given a linked list and a positive integer k, rotate the list to the right by k places.

For example, given the linked list 

7 -> 7 -> 3 -> 5 and k = 2, 
it should become 3 -> 5 -> 7 -> 7.

Given the linked list 
1 -> 2 -> 3 -> 4 -> 5 and k = 3

 it should become 3 -> 4 -> 5 -> 1 -> 2
"""

shift to right by k places
matter of detecting who is the front and who is the back 


a - b - c - d - e -f- g 
g - a - b - c - d - e - f
f - g - a - b - c - d - e
e - f - g - a - b - c - d 


root = last_idx -k + 1
 
last - k = last
last - k + 1 = first
first = 

last.next = first

