### Algorithm
- if it is only one element in the list it is already sorted, return.
- divide the list recursively into two halves until it can no more be divided.
- merge the smaller lists into new list in sorted order.

### Pseudocode

<u>recursively divide and then merge</u>
```
procedure mergesort( var a as array )
  if ( n == 1 ) return a

  var l1 as array = a[0] ... a[n/2]
  var l2 as array = a[n/2+1] ... a[n]

  l1 = mergesort( l1 )
  l2 = mergesort( l2 )

  return merge( l1, l2 )
end procedure
```

<u>merge logic</u>
- this works because the smallest group is sorted
- hence when we do comparison, we are sure that the following will be larger
- it will sort on its own after getting more elements
```
procedure merge( var a as array, var b as array )

  var c as array
  while ( a and b have elements )
    if ( a[0] > b[0] )
      add b[0] to the end of c
      remove b[0] from b
    else
      add a[0] to the end of c
      remove a[0] from a
    end if
  end while

  while ( a has elements )
    add a[0] to the end of c
    remove a[0] from a
  end while

  while ( b has elements )
    add b[0] to the end of c
    remove b[0] from b
  end while

  return c

end procedure
```