package menu

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Johanx22x/multicore-project/internal/ansi"
	"github.com/Johanx22x/multicore-project/internal/jsonm"
	"github.com/Johanx22x/multicore-project/internal/scrap"
)

// Clear the console screen
func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func dataInfo() {
    fmt.Printf(`%s
╔═══╗      ╔╗          ╔══╗     ╔═╗    
╚╗╔╗║     ╔╝╚╗         ╚╣╠╝     ║╔╝    
 ║║║║╔══╗ ╚╗╔╝╔══╗      ║║ ╔═╗ ╔╝╚╗╔══╗
 ║║║║╚ ╗║  ║║ ╚ ╗║      ║║ ║╔╗╗╚╗╔╝║╔╗║
╔╝╚╝║║╚╝╚╗ ║╚╗║╚╝╚╗    ╔╣╠╗║║║║ ║║ ║╚╝║
╚═══╝╚═══╝ ╚═╝╚═══╝    ╚══╝╚╝╚╝ ╚╝ ╚══╝%s%s`, ansi.Red(), ansi.Reset(), ansi.NewLine())
}

func thankYou() {
    fmt.Printf(`%s
╔═══╗                             ╔╗       ╔╗        
║╔═╗║                             ║║      ╔╝╚╗       
║╚══╗╔══╗╔══╗    ╔╗ ╔╗╔══╗╔╗╔╗    ║║ ╔══╗ ╚╗╔╝╔══╗╔═╗
╚══╗║║╔╗║║╔╗║    ║║ ║║║╔╗║║║║║    ║║ ╚ ╗║  ║║ ║╔╗║║╔╝
║╚═╝║║║═╣║║═╣    ║╚═╝║║╚╝║║╚╝║    ║╚╗║╚╝╚╗ ║╚╗║║═╣║║ 
╚═══╝╚══╝╚══╝    ╚═╗╔╝╚══╝╚══╝    ╚═╝╚═══╝ ╚═╝╚══╝╚╝ 
                 ╔═╝║                                
                 ╚══╝                                %s%s`, ansi.Red(), ansi.Reset(), ansi.NewLine())
}

func welcomeMessage() {
    fmt.Printf(`%s
╔═╗╔═╗    ╔╗  ╔╗                      ╔═══╗                          
║║╚╝║║    ║║ ╔╝╚╗                     ║╔═╗║                          
║╔╗╔╗║╔╗╔╗║║ ╚╗╔╝╔╗╔══╗╔══╗╔═╗╔══╗    ║╚══╗╔══╗╔═╗╔══╗ ╔══╗╔╗╔═╗ ╔══╗
║║║║║║║║║║║║  ║║ ╠╣║╔═╝║╔╗║║╔╝║╔╗║    ╚══╗║║╔═╝║╔╝╚ ╗║ ║╔╗║╠╣║╔╗╗║╔╗║
║║║║║║║╚╝║║╚╗ ║╚╗║║║╚═╗║╚╝║║║ ║║═╣    ║╚═╝║║╚═╗║║ ║╚╝╚╗║╚╝║║║║║║║║╚╝║
╚╝╚╝╚╝╚══╝╚═╝ ╚═╝╚╝╚══╝╚══╝╚╝ ╚══╝    ╚═══╝╚══╝╚╝ ╚═══╝║╔═╝╚╝╚╝╚╝╚═╗║
                                                       ║║        ╔═╝║
                                                       ╚╝        ╚══╝%s%s`, ansi.Red(), ansi.Reset(), ansi.NewLine())
}

func loading() {
    fmt.Printf(`%s
╔╗              ╔╗                
║║              ║║                
║║   ╔══╗╔══╗ ╔═╝║╔╗╔═╗ ╔══╗      
║║ ╔╗║╔╗║╚ ╗║ ║╔╗║╠╣║╔╗╗║╔╗║      
║╚═╝║║╚╝║║╚╝╚╗║╚╝║║║║║║║║╚╝║╔╗╔╗╔╗
╚═══╝╚══╝╚═══╝╚══╝╚╝╚╝╚╝╚═╗║╚╝╚╝╚╝
                        ╔═╝║      
                        ╚══╝      %s%s`, ansi.Red(), ansi.Reset(), ansi.NewLine())
}

func loaded() {
    fmt.Printf(`%s
╔╗              ╔╗      ╔╗
║║              ║║      ║║
║║   ╔══╗╔══╗ ╔═╝║╔══╗╔═╝║
║║ ╔╗║╔╗║╚ ╗║ ║╔╗║║╔╗║║╔╗║
║╚═╝║║╚╝║║╚╝╚╗║╚╝║║║═╣║╚╝║
╚═══╝╚══╝╚═══╝╚══╝╚══╝╚══╝%s%s`, ansi.Red(), ansi.Reset(), ansi.NewLine())
}

func sorryMessage() {
    fmt.Printf(`%s
╔═══╗                     
║╔═╗║                     
║╚══╗╔══╗╔═╗╔═╗╔╗ ╔╗      
╚══╗║║╔╗║║╔╝║╔╝║║ ║║      
║╚═╝║║╚╝║║║ ║║ ║╚═╝║╔╗╔╗╔╗
╚═══╝╚══╝╚╝ ╚╝ ╚═╗╔╝╚╝╚╝╚╝
               ╔═╝║       
               ╚══╝       %s%s`, ansi.Red(), ansi.Reset(), ansi.NewLine())
}

func workersMessage() {
    fmt.Printf(`%s
╔═╗ ╔╗              ╔╗╔╗╔╗       ╔╗             
║║╚╗║║              ║║║║║║       ║║             
║╔╗╚╝║╔══╗╔╗╔╗╔╗    ║║║║║║╔══╗╔═╗║║╔╗╔══╗╔═╗╔══╗
║║╚╗║║║╔╗║║╚╝╚╝║    ║╚╝╚╝║║╔╗║║╔╝║╚╝╝║╔╗║║╔╝║══╣
║║ ║║║║║═╣╚╗╔╗╔╝    ╚╗╔╗╔╝║╚╝║║║ ║╔╗╗║║═╣║║ ╠══║
╚╝ ╚═╝╚══╝ ╚╝╚╝      ╚╝╚╝ ╚══╝╚╝ ╚╝╚╝╚══╝╚╝ ╚══╝ 
 %s%s`, ansi.Red(), ansi.Reset(), ansi.NewLine())
}

// Display menu options
func displayOptions() {
    fmt.Printf("(%s0%s) - Exit\n", ansi.BlueUnderline(), ansi.Reset())
    fmt.Printf("(%s1%s) - Display options\n", ansi.BlueUnderline(), ansi.Reset())
    fmt.Printf("(%s2%s) - Obtain CSV table\n", ansi.BlueUnderline(), ansi.Reset())
    fmt.Printf("(%s3%s) - Obtain metadata from websites\n", ansi.BlueUnderline(), ansi.Reset())
    fmt.Printf("(%s4%s) - Analyze obtained data from websites\n", ansi.BlueUnderline(), ansi.Reset())
    fmt.Printf("(%s5%s) - Change the number of workers (default 15)\n", ansi.BlueUnderline(), ansi.Reset())
    fmt.Println()
}

// Get input, utilitie for the main menu
func getInput() (opt int, err error) {
    _, err = fmt.Scan(&opt)
    if err != nil {
        err = fmt.Errorf("Please choose a valid option!")
        return 0, err
    }
    return
}

// Chnage the current amount of worker
func changeWorkers() (int, error) {
    clearScreen()
    workersMessage()

    // Sub menu to set the workers
    for {
        fmt.Printf("Enter the new number of workers (%s0%s exit to principal menu): ", ansi.Cyan(), ansi.Reset())
        newWorkers, err := getInput()
        if err != nil {
            // NOTE: Manage input error
            clearScreen()
            sorryMessage()
            fmt.Printf("\n%sNo valid input, must be an integer!%s\n\n", ansi.CyanUnderline(), ansi.Reset())
            continue
        } 
        if newWorkers == 0 { 
            // NOTE: No changes in the workers amount
            err := fmt.Errorf("%sThe amount of workers did not change!%s", ansi.CyanUnderline(), ansi.Reset())
            return 0, err 
        }
        if newWorkers < 1 || newWorkers > 100 {
            // NOTE: No valid number
            clearScreen()
            sorryMessage()
            fmt.Printf("\n%sThis is a bad number, please choose a number between %s1%s%s and %s100%s%s!%s\n\n", ansi.Cyan(), ansi.BlueUnderline(), ansi.Reset(), ansi.Cyan(), ansi.BlueUnderline(), ansi.Reset(), ansi.Cyan(), ansi.Reset())
            continue
        }
        clearScreen()
        welcomeMessage()
        fmt.Printf("%sNow you have %d workers!%s\n\n", ansi.Blue(), newWorkers, ansi.Reset())

        return newWorkers, nil
    }
}

func Menu() {
    // Create the chart folder 
    os.Mkdir("./Data/Chart/", os.ModePerm);

    // Number of workers in concurrent calls
    // Can be modified
    workers := 15

    clearScreen()
    welcomeMessage()
    for {
        fmt.Printf("Choose an option (%s0%s exit / %s1%s show options): ", ansi.Cyan(), ansi.Reset(), ansi.Cyan(), ansi.Reset())

        // Read the desired option by the user
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
            // NOTE: Exit
            clearScreen()
            thankYou()
            fmt.Printf("\n%sThank you for use our system.\n\nHave a nice day!%s\n", ansi.Yellow(), ansi.Reset())
            return
        case 1:
            // NOTE: Show the options of the program
            clearScreen()
            welcomeMessage()
            displayOptions()
        case 2:
            // NOTE: Load the CSV table with the 1000 most visited websites
            clearScreen()
            loading()
            fmt.Println("\nLoading table...")
            scrap.ObtainCSV()
            clearScreen()
            loaded()
            fmt.Printf("\n%sTable loaded!%s\n", ansi.Blue(), ansi.Reset())
            fmt.Println()
        case 3:
            // NOTE: Obtain the medata of every website
            clearScreen()
            loading()
            fmt.Println("\nObtaining metadata...")
            jsonm.SaveMetadata(workers)
            clearScreen()
            loaded()
            fmt.Printf("\n%sMetadata obtained!%s\n", ansi.Blue(), ansi.Reset())
            fmt.Println()
        case 4:
            // NOTE: Obtain data information and create the charts
            clearScreen()
            dataInfo()
            jsonm.ShowWebsitesInfo()
        case 5:
            // NOTE: Change the number of workers that the system actually has
            newWorkers, err := changeWorkers()
            if err != nil {
                clearScreen()
                welcomeMessage()
                fmt.Printf("\n%s%s%s\n\n", ansi.Red(), err, ansi.Reset())
            } else {
                workers = newWorkers
            }
        default:
            // NOTE: default case in case of any error
            clearScreen()
            welcomeMessage()
            fmt.Printf("%sPlease choose a valid option!%s\n\n", ansi.Red(), ansi.Reset())
            displayOptions()
        }
    }
}
