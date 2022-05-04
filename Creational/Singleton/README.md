# Pattern Singletone

Singleton is a creational pattern(object level).

It is helpful in cases where we need to have exactly **one** instance of something, e.g.
database or service.

Singleton guarantees that there is only one instance of some object.

In Go Singleton typically created using `sync.Once`. 
It is good because it is both lazy and threadsafe.
([Example](./singleton.go))