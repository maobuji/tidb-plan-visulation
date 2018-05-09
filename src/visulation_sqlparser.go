package plan

import (
	"strings"
	"github.com/pingcap/tidb/ast"
	"fmt"
	"reflect"
	"github.com/pingcap/tidb/plan/visulation/util"
)


var number int=0
var nodeMap = make(map[ast.Node]string)


var objectSlice =make([]string,0)
var relationSlice=make([]string,0)

var nodeStack=stack.NewStack()

type visitor struct{}

func (v *visitor) Enter(in ast.Node) (out ast.Node, skipChildren bool) {


	if _,ok :=nodeMap[in];!ok{
		number++
		nodeName:= fmt.Sprintf("%s_%d", reflect.TypeOf(in),number)
		nodeName=strings.Replace(nodeName, "*", "", -1)
		nodeMap[in]=nodeName
		objectSlice =append(objectSlice, nodeName)

	}

	if(nodeStack.Len()!=0){
		parent :=nodeStack.Peak().(ast.Node)

		if parentName,ok :=nodeMap[parent];ok{
			if childName,ok :=nodeMap[in];ok{
				relationString:=fmt.Sprintf("%s-->%s\n",parentName,childName)
				relationSlice= append(relationSlice,relationString)
			}
		}
	}

	nodeStack.Push(in)
	return in, false
}

func (v *visitor) Leave(in ast.Node) (out ast.Node, ok bool) {

	if(nodeStack.Peak()==in){
		nodeStack.Pop()
	}

	return in, true
}


func NodeAnalysis(sql string,node ast.StmtNode)  {

	if(!strings.Contains(sql, "#uml")) {
		return
	}

	if _,ok :=nodeMap[node];!ok{
		number++
		nodeName:= fmt.Sprintf("%s_%d", reflect.TypeOf(node),number)
		nodeName=strings.Replace(nodeName, "*", "", -1)
		nodeMap[node]=nodeName
		objectSlice =append(objectSlice, nodeName)

	}


	v := visitor{}
	node.Accept(&v)


	fmt.Printf("Node Analysis start********************************************************************\n")

	fmt.Printf("@startuml\n")

	for _, value := range objectSlice {

		fmt.Printf("object %s\n", value)

	}

	objectSlice=make([]string,0)

	for _, value := range relationSlice {

		fmt.Printf("%s", value)

	}

	relationSlice=make([]string,0)
	nodeMap = make(map[ast.Node]string)
	nodeStack.Empty()
	number=0

	fmt.Printf("@enduml\n")
	fmt.Printf("Node Analysis end********************************************************************\n")

}

