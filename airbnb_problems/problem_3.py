"""This problem was asked by Airbnb.

Given a list of words, find all pairs of unique 
indices such that the concatenation of the two words is a palindrome.

For example, given the list ["code", "edoc", "da", "d"], return [(0, 1), (1, 0), (2, 3)].
"""

"""
function findPalindromePairs(words):
    result = []  // List to store the pairs of indices
    wordMap = {}  // Map to store each word and its index

    // Step 1: Populate the word map
    for index, word in enumerate(words):
        wordMap[word] = index

    // Step 2: Iterate through each word and check for palindrome pairs
    for index, word in enumerate(words):
        for i in range(len(word) + 1):  // +1 to include empty suffix
            prefix = word[:i]
            suffix = word[i:]

            // Case 1: If the prefix is a palindrome, check for the reverse of the suffix
            if isPalindrome(prefix):
                reversedSuffix = reverseString(suffix)
                if reversedSuffix in wordMap and wordMap[reversedSuffix] != index:
                    result.append([wordMap[reversedSuffix], index])

            // Case 2: If the suffix is a palindrome, check for the reverse of the prefix
            // Ensure suffix is not empty to avoid duplicates
            if i > 0 and isPalindrome(suffix):
                reversedPrefix = reverseString(prefix)
                if reversedPrefix in wordMap and wordMap[reversedPrefix] != index:
                    result.append([index, wordMap[reversedPrefix]])

    return result

function isPalindrome(s):
    return s == reverseString(s)

function reverseString(s):
    return s[::-1]  // Reverse the string
"""

Explanation
Populate the Word Map:

Create a dictionary where keys are the words and values are their indices. This allows O(1) lookups for reversed words.
Iterate Through Each Word:

Split the word into all possible prefixes and suffixes.
For example, for "code", the splits are:
Prefixes: ["", "c", "co", "cod", "code"]
Suffixes: ["code", "ode", "de", "e", ""]
Check Palindrome Conditions:

Case 1: If the prefix is a palindrome, then the reversed suffix must exist in wordMap for it to form a valid pair.
Case 2: If the suffix is a palindrome, then the reversed prefix must exist in wordMap.
Avoid Self-Matches:

Ensure wordMap[reversedSuffix] != index and wordMap[reversedPrefix] != index to avoid matching a word with itself.
Handle Edge Cases:

Empty strings can pair with any palindrome, so include them in the checks.
Avoid duplicates by ensuring proper indexing and non-empty checks for certain cases.
Complexity
Time Complexity: O(n * k²), where n is the number of words and k is the average length of a word (due to splitting into prefixes and suffixes).
Space Complexity: O(n * k) for the word map and temporary string operations.
Let me know if you’d like further clarification or an actual implementation in Go!