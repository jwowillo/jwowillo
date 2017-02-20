# Land Grab

`Land Grab` is a turn-based strategy game where a player competes against other
players by using units and towers to claim the most land. Artificial
intelligence is emphasized and multiple players are supported. Navigating turns
is doable with key-controls only so turns may end quickly for experienced
players. Taking shorter turns is rewarded.

The game will be separated into two major components: the front and back ends.
The front-end is a JavaScript client which uses canvases to show game state and
allow inputs. The front-end communicates to the back-end with websockets. The
back-end is a `trim` web-application which exposes websocket endpoints to manage
game state and make AI-player decisions.

The game and players are defined in a package separate from the web-application.
AI-player implementations which fulfill a `Player` interface are provided in
another separate package. Algorithms used by AI-players are provided in a
package in another separate package.

All packages are acceptance-tested. AI-players are benchmarked since speed will
be an aspect of their value. The game is unit tested since if the game-rules are
correct, players will be forced to only make valid moves. Acceptance-testing for
the client will involve defining and checking set of views which should be
manually observed.

Angular shouldn't be used for this project. Use
https://github.com/lean/phaser-es6-webpack as example. Refactor into seed.
