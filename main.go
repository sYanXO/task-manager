package main
import(
	"fmt"
	"os"
)
type Task struct{
	ID int
	Name string
}

var tasks []Task

func addTask(name string) {
		task:= Task{
				ID: len(tasks)+1,
				Name : name,
		}	
		tasks = append(tasks,task)
}	



func main(){
	fmt.Println("This is a task manager . Welcome v0")



	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: add | list | run")
		return
	}

	command := args[1]
	fmt.Println("command :",command)

	switch command {
		case "add" :
			if len(args) < 3 {
			fmt.Println("Usage: add <task_name>")
			return
		
		}
	addTask(args[2])
	fmt.Println("Task added")

/////////////////////////////////////////////////////////////


	case "list" :
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
				return
	}
	for _, task := range tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Name)
	}
/////////////////////////////////////////////////////////////////////
		
	case "run" :
		fmt.Println("Running tasks (sequentially)")
		for _, task := range tasks {
			fmt.Println("Executing:", task.Name)
		}
////////////////////////////////////////////////////////////////////
	default:
		fmt.Println("Unrealised command : ",command)
	}
}

