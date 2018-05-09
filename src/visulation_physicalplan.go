package plan

import (
	"strings"
	"fmt"
	"reflect"
)


var physicalNodeNameMap = make(map[string]string)

func PhysicalPlanAnalysis(sql string,physical PhysicalPlan) {


	if(!strings.Contains(sql, "#uml")) {
		return
	}

	fmt.Printf("PhysicalPlan Analysis start********************************************************************\n")

	fmt.Println("@startuml")
	physicalNodeNameMap = make(map[string]string)
	physicalPlanDetailAnalysis("",physical);
	fmt.Println("@enduml")


	fmt.Printf("PhysicalPlan Analysis end********************************************************************\n")

}

func physicalPlanDetailAnalysis( parentName string,physical PhysicalPlan) {


	name:=physical.ExplainID();
	nodeType:=reflect.TypeOf(physical).String()
	if _,ok :=physicalNodeNameMap[name];!ok{

		fmt.Println("object " + name)
		fmt.Println(name + " : type = " + nodeType)

		if(nodeType!="*plan.PhysicalSelection"&&nodeType!="*plan.PhysicalTableScan"&&nodeType!="*plan.PhysicalIndexScan") {
			fmt.Println(name + " : task = root")
		}else{
			fmt.Println(name + " : task = cop")
		}
		//fmt.Println(name + " : info = "+physical.ExplainInfo())

		physicalNodeNameMap[name]=name
	}


	if(parentName!=""){
		fmt.Println(parentName+"-->"+name)
	}


	switch nodes:=physical.(type){
	   case *PhysicalTableReader:
		   for _,childLogic := range nodes.TablePlans {
			   physicalPlanDetailAnalysis(name,childLogic)

		   }
		   break
	   case *PhysicalIndexLookUpReader:
		   for _,childLogic := range nodes.TablePlans {
			   physicalPlanDetailAnalysis(name,childLogic)
		   }

		   for _,childLogic := range nodes.IndexPlans {
			   physicalPlanDetailAnalysis(name,childLogic)
		   }
		   break

	case *PhysicalIndexReader:
		for _,childLogic := range nodes.IndexPlans {
			physicalPlanDetailAnalysis(name,childLogic)
		}
		break



	default:
		for _, childLogic := range physical.Children() {
			physicalPlanDetailAnalysis(name, childLogic)

		}
	}
}
