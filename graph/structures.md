## Structures

Structures which are useful in defining and working with graphs.

### Vertex

```
NewVertex(Label) Vertex

Vertex
  Label() Label
```

Vertex represents a single point in a graph. Vertices have a label which they
are identified by outside of Graphs.

### Colors

```
Color enum
  Navy
  Orange
  Red
  Black
  Green
  Yellow
```

Color which a structure may be assigned. Colors are useful for keeping track of
information about parts of a Graph or finding structure within the Graph.

### VertexSet

```
NewVertexSet() VertexSet

VertexSet
  Add(Vertex)
  Contains(Vertex) Boolean
  Remove(Vertex)
  Len() Integer
  Items() List<Vertex>
```

VertexSet is a set of Labels.

### Edge

```
NewEdge(Vertex, Vertex) Edge
NewWeightedEdge(Vertex, Vertex, Float) Edge

Edge
  Start() Vertex
  End() Vertex
  Weight() Float
```

Edge is a connection spanning two Vertices. Can be either directed or undirected
depending on context. New Edges can be created which are either unweighted or
weighted.

### EdgeSet

```
NewEdgeSet(VertexSet) EdgeSet

Edgeset
  Add(Edge) Error
  Contains(Edge) Boolean
  Remove(Edge)
  Len() Integer
  Items() List<Edge>
```

EdgeSet is a set of Edges. Must be constrained by a VertexSet so that Edges may
only go between existing Vertices. An Error is raised if Add is given an Edge
which spans to a Vertex which doesn't exist in the VertexSet.

### VertexEdgePair

```
NewVertexEdgePair(v Vertex, e Edge) VertexEdgePair

VertexEdgePair
  Edge() Edge
  Vertex() Vertex
```

VertexEdgePair is a structure which represents information about one end of an
Edge. Specifically, it contains the Edge it is describing along with a Vertex
on one end of the Edge.

### Graph

```
New(VertexSet, EdgeSet) Graph
NewDirected(VertexSet, EdgeSet) Graph

Graph
  VertexSet() VertexSet
  EdgeSet() EdgeSet
  AdjacencyMatrix() containers.Matrix
  IntsToLabels() Map<Integer, Label>
  LabelsToInts() Map<Label, Integer>
  AdjacentTo(Label) VertexSet
  IncidentTo(Label) EdgeSet

```

Graph is defined as a VertexSet and an EdgeSet. The VertexSet defines the
Vertices in the Graph and the EdgeSet defines the connections amongst Vertices.
Undirected Graphs are created with New while DirectedGraphs are created with
NewDirected.

VertexSet and EdgeSet return the sets used to build the Graph. AdjacencyMatrix
and WeightsMatrix return the Graph's internal representation. This
representation involves mapping Vertex labels to Integer positions in the
matrices. The mapping from Labels to Integers and from Integers to Labels can be
retrieved with IntsToLabels and LabelsToInts. Use AdjacentTo to find Vertices
adjacent to the Vertex with the given Label. Use IncidentTo to find Edges
incident to the Vertex with the given Label. SubgraphInducedBy a list of Labels
returns the subgraph containing the Vertices with those Labels, along with
incident Edges travelling only to other Vertices in the subgraph.
