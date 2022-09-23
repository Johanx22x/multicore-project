package menu

import (
	"fmt"
    "os"
    "os/exec"

	"github.com/Johanx22x/multicore-project/internal/jsonm"
	"github.com/Johanx22x/multicore-project/internal/scrap"
)

const WORKERS = 15

func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func dataInfo() {
    fmt.Println(`
╔═══╗      ╔╗          ╔══╗     ╔═╗    
╚╗╔╗║     ╔╝╚╗         ╚╣╠╝     ║╔╝    
 ║║║║╔══╗ ╚╗╔╝╔══╗      ║║ ╔═╗ ╔╝╚╗╔══╗
 ║║║║╚ ╗║  ║║ ╚ ╗║      ║║ ║╔╗╗╚╗╔╝║╔╗║
╔╝╚╝║║╚╝╚╗ ║╚╗║╚╝╚╗    ╔╣╠╗║║║║ ║║ ║╚╝║
╚═══╝╚═══╝ ╚═╝╚═══╝    ╚══╝╚╝╚╝ ╚╝ ╚══╝`)
}

func thankYou() {
    fmt.Println(`
╔═══╗                             ╔╗       ╔╗        
║╔═╗║                             ║║      ╔╝╚╗       
║╚══╗╔══╗╔══╗    ╔╗ ╔╗╔══╗╔╗╔╗    ║║ ╔══╗ ╚╗╔╝╔══╗╔═╗
╚══╗║║╔╗║║╔╗║    ║║ ║║║╔╗║║║║║    ║║ ╚ ╗║  ║║ ║╔╗║║╔╝
║╚═╝║║║═╣║║═╣    ║╚═╝║║╚╝║║╚╝║    ║╚╗║╚╝╚╗ ║╚╗║║═╣║║ 
╚═══╝╚══╝╚══╝    ╚═╗╔╝╚══╝╚══╝    ╚═╝╚═══╝ ╚═╝╚══╝╚╝ 
                 ╔═╝║                                
                 ╚══╝                                `)
}

func welcomeMessage() {
    fmt.Println(`
╔═╗╔═╗    ╔╗  ╔╗                      ╔═══╗                          
║║╚╝║║    ║║ ╔╝╚╗                     ║╔═╗║                          
║╔╗╔╗║╔╗╔╗║║ ╚╗╔╝╔╗╔══╗╔══╗╔═╗╔══╗    ║╚══╗╔══╗╔═╗╔══╗ ╔══╗╔╗╔═╗ ╔══╗
║║║║║║║║║║║║  ║║ ╠╣║╔═╝║╔╗║║╔╝║╔╗║    ╚══╗║║╔═╝║╔╝╚ ╗║ ║╔╗║╠╣║╔╗╗║╔╗║
║║║║║║║╚╝║║╚╗ ║╚╗║║║╚═╗║╚╝║║║ ║║═╣    ║╚═╝║║╚═╗║║ ║╚╝╚╗║╚╝║║║║║║║║╚╝║
╚╝╚╝╚╝╚══╝╚═╝ ╚═╝╚╝╚══╝╚══╝╚╝ ╚══╝    ╚═══╝╚══╝╚╝ ╚═══╝║╔═╝╚╝╚╝╚╝╚═╗║
                                                       ║║        ╔═╝║
                                                       ╚╝        ╚══╝`)
}

func loading() {
    fmt.Println(`
╔╗              ╔╗                
║║              ║║                
║║   ╔══╗╔══╗ ╔═╝║╔╗╔═╗ ╔══╗      
║║ ╔╗║╔╗║╚ ╗║ ║╔╗║╠╣║╔╗╗║╔╗║      
║╚═╝║║╚╝║║╚╝╚╗║╚╝║║║║║║║║╚╝║╔╗╔╗╔╗
╚═══╝╚══╝╚═══╝╚══╝╚╝╚╝╚╝╚═╗║╚╝╚╝╚╝
                        ╔═╝║      
                        ╚══╝      `)
}

func loaded() {
    fmt.Println(`
╔╗              ╔╗      ╔╗
║║              ║║      ║║
║║   ╔══╗╔══╗ ╔═╝║╔══╗╔═╝║
║║ ╔╗║╔╗║╚ ╗║ ║╔╗║║╔╗║║╔╗║
║╚═╝║║╚╝║║╚╝╚╗║╚╝║║║═╣║╚╝║
╚═══╝╚══╝╚═══╝╚══╝╚══╝╚══╝`)
}

func displayOptions() {
    fmt.Println("2 - Obtain CSV table")
    fmt.Println("3 - Obtain metadata from websites")
    fmt.Println("4 - Display how much websites are in English or Spanish")
    fmt.Println()
}

func getInput() (opt int, err error) {
    _, err = fmt.Scan(&opt)
    if err != nil {
        err = fmt.Errorf("Please enter a valid option!")
        return 0, err
    }
    return
}

func Menu() {
    clearScreen()
    welcomeMessage()
    for {
        fmt.Println("What do you want to do? (0 exit / 1 show options)")

        opt, err := getInput()
        if err != nil {
            clearScreen()
            welcomeMessage()
            fmt.Printf("\n%s\n\n", err)
            displayOptions()
            continue
        }

        switch opt {
        case 0:
            clearScreen()
            thankYou()
            fmt.Println("\nThank you for use our system.\n\nHave a nice day!")
            return
        case 1:
            clearScreen()
            welcomeMessage()
            displayOptions()
        case 2:
            clearScreen()
            loading()
            fmt.Println("\nLoading table...")
            scrap.ObtainCSV()
            clearScreen()
            loaded()
            fmt.Println("\nTable loaded!")
            fmt.Println()
        case 3:
            clearScreen()
            loading()
            fmt.Println("\nObtaining metadata...")
            jsonm.SaveMetadata(WORKERS)
            clearScreen()
            loaded()
            fmt.Println("\nMetadata obtained!")
            fmt.Println()
        case 4:
            clearScreen()
            dataInfo()
            jsonm.ShowWebsitesInfo()
        }
    }
}
