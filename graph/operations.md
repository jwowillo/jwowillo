## Operations

Many useful operations can be performed on graphs. These range from linear or
constant time algorithms to attempted solutions at NP-hard problems.

### MaximalDegree

`MaximalDegree(Graph) Ingeter`

MaximalDegree returns the degree of the Vertex with the highest degree in the
Graph. The degree of a graph is defined as the number of edges incident to it.

### MinimalDegree

`MaximalDegree(Graph) Ingeter`

MinimalDegree returns the degree of the Vertex with the lowest degree in the
Graph.

### IsTree

`IsTree(Graph) Boolean`

IsTree returns true iff the given Graph has no cycles.

### IsBipartite

`IsBipartite(Graph) Boolean`

IsBipartite returns true iff the given Graph has a bipartite partitioning.

### BipartitePartitioning

`BipartitePartitioning(*Graph) (VertexSet, VertexSet, error)`

BipartitePartitioning returns a bipartite partitioning of the Graph as two
VertexSets if one exists, and an ErrNotBipartite error otherwise.

A bipartite partitioning is a division of a Graph's Vertices into two sets such
that no two Vertices in the same set share an Edge.

### IsConnected

`IsConnected(Graph) Boolean`

IsConnected returns true iff the given Graph is connected. As defined here, to
be connected, every Vertex must be reachable from every other Vertex using
existing edges.

### IsPlanar

`IsPlanar(Graph) Boolean`

IsPlanar returns true iff the Graph is planar. This means the Graph can be drawn
on a plane without requiring any Edges to overlap.

### IsSimple

`IsSimple(Graph) Boolean`

IsSimple returns true iff the Graph is simple. For a Graph to be simple, it
can't have any loops, parallel edges, be weighted, or directed.

### IsWeighted

`IsWeighted(Graph) Boolean`

IsWeighted returns true iff the Graph is weighted. This means it contains at
least one edge with a cost associated with it.

### ChromaticNumber

`ChromaticNumber(Graph) Integer`

ChromaticNumber returns the lowest number of colors needed to color the Vertices
of the Graph such that no two adjacent Vertices share the same color.

### WalkCount

`WalkCount(Graph, Label, Label, Integer) Integer`

WalkCount between the Vertices represented by the given Labels in the Graph. The
walks are required to be the given length.

### Path

`Path(Graph, Label, Label) List<Label>`

Path between Vertices represented by the given Labels in the Graph represented
as a list of intermediary Vertices on the path. An error is returned if no path
exists.

### Complement

`Path(Graph) Graph`

Complement returns the Graph created by deleting all of the given Graph's Edges
from the complete Graph with the same number of Vertices as the given Graph.
Edge weights are ignored.

### SubgraphInducedBy

`SubgraphInducedBy(Graph, List<Label>) Graph`

SubgraphInducedBy returns a Graph which is a subgraph of the given one
containing the Vertices represented by the given List of Labels along with all
edges incident to Vertices in the List of Labels.
