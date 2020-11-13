
## Builder Design Pattern

### Motivation
1. Some objects are simple and you can just create them using a constructor call
   or you can just call them by initializing them by a simple way.
2. Other objects require a lot of ceremony.
3. Having a factory function with 10 args is not productive.
   which means we are asking user to initialize so many things by himself.
*  => Instead we should construct the object piece wise rather than in a single factory call.
4. so the builder pattern provides an API for constructuing the object step-by-step rather
   than doing everything at once.

#### Summary
* When piecewise object constructuion is complicated, provide an API for doing it
  succinctly. This ensure the construction is understandable and not that complicated which happens all at once.

