## containers

containers contains implementations of common data structures.

### Queue

```
NewQueue() Queue

Queue
  Push(Object)
  Pop() Object
  Len() Integer
```

Queue is a first-in first-out data structure. Push places an Object at the back
of the Queue and Pop removes and returns the item at the front of the Queue. Len
returns the number of elements in the Queue.

### Stack

```
NewStack() Stack

Stack
  Push(Object)
  Pop() Object
  Len() Integer
```

Stack is a last-in first-out data structure. Push places an Object on the top
of the Stack and Pop removes and returns the item on the top of the Stack. Len
returns the number of elements in the Stack.

### Set

```
NewSet() Set

Set
  Add(Object)
  Remove(Object)
  Contains(Object) Boolean
  Items() List<Object>
  Len() Integer
```

Set is a unique collection of Objects. Add places a new Object in the Set if it
doesn't already exist in the Set. Remove removes the Object from the Set if it
exists in the Set. Contains returns true iff the Object is in the Set. Items
returns a List<Object> of Objects in the Set. Len returns the number of Objects
in the Set.

### Matrix

```
NewMatrix(Integer, Integer) Matrix
NewMatrix(Integer) Matrix

Matrix = List<Vector>
  Row(Integer) Vector
  Column(Integer) Vector
  String() String

Multiply(Matrix, Matrix) (Matrix, Error)
Transpose(Matrix) Matrix
```

Matrix is an mxn array of Float values. NewMatrix makes a Matrix with
dimensions mxn m is the first Integer and n is the second Integer. 
NewSquareMatrix makes an nxn Matrix.

Rows and Columns of Matrices can be accessed be position, returning a Vector.
Matrices also have String representations.

An nxm Matrix can be transposed to an mxn Matrix. A mxy Matrix can be
multiplied with a yxn Matrix to give an mxn Matrix.

### Vector

```
Vector = List<Float>

InnerProduct(Vector, Vector) (Float, Error)
```

Vector represents a row in a Matrix. Can take the InnerProduct of two Vectors
if they have the same length.

### PriorityQueue

```
NewPriorityQueue() PriorityQueue

PriorityQueue
  Push(Object, Float)
  Pop() Object
  Len() Integer
```

PriorityQueue is a Queue where the lowest priority elements are always at the
front. Push places the Object in the PriorityQueue with the given priority.

### DisjointSets

```
NewDisjointSets(Integer) DisjointSets

DisjointSets
  Union(Integer, Integer)
  Find(Integer)
  Sets() List<Set>
```

DisjointSets is a collection of Sets which share no Objects. Union merges the
Sets containing the two Integers. Find finds the representative of the Set
containing the Integer. Sets returns a List of the Sets which comprise the
DisjointSets.
