package main

import "fmt"


type Notifier interface {
    Send(message string)
}


type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
    fmt.Println("Sending Email:", message)
}


type LoggingNotifier struct {
    notifier Notifier
}

func (l *LoggingNotifier) Send(message string) {
    fmt.Println("Logging: About to send notification")
    l.notifier.Send(message)
    fmt.Println("Logging: Notification sent")
}

func main() {
    
    emailNotifier := &EmailNotifier{}

    loggingNotifier := &LoggingNotifier{notifier: emailNotifier}

    loggingNotifier.Send("Hello, Decorator Pattern!")
}
