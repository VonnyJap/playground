#### Pass by Value
- the function makes a local copy
- the original data is not changed
- example: the data is used for calculation but it is not desired to store it back to the original variable

#### Pass by Reference
- using pointer
- the value of original data that get passed by reference can be changed

#### Notes
- When a pointer is assigned to an address of a var, it is just an address that we referring to
- When the pointer is changed - it is actually changing pointer to point to another address
- In this case it will not change the value of the original data

https://www.digitalocean.com/community/conceptual_articles/understanding-pointers-in-go