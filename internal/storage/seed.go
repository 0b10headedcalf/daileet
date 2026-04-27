package storage

import (
	"database/sql"
	"fmt"

	"github.com/0b10headedcalf/daileet/internal/models"
)

// SeedBlind75 inserts the canonical Blind 75 list into the DB.
func SeedBlind75(db *sql.DB) error {
	problems := []models.Problem{
		// Arrays & Hashing
		{Title: "Contains Duplicate", TitleSlug: "contains-duplicate", Difficulty: models.Easy, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/contains-duplicate/"},
		{Title: "Valid Anagram", TitleSlug: "valid-anagram", Difficulty: models.Easy, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/valid-anagram/"},
		{Title: "Two Sum", TitleSlug: "two-sum", Difficulty: models.Easy, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/two-sum/"},
		{Title: "Group Anagrams", TitleSlug: "group-anagrams", Difficulty: models.Medium, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/group-anagrams/"},
		{Title: "Top K Frequent Elements", TitleSlug: "top-k-frequent-elements", Difficulty: models.Medium, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/top-k-frequent-elements/"},
		{Title: "Product of Array Except Self", TitleSlug: "product-of-array-except-self", Difficulty: models.Medium, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/product-of-array-except-self/"},
		{Title: "Valid Sudoku", TitleSlug: "valid-sudoku", Difficulty: models.Medium, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/valid-sudoku/"},
		{Title: "Encode and Decode Strings", TitleSlug: "encode-and-decode-strings", Difficulty: models.Medium, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/encode-and-decode-strings/"},
		{Title: "Longest Consecutive Sequence", TitleSlug: "longest-consecutive-sequence", Difficulty: models.Medium, Pattern: "Arrays & Hashing", URL: "https://leetcode.com/problems/longest-consecutive-sequence/"},

		// Two Pointers
		{Title: "Valid Palindrome", TitleSlug: "valid-palindrome", Difficulty: models.Easy, Pattern: "Two Pointers", URL: "https://leetcode.com/problems/valid-palindrome/"},
		{Title: "Two Sum II Input Array Is Sorted", TitleSlug: "two-sum-ii-input-array-is-sorted", Difficulty: models.Medium, Pattern: "Two Pointers", URL: "https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/"},
		{Title: "3Sum", TitleSlug: "3sum", Difficulty: models.Medium, Pattern: "Two Pointers", URL: "https://leetcode.com/problems/3sum/"},
		{Title: "Container With Most Water", TitleSlug: "container-with-most-water", Difficulty: models.Medium, Pattern: "Two Pointers", URL: "https://leetcode.com/problems/container-with-most-water/"},
		{Title: "Trapping Rain Water", TitleSlug: "trapping-rain-water", Difficulty: models.Hard, Pattern: "Two Pointers", URL: "https://leetcode.com/problems/trapping-rain-water/"},

		// Sliding Window
		{Title: "Best Time to Buy and Sell Stock", TitleSlug: "best-time-to-buy-and-sell-stock", Difficulty: models.Easy, Pattern: "Sliding Window", URL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock/"},
		{Title: "Longest Substring Without Repeating Characters", TitleSlug: "longest-substring-without-repeating-characters", Difficulty: models.Medium, Pattern: "Sliding Window", URL: "https://leetcode.com/problems/longest-substring-without-repeating-characters/"},
		{Title: "Longest Repeating Character Replacement", TitleSlug: "longest-repeating-character-replacement", Difficulty: models.Medium, Pattern: "Sliding Window", URL: "https://leetcode.com/problems/longest-repeating-character-replacement/"},
		{Title: "Permutation In String", TitleSlug: "permutation-in-string", Difficulty: models.Medium, Pattern: "Sliding Window", URL: "https://leetcode.com/problems/permutation-in-string/"},
		{Title: "Minimum Window Substring", TitleSlug: "minimum-window-substring", Difficulty: models.Hard, Pattern: "Sliding Window", URL: "https://leetcode.com/problems/minimum-window-substring/"},
		{Title: "Sliding Window Maximum", TitleSlug: "sliding-window-maximum", Difficulty: models.Hard, Pattern: "Sliding Window", URL: "https://leetcode.com/problems/sliding-window-maximum/"},

		// Stack
		{Title: "Valid Parentheses", TitleSlug: "valid-parentheses", Difficulty: models.Easy, Pattern: "Stack", URL: "https://leetcode.com/problems/valid-parentheses/"},
		{Title: "Min Stack", TitleSlug: "min-stack", Difficulty: models.Medium, Pattern: "Stack", URL: "https://leetcode.com/problems/min-stack/"},
		{Title: "Evaluate Reverse Polish Notation", TitleSlug: "evaluate-reverse-polish-notation", Difficulty: models.Medium, Pattern: "Stack", URL: "https://leetcode.com/problems/evaluate-reverse-polish-notation/"},
		{Title: "Generate Parentheses", TitleSlug: "generate-parentheses", Difficulty: models.Medium, Pattern: "Stack", URL: "https://leetcode.com/problems/generate-parentheses/"},
		{Title: "Daily Temperatures", TitleSlug: "daily-temperatures", Difficulty: models.Medium, Pattern: "Stack", URL: "https://leetcode.com/problems/daily-temperatures/"},
		{Title: "Car Fleet", TitleSlug: "car-fleet", Difficulty: models.Medium, Pattern: "Stack", URL: "https://leetcode.com/problems/car-fleet/"},
		{Title: "Largest Rectangle In Histogram", TitleSlug: "largest-rectangle-in-histogram", Difficulty: models.Hard, Pattern: "Stack", URL: "https://leetcode.com/problems/largest-rectangle-in-histogram/"},

		// Binary Search
		{Title: "Binary Search", TitleSlug: "binary-search", Difficulty: models.Easy, Pattern: "Binary Search", URL: "https://leetcode.com/problems/binary-search/"},
		{Title: "Search a 2D Matrix", TitleSlug: "search-a-2d-matrix", Difficulty: models.Medium, Pattern: "Binary Search", URL: "https://leetcode.com/problems/search-a-2d-matrix/"},
		{Title: "Koko Eating Bananas", TitleSlug: "koko-eating-bananas", Difficulty: models.Medium, Pattern: "Binary Search", URL: "https://leetcode.com/problems/koko-eating-bananas/"},
		{Title: "Search In Rotated Sorted Array", TitleSlug: "search-in-rotated-sorted-array", Difficulty: models.Medium, Pattern: "Binary Search", URL: "https://leetcode.com/problems/search-in-rotated-sorted-array/"},
		{Title: "Find Minimum In Rotated Sorted Array", TitleSlug: "find-minimum-in-rotated-sorted-array", Difficulty: models.Medium, Pattern: "Binary Search", URL: "https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/"},
		{Title: "Time Based Key Value Store", TitleSlug: "time-based-key-value-store", Difficulty: models.Medium, Pattern: "Binary Search", URL: "https://leetcode.com/problems/time-based-key-value-store/"},
		{Title: "Median of Two Sorted Arrays", TitleSlug: "median-of-two-sorted-arrays", Difficulty: models.Hard, Pattern: "Binary Search", URL: "https://leetcode.com/problems/median-of-two-sorted-arrays/"},

		// Linked List
		{Title: "Reverse Linked List", TitleSlug: "reverse-linked-list", Difficulty: models.Easy, Pattern: "Linked List", URL: "https://leetcode.com/problems/reverse-linked-list/"},
		{Title: "Merge Two Sorted Lists", TitleSlug: "merge-two-sorted-lists", Difficulty: models.Easy, Pattern: "Linked List", URL: "https://leetcode.com/problems/merge-two-sorted-lists/"},
		{Title: "Linked List Cycle", TitleSlug: "linked-list-cycle", Difficulty: models.Easy, Pattern: "Linked List", URL: "https://leetcode.com/problems/linked-list-cycle/"},
		{Title: "Reorder List", TitleSlug: "reorder-list", Difficulty: models.Medium, Pattern: "Linked List", URL: "https://leetcode.com/problems/reorder-list/"},
		{Title: "Remove Nth Node From End of List", TitleSlug: "remove-nth-node-from-end-of-list", Difficulty: models.Medium, Pattern: "Linked List", URL: "https://leetcode.com/problems/remove-nth-node-from-end-of-list/"},
		{Title: "Copy List With Random Pointer", TitleSlug: "copy-list-with-random-pointer", Difficulty: models.Medium, Pattern: "Linked List", URL: "https://leetcode.com/problems/copy-list-with-random-pointer/"},
		{Title: "Add Two Numbers", TitleSlug: "add-two-numbers", Difficulty: models.Medium, Pattern: "Linked List", URL: "https://leetcode.com/problems/add-two-numbers/"},
		{Title: "Find The Duplicate Number", TitleSlug: "find-the-duplicate-number", Difficulty: models.Medium, Pattern: "Linked List", URL: "https://leetcode.com/problems/find-the-duplicate-number/"},
		{Title: "LRU Cache", TitleSlug: "lru-cache", Difficulty: models.Medium, Pattern: "Linked List", URL: "https://leetcode.com/problems/lru-cache/"},
		{Title: "Merge K Sorted Lists", TitleSlug: "merge-k-sorted-lists", Difficulty: models.Hard, Pattern: "Linked List", URL: "https://leetcode.com/problems/merge-k-sorted-lists/"},
		{Title: "Reverse Nodes In K Group", TitleSlug: "reverse-nodes-in-k-group", Difficulty: models.Hard, Pattern: "Linked List", URL: "https://leetcode.com/problems/reverse-nodes-in-k-group/"},

		// Trees
		{Title: "Invert Binary Tree", TitleSlug: "invert-binary-tree", Difficulty: models.Easy, Pattern: "Trees", URL: "https://leetcode.com/problems/invert-binary-tree/"},
		{Title: "Maximum Depth of Binary Tree", TitleSlug: "maximum-depth-of-binary-tree", Difficulty: models.Easy, Pattern: "Trees", URL: "https://leetcode.com/problems/maximum-depth-of-binary-tree/"},
		{Title: "Diameter of Binary Tree", TitleSlug: "diameter-of-binary-tree", Difficulty: models.Easy, Pattern: "Trees", URL: "https://leetcode.com/problems/diameter-of-binary-tree/"},
		{Title: "Balanced Binary Tree", TitleSlug: "balanced-binary-tree", Difficulty: models.Easy, Pattern: "Trees", URL: "https://leetcode.com/problems/balanced-binary-tree/"},
		{Title: "Same Tree", TitleSlug: "same-tree", Difficulty: models.Easy, Pattern: "Trees", URL: "https://leetcode.com/problems/same-tree/"},
		{Title: "Subtree of Another Tree", TitleSlug: "subtree-of-another-tree", Difficulty: models.Easy, Pattern: "Trees", URL: "https://leetcode.com/problems/subtree-of-another-tree/"},
		{Title: "Lowest Common Ancestor of a Binary Search Tree", TitleSlug: "lowest-common-ancestor-of-a-binary-search-tree", Difficulty: models.Medium, Pattern: "Trees", URL: "https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/"},
		{Title: "Binary Tree Level Order Traversal", TitleSlug: "binary-tree-level-order-traversal", Difficulty: models.Medium, Pattern: "Trees", URL: "https://leetcode.com/problems/binary-tree-level-order-traversal/"},
		{Title: "Binary Tree Right Side View", TitleSlug: "binary-tree-right-side-view", Difficulty: models.Medium, Pattern: "Trees", URL: "https://leetcode.com/problems/binary-tree-right-side-view/"},
		{Title: "Count Good Nodes In Binary Tree", TitleSlug: "count-good-nodes-in-binary-tree", Difficulty: models.Medium, Pattern: "Trees", URL: "https://leetcode.com/problems/count-good-nodes-in-binary-tree/"},
		{Title: "Validate Binary Search Tree", TitleSlug: "validate-binary-search-tree", Difficulty: models.Medium, Pattern: "Trees", URL: "https://leetcode.com/problems/validate-binary-search-tree/"},
		{Title: "Kth Smallest Element In a Bst", TitleSlug: "kth-smallest-element-in-a-bst", Difficulty: models.Medium, Pattern: "Trees", URL: "https://leetcode.com/problems/kth-smallest-element-in-a-bst/"},
		{Title: "Construct Binary Tree From Preorder And Inorder Traversal", TitleSlug: "construct-binary-tree-from-preorder-and-inorder-traversal", Difficulty: models.Medium, Pattern: "Trees", URL: "https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/"},
		{Title: "Binary Tree Maximum Path Sum", TitleSlug: "binary-tree-maximum-path-sum", Difficulty: models.Hard, Pattern: "Trees", URL: "https://leetcode.com/problems/binary-tree-maximum-path-sum/"},
		{Title: "Serialize And Deserialize Binary Tree", TitleSlug: "serialize-and-deserialize-binary-tree", Difficulty: models.Hard, Pattern: "Trees", URL: "https://leetcode.com/problems/serialize-and-deserialize-binary-tree/"},

		// Heap / Priority Queue
		{Title: "Kth Largest Element In a Stream", TitleSlug: "kth-largest-element-in-a-stream", Difficulty: models.Easy, Pattern: "Heap / Priority Queue", URL: "https://leetcode.com/problems/kth-largest-element-in-a-stream/"},
		{Title: "Last Stone Weight", TitleSlug: "last-stone-weight", Difficulty: models.Easy, Pattern: "Heap / Priority Queue", URL: "https://leetcode.com/problems/last-stone-weight/"},
		{Title: "K Closest Points to Origin", TitleSlug: "k-closest-points-to-origin", Difficulty: models.Medium, Pattern: "Heap / Priority Queue", URL: "https://leetcode.com/problems/k-closest-points-to-origin/"},
		{Title: "Kth Largest Element In An Array", TitleSlug: "kth-largest-element-in-an-array", Difficulty: models.Medium, Pattern: "Heap / Priority Queue", URL: "https://leetcode.com/problems/kth-largest-element-in-an-array/"},
		{Title: "Task Scheduler", TitleSlug: "task-scheduler", Difficulty: models.Medium, Pattern: "Heap / Priority Queue", URL: "https://leetcode.com/problems/task-scheduler/"},
		{Title: "Design Twitter", TitleSlug: "design-twitter", Difficulty: models.Medium, Pattern: "Heap / Priority Queue", URL: "https://leetcode.com/problems/design-twitter/"},
		{Title: "Find Median From Data Stream", TitleSlug: "find-median-from-data-stream", Difficulty: models.Hard, Pattern: "Heap / Priority Queue", URL: "https://leetcode.com/problems/find-median-from-data-stream/"},

		// Backtracking
		{Title: "Subsets", TitleSlug: "subsets", Difficulty: models.Medium, Pattern: "Backtracking", URL: "https://leetcode.com/problems/subsets/"},
		{Title: "Combination Sum", TitleSlug: "combination-sum", Difficulty: models.Medium, Pattern: "Backtracking", URL: "https://leetcode.com/problems/combination-sum/"},
		{Title: "Permutations", TitleSlug: "permutations", Difficulty: models.Medium, Pattern: "Backtracking", URL: "https://leetcode.com/problems/permutations/"},
		{Title: "Subsets II", TitleSlug: "subsets-ii", Difficulty: models.Medium, Pattern: "Backtracking", URL: "https://leetcode.com/problems/subsets-ii/"},
		{Title: "Combination Sum II", TitleSlug: "combination-sum-ii", Difficulty: models.Medium, Pattern: "Backtracking", URL: "https://leetcode.com/problems/combination-sum-ii/"},
		{Title: "Word Search", TitleSlug: "word-search", Difficulty: models.Medium, Pattern: "Backtracking", URL: "https://leetcode.com/problems/word-search/"},
		{Title: "Palindrome Partitioning", TitleSlug: "palindrome-partitioning", Difficulty: models.Medium, Pattern: "Backtracking", URL: "https://leetcode.com/problems/palindrome-partitioning/"},
		{Title: "Letter Combinations of a Phone Number", TitleSlug: "letter-combinations-of-a-phone-number", Difficulty: models.Medium, Pattern: "Backtracking", URL: "https://leetcode.com/problems/letter-combinations-of-a-phone-number/"},
		{Title: "N-Queens", TitleSlug: "n-queens", Difficulty: models.Hard, Pattern: "Backtracking", URL: "https://leetcode.com/problems/n-queens/"},

		// Tries
		{Title: "Implement Trie Prefix Tree", TitleSlug: "implement-trie-prefix-tree", Difficulty: models.Medium, Pattern: "Tries", URL: "https://leetcode.com/problems/implement-trie-prefix-tree/"},
		{Title: "Design Add And Search Words Data Structure", TitleSlug: "design-add-and-search-words-data-structure", Difficulty: models.Medium, Pattern: "Tries", URL: "https://leetcode.com/problems/design-add-and-search-words-data-structure/"},
		{Title: "Word Search II", TitleSlug: "word-search-ii", Difficulty: models.Hard, Pattern: "Tries", URL: "https://leetcode.com/problems/word-search-ii/"},

		// Graphs
		{Title: "Number of Islands", TitleSlug: "number-of-islands", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/number-of-islands/"},
		{Title: "Clone Graph", TitleSlug: "clone-graph", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/clone-graph/"},
		{Title: "Max Area of Island", TitleSlug: "max-area-of-island", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/max-area-of-island/"},
		{Title: "Pacific Atlantic Water Flow", TitleSlug: "pacific-atlantic-water-flow", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/pacific-atlantic-water-flow/"},
		{Title: "Surrounded Regions", TitleSlug: "surrounded-regions", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/surrounded-regions/"},
		{Title: "Rotting Oranges", TitleSlug: "rotting-oranges", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/rotting-oranges/"},
		{Title: "Walls and Gates", TitleSlug: "walls-and-gates", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/walls-and-gates/"},
		{Title: "Course Schedule", TitleSlug: "course-schedule", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/course-schedule/"},
		{Title: "Course Schedule II", TitleSlug: "course-schedule-ii", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/course-schedule-ii/"},
		{Title: "Redundant Connection", TitleSlug: "redundant-connection", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/redundant-connection/"},
		{Title: "Number of Provinces", TitleSlug: "number-of-provinces", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/number-of-provinces/"},
		{Title: "Graph Valid Tree", TitleSlug: "graph-valid-tree", Difficulty: models.Medium, Pattern: "Graphs", URL: "https://leetcode.com/problems/graph-valid-tree/"},
		{Title: "Word Ladder", TitleSlug: "word-ladder", Difficulty: models.Hard, Pattern: "Graphs", URL: "https://leetcode.com/problems/word-ladder/"},

		// Advanced Graphs
		{Title: "Reconstruct Itinerary", TitleSlug: "reconstruct-itinerary", Difficulty: models.Hard, Pattern: "Advanced Graphs", URL: "https://leetcode.com/problems/reconstruct-itinerary/"},
		{Title: "Min Cost to Connect All Points", TitleSlug: "min-cost-to-connect-all-points", Difficulty: models.Medium, Pattern: "Advanced Graphs", URL: "https://leetcode.com/problems/min-cost-to-connect-all-points/"},
		{Title: "Network Delay Time", TitleSlug: "network-delay-time", Difficulty: models.Medium, Pattern: "Advanced Graphs", URL: "https://leetcode.com/problems/network-delay-time/"},
		{Title: "Swim In Rising Water", TitleSlug: "swim-in-rising-water", Difficulty: models.Hard, Pattern: "Advanced Graphs", URL: "https://leetcode.com/problems/swim-in-rising-water/"},
		{Title: "Alien Dictionary", TitleSlug: "alien-dictionary", Difficulty: models.Hard, Pattern: "Advanced Graphs", URL: "https://leetcode.com/problems/alien-dictionary/"},
		{Title: "Cheapest Flights Within K Stops", TitleSlug: "cheapest-flights-within-k-stops", Difficulty: models.Medium, Pattern: "Advanced Graphs", URL: "https://leetcode.com/problems/cheapest-flights-within-k-stops/"},

		// 1-D DP
		{Title: "Climbing Stairs", TitleSlug: "climbing-stairs", Difficulty: models.Easy, Pattern: "1-D DP", URL: "https://leetcode.com/problems/climbing-stairs/"},
		{Title: "House Robber", TitleSlug: "house-robber", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/house-robber/"},
		{Title: "House Robber II", TitleSlug: "house-robber-ii", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/house-robber-ii/"},
		{Title: "Longest Palindromic Substring", TitleSlug: "longest-palindromic-substring", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/longest-palindromic-substring/"},
		{Title: "Palindromic Substrings", TitleSlug: "palindromic-substrings", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/palindromic-substrings/"},
		{Title: "Decode Ways", TitleSlug: "decode-ways", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/decode-ways/"},
		{Title: "Coin Change", TitleSlug: "coin-change", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/coin-change/"},
		{Title: "Maximum Product Subarray", TitleSlug: "maximum-product-subarray", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/maximum-product-subarray/"},
		{Title: "Word Break", TitleSlug: "word-break", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/word-break/"},
		{Title: "Longest Increasing Subsequence", TitleSlug: "longest-increasing-subsequence", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/longest-increasing-subsequence/"},
		{Title: "Partition Equal Subset Sum", TitleSlug: "partition-equal-subset-sum", Difficulty: models.Medium, Pattern: "1-D DP", URL: "https://leetcode.com/problems/partition-equal-subset-sum/"},

		// 2-D DP
		{Title: "Unique Paths", TitleSlug: "unique-paths", Difficulty: models.Medium, Pattern: "2-D DP", URL: "https://leetcode.com/problems/unique-paths/"},
		{Title: "Longest Common Subsequence", TitleSlug: "longest-common-subsequence", Difficulty: models.Medium, Pattern: "2-D DP", URL: "https://leetcode.com/problems/longest-common-subsequence/"},
		{Title: "Best Time to Buy and Sell Stock With Cooldown", TitleSlug: "best-time-to-buy-and-sell-stock-with-cooldown", Difficulty: models.Medium, Pattern: "2-D DP", URL: "https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/"},
		{Title: "Coin Change II", TitleSlug: "coin-change-ii", Difficulty: models.Medium, Pattern: "2-D DP", URL: "https://leetcode.com/problems/coin-change-ii/"},
		{Title: "Target Sum", TitleSlug: "target-sum", Difficulty: models.Medium, Pattern: "2-D DP", URL: "https://leetcode.com/problems/target-sum/"},
		{Title: "Interleaving String", TitleSlug: "interleaving-string", Difficulty: models.Medium, Pattern: "2-D DP", URL: "https://leetcode.com/problems/interleaving-string/"},
		{Title: "Edit Distance", TitleSlug: "edit-distance", Difficulty: models.Medium, Pattern: "2-D DP", URL: "https://leetcode.com/problems/edit-distance/"},
		{Title: "Regular Expression Matching", TitleSlug: "regular-expression-matching", Difficulty: models.Hard, Pattern: "2-D DP", URL: "https://leetcode.com/problems/regular-expression-matching/"},

		// Greedy
		{Title: "Maximum Subarray", TitleSlug: "maximum-subarray", Difficulty: models.Medium, Pattern: "Greedy", URL: "https://leetcode.com/problems/maximum-subarray/"},
		{Title: "Jump Game", TitleSlug: "jump-game", Difficulty: models.Medium, Pattern: "Greedy", URL: "https://leetcode.com/problems/jump-game/"},
		{Title: "Jump Game II", TitleSlug: "jump-game-ii", Difficulty: models.Medium, Pattern: "Greedy", URL: "https://leetcode.com/problems/jump-game-ii/"},
		{Title: "Gas Station", TitleSlug: "gas-station", Difficulty: models.Medium, Pattern: "Greedy", URL: "https://leetcode.com/problems/gas-station/"},
		{Title: "Hand of Straights", TitleSlug: "hand-of-straights", Difficulty: models.Medium, Pattern: "Greedy", URL: "https://leetcode.com/problems/hand-of-straights/"},
		{Title: "Merge Triplets to Form Target Triplet", TitleSlug: "merge-triplets-to-form-target-triplet", Difficulty: models.Medium, Pattern: "Greedy", URL: "https://leetcode.com/problems/merge-triplets-to-form-target-triplet/"},
		{Title: "Partition Labels", TitleSlug: "partition-labels", Difficulty: models.Medium, Pattern: "Greedy", URL: "https://leetcode.com/problems/partition-labels/"},
		{Title: "Valid Parenthesis String", TitleSlug: "valid-parenthesis-string", Difficulty: models.Medium, Pattern: "Greedy", URL: "https://leetcode.com/problems/valid-parenthesis-string/"},

		// Intervals
		{Title: "Insert Interval", TitleSlug: "insert-interval", Difficulty: models.Medium, Pattern: "Intervals", URL: "https://leetcode.com/problems/insert-interval/"},
		{Title: "Merge Intervals", TitleSlug: "merge-intervals", Difficulty: models.Medium, Pattern: "Intervals", URL: "https://leetcode.com/problems/merge-intervals/"},
		{Title: "Non Overlapping Intervals", TitleSlug: "non-overlapping-intervals", Difficulty: models.Medium, Pattern: "Intervals", URL: "https://leetcode.com/problems/non-overlapping-intervals/"},
		{Title: "Meeting Rooms", TitleSlug: "meeting-rooms", Difficulty: models.Easy, Pattern: "Intervals", URL: "https://leetcode.com/problems/meeting-rooms/"},
		{Title: "Meeting Rooms II", TitleSlug: "meeting-rooms-ii", Difficulty: models.Medium, Pattern: "Intervals", URL: "https://leetcode.com/problems/meeting-rooms-ii/"},
		{Title: "Minimum Interval to Include Each Query", TitleSlug: "minimum-interval-to-include-each-query", Difficulty: models.Hard, Pattern: "Intervals", URL: "https://leetcode.com/problems/minimum-interval-to-include-each-query/"},

		// Math & Geometry
		{Title: "Rotate Image", TitleSlug: "rotate-image", Difficulty: models.Medium, Pattern: "Math & Geometry", URL: "https://leetcode.com/problems/rotate-image/"},
		{Title: "Spiral Matrix", TitleSlug: "spiral-matrix", Difficulty: models.Medium, Pattern: "Math & Geometry", URL: "https://leetcode.com/problems/spiral-matrix/"},
		{Title: "Set Matrix Zeroes", TitleSlug: "set-matrix-zeroes", Difficulty: models.Medium, Pattern: "Math & Geometry", URL: "https://leetcode.com/problems/set-matrix-zeroes/"},
		{Title: "Happy Number", TitleSlug: "happy-number", Difficulty: models.Easy, Pattern: "Math & Geometry", URL: "https://leetcode.com/problems/happy-number/"},
		{Title: "Plus One", TitleSlug: "plus-one", Difficulty: models.Easy, Pattern: "Math & Geometry", URL: "https://leetcode.com/problems/plus-one/"},
		{Title: "Pow(x, n)", TitleSlug: "powx-n", Difficulty: models.Medium, Pattern: "Math & Geometry", URL: "https://leetcode.com/problems/powx-n/"},
		{Title: "Multiply Strings", TitleSlug: "multiply-strings", Difficulty: models.Medium, Pattern: "Math & Geometry", URL: "https://leetcode.com/problems/multiply-strings/"},
		{Title: "Detect Squares", TitleSlug: "detect-squares", Difficulty: models.Medium, Pattern: "Math & Geometry", URL: "https://leetcode.com/problems/detect-squares/"},

		// Bit Manipulation
		{Title: "Single Number", TitleSlug: "single-number", Difficulty: models.Easy, Pattern: "Bit Manipulation", URL: "https://leetcode.com/problems/single-number/"},
		{Title: "Number of 1 Bits", TitleSlug: "number-of-1-bits", Difficulty: models.Easy, Pattern: "Bit Manipulation", URL: "https://leetcode.com/problems/number-of-1-bits/"},
		{Title: "Counting Bits", TitleSlug: "counting-bits", Difficulty: models.Easy, Pattern: "Bit Manipulation", URL: "https://leetcode.com/problems/counting-bits/"},
		{Title: "Reverse Bits", TitleSlug: "reverse-bits", Difficulty: models.Easy, Pattern: "Bit Manipulation", URL: "https://leetcode.com/problems/reverse-bits/"},
		{Title: "Missing Number", TitleSlug: "missing-number", Difficulty: models.Easy, Pattern: "Bit Manipulation", URL: "https://leetcode.com/problems/missing-number/"},
		{Title: "Sum of Two Integers", TitleSlug: "sum-of-two-integers", Difficulty: models.Medium, Pattern: "Bit Manipulation", URL: "https://leetcode.com/problems/sum-of-two-integers/"},
		{Title: "Reverse Integer", TitleSlug: "reverse-integer", Difficulty: models.Medium, Pattern: "Bit Manipulation", URL: "https://leetcode.com/problems/reverse-integer/"},
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO problems (title, title_slug, difficulty, pattern, url, due_date, status)
		VALUES (?, ?, ?, ?, ?, NULL, 'new')
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range problems {
		if _, err := stmt.Exec(p.Title, p.TitleSlug, string(p.Difficulty), p.Pattern, p.URL); err != nil {
			return fmt.Errorf("insert %s: %w", p.TitleSlug, err)
		}
	}

	return tx.Commit()
}

// GetSession retrieves the stored LeetCode session cookie from config.
func GetSession(db *sql.DB) (string, error) {
	var val sql.NullString
	err := db.QueryRow("SELECT value FROM config WHERE key = 'leetcode_session'").Scan(&val)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return val.String, nil
}

// SetSession stores the LeetCode session cookie.
func SetSession(db *sql.DB, session string) error {
	_, err := db.Exec(`INSERT INTO config (key, value) VALUES ('leetcode_session', ?)
		ON CONFLICT(key) DO UPDATE SET value = excluded.value`, session)
	return err
}
