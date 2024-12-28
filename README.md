# Go Slice Modification in Range Loop
This example demonstrates a common error in Go where modifications to a slice inside a `for...range` loop do not affect the original slice.  This is because `range` creates a copy of the slice element. 

**Bug:** The `myFunc` function intends to modify the elements of the input slice, but it fails to do so because it modifies a copy. 

**Solution:** Use a `for` loop with an index to iterate and modify the original slice.