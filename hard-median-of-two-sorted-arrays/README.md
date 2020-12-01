# Median of Two Sorted Arrays [Hard]

### Task
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
**Follow up:** The overall run time complexity should be O(log (m+n)).

**Example 1:**

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
```

**Example 2:**

```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

**Example 3:**

```
Input: nums1 = [0,0], nums2 = [0,0]
Output: 0.00000
```

**Example 4:**

```
Input: nums1 = [], nums2 = [1]
Output: 1.00000
```

**Example 5:**

```
Input: nums1 = [2], nums2 = []
Output: 2.00000
```

**Constraints:**

```
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
```


### Solution
We know that arrays are sorted, therefore complexity of O(log(m+n)) shouldn't bother us.
First idea that I thought of was merging two arrays and then find the median. But wait, we know
that median is in the center of this imaginary array, so we just need first half ot it and get 
one or elements to calculate the median.
If you look at examples you will find out that we when the length of both arrays is even we need 
to find 2 elements, and when its length is odd then 1 element. In the case of even length the
second element always has the index which is equal to `totalLength / 2` and the first one has index 
equal to `totalLength / 2 - 1`. For example for the merged array:

```
[1,2,3,4]
 0 1 2 3
   | |
```
Length is `4` and the index of the second element equal `4/2=2`. Thus, we need element `3` and `2`.

For the merged array with odd length, for instance with length `3` there is only one element which is 
the median itself. 

```
[1,2,3]
 0 1 3
   |
```
Length is `3` and the index of the median element equal `3/2=1`.

Knowing that let's write the code to get those values:

```go
totalLen := len(nums1) + len(nums2)
halfLen := totalLen / 2
mod := totalLen % 2
poss := make(map[int]bool)
```

Now we should remember which indexes should be summed up to get the median:

```go
if mod == 0 {
    poss[halfLen] = true
    poss[halfLen-1] = true
} else {
    poss[halfLen] = true
}
```

The length of two arrays might not be equal, therefore we should prepare two variables
to iterate through them:

```go
var i1, i2 int
```

The variable to store the sum of the median elements, because it could be one or two of them:

```go
var sum float64
```

We prepared everything to iterate through two arrays and get elements we need. Remeber, that we 
need only the element with index of `halfLen` and element before it if the total length of the 
imaginary merged array is even.
Main logic:

```go
// we start the iteration to the middle of our imaginary 
// merged array
for i := 0; i <= halfLen; i++ {
    var cur int
    // both elements might become element with index i
    if i1 < len(nums1) && i2 < len(nums2) {
        // choose the smallest one and increment relevant
        // index by 1
        if nums1[i1] < nums2[i2] {
            cur = nums1[i1]
            i1++
        } else {
            cur = nums2[i2]
            i2++
        }
    // the second array is empty at this moment
    // so we work only with first array
    } else if i1 < len(nums1) {
        cur = nums1[i1]
        i1++
    // the first array is empty at this moment
    // so we work only with first array
    } else {
        cur = nums2[i2]
        i2++
    }
    // check if we reached the median indexes
    // if true add the elements to the sum
    if _, ok := poss[i]; ok {
        sum += float64(cur)
    }
}
``` 

We got the sum of median elements. Just return the result:

```go
return sum/float64(len(poss))
```