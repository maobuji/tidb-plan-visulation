package plan

import (
	"fmt"
	"strings"
	"reflect"
	"github.com/pingcap/tidb/plan/visulation/util"
)






var logicNodeStack=stack.NewStack()


func LogicalPlanAnalysis(sql string,logic LogicalPlan) {


	if(!strings.Contains(sql, "#uml")) {
		return
	}

	fmt.Printf("LogicalPlan Analysis start********************************************************************\n")



	fmt.Println("@startuml")
	logicalPlanDetailAnalysis("",logic);
	fmt.Println("@enduml")


	fmt.Printf("LogicalPlan Analysis end********************************************************************\n")

}

func logicalPlanDetailAnalysis(parentName string,logicPlan LogicalPlan) {

    name:=logicPlan.ExplainID();
	fmt.Println("object "+name)
	fmt.Println(name+" : type = "+reflect.TypeOf(logicPlan).String())

	if(parentName!=""){
		fmt.Println(parentName+"-->"+name)
	}

	for _,childLogicPlan := range logicPlan.Children() {
		logicalPlanDetailAnalysis(name,childLogicPlan)
	}

}
