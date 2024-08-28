# [WASM demo (click to run 🚀)](https://zeozeozeo.github.io/ebitengine-microui-go/demo)

Demo of running Ebitengine + ebitengine-microui-go on the web.

# Building this demo

1. Clone the repository
2. Navigate into the demo directory: `cd examples/demo`
3. Build the demo for WebAssembly:

    On Linux:

    ```
    env GOOS=js GOARCH=wasm go build -o demo.wasm .
    ```

    On Windows PowerShell:

    ```
    $Env:GOOS = 'js'
    $Env:GOARCH = 'wasm'
    go build -o yourgame.wasm .
    Remove-Item Env:GOOS
    Remove-Item Env:GOARCH
    ```

4. Copy `wasm_exec.js` into the current directory:

    On Linux:

    ```
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
    ```

    On Windows PowerShell:

    ```
    $goroot = go env GOROOT
    cp $goroot\misc\wasm\wasm_exec.js .
    ```

5. Create this HTML file

    ```html
    <!DOCTYPE html>
    <script src="wasm_exec.js"></script>
    <script>
        // Polyfill
        if (!WebAssembly.instantiateStreaming) {
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
                const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }

        const go = new Go();
        WebAssembly.instantiateStreaming(
            fetch("demo.wasm"),
            go.importObject
        ).then((result) => {
            go.run(result.instance);
        });
    </script>
    ```

6. Start a local HTTP server and open the page in your browser

If you want to embed the game into another page, use iframes (assuming that `main.html` is the name of the above HTML file):

```html
<!DOCTYPE html>
<iframe src="main.html" width="640" height="480" allow="autoplay"></iframe>
```
