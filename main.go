package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)
var mu sync.Mutex
var wg sync.WaitGroup

type Task struct{
	ID int
	Name string
	Status string
}

var tasks []Task

var TaskQueue = make(chan *Task) // create a channel named TaskQueue that can only parse and store Task type data 

func addTask(name string) {
	mu.Lock()
	defer mu.Unlock()



	task:= Task{
		ID: len(tasks)+1,
		Name : name,
		Status : "Pending",
	}	
	tasks = append(tasks,task)
}	

func worker(id int, ctx context.Context, tasks <-chan *Task){
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task.ID)
		mu.Lock()
		task.Status = "Running"
		mu.Unlock()

		time.Sleep(2*time.Second)
		
		mu.Lock()
		task.Status = "Completed"
		mu.Unlock()
		fmt.Printf("Worker %d finished task %d\n", id, task.ID)


		wg.Done()
	}

}

func main(){
	fmt.Println("This is a task manager . Welcome v0")


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i:=1; i<=3; i++ {
		go worker(i,ctx,TaskQueue)
	}

	scanner := bufio.NewScanner(os.Stdin)


	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()

		if line == ""{
			continue
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
			mu.Lock()
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
				mu.Unlock()
				continue
			}
			for _, task := range tasks {
				fmt.Printf("%d: %s [%s]\n", task.ID, task.Name, task.Status)
			}

			mu.Unlock()
			/////////////////////////////////////////////////////////////////////

		case "run" :
			mu.Lock()
			if len(tasks) == 0{
				fmt.Println("There are no tasks to continue")
				mu.Unlock()
				continue
			}
			taskSnapshot := make([]*Task, 0, len(tasks))

			for i := range tasks{
				taskSnapshot = append(taskSnapshot, &tasks[i])
				wg.Add(1)
			}
			mu.Unlock()

			for _, task := range taskSnapshot{
						TaskQueue <- task
			}

			fmt.Println("task dispatched to workers")

			wg.Wait()

			fmt.Println("All tasks Completed")
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

