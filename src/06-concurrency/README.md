# Concurrency
In this section, you will learn about what makes Go special. Goroutines and channels are the building blocks of the Go concurrency model. Upon completion of this section you should have a basic understanding of the following:

* Goroutines
* Channels - Blocking, and buffered.
* Channels with range - Channels can be a closable stream of objects.
* Select - It’s kind of like a switch statement for channels.

There are additional synchronization primitives and types contained in the `sync` and `sync/atomic` packages. In general, you should only use these packages for low-level code and prefer the higher level abstraction provided by the built-in goroutines and channels.
