## ++x
x is incremented prior to assignment

```
// If y = ++x, the variable x will be incremented before assigning its value to y.
var x = 0;
var y = 0;
function preIncrement(){
    while(x < 3){
      y =  ++x;
        console.log("y = " + y);
    }
      x++;
}
preIncrement();
```
```
y = 1
y = 2
y = 3
```
```
let x = 0;
console.log(x); //outputs 0
console.log(++x); //outputs 1
```

## x++
x is incremented after the assignment

```
// If y = x++, the variable x will be incremented after assigning its value to y.
var x = 0;
var y = 0;
function postIncrement(){
    while(x < 3){
      y =  x++;
        console.log("y = " + y);
    }
      x++;
}
postIncrement();
```
```
y = 0
y = 1
y = 2
```
```
let x = 0;
console.log(x); //outputs 0
console.log(x++); //outputs 0
```