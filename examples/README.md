## Examples with custom Gio Components

Every example has a main program. Each of them will rely on one or more custom components to solve a (simple) UI task.

The nomenclature somewhat resembles the HTML/CSS/JS model. One reason for that is that I find it easier to reason about the workings of Gio thinking of a browser UI. But by no means is the nomenclature relevant for having a user-level component model.

* The [count](https://github.com/npillmayer/giocomp/tree/main/examples/count)/[asynccount](https://github.com/npillmayer/giocomp/tree/main/examples/asynccount) examples have trivial functionality. Their purpose is to get a feeling for the challenges of encapsulating components with Gio.

* The [todo](https://github.com/npillmayer/giocomp/tree/main/examples/todo) example is a re-implementation of an app implemented in the course of [a great Mozilla-tutorial](https://developer.mozilla.org/en-US/docs/Learn/Tools_and_testing/Client-side_JavaScript_frameworks/Svelte_getting_started) for Svelte. 

For each example, subfolders will contain the component's code. Some will as well have a subfolder for “domain objects”, which should never be infected with UI code.

### Approach

All examples are certain to contain a lot of performance-bugs in terms of how one would efficiently use Gio. However, for now this is of no importance to me, as my focus is on user-level expressiveness. That said, the underlying model should respect Gio's idea that rendering-related function should be zero-alloc, i.e. the “DOM” isn't instantiated but rather a hierarchy of function calls. I modfied it slightly with constructing an expressive tree of functions first which are then are called in a second step.