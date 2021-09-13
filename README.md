## NU - The LambdaNEW programming language, runtime, REPL and VM

LambdaNEW and NU, it's programming language and runtime, inherit many ideas from the original LambdaMOO [1][2] system. A lot has changed since LambdaMOO's origin in the mid/late 1990's and this project is an experiment to re-implement many of those ideas with today's technology.

What this repo is:

* A definition of NU, the LambdaNEW programming language
* An implementation of a REPL for NU
* The NU runtime (and maybe VM, tbd)

What this is NOT:

* a blockchain/coin/NFT/Game/Metaverse
* a node/server implementing the above

## Overview

## Basic Terminology

### Objects, Verbs and Properties

NU consists of objects.

NU objects have properties and verbs.

Properties are values stored on the object and referenced in code via the "." operator, for example:

```shell
someVariable = object.propertyName
```

Verbs are code stored on the object and invoked in code via the "." operator, for example:

```shell
someVariable = object.verbName()
```

### Object Orientation

MOO stands for "Mud, Object-Oriented". MOO's object-oriented approach is slightly different from many object-oriented programming langauges. In most object-oriented languages, there is a division between object definition (the class) and instances of the object in the system. The object definition exists in the source code, but the executable code never deals with them directly. Instead you create an "instance" of a given class and use this instance.

In the MOO world, the object is defined by example. You create an object instance in the system and then dynamically add verbs and properties to make your "prototype". Then you can create a new object that is a "decendant" from the original object. The new object in turn can be dynamically modified with more verbs and properties, and then you can create more objects that descend from the second object.

NU follows MOO's approach to OO.

### Types of Objects

All objects implement the same basic concepts:

* Objects have Verbs and Properties
* Objects can contain other objects
* Each object has exactly one parent object, except for the `root` object
* Each object has a set of defined (built-in) attributes and capabilities that can not be removed or modified.

Each object is of one of the following types:

* Node
* Place
* Location
* Thing
* Actor, with sub-type Person, Bot, Agent or Twin

`Node` is the `root` object of the NU instance/runtime. Node contains all other objects in the runtime.

A `Place` represents a logical or abstract location e.g. a room or building.

A Place can have a `Location`. Location represent a reference to the place in the physical world, e.g. by providing an address or some coordinates.

A `Thing` represents anything in a Node that can be interacted with.

`Actor` as object type is never instantiated directly but represents an abstract type for objects that interact with other objects in the Node. `Actor` can be one of the following:

* `Person`, representing a real person or user
* `Bot`, a local software agent
* `Agent`, a bot, acting on behave of a `Person`
* `Twin`, a *projection* of a Person from an external Node into the current Node

### Containment rules

* All objects except Node have exactly 1 parent object
* Objects can not be contained in more than one object (same as above rule, just to be sure ...)
* Node contains 0..n Places, Things and Actors
* Place contains 0..n Places, Things and Actors
* Place has 0 or 1 Location
* Location can not contain other objects
* Thing contains 0..n Things
* Person contains 0..n  Things
* Person has to be contained in a Place
* Bot, Agent and Twin can not contain other objects
* Bot, Agent and Twin can be in Node or in Place


### Object migration

* Place, Thing, Person, Bot and Agent can `migrate` from one Node to another Node
* Node and Twin `CAN NOT` migrate from one Node to another Node
* Objects contained in another Object migrate with the parent object
* Access controlls, object permissions or ownership rules might prevent object migration


### Access controll, object permissions or ownership rules

TBD


## References

* [1: ProgrammersManual.html](http://www.moo-cows.com/docs/manuals/ProgrammersManual.html)
* [2: moo-programmers-manual-updated.md](https://github.com/lambdanew/lambda-moo-programming/blob/master/tutorials/moo-programmers-manual-updated.md)