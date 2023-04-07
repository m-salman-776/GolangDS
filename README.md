# GolangDS

### Why Go is fast ###
Statically typed which helps to mitigate run time errors. Fastest compilation time as compared to other language.

***Go vs Java***
Java compiles on VM its code need to be converted to byte code before passing through compiler which is slower process.
Go don't rely on VM for code compilation it directly compiled from binary which is why it is much faster.

***Automatic Garbage Collection***

***Wait Group*** are for synchronisation b/w go routines
Whenever new variable created it takes memory and CPU cycles large number of unused variables can cause latency in the system and make garbage collection difficult for the CPU.

### Memory Management in Go ###

***new()***
allocated but not initialised
u get memory address
zeroed storage - u cant put any data

***make()***
allocated but not initialised
u get memory address
zeroed storage - u put any data


***Understanding Allocation*** : Stack and Heap 

Sharing down typically stays on stack : passing pointer in function

Sharing up typically escape to heap : returning pointers

typically ? because only compiler knows

When possible the Go compiler will allocate variables that are location to a function in that function's stack frame
However if the compiler can not prove that the variable is not referenced after the function returns then the 
compiler must allocate the variable on the heap to avoid  dangling pointers

***When values are constructed on the heap***
When values possibly be referenced after the function that constructed the values returns.
When the compiler determines a value is too large to fit on the stack.
When the compiler doesn't know the size of values at compile time.
Value stored in interface variables