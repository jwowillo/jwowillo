## search

search contains Agents which solve problems by searching for a solution.

### Solution

`Solution = List<solver.Action>`

Solution to search Problems is a sequence of solver.Actions which, if performed
in order, will traverse from a start solver.State to an end solver.State.

### Agent

```
Agent interface
  Solve(solver.Problem) Solution
```

Agent Solves solver.Problems by returning Solutions.

### UninformedAgent

```
UninformedAgent interface
  Solve(solver.UninformedProblem) Solution
```

UninformedAgent Solves solver.UninformedProblems by returning Solutions.

### InformedAgent

```
InformedAgent interface
  Solve(solver.InformedProblem) Solution
```

InformedAgent Solves solver.InformedProblems by returning Solutions.

### KnownGoalAgent

```
KnownGoalAgent interface
  Solve(solver.KnownGoalProblem) Solution
```

KnownGoalAgent Solves solver.KnownGoalProblems by returning Solutions.

### BreadthFirstAgent

```
NewBreadthFirstAgent() BreadthFirstAgent

BreadthFirstAgent
  Solve(solver.Problem) Solution
```

BreadthFirstAgent is an Agent which Solves solver.Problems by searching all
solver.States at a certain search-tree depth before proceeding to the next
depth.

### DepthFirstAgent

```
NewDepthFirstAgent() DepthFirstAgent

DepthFirstAgent
  Solve(solver.Problem) Solution
```

DepthFirstAgent is an Agent which Solves solver.Problems by searching an entire
path in the search-tree before trying another one.

### DepthLimitedAgent

```
NewDepthLimitedAgent(Integer) DepthLimitedAgent

DepthLimitedAgent
  Solve(solver.Problem) Solution
```

DepthLimitedAgent is an Agent which Solves solver.Problems in the same way as
the DepthFirstAgent, but gives up after reaching a fixed depth.

### IterativeDeepeningAgent

```
NewIterativeDeepeningAgent() IterativeDeepningAgent

IterativeDeepeningAgent
  Solve(solver.Problem) Solution
```

IterativeDeepeningAgent is an Agent which Solves solver.Problems by performing
like a DepthLimitedAgent with depth-limits increasing from zero to infinity
until a goal is found.

### UniformCostAgent

```
UniformCostAgent() UniformCostAgent

UniformCostAgent
  Solve(solver.UninformedProblem) Solution
```

UniformCostAgent Solves solver.UninformedProblems by expanding solver.States
with the cheapest path cost first.

### GreedyBestFirstAgent

```
GreedyBestFirstAgent() GreedyBestFirstAgent

GreedyBestFirstAgent
  Solve(solution.InformedProblem) Solution
```

GreedyBestFirstAgent Solves solver.InformedProblems by expanding solver.States
with the cheapest heuristic costs first.

### AStarAgent

```
NewAStarAgent() AStarAgent

AStarAgent
  Solve(solution.InformedProblem) Solution
```

AStarAgent Solves solver.InformedProblems by expanding solver.States with the
cheapest combination of path and heuristic costs first.
