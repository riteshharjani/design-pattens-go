## Motivation
1. For components may go in and out of a system at any time meaning it could be
   created or destroyed.
a. Chat room participants where people can come in and go out.
b. Players in an MMORPG

2. It makes no sense for them to have direct references(pointers) to one another
   Pointer references from one to another should not be made.
a. since, those references may all go dead

### Solution:
1. Have them all refer to some central component that facilitates communication

### In Short
A component that faciilitates communication between other components w/o them
necessarily being aware of each other or having direct links(references) to one
another.

### Sumarry
* Create the mediator as kind of a central component, and then each object in the system
  points to the mediator.
a. e.g. assign a field in factor function
* Mediator engages in bidrectional communication with it's connected components.
* Mediator has methods the components can call
* Components have methods the mediator can call
* Event processing (e.g. Rx) libs make communication easier to implement
