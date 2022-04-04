# User Datagram Protocol
___
UDP fungerer tpå toppen av IP for å overføre datagrammer over et nettverk. UDP krever
ikke at kilden og destinasjonen etablerer et treveis håndtrykk før overføring
sånn som i TCP. I tillegg er det ikke behov for en (peer-to-peer)ende-til-ende-tilkobling.

Siden UDP unngår overhead forbundet med tilkoblinger, feilkontroller og reoverføring
av manglede data, er den egnet for sanntids- eller høyytelsesapplikasjoner som ikke
krever dataverifisering eller korrigering. Hvis en verifisering er nødvendig, kan det
utføres på applikajsonslaget.

## Fordeler og ulemper med UDP
___
* ingen reoverføringsforsinkelser - Egnet for tidssensitivie applikasjoner.
* Hastighet - Hastigheten er nyttig ved spørring der datapakker er små.
* Egnet for sendinger - mangel på peer-to-peer kommunikasjon gjør det egnet for sendinger.

UDPs manglede tilkoblingskrav og dataverifisering skaper en rekke problemer med
overføring av pakker. Blant annet:
* Ingen garantert bestilling av pakker.
* Ingen bekreftelse på at datamaskinen er klar til å motta meldingen.
* Ingen beskyttelse mot dupliserte pakker.
* Ingen garanti for at destinasjonen vil motta alle overførte byte.
UDP gir imidlertid en kontrollsum for å verifisere individuelle pakkeintegritet.

***Hentet fra:***
‘What Is UDP | From Header Structure to Packets Used in DDoS Attacks | Imperva’. Learning Center, https://www.imperva.com/learn/ddos/udp-user-datagram-protocol/. Accessed 4 Apr. 2022.