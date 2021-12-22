if [ "$1" == "currency-service" ]
then
  printf "\n===== Init Swagger Currency Service =====\n\n" && \
  swag init -d ./services/currency-service -o ./services/currency-service/docs --parseDependency
else
  printf "\n===== Init Swagger Currency Service =====\n\n" && \
  swag init -d ./services/currency-service -o ./services/currency-service/docs --parseDependency
fi