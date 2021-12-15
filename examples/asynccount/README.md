## Async Count – an asynchronous Counter

<img alt="UI for an anychronous counter" src="http://npillmayer.github.io/UAX/img/async-counter.png"
    width="220">

This example is about finding a way to offload domain-model actions. Operations on domain objects generally are too costly to be executed on the rendering-goroutine. We need an expressive way of offloading these and have a transparent event routing model to keep user-level code understandable. We strive for compactness and readability for the component's code.

Pressing the button will simulate a delay typical for an expensive operation on domain level. The UI should remain responsive, until the proessing is done and the count is updated.
