The LambdaNEW programming language, runtime, REPL and VM

## LambdaNEW programming language, runtime, REPL and VM

### About LambdaNEW

As the name might suggest, LambdaNEW is an experiment to recreate the LambdaMOO experience with today's technology and possibilities. The mobile internet, the wide adoption of messenger apps and ideas from the blockchain community (I see you rolling your eyes ...) influence LambdaNEW. 

Like the old LambdaMOO, it is built around the idea of a relatively small *runtime* i.e. the server or `Node` and an extensible *database* and *programming language*.


### Background and some history

LambdaMOO is/was a network-accessible, multi-user, programmable, interactive system well-suited to the construction of text-based adventure games, conferencing systems, and other collaborative software. Its most common use, however, is as a multi-participant, low-bandwidth virtual reality, or in 2021s terms, 'a metaverse'.

Participants (usually referred to as players) used to connect to LambdaMOO using Telnet or some other, more specialized, client program. Upon connection, they were usually presented with a welcome message explaining how to either create a new character or connect to an existing one. Characters are the embodiment of players in the virtual reality that is LambdaMOO.

**Note:** _MOOs predate the current app and smartphone centric internet by at least two decades and no one really connects to a MOO over telnet these days. Today, the most common client would be either a web app or maybe an extension/plug-in into a common messenger app like Telegram or Slack._

Having connected to a character, players then give one-line commands that are parsed and interpreted by LambdaMOO as appropriate. Such commands may cause changes in the virtual reality, such as the location of a character, or may simply report on the current state of that reality, such as the appearance of some object.

Example:

```shell
go north
look
say "Hello everyone !"
````

The job of interpreting those commands is shared between the two major components in the LambdaMOO system: the server and the database. The server is a program, written in a standard programming language, that manages the network connections, maintains queues of commands and other tasks to be executed, controls all access to the database, and executes other programs written in the MOO programming language. The database contains representations of all the objects in the virtual reality, including the MOO programs that the server executes to give those objects their specific behaviors.

Almost every command is parsed by the server into a call on a MOO procedure, or verb, that actually does the work. Thus, programming in the MOO language is a central part of making non-trivial extensions to the database and thus, the virtual reality.

**Note:** _Most people tend to associate the term "database" with a "relational database", and most relational database programs tend to keep most of their data in disk storage. The purely technical meaning of "database" is "an organized collection of information." LambdaMOO's database is not relational, it is an object database, and it is kept entirely in-memory. The only reason I'm pointing this out is to head off any chance of you confusing the moo database for a relational database._


## Overview

## Basic Concepts in LambdaNEW

**Note:** `LambdaNEW` is the name of the project and the entire *system*. `NU` or `Nu` is the name of the programming language and runtime.

### Object Orientation

MOO stands for "Mud, Object-Oriented". MOO's object-oriented approach is slightly different from many object-oriented programming langauges. In most object-oriented languages, there is a division between object definition (the class) and instances of the object in the system. The object definition exists in the source code, but the executable code never deals with them directly. Instead you create an "instance" of a given class and use this instance.

In the MOO world, the object is defined by example. You create an object instance in the system and then dynamically add verbs and properties to make your "prototype". Then you can create a new object that is a "decendant" from the original object. The new object in turn can be dynamically modified with more verbs and properties, and then you can create more objects that descend from the second object.

NU implements MOO's approach to OO.

### Objects, Verbs and Properties

* NU consists of objects.
* NU objects have properties and verbs.
* NU implements a simple, generic programming language to manipulate objects.


`Properties` are values stored on the object and referenced in code via the "." operator, for example:

```shell
someVariable = object.propertyName
```

`Verbs` are code stored on the object and invoked in code via the "." operator, for example:

```shell
someVariable = object.verbName()
```


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
* Actor, with following sub-types
*   * Person
*   * Bot
*   * Agent
*   * Twin

`Node` is the `root` object of the NU instance/runtime. `Node` contains all other objects in the runtime.

A `Place` represents a logical or abstract location e.g. a room or building.

A Place can have a `Location`. Location represent a reference to the place in the physical world, e.g. by providing an address or some coordinates.

A `Thing` represents anything in a Node that can be interacted with.

`Actor` as object type is never instantiated directly but represents an abstract type for objects that interact with other objects in the Node. `Actor` can be one of the following:

* `Person`, representing a natural person or legal entity
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


### Object Identity and Ownership of objects

* Node and Person have an `Identity`
* Identity is represented by a unique, unmutable public/private key pair.
* Node and Person have an `address` derieved from their public key.


* All objects have an `Owner`
* All objects have a `Creator`
* Node and Person are their own owner
* All other types of objects are by default owned by the Node or Person who created them.
* An object's creator is unmutable while the object's owner might change over time.


### Object migration

* Place, Thing, Person, Bot and Agent can *migrate* from one Node to another Node
* Node and Twin **CAN NOT** migrate from one Node to another Node
* Objects contained in another bbject migrate with the parent object
* Access controlls, object permissions or ownership rules might prevent object migration


### Access controll, object permissions or ownership rules

TBD