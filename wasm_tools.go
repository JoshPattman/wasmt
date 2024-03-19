// package wasmt simplifies the process of inserting wasm into go std templates, using gin
package wasmt

import (
	"fmt"
	"html/template"

	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed wasm_run_script.js
var wasmRunScript string

//go:embed wasm_runtime.js
var wasmRuntimeScript string

var wasmRuntimeUrl = ""

// Setup configures the router to automatically server require wasm stuff.
// It will serve the wasm runtime at /wasm_runtime.js
func Setup(r *gin.Engine) {
	SetupWithRuntimeURL(r, "/wasm_runtime.js")
}

// SetupWithRuntimeURL configures the router to automatically server require wasm stuff.
// It will serve the wasm runtime at url
func SetupWithRuntimeURL(r *gin.Engine, url string) {
	wasmRuntimeUrl = url
	r.GET(wasmRuntimeUrl, func(c *gin.Context) {
		c.Header("Content-Type", "application/javascript")
		c.String(200, wasmRuntimeScript)
	})
}

// NewTemplate is a drop-in replacement for template.New(name),
// but it adds support for embedding wasm easily.
// Use {{insertWasmRuntime}} somewhere in the template,
// before calling {{insertWasmScript "/path/to/script.wasm"}} to insert the wasms scripts you want.
func NewTemplate(name string) *template.Template {
	return template.New(name).Funcs(template.FuncMap{
		"insertWasmRuntime": insertWasmRuntime,
		"insertWasmScript":  insertWasmScript,
	})
}

func insertWasmRuntime() template.HTML {
	return template.HTML(fmt.Sprintf("<script src=%s></script>", wasmRuntimeUrl))
}

func insertWasmScript(scriptURL string) template.HTML {
	script := fmt.Sprintf(wasmRunScript, scriptURL)
	return template.HTML(fmt.Sprintf("<script>%s</script>", script))
}
