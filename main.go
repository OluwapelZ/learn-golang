// package main

// // import "fmt"
// import (
// 	_"github.com/go-chi/chi"
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello World!")
// }

package main

import (
	"fmt"
    "strconv"
    "github.com/astaxie/beego"
)


func main() {
    /* This would match routes like the following:
       /sum/3/5
       /product/6/23
       ...
    */
	beego.Router("/health", &healthCheckController{})
    beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
    beego.Run()
}

type healthCheckController struct {
	beego.Controller
}

type mainController struct {
    beego.Controller
}

func (c *healthCheckController) Get() {
	fmt.Print("Hello World!");
	c.Data["result"] = "Hello World!";
}


func (c *mainController) Get() {

    //Obtain the values of the route parameters defined in the route above    
    operation := c.Ctx.Input.Param(":operation")
    num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
    num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

    //Set the values for use in the template
    c.Data["operation"] = operation
    c.Data["num1"] = num1
    c.Data["num2"] = num2
    c.TplName = "result.html"

    // Perform the calculation depending on the 'operation' route parameter
    switch operation {
    case "sum":
        c.Data["result"] = add(num1, num2)
    case "product":
        c.Data["result"] = multiply(num1, num2)
    default:
        c.TplName = "invalid-route.html"
    }
}

func add(n1, n2 int) int {
    return n1 + n2
}

func multiply(n1, n2 int) int {
    return n1 * n2
}