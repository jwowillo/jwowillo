## solver

solver assists in solving and visualizing artificial intelligence problems.

It is organized into problems and agents which solve different problems. This
allows general problems to be defined and sent to the correct agent.

### State

`State = Object`

State of the Problem.

### Action

`Action = Object`

Action to travel from one State to another.

### Solution

`Solution = Object`

Solution to a Problem.

### Problem

```
Problem Interface
  Initial() State
  Actions(State) List<Action>
  Result(State, Action) State
  IsGoal(State) boolean
```

Problem is anything which has an Initial state, Actions available at States,
Results of applying an Action to a State, and a State deemed good enough to
return true with IsGoal.

### Agent

```
Agent Interface
  Solve(Problem) Solution
```

Agent solves a Problem by returning a Solution. A nil Solution means one doesn't
exist. The interface is defined so that other Agents can follow the pattern by
convention.

### UninformedProblem

```
UninformedProblem Interface
  Problem
  Cost(State, State, Action) Float
```

UninformedProblem is a Problem where Actions to get from one State to another
have Cost but there is no concept of being closer or further from a goal State.

### InformedProblem

```
InformedProblem Interface
  UninformedProblem
  Heuristic(State) Float
```

InformedProblem is an UninformedProblem with the added stipulation that there is
some concept of distance to goal State.

### KnownGoalProblem

```
KnownGoalProblem Interface
  UninformedProblem
  Predecessors(State) List<State>
  Goal() State
```

KnownGoalProblem is a Problem where the Goal State is known and Predecessors to
States can be found.
