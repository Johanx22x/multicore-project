package ansi

// This package is an utility for ansii colors

func Black() string {
    return "\u001b[30m"
}

func Red() string {
    return "\u001b[31m"
}

func Green() string {
    return "\u001b[32m"
}

func Yellow() string {
    return "\u001b[33m"
}

func Blue() string {
    return "\u001b[34m"
}

func BlueUnderline() string {
    return "\u001b[4;34m"
}

func Purple() string {
    return "\u001b[35m"
}

func Cyan() string {
    return "\u001b[36m"
}

func CyanUnderline() string {
    return "\u001b[4;36m"
}

func White() string {
    return "\u001b[37m"
}

func Reset() string {
    return "\u001b[0m"
}

func NewLine() string {
    return "\n"
}
