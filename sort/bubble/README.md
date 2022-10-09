### Algorithm
- Starts with very first two elements
- Compare them and check which one is greater
- Swap the elements as needed
- Double for loop
- Each loop ends at the second last element

```
begin BubbleSort(list)

  for all elements of list
    if list[i] > list[i+1]
      swap(list[i], list[i+1])
    end if
  end for

  return list

end BubbleSort
```