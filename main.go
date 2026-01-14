package main
import(
	"fmt"
	"os"
	"time"
	"bufio"
	"strings"
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




	for i:=1; i<=3; i++ {
		go worker(i,TaskQueue)
	}

	scanner := bufio.NewScanner(os.Stdin)


	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()

		if line == ""{
			continue;
		}

		parts := strings.Fields(line)
		command := parts[0]

		switch command {
		case "add" :
			if len(parts) < 2{
				fmt.Println("Usage :  Add <task_name>")
				continue
			}
			addTask(parts[1])
			fmt.Println("Task added")

			/////////////////////////////////////////////////////////////


		case "list" :
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
				continue
			}
			for _, task := range tasks {
				fmt.Printf("%d: %s [%s]\n", task.ID, task.Name, task.Status)
			}
			/////////////////////////////////////////////////////////////////////

		case "run" :
			if len(tasks) == 0{
				fmt.Println("There are no tasks to continue")
				continue
			}
			for i := range tasks{
				TaskQueue <- &tasks[i]
			}

			fmt.Println("task dispatched to workers")


			////////////////////////////////////////////////////////////////////
			//
			//

			case "exit":
				fmt.Println("Goodbye!")
				return

		default:
			fmt.Println("Unrealised command : ",command)
		}
	}
}

