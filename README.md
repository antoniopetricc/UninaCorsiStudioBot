
<img width="6912" height="3456" alt="UninaCorsiStudioBot" src="https://github.com/user-attachments/assets/a74e0686-ab8a-49c4-80b1-faf44ae8ff61" />

# UninaCorsiStudioBot

Un bot telegram per consultare **corsi**, **dipartimenti** e **informazioni** per l'<a href='https://unina.it'>Università degli Studi di Napoli Federico II</a>. 

Sviluppato in **Go** con [GoBotAPI](https://github.com/GoBotApiOfficial/gobotapi).

## Caratteristiche principali
- 📚 Ricerca dei **corsi di studio** con informazioni dettagliate  
- 🏛️ Visualizzazione dei **dipartimenti** universitari  
- ℹ️ Accesso rapido a **informazioni utili** riguardo agli orari delle lezioni (PROSSIMAMENTE!!!)
- ⚡ Risposte rapide e navigazione semplice direttamente da Telegram  

## Installazione

1. Clona il repository:

```bash
git clone https://github.com/antoniopetricc/UninaCorsiStudioBot.git
cd UninaCorsiStudioBot
````

2. Crea un file `.env` con le variabili di configurazione (token del bot, ecc...):

```env
BOT_API_TOKEN=TOKEN DA @BotFather
LOAD_DEPARTMENTS=true
LOAD_COURSES=true
```

### Avvio con Docker

1. Avvia il container:

```bash
docker compose up -d
```

