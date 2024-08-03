# Installation 

- install docker && docker compose
- clone services
- copy .env.ecample to .env and set configs
- first you must run auth-service bacouse of networking connection
- run `docker compose build && docker compose up -d` 
- you should try two times becouse (I didn't add shell script for waiting databse container) or (I didn't separate database to other compose)


## Unfortunately 

1. I could not finish services fully functional which I thought. I was going to add NATS message-broker and api-gateway. But I have time shoratge. I have spent working on this project only nights after work. Maybe this playing cursial role on this o    ocassion 

## However 
1. The services are very easy to extension and change any plagin or database easly without changing current core functionality. (Used Clean archetecture(some bit :), and and more good metadalogies)