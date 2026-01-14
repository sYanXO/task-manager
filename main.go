package main
import(
	"fmt"
	"os"
	"time"
)
type Task struct{
	ID int
	Name string
	Status string
}

var tasks []Task

var TaskQueue = make(chan *Task) // create a channel named TaskQueue that can only parse and store Task type data 

func addTask(name string) {
		task:= Task{
				ID: len(tasks)+1,
				Name : name,
				Status : "Pending",
		}	
		tasks = append(tasks,task)
}	

func worker(id int, tasks <-chan *Task){
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task.ID)
		task.Status = "Running"
		

	time.Sleep(2*time.Second)

	task.Status = "Completed"
	fmt.Printf("Worker %d finished task %d\n", id, task.ID)
}

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

	for i:=1; i<=3; i++ {
		go worker(i,TaskQueue)
	}

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
		if len(tasks) == 0{
			fmt.Println("There are no tasks to run")
			return
		}
		for i := range tasks{
				TaskQueue <- &tasks[i]
			}

			fmt.Println("task dispatched to workers")

			time.Sleep(5*time.Second)
////////////////////////////////////////////////////////////////////
	default:
		fmt.Println("Unrealised command : ",command)
	}
}

