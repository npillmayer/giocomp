## Experiments for Components in Gio UI

[Gio](https://gioui.org/) is a great UI-framework in Go.  It follows a paradigm of [immediate mode](https://gioui.org/doc/architecture) UI, an approach which is supposed to make it easier to reason about state changes.

For more complex UI layouts it is advantageous to be able to factor the UI into manageable entities. [Svelte](https://svelte.dev/), a recent framework for Javascript, does this for the reactive UI paradigm, and using it feels pretty comfortable. I like Svelte's approach for a couple of reasons:

* Locality of code for behaviour and looks of a component
* Direct correlation between view delegate objects and display
* Sane event propagation model
* All kinds of smart helpers

With Svelte, much of this is supported by a compiler, which produces efficient code to alter the DOM. In Go code-generation is common, but I prefer being able to write expressive code on top of smart base objects.

This repo is a test if expressive user-level UI-code can be written for Gio. This is a super rough first sketch, and I will eventually build upon it to grow into something useful.

### Examples

The user-level code resides in the [examples folder](https://github.com/npillmayer/giocomp/tree/main/examples), together with a description of what I learned implementing them.
