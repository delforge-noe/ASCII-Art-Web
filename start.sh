#!/bion/bash

go run main.go &

sleep 3

open http://localhost:8080

echo -e "\E[32m serv lancé, Ctrl+C pour quitter \E[0m"
wait

#je suis con j'ai encore oublié de push au fur et à mesure