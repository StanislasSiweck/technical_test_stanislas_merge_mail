# Test technique Seqone: microservice email

## Objectif

Mettre en place un microservice de gestion de messages golang permettant l'envoi par mail de divers messages reçus. La subtilité est que l'on veut pas spammer, donc il faudrait envoyer régulièrement un mail résumant l'ensemble des messages reçus depuis le dernier envoi.

### Specs

- Une route api rest qui permet de post un message.
- Envoi d'un mail toutes les 5 minutes qui regroupe tous les messages reçus entre temps.
- Pour le moment on a qu'un seul destinataire.
- **N'ayant pour le moment pas de serveur smtp en place, il faudrait que le service utilise une `interface` qui pour le moment se contente de logger les messages dans la console. Elle doit être facilement remplaçable**.
- En bonus si le temps est disponible, un packaging docker et des tests. CI/CD si t'as vraiment le temps :D.

### Tips

- Pour installer Go: https://go.dev/doc/install
- Pour apprendre Go: https://go.dev/tour/list
- Pour t'aider un peu, nous avons commencer l'implementation. tu peux completer les TODO. Cependant tu peux aussi partir de zero si tu prefères.
- Tu peux utiliser les tickets gitlab pour documenter ton process de développement, expliquer des difficultés ou nous poser des questions.
- N'hésite pas à faire une merge request.
- Si tu ne parviens pas à tout faire en quelques heures, ce n'est pas grave, explique-nous pourquoi, mais fear not ce n'est pas bloquant pour le recrutement !
