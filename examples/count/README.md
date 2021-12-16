## Count – a simple Counter

<img alt="UI for a simple counter" src="http://npillmayer.github.io/UAX/img/simple-counter.png"
    width="220">

This little example is mainly about exploring a way of expressing the visual representation in a more concise way.

### Building the View Tree

Gio is very flexible and it is possible to do *anything* related UI graphics with it. However, building the visual tree assembles nested creation of objects and function calls (often closures), which quickly becomes tedious.

One way of simplifying tedious nested function calls is to make is possible to chain them. So instead of 

```Go
    Layout(context) dimensions {
        property1 := …
        container{axis: V}.Layout(context, property1) {
            widget1(context),
            func(context) dimensions {
                property2 := …
                text := "…"
                widget2(context, property2, text)
            }
        }
    }
```

we may prefer typing something like this:

```Go
    layout().Content(
        container().Axis(V).Prop(…).Content(
            widget1(),
            widget2().Prop(…).Text("…")
        )
    )
```

(*Both snippets in pseudo-code.*)

Making it easier for users to create the widget tree requires reducing complexity by reducing options. It is necessary to strike a balance between flexibility and expressiveness. The HTML/CSS model is widely used and may be an adequate set of functionality. A requirement is, however, to always make it easy to opt out and fall back to Gio's full capacity.

### Encapsulation

The counter box is a user-level component. It encapsulates widget behaviour and visuals in one place. Its interface is provided by a UI delegate binding to the domain object (a simple `int` in this case).


