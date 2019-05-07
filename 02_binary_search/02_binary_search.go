package main

import "fmt"

func main() {
  // int_slice := []int{1,2,10,15,20}
  int_slice := []int{1,2,3,4,5,6,7,8,9,10,12,23,34,45,56,67,78,89,90,123,234,345,456,567,678,789,890}
  // int_slice := []int{1,2,3,4,5,6,6,10,15,15,17,18,19,30}
  rec_find := recBinarySearch(int_slice, 10, 0, len(int_slice)-1)
  loop_find := loopBinarySearch(int_slice, 10, 0, len(int_slice)-1)
  fmt.Println("recBinarySearch found: ", rec_find, "loopBinarySearch found:", loop_find)
}

func recBinarySearch(subject []int, target, low, high int) int {
  if high <= low { return -1 }
  idx := int((high+low)/2)
  if subject[idx] > target { return recBinarySearch(subject, target, low, idx) }
  if subject[idx] < target { return recBinarySearch(subject, target, idx+1, high) }
  return idx
}

func loopBinarySearch(subject []int, target, low, high int) int {
  idx := int((high+low)/2)
  for {
    if high <= low { return -1 }
    if subject[idx] == target { break }
    if subject[idx] > target { high = idx }
    if subject[idx] < target { low = idx+1 }
    idx = int((high+low)/2)
  }
  return idx
}
