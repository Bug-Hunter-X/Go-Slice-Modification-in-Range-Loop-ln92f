func myFunc(a []int) { 
  for i := 0; i < len(a); i++ {
    a[i] *= 2 // Modifies the original slice
    fmt.Println(a[i])
  }
} 