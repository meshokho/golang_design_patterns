# Pattern Adapter

Adapter is a **structural** pattern(class-level).

It allows to use existing code, e.g. class with an interface that is unappropriated for the new system.
In order to not change the existing class, it is possible to create an Adapter.
Adapter has similar function as a real-world adapters(USB Type-C to USB-Mini for example).

It takes an existing interface Target and a new interface Adaptee and somehow connects them between each other.

Simple [[example]](adapter.go)