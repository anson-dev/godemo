package main

import (
    "fmt"
    "net"
    "log"
)


func handleConn(c net.Conn) {
    defer c.Close()
    for {
        // read from the connection
        var buf = make([]byte, 10)
        log.Println("start to read from conn")
        n, err := c.Read(buf)
        if err != nil {
            log.Println("conn read error:", err)
            return
        }
        log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
    }
}

func main() {
    l, err := net.Listen("tcp", ":8888")
    if err != nil {
        fmt.Println("listen error:", err)
        return
    }

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println("accept error:", err)
            break
        }
        // start a new goroutine to handle
        // the new connection.
        go handleConn(c)
    }
}
