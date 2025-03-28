package taskmanager

import "fmt"

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

var tasks = []Task

var idCounter  int

//creacion de cada funcion

func createTask() {
	
	var title string
	var description string
	fmt.Println("Ingrese el titulo de la tarea: ")
	fmt.Scanln(&title)
	fmt.Println("Ingrese la descripcion de la tarea: ")
	fmt.Scanln(&description)

	idCounter++
	newTask := Task{
		ID: idCounter,
		Title: title,
		Description: description,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	fmt.Println("Tarea creada exitosamente!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas disponibles")
		return
	}

	fmt.Println("Tareas disponibles:")
	for _, task := range tasks {

		status := "Pendiente"
		if task.Completed {
			status = "Completada"
		}
		fmt.Printf("ID: %d, Titulo: %s, Descripcion: %s, Estado: %s\n", task.ID, task.Title, task.Description, status)
		
			
		}
}

func updateTask() {
	var id int
	fmt.Println("Ingrese el ID de la tarea a actualizar: ")
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id {
			task[i].Completed = !task[i].Completed
			fmt.Println("Tarea actualizada exitosamente!")
			return
		}
	}
	fmt.Println("Tarea no encontrada")
}

func deleteTask() {
	var id int
	fmt.Println("Ingrese el ID de la tarea a eliminar: ")
	fmt.Scanln(&id)

	for i, task := range tasks {
		
	}
	if task.ID == id {
		tasks = append(tasks[:i], tasks[i+1:]...)
		fmt.Println("Tarea eliminada exitosamente!")
		return
	}
	fmt.Println("Tarea no encontrada")
}

func main() {
	fmt.Println("MENU PRINCIPAL")

	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1. Crear tarea")
		fmt.Println("2. Listar tareas")
		fmt.Println("3. Actualizar tarea")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")

		fmt.Print("Seleccione una opcion: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createTask()
		case 2:
			listTasks()
		case 3:
			updateTask()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
