# Pattern Prototype

Prototype is a creational pattern(object level).

It allows to create new objects by copying of early-created objects - **prototypes**.
The problem is that if you're copying some struct, you need to take care of pointers,
because changes in copy may affect original(prototype).

There is a smart and easy way to do this in Go. It is possible to use encoders/decoders, 
e.g. byte or JSON to automatically create deep copies of an object. ([Example](./prototype.go))

Also is it possible to use prototype interface ([Example](./prototypeWithInterface.go)).