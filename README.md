# DefKit
Component based scripts for [Defold](www.defold.com) engine inspired by the  [Game Maker](http://www.yoyogames.com/gamemaker) studio visual scripting.

---

## [Events]((https://docs.yoyogames.com/source/dadiospice/000_using%20gamemaker/events/index.html)
Events are input from different sources to which you can react and connect one or many actions. These range from user-input such as keys being pressed, instance lifecycle such as create and destroy, or collision betwen different instances to name a few.

### General
* on_create : invoked directly on create of the object.
* on_destroy : invoked when the object is told to be removed.

### Input
These scripts allow users to connect logic to input, as long as the referenced action_id is defined in the input definition file.

* on_button : repeated for as long as the button is held down.
* on_button_pressed : invoked once as the button is pressed down.
* on_button_released : invoked once as the button is released after being pressed.

### Contact
These scripts allow users to connect logic to [contact between objects](http://www.defold.com/manuals/physics/#_collision_objects), as long as the objects both have collision-objects and masks are correctly setup.

* on_collision : repeated as long as the object matching the collision group is in contact.

### Time
These script allow users to connect logic to events related to time, either user-defined alarms or the normal update-cycle.

* on_alarm : invoked when the set amount of seconds have passed since the user-defined timer was started.
* on_Step : invoked each cycle/frame.


## [Actions](https://docs.yoyogames.com/source/dadiospice/000_using%20gamemaker/actions/index.html)
Actions are small blocks of logic that allow interraction in a multitude of ways; movement, creation and destruction to name a few.

__Shared Information__

Event : All Actions come with an "event" url parameter, this is how you connect which Event is the cause of this Action.

Target : Many but not all Actions also come with a Self, Other or Object Type parameter

Relative: Let the set value be influenced by context based values, depending on Action this could be worl-position or existing velocity among other things.

Variables: Many of the input-values also come with a "Var" option which is short for Variable. These enable passing dynamic values, either predefined or controlled using set_variable.

### Basic
* create : tell a Factory to create one instance of it's targetted game-object, at the given location.
* destroy : tell the targetted game-object to be destroyed instantly.

### Move
The actions let the user change position over time for targetted object, if set to relative the existing movement-speed will be taken in-to account.

* move_fixed : given one or more directions, North, West, South, East or diagonals inbetween. Randoms between the given directions and move at the assigned speed.
* move_free : given an angle value, move in that direction at the given speed.

### Basic2

* set_alarm : tell the specified alarm event to fire after a set amount of seconds.
* set_score : set the current score value to the defined amount, or influence the existing by relative.
* set_variable : set the specified variable to the defined number value, or influence the existing by relative.
* draw_score : render the current score at the specified coordinate on the screen.

### Controll

* test_chance : roll a dice with set amount of sides to see if it gives a 1, then invoke other actions.
* test_instance_count : count the amount of objects to see if they are "equal (0)", "less than (-1)" or "more than(1)" the specified amount, then invoke other actions.


## Examples
