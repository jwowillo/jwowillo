## Testing

Tests are performed on labelled known graphs. Different features are checked or
operations are performed and they are compared to the labelled values. Tests
fail if the values do not match.

### Syntax

Tests are added by adding txt files to the graphs directory. Each txt file gets
converted to a test case. The txt files are meant to describe graphs, and have
the syntax

```
name:
<name>

vertex set:
{<vertex 1 label>, ...}

edge set:
{<vertex 1 label> <edge> <vertex 2 label>, ...}
// edge can be --- (?weight), --> (?weight) where the first signifies an
// undirected edge with optional weight and the seceond signifies a directed
// edge with optional weight.

adjacency matrix:
                 <vertex 1 label> <vertex 2 label> ...
<vertex 1 label>         ?                ?
<vertex 2 label>         ?                ?
.
.
.

adjacencies:
<vertex 1 label> {<adjacent vertex label>, ...}
.
.
.

incidences:
<vertex 1 label> {<vertex 1 label> <incident edge> <adjacent vertex label>, ...}
.
.
.

subgraph induced by:
{<vertex 1 label>, ...} {<vertex a label> <edge> <vertex b label>, ...}
.
.
.

```

Graph examples below describe graphs using a syntax like to the one described.

### Complete Graph

Complete graph on four vertices, or K<sub>4</sub>.

![Complete Graph]({{ cdn }}/graphs/complete_graph.png)

Adjacencies:

* a: b, d
* b: a, c, d
* c: b, d
* d: a, b, c

Adjacency Matrix:

|        |  a  |  b  |  c  |  d  |
|  ---   | --- | --- | --- | --- |
|<b>a</b>|  0  |  1  |  0  |  1  |
|<b>b</b>|  1  |  0  |  1  |  1  |
|<b>c</b>|  0  |  1  |  0  |  1  |
|<b>d</b>|  1  |  1  |  1  |  0  |

### Empty Graph

Empty graph with two vertices and no edges.

![Empty Graph]({{ cdn }}/graphs/empty_graph.png)

Adjacencies:

* a:
* b:

Adjacency Matrix:

|        |  a  |  b  |
|  ---   | --- | --- |
|<b>a</b>|  0  |  0  |
|<b>b</b>|  0  |  0  |

### Isolated Vertex Graph

Graph with one vertex attached to C<sub>4.

![Isolated Vertex Graph]({{ cdn }}/graphs/isolated_vertex_graph.png)

Adjacencies:

* a: b
* b: a, c, e
* c: b, d
* d: c, e
* e: b, d

Adjacency Matrix:

|        |  a  |  b  |  c  |  d  |  e  |
|  ---   | --- | --- | --- | --- | --- |
|<b>a</b>|  0  |  1  |  0  |  0  |  0  |
|<b>b</b>|  1  |  0  |  1  |  0  |  1  |
|<b>c</b>|  0  |  1  |  0  |  1  |  0  |
|<b>d</b>|  0  |  0  |  1  |  0  |  1  |
|<b>e</b>|  0  |  1  |  0  |  1  |  0  |

### Loop Graph

Non simple graph with a loop.

![Loop Graph]({{ cdn }}/graphs/loop_graph.png)

Adjacencies:

* a: a, b
* b: a

Adjacency Matrix:

|        |  a  |  b  |
|  ---   | --- | --- |
|<b>a</b>|  2  |  1  |
|<b>b</b>|  1  |  0  |

### Parallel Graph

Graph with parallel edges between two vertices.

![Parallel Graph]({{ cdn }}/graphs/parallel_graph.png)

Adjacencies:

* a: b, b
* b: a, a

Adjacency Matrix:

|        |  a  |  b  |
|  ---   | --- | --- |
|<b>a</b>|  0  |  2  |
|<b>b</b>|  2  |  0  |

### Unconnected Directed Graph

Directed graph where not all vertices are reachable from other vertices.

![Unconnected Directed Graph]({{ cdn }}/graphs/unconnected_directed_graph.png)

Adjacency Matrix:

|        |  a  |  b  |  c  |  d  |
|  ---   | --- | --- | --- | --- |
|<b>a</b>|  0  |  0  |  0  |  0  |
|<b>b</b>|  1  |  0  |  0  |  0  |
|<b>c</b>|  1  |  0  |  0  |  1  |
|<b>d</b>|  0  |  1  |  0  |  0  |
