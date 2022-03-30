#Sending and Receving Data
___
Lese data fra en netverktilkobling og skrive data til en netverkstilkobling  
er gjøres på samme måte man leser av og sender data til filer. (net.Conn implementerer io.ReadWriteClose interface)  
Man setter dataen i en fixed-sized buffer.  
Man bruker en bufio.Scanner for å lese dataen fra en netverkstilkobling  
til man møter en spesifikk seperator. 

###Reading Data into a Fixed Buffer
io.Reader interface som er implementert i net.Conn tillater deg til å lese data  
fra en netverkstilkobling.

Read metoden vil fylle bufferen med data til den har nådd sin kapasitet,  
dermed er det ikke sikkert at man får alt. Det vil si at payloaden som skal bli sendt  
er større enn bufferen. 
Løsningen er bufferen går i en løkke for å hente segmenter av "payloaden" helt til  
man får en io.EOF error.

***Eksempel***  
```go

buf := make([]byte, 3 1<<19) // 512 KB
payload := make([]byte, 1<<24) // 16 MB

// buff < payload
for {
    n, err := conn.Read(buf)
        if err != nil {
            if err != io.EOF {
                t.Error(err)
            }
            break
        }
    t.Logf("read %d bytes", n)
    }
```

###Delimited Reading by Using a Scanner
bufio.Scanner tillater å lese separert data. Les mer på ( https://pkg.go.dev/bufio )