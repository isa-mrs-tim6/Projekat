# Projekat
Kombinovani projekat tima 6 iz predmeta "Internet softverske arhitekture" i "Metodologije razvoja softvera".

Za pokretanje projekta potrebno je pratiti dole napisano uputstvo.

Deployment projekta odrađen je posredstvom AWS servisa, i može se naći na: http://ec2-18-196-100-235.eu-central-1.compute.amazonaws.com:5000/

## Pokretanje backend aplikacije
Fajlovi backend aplikacije nalaze se u backend/bin direktorijumu.

Priloženi su Windows i Linux kompajlirani fajlovi.

Za korišćenje aplikacije, potrebno je instalacije PostgreSQL baze.

Komunikacija sa bazom se odvija na portu 5432, na instanci baze pod nazivom "postgres"

Pri pokretanju moguće je uneti parametre pokretanja:
* username, Korisničko ime za pristup bazi
* password, Šifra pristupa bazi
* recreate, Polje koje rekreira tabele u bazi ako se postavi na TRUE, default opcija je FALSE
* demo, Polje koje dodaje mockup objekte u bazu ako se postavi na TRUE, default opcije je FALSE
* emailDomain, Email nalog koji aplikacija koristi za slanje mailova
* emailPassword, Šifra email naloga koji se koristi za slanje mailova

## Pokretanje frontend aplikacije
Pre pokretanja, potrebno je instalirati Node.js okruženje sa npm packet managerom.

Nakon instalacije, iz direktorijuma frontend/vueApp pokrenuti komande:

```
npm install
npm run serve
```
