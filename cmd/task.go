package cmd

import (
	"fmt"
	"strings"
	"github.com/mbndr/figlet4go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
    gorm.Model
	Id     uint
    Tittle     string  // Exported fields should start with an uppercase letter
    Description string  
    IsFinished bool
}

func dbDispenser() *gorm.DB{
	db_name := "root"
	db_password := "Pass@123"
	dsn := db_name+":"+db_password+"@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//fmt.Println(dsn)
	db , error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error!=nil {
		fmt.Println("failed to connect to database")
	}
	db.AutoMigrate(&Task{})
	return db
}

func add(tittle string, desc string) string{
	db := dbDispenser()
	db.Create(&Task{Tittle: tittle,Description: desc,IsFinished: false})
	message := "Task added sucessfully!"
	return message
}

func viewall() {
	var tasks []Task
	db := dbDispenser()
	db.Find(&tasks)
	for i, task:= range tasks{
		ascii := figlet4go.NewAsciiRender()
		options := figlet4go.NewRenderOptions()
		options.FontColor = []figlet4go.Color{
			// Colors can be given by default ansi color codes...
			figlet4go.ColorGreen,
			figlet4go.ColorYellow,
			figlet4go.ColorCyan,
			// ...or by an hex string...
			//figlet4go.NewTrueColorFromHexString("885DBA"),
			// ...or by an TrueColor object with rgb values
			//figlet4go.TrueColor{136, 93, 186},
		}

		
		// The underscore would be an error
		var b strings.Builder
    	b.WriteString(fmt.Sprintf("%d", i+1))
		num := "Task "+b.String()
		fmt.Printf("\n------------------------------------------------\n")
		renderStr, _ := ascii.RenderOpts(num, options)
		fmt.Print(renderStr)
        fmt.Printf("|| ID: %d\n", task.Id)
        fmt.Printf("|| Tittle: %s", task.Tittle)
        fmt.Printf("|| Description: %s", task.Description)
        fmt.Printf("|| Completed?: ")
		if task.IsFinished {
        	fmt.Printf(" Yep!")
		}
		if !task.IsFinished {
			fmt.Printf(" Nope!")
		}
		fmt.Printf("\n------------------------------------------------\n")

	}
}

func deleteT(id int) {
	db := dbDispenser()
	var task Task
	db.Delete(&task, id)
}

func update(id int) string {
	var message string
	db := dbDispenser()
	var task Task
	db.First(&task, "id = ?", id)
	db.Model(&task).Update("IsFinished",true)
	return message
} 



