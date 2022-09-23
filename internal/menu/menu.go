package menu

import (
	"fmt"
    "os"
    "os/exec"

	"github.com/Johanx22x/multicore-project/internal/jsonm"
	"github.com/Johanx22x/multicore-project/internal/scrap"
)

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

func sorryMessage() {
    fmt.Println(`
╔═══╗                     
║╔═╗║                     
║╚══╗╔══╗╔═╗╔═╗╔╗ ╔╗      
╚══╗║║╔╗║║╔╝║╔╝║║ ║║      
║╚═╝║║╚╝║║║ ║║ ║╚═╝║╔╗╔╗╔╗
╚═══╝╚══╝╚╝ ╚╝ ╚═╗╔╝╚╝╚╝╚╝
               ╔═╝║       
               ╚══╝       `)
}

func workersMessage() {
    fmt.Println(`
╔═╗ ╔╗              ╔╗╔╗╔╗       ╔╗             
║║╚╗║║              ║║║║║║       ║║             
║╔╗╚╝║╔══╗╔╗╔╗╔╗    ║║║║║║╔══╗╔═╗║║╔╗╔══╗╔═╗╔══╗
║║╚╗║║║╔╗║║╚╝╚╝║    ║╚╝╚╝║║╔╗║║╔╝║╚╝╝║╔╗║║╔╝║══╣
║║ ║║║║║═╣╚╗╔╗╔╝    ╚╗╔╗╔╝║╚╝║║║ ║╔╗╗║║═╣║║ ╠══║
╚╝ ╚═╝╚══╝ ╚╝╚╝      ╚╝╚╝ ╚══╝╚╝ ╚╝╚╝╚══╝╚╝ ╚══╝ 
 `)
}

func displayOptions() {
    fmt.Println("2 - Obtain CSV table")
    fmt.Println("3 - Obtain metadata from websites")
    fmt.Println("4 - Analyze obtained data from websites")
    fmt.Println("5 - Change the number of workers (default 15)")
    fmt.Println("6 - Launch data charts local server")
    fmt.Println()
}

func getInput() (opt int, err error) {
    _, err = fmt.Scan(&opt)
    if err != nil {
        err = fmt.Errorf("Please choose a valid option!")
        return 0, err
    }
    return
}

func changeWorkers() (int, error) {
    clearScreen()
    workersMessage()
    for {
        fmt.Println("Enter the new number of workers (0 exit to principal menu):")
        newWorkers, err := getInput()
        if err != nil {
            clearScreen()
            sorryMessage()
            fmt.Printf("\nNo valid input, must be an integer!\n\n")
            continue
        } 
        if newWorkers == 0 { 
            err := fmt.Errorf("The amount of workers did not change!")
            return 0, err 
        }
        if newWorkers < 1 || newWorkers > 100 {
            clearScreen()
            sorryMessage()
            fmt.Printf("\nThis is a bad number, please choose a number between 1 and 100!\n\n")
            continue
        }
        clearScreen()
        welcomeMessage()
        fmt.Printf("Now you have %d workers!\n\n", newWorkers)
        return newWorkers, nil
    }
}

func Menu() {
    // Number of workers in concurrent calls
    workers := 15

    clearScreen()
    welcomeMessage()
    for {
        fmt.Println("What do you want to do? (0 exit / 1 show options)")

        opt, err := getInput()
        if err != nil {
            clearScreen()
            welcomeMessage()
            fmt.Printf("%s\n\n", err)
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
            // TODO: implement the time in which the data is obtained
            clearScreen()
            loading()
            fmt.Println("\nObtaining metadata...")
            jsonm.SaveMetadata(workers)
            clearScreen()
            loaded()
            fmt.Println("\nMetadata obtained!")
            fmt.Println()
        case 4:
            clearScreen()
            dataInfo()
            jsonm.ShowWebsitesInfo()
        case 5:
            newWorkers, err := changeWorkers()
            if err != nil {
                clearScreen()
                welcomeMessage()
                fmt.Printf("%s\n\n", err)
            } else {
                workers = newWorkers
            }
        default:
            clearScreen()
            welcomeMessage()
            fmt.Printf("Please choose a valid option!\n\n")
            displayOptions()
        }
    }
}
