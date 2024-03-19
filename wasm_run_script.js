const go = new Go();
WebAssembly.instantiateStreaming(fetch("%s"), go.importObject)
    .then((result) => {
        go.run(result.instance);
    });