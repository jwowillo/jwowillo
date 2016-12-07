## Errors

Several errors can occur while working with graphs. They indicate graphs lack
certain properties or something impossible was attempted with them.

<b>`ErrBadEdge`</b>: Error is raised if an edge was added to an EdgeSet which
contained Vertices not included in the EdgeSet's VertexSet.

<b>`ErrNotBipartite`</b>: Error is raised if a bipartite partitioning of a Graph
is requested but the Graph is not bipartite.
