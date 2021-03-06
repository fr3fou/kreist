# kreist

![](./demo.png)

A simple car game with Raylib (WIP)

## TODO

_in no particular order_

- [x] Proper level rendering (right now it renders an exported image from Tiled)
  - [x] ~~Implement custom raylib renderer for `lafriks/go-tiled`~~ not going to be needed
- [ ] Collision layer
- [ ] Refactor
- [ ] Scale car / world
- [ ] Adjust steering to be more realistic
  - [x] Fixed wheelbase length
- [ ] Braking
- [ ] Different surfaces should have different physics
- [x] Better name - thanks https://github.com/impzero
- [ ] Sounds
- [ ] More Levels
- [ ] Multiplayer
- [ ] Nicknames
- [ ] Main Menu
- [ ] Level selection
- [ ] Car selection
  - [ ] Different cars should have different specifications
    - [ ] Store them in JSON
- [ ] Logo
- [x] CI Building
- [ ] Replace dt computation with `rl.GetFrameTime()`

## References

- <http://engineeringdotnet.blogspot.com/2010/04/simple-2d-car-physics-in-games.html>
- <http://kidscancode.org/godot_recipes/2d/car_steering/>
- <https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#object>
- <https://www.toptal.com/game/video-game-physics-part-i-an-introduction-to-rigid-body-dynamics>
- <https://www.toptal.com/game/video-game-physics-part-ii-collision-detection-for-solid-objects>
